package images

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/martyria/martyria/internal/models"
)

// Service orchestrates image discovery, downloading, and storage.
type Service struct {
	Pool      *pgxpool.Pool
	Wikimedia *WikimediaClient
	ImageDir  string
	BaseURL   string
}

// NewService creates an image service.
func NewService(pool *pgxpool.Pool, imageDir, baseURL string) *Service {
	return &Service{
		Pool:      pool,
		Wikimedia: NewWikimediaClient(),
		ImageDir:  imageDir,
		BaseURL:   baseURL,
	}
}

// FetchForAuthor searches Wikimedia Commons for images of an author,
// downloads the best candidates, and stores them in the database.
func (s *Service) FetchForAuthor(ctx context.Context, author models.Author) (int, error) {
	// Check if we already have images
	var count int
	s.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM images WHERE author_id = $1", author.ID).Scan(&count)
	if count > 0 {
		log.Printf("Author %s already has %d images, skipping", author.Slug, count)
		return count, nil
	}

	// Strategy 1: Search by Wikimedia category (best results)
	var results []WikimediaImage
	if author.WikimediaCategory != nil && *author.WikimediaCategory != "" {
		imgs, err := s.Wikimedia.SearchByCategory(ctx, *author.WikimediaCategory, 10)
		if err != nil {
			log.Printf("Category search failed for %s: %v", author.Slug, err)
		} else {
			results = append(results, imgs...)
		}
	}

	// Strategy 2: Fall back to name search
	if len(results) == 0 {
		query := author.Name
		if author.Title != nil {
			query = *author.Title + " " + author.Name
		}
		imgs, err := s.Wikimedia.SearchByName(ctx, query, 5)
		if err != nil {
			log.Printf("Name search failed for %s: %v", author.Slug, err)
		} else {
			results = append(results, imgs...)
		}
	}

	if len(results) == 0 {
		log.Printf("No images found for %s", author.Slug)
		return 0, nil
	}

	// Score and rank results
	scored := scoreImages(results, author)

	// Download and store top results (up to 3)
	maxStore := 3
	if len(scored) < maxStore {
		maxStore = len(scored)
	}

	stored := 0
	for i := 0; i < maxStore; i++ {
		img := scored[i]

		// Download full image
		localPath, err := s.Wikimedia.DownloadImage(ctx, img.image.URL, s.ImageDir, author.Slug)
		if err != nil {
			log.Printf("Failed to download image for %s: %v", author.Slug, err)
			continue
		}

		// Download thumbnail
		var thumbPath *string
		if img.image.ThumbURL != "" {
			tp, err := s.Wikimedia.DownloadThumbnail(ctx, img.image.ThumbURL, s.ImageDir, author.Slug)
			if err == nil {
				thumbPath = &tp
			}
		}

		// Detect style from title/description
		style := detectStyle(img.image.Title)

		// Store in database
		isPrimary := (i == 0) // First (highest-scored) image is primary
		_, err = s.Pool.Exec(ctx, `
			INSERT INTO images (author_id, source_type, source_url, source_attribution, source_license,
				style, width, height, mime_type, local_path, thumbnail_path,
				is_ai_generated, is_primary, quality_score)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		`,
			author.ID,
			models.ImageSourceWikimedia,
			img.image.DescriptionURL,
			img.image.Attribution,
			img.image.License,
			style,
			img.image.Width,
			img.image.Height,
			img.image.MimeType,
			localPath,
			thumbPath,
			false,      // is_ai_generated
			isPrimary,  // is_primary
			img.score,  // quality_score
		)
		if err != nil {
			log.Printf("Failed to store image for %s: %v", author.Slug, err)
			continue
		}
		stored++
	}

	log.Printf("Stored %d images for %s (from %d candidates)", stored, author.Slug, len(results))
	return stored, nil
}

// FetchAllAuthors fetches images for all authors that don't have any yet.
func (s *Service) FetchAllAuthors(ctx context.Context) error {
	rows, err := s.Pool.Query(ctx, `
		SELECT a.id, a.slug, a.name, a.title, a.era, a.tradition,
			a.canonized, a.copyright_status, a.wikimedia_category
		FROM authors a
		WHERE NOT EXISTS (SELECT 1 FROM images WHERE author_id = a.id)
		ORDER BY a.id
	`)
	if err != nil {
		return fmt.Errorf("list authors without images: %w", err)
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		a := models.Author{}
		if err := rows.Scan(
			&a.ID, &a.Slug, &a.Name, &a.Title, &a.Era, &a.Tradition,
			&a.Canonized, &a.CopyrightStatus, &a.WikimediaCategory,
		); err != nil {
			return fmt.Errorf("scan author: %w", err)
		}
		authors = append(authors, a)
	}

	log.Printf("Fetching images for %d authors...", len(authors))

	for _, author := range authors {
		if _, err := s.FetchForAuthor(ctx, author); err != nil {
			log.Printf("Error fetching images for %s: %v", author.Slug, err)
			// Continue with next author, don't abort
		}
	}

	return nil
}

// --- Scoring ---

type scoredImage struct {
	image WikimediaImage
	score int
}

// scoreImages ranks images by relevance for a given author.
func scoreImages(images []WikimediaImage, author models.Author) []scoredImage {
	var scored []scoredImage
	for _, img := range images {
		score := 0
		titleLower := strings.ToLower(img.Title)
		nameLower := strings.ToLower(author.Name)

		// Name appears in title → very relevant
		if strings.Contains(titleLower, nameLower) {
			score += 30
		}

		// Icon/fresco/mosaic preferred for saints
		if author.Canonized {
			if strings.Contains(titleLower, "icon") {
				score += 25
			}
			if strings.Contains(titleLower, "fresco") {
				score += 15
			}
			if strings.Contains(titleLower, "mosaic") {
				score += 15
			}
		}

		// Good resolution (but not excessively large)
		if img.Width >= 400 && img.Width <= 4000 {
			score += 10
		}
		if img.Width >= 800 {
			score += 5
		}

		// Prefer square-ish or portrait orientation (better for icons)
		if img.Height > 0 && img.Width > 0 {
			ratio := float64(img.Height) / float64(img.Width)
			if ratio >= 0.8 && ratio <= 1.5 {
				score += 10
			}
		}

		// Known good license
		licenseLower := strings.ToLower(img.License)
		if strings.Contains(licenseLower, "public domain") || strings.Contains(licenseLower, "pd") {
			score += 10
		}
		if strings.Contains(licenseLower, "cc") {
			score += 5
		}

		// Prefer JPEG for photos, PNG for icons
		if img.MimeType == "image/jpeg" {
			score += 3
		}

		// Penalize SVG/non-raster
		if img.MimeType == "image/svg+xml" {
			score -= 10
		}

		scored = append(scored, scoredImage{image: img, score: score})
	}

	// Sort by score descending
	for i := 0; i < len(scored); i++ {
		for j := i + 1; j < len(scored); j++ {
			if scored[j].score > scored[i].score {
				scored[i], scored[j] = scored[j], scored[i]
			}
		}
	}

	return scored
}

// detectStyle guesses the art style from the file title.
func detectStyle(title string) models.ImageStyle {
	t := strings.ToLower(title)
	switch {
	case strings.Contains(t, "icon"):
		return models.StyleByzantineIcon
	case strings.Contains(t, "fresco"):
		return models.StyleFresco
	case strings.Contains(t, "mosaic"):
		return models.StyleMosaic
	case strings.Contains(t, "manuscript") || strings.Contains(t, "miniature"):
		return models.StyleManuscript
	case strings.Contains(t, "engrav") || strings.Contains(t, "woodcut"):
		return models.StyleEngraving
	case strings.Contains(t, "photo"):
		return models.StylePhotograph
	case strings.Contains(t, "paint") || strings.Contains(t, "oil"):
		return models.StyleOilPainting
	default:
		return models.StyleOther
	}
}

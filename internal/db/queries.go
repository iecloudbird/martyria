package db

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/martyria/martyria/internal/models"
)

// --- Authors ---

func (d *DB) GetAuthor(ctx context.Context, slug string) (*models.Author, error) {
	a := &models.Author{}
	err := d.Pool.QueryRow(ctx, `
		SELECT a.id, a.slug, a.name, a.name_original, a.title,
			a.born_year, a.died_year, a.era, a.tradition,
			a.bio, a.bio_short, a.canonized, a.canonized_date, a.canonized_by,
			a.feast_day_orthodox, a.feast_day_catholic, a.copyright_status,
			a.wikipedia_url, a.wikimedia_category,
			a.created_at, a.updated_at,
			(SELECT COUNT(*) FROM quotes WHERE author_id = a.id) as quote_count
		FROM authors a WHERE a.slug = $1
	`, slug).Scan(
		&a.ID, &a.Slug, &a.Name, &a.NameOriginal, &a.Title,
		&a.BornYear, &a.DiedYear, &a.Era, &a.Tradition,
		&a.Bio, &a.BioShort, &a.Canonized, &a.CanonizedDate, &a.CanonizedBy,
		&a.FeastDayOrthodox, &a.FeastDayCatholic, &a.CopyrightStatus,
		&a.WikipediaURL, &a.WikimediaCategory,
		&a.CreatedAt, &a.UpdatedAt,
		&a.QuoteCount,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get author: %w", err)
	}
	return a, nil
}

func (d *DB) ListAuthors(ctx context.Context, f models.AuthorFilter) ([]models.Author, int64, error) {
	if f.Page < 1 {
		f.Page = 1
	}
	if f.PerPage < 1 || f.PerPage > 100 {
		f.PerPage = 20
	}

	where := []string{"1=1"}
	args := []interface{}{}
	argN := 1

	if f.Era != "" {
		where = append(where, fmt.Sprintf("a.era = $%d", argN))
		args = append(args, f.Era)
		argN++
	}
	if f.Tradition != "" {
		where = append(where, fmt.Sprintf("a.tradition = $%d", argN))
		args = append(args, f.Tradition)
		argN++
	}
	if f.Search != "" {
		where = append(where, fmt.Sprintf("(a.name ILIKE $%d OR a.bio_short ILIKE $%d)", argN, argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}

	whereClause := strings.Join(where, " AND ")

	// Count
	var total int64
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM authors a WHERE %s", whereClause)
	if err := d.Pool.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count authors: %w", err)
	}

	// Fetch
	offset := (f.Page - 1) * f.PerPage
	query := fmt.Sprintf(`
		SELECT a.id, a.slug, a.name, a.name_original, a.title,
			a.born_year, a.died_year, a.era, a.tradition,
			a.bio_short, a.canonized, a.copyright_status,
			a.feast_day_orthodox, a.feast_day_catholic,
			a.created_at, a.updated_at,
			(SELECT COUNT(*) FROM quotes WHERE author_id = a.id) as quote_count
		FROM authors a
		WHERE %s
		ORDER BY a.born_year ASC NULLS LAST, a.name ASC
		LIMIT $%d OFFSET $%d
	`, whereClause, argN, argN+1)
	args = append(args, f.PerPage, offset)

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list authors: %w", err)
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		a := models.Author{}
		if err := rows.Scan(
			&a.ID, &a.Slug, &a.Name, &a.NameOriginal, &a.Title,
			&a.BornYear, &a.DiedYear, &a.Era, &a.Tradition,
			&a.BioShort, &a.Canonized, &a.CopyrightStatus,
			&a.FeastDayOrthodox, &a.FeastDayCatholic,
			&a.CreatedAt, &a.UpdatedAt,
			&a.QuoteCount,
		); err != nil {
			return nil, 0, fmt.Errorf("scan author: %w", err)
		}
		authors = append(authors, a)
	}

	return authors, total, nil
}

// --- Quotes ---

func (d *DB) GetQuote(ctx context.Context, id int64) (*models.Quote, error) {
	q := &models.Quote{}
	a := &models.Author{}
	err := d.Pool.QueryRow(ctx, `
		SELECT q.id, q.author_id, q.text, q.text_original, q.language,
			q.source_work, q.source_chapter, q.source_publisher, q.source_page, q.source_url,
			q.license, q.verified, q.created_at, q.updated_at,
			a.id, a.slug, a.name, a.era, a.tradition, a.bio_short, a.copyright_status
		FROM quotes q
		JOIN authors a ON a.id = q.author_id
		WHERE q.id = $1
	`, id).Scan(
		&q.ID, &q.AuthorID, &q.Text, &q.TextOriginal, &q.Language,
		&q.SourceWork, &q.SourceChapter, &q.SourcePublisher, &q.SourcePage, &q.SourceURL,
		&q.License, &q.Verified, &q.CreatedAt, &q.UpdatedAt,
		&a.ID, &a.Slug, &a.Name, &a.Era, &a.Tradition, &a.BioShort, &a.CopyrightStatus,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get quote: %w", err)
	}
	q.Author = a
	q.Attribution = buildAttribution(q, a)

	// Fetch topics
	topics, err := d.getQuoteTopics(ctx, q.ID)
	if err != nil {
		return nil, err
	}
	q.Topics = topics

	return q, nil
}

func (d *DB) GetRandomQuote(ctx context.Context, f models.QuoteFilter) (*models.Quote, error) {
	where, args := buildQuoteWhere(f)
	query := fmt.Sprintf(`
		SELECT q.id, q.author_id, q.text, q.text_original, q.language,
			q.source_work, q.source_chapter, q.source_publisher, q.source_page, q.source_url,
			q.license, q.verified, q.created_at, q.updated_at,
			a.id, a.slug, a.name, a.era, a.tradition, a.bio_short, a.copyright_status
		FROM quotes q
		JOIN authors a ON a.id = q.author_id
		%s
		ORDER BY RANDOM()
		LIMIT 1
	`, where)

	q := &models.Quote{}
	a := &models.Author{}
	err := d.Pool.QueryRow(ctx, query, args...).Scan(
		&q.ID, &q.AuthorID, &q.Text, &q.TextOriginal, &q.Language,
		&q.SourceWork, &q.SourceChapter, &q.SourcePublisher, &q.SourcePage, &q.SourceURL,
		&q.License, &q.Verified, &q.CreatedAt, &q.UpdatedAt,
		&a.ID, &a.Slug, &a.Name, &a.Era, &a.Tradition, &a.BioShort, &a.CopyrightStatus,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("random quote: %w", err)
	}
	q.Author = a
	q.Attribution = buildAttribution(q, a)
	return q, nil
}

func (d *DB) ListQuotes(ctx context.Context, f models.QuoteFilter) ([]models.Quote, int64, error) {
	if f.Page < 1 {
		f.Page = 1
	}
	if f.PerPage < 1 || f.PerPage > 100 {
		f.PerPage = 20
	}

	where, args := buildQuoteWhere(f)

	var total int64
	countQ := fmt.Sprintf("SELECT COUNT(*) FROM quotes q JOIN authors a ON a.id = q.author_id %s", where)
	if err := d.Pool.QueryRow(ctx, countQ, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count quotes: %w", err)
	}

	offset := (f.Page - 1) * f.PerPage
	argN := len(args) + 1
	query := fmt.Sprintf(`
		SELECT q.id, q.author_id, q.text, q.language,
			q.source_work, q.source_chapter, q.license, q.verified,
			q.created_at, q.updated_at,
			a.id, a.slug, a.name, a.era, a.tradition, a.copyright_status
		FROM quotes q
		JOIN authors a ON a.id = q.author_id
		%s
		ORDER BY q.id ASC
		LIMIT $%d OFFSET $%d
	`, where, argN, argN+1)
	args = append(args, f.PerPage, offset)

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list quotes: %w", err)
	}
	defer rows.Close()

	var quotes []models.Quote
	for rows.Next() {
		q := models.Quote{}
		a := models.Author{}
		if err := rows.Scan(
			&q.ID, &q.AuthorID, &q.Text, &q.Language,
			&q.SourceWork, &q.SourceChapter, &q.License, &q.Verified,
			&q.CreatedAt, &q.UpdatedAt,
			&a.ID, &a.Slug, &a.Name, &a.Era, &a.Tradition, &a.CopyrightStatus,
		); err != nil {
			return nil, 0, fmt.Errorf("scan quote: %w", err)
		}
		q.Author = &a
		q.Attribution = buildAttribution(&q, &a)
		quotes = append(quotes, q)
	}
	return quotes, total, nil
}

func (d *DB) GetDailyQuote(ctx context.Context, date time.Time) (*models.Quote, *string, error) {
	dateStr := date.Format("2006-01-02")

	// Check for a scheduled daily quote
	var quoteID int64
	var reason *string
	err := d.Pool.QueryRow(ctx,
		"SELECT quote_id, reason FROM daily_quotes WHERE date = $1", dateStr,
	).Scan(&quoteID, &reason)

	if err == pgx.ErrNoRows {
		// Fallback: deterministic pseudo-random based on date
		var total int64
		d.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM quotes WHERE verified = true").Scan(&total)
		if total == 0 {
			return nil, nil, nil
		}

		// Simple hash of date to pick a quote
		dayNum := date.YearDay() + date.Year()*366
		offset := dayNum % int(total)

		err = d.Pool.QueryRow(ctx, `
			SELECT id FROM quotes WHERE verified = true ORDER BY id LIMIT 1 OFFSET $1
		`, offset).Scan(&quoteID)
		if err != nil {
			return nil, nil, fmt.Errorf("daily quote fallback: %w", err)
		}
	} else if err != nil {
		return nil, nil, fmt.Errorf("daily quote: %w", err)
	}

	quote, err := d.GetQuote(ctx, quoteID)
	if err != nil {
		return nil, nil, err
	}
	return quote, reason, nil
}

// --- Topics ---

func (d *DB) ListTopics(ctx context.Context) ([]models.Topic, error) {
	rows, err := d.Pool.Query(ctx, `
		SELECT t.id, t.slug, t.name, t.description,
			(SELECT COUNT(*) FROM quote_topics qt WHERE qt.topic_id = t.id) as quote_count
		FROM topics t
		ORDER BY t.name ASC
	`)
	if err != nil {
		return nil, fmt.Errorf("list topics: %w", err)
	}
	defer rows.Close()

	var topics []models.Topic
	for rows.Next() {
		t := models.Topic{}
		if err := rows.Scan(&t.ID, &t.Slug, &t.Name, &t.Description, &t.QuoteCount); err != nil {
			return nil, fmt.Errorf("scan topic: %w", err)
		}
		topics = append(topics, t)
	}
	return topics, nil
}

// --- Images ---

func (d *DB) GetPrimaryImage(ctx context.Context, authorID int64) (*models.Image, error) {
	img := &models.Image{}
	err := d.Pool.QueryRow(ctx, `
		SELECT id, author_id, source_type, source_url, source_attribution, source_license,
			style, width, height, mime_type, local_path, thumbnail_path,
			is_ai_generated, is_primary, quality_score, created_at
		FROM images
		WHERE author_id = $1 AND is_primary = true
		LIMIT 1
	`, authorID).Scan(
		&img.ID, &img.AuthorID, &img.SourceType, &img.SourceURL, &img.SourceAttribution, &img.SourceLicense,
		&img.Style, &img.Width, &img.Height, &img.MimeType, &img.LocalPath, &img.ThumbnailPath,
		&img.IsAIGenerated, &img.IsPrimary, &img.QualityScore, &img.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get primary image: %w", err)
	}
	return img, nil
}

// --- Helpers ---

func buildQuoteWhere(f models.QuoteFilter) (string, []interface{}) {
	where := []string{"1=1"}
	args := []interface{}{}
	argN := 1

	if f.AuthorSlug != "" {
		where = append(where, fmt.Sprintf("a.slug = $%d", argN))
		args = append(args, f.AuthorSlug)
		argN++
	}
	if f.TopicSlug != "" {
		where = append(where, fmt.Sprintf("EXISTS (SELECT 1 FROM quote_topics qt JOIN topics t ON t.id = qt.topic_id WHERE qt.quote_id = q.id AND t.slug = $%d)", argN))
		args = append(args, f.TopicSlug)
		argN++
	}
	if f.Era != "" {
		where = append(where, fmt.Sprintf("a.era = $%d", argN))
		args = append(args, f.Era)
		argN++
	}
	if f.Tradition != "" {
		where = append(where, fmt.Sprintf("a.tradition = $%d", argN))
		args = append(args, f.Tradition)
		argN++
	}
	if f.Verified != nil {
		where = append(where, fmt.Sprintf("q.verified = $%d", argN))
		args = append(args, *f.Verified)
		argN++
	}
	if f.Language != "" {
		where = append(where, fmt.Sprintf("q.language = $%d", argN))
		args = append(args, f.Language)
		argN++
	}

	return "WHERE " + strings.Join(where, " AND "), args
}

func buildAttribution(q *models.Quote, a *models.Author) *string {
	if a.CopyrightStatus != models.CopyrightFairUse {
		return nil
	}

	parts := []string{}
	if q.SourceWork != nil {
		parts = append(parts, fmt.Sprintf("From '%s'", *q.SourceWork))
	}
	if q.SourcePublisher != nil {
		parts = append(parts, *q.SourcePublisher)
	}
	if len(parts) == 0 {
		return nil
	}
	attr := strings.Join(parts, ", ")
	return &attr
}

func (d *DB) getQuoteTopics(ctx context.Context, quoteID int64) ([]models.Topic, error) {
	rows, err := d.Pool.Query(ctx, `
		SELECT t.id, t.slug, t.name FROM topics t
		JOIN quote_topics qt ON qt.topic_id = t.id
		WHERE qt.quote_id = $1
		ORDER BY t.name
	`, quoteID)
	if err != nil {
		return nil, fmt.Errorf("get quote topics: %w", err)
	}
	defer rows.Close()

	var topics []models.Topic
	for rows.Next() {
		t := models.Topic{}
		if err := rows.Scan(&t.ID, &t.Slug, &t.Name); err != nil {
			return nil, fmt.Errorf("scan topic: %w", err)
		}
		topics = append(topics, t)
	}
	return topics, nil
}

// TotalPages helper
func TotalPages(total int64, perPage int) int {
	return int(math.Ceil(float64(total) / float64(perPage)))
}

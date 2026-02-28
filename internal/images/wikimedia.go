package images

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// WikimediaClient searches and downloads images from Wikimedia Commons.
type WikimediaClient struct {
	HTTPClient *http.Client
	UserAgent  string // Wikimedia requires a descriptive User-Agent
}

// WikimediaImage represents a single image result from the Commons API.
type WikimediaImage struct {
	Title         string // e.g. "File:Ikon of Saint John Chrysostom.jpg"
	PageID        int
	URL           string // Direct image URL
	ThumbURL      string // Thumbnail URL
	Width         int
	Height        int
	MimeType      string
	License       string
	Attribution   string
	DescriptionURL string // Page on Commons
}

// NewWikimediaClient creates a client for the Wikimedia Commons API.
func NewWikimediaClient() *WikimediaClient {
	return &WikimediaClient{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		UserAgent:  "Martyria/0.1 (https://github.com/iecloudbird/martyria; christian-heritage-project) Go-http-client",
	}
}

// SearchByCategory finds images in a Wikimedia Commons category.
// This is the primary method — author records store their wikimedia_category.
func (wc *WikimediaClient) SearchByCategory(ctx context.Context, category string, limit int) ([]WikimediaImage, error) {
	if category == "" {
		return nil, nil
	}
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	// Step 1: Get file pages in the category
	params := url.Values{
		"action":  {"query"},
		"format":  {"json"},
		"list":    {"categorymembers"},
		"cmtitle": {"Category:" + category},
		"cmtype":  {"file"},
		"cmlimit": {fmt.Sprintf("%d", limit)},
	}

	body, err := wc.apiGet(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("category members: %w", err)
	}

	var catResp struct {
		Query struct {
			CategoryMembers []struct {
				PageID int    `json:"pageid"`
				Title  string `json:"title"`
			} `json:"categorymembers"`
		} `json:"query"`
	}
	if err := json.Unmarshal(body, &catResp); err != nil {
		return nil, fmt.Errorf("parse category response: %w", err)
	}

	if len(catResp.Query.CategoryMembers) == 0 {
		return nil, nil
	}

	// Step 2: Get image info for all files
	titles := make([]string, len(catResp.Query.CategoryMembers))
	for i, m := range catResp.Query.CategoryMembers {
		titles[i] = m.Title
	}

	return wc.getImageInfo(ctx, titles)
}

// SearchByName searches Wikimedia Commons for images matching a query string.
// Fallback when no category is set on the author.
func (wc *WikimediaClient) SearchByName(ctx context.Context, query string, limit int) ([]WikimediaImage, error) {
	if query == "" {
		return nil, nil
	}
	if limit <= 0 || limit > 20 {
		limit = 5
	}

	params := url.Values{
		"action":   {"query"},
		"format":   {"json"},
		"generator": {"search"},
		"gsrsearch": {query + " icon OR saint OR fresco OR painting"},
		"gsrnamespace": {"6"}, // File namespace
		"gsrlimit":  {fmt.Sprintf("%d", limit)},
		"prop":      {"imageinfo"},
		"iiprop":    {"url|size|mime|extmetadata"},
		"iiurlwidth": {"800"},
	}

	body, err := wc.apiGet(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("search: %w", err)
	}

	return wc.parseImageInfoResponse(body)
}

// DownloadImage downloads an image to the local filesystem.
// Returns the local path relative to imageDir.
func (wc *WikimediaClient) DownloadImage(ctx context.Context, imageURL, imageDir, authorSlug string) (string, error) {
	// Create author directory
	authorDir := filepath.Join(imageDir, authorSlug)
	if err := os.MkdirAll(authorDir, 0o755); err != nil {
		return "", fmt.Errorf("mkdir: %w", err)
	}

	// Determine filename from URL
	parsed, err := url.Parse(imageURL)
	if err != nil {
		return "", fmt.Errorf("parse URL: %w", err)
	}
	filename := filepath.Base(parsed.Path)
	// Sanitize filename
	filename = strings.ReplaceAll(filename, " ", "_")
	localPath := filepath.Join(authorSlug, filename)
	fullPath := filepath.Join(imageDir, localPath)

	// Skip if already downloaded
	if _, err := os.Stat(fullPath); err == nil {
		log.Printf("Image already exists: %s", localPath)
		return localPath, nil
	}

	// Download
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageURL, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("User-Agent", wc.UserAgent)

	resp, err := wc.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download status %d for %s", resp.StatusCode, imageURL)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("create file: %w", err)
	}
	defer file.Close()

	written, err := io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("write file: %w", err)
	}

	log.Printf("Downloaded %s (%d bytes)", localPath, written)
	return localPath, nil
}

// DownloadThumbnail downloads a thumbnail version of an image.
func (wc *WikimediaClient) DownloadThumbnail(ctx context.Context, thumbURL, imageDir, authorSlug string) (string, error) {
	authorDir := filepath.Join(imageDir, authorSlug, "thumbs")
	if err := os.MkdirAll(authorDir, 0o755); err != nil {
		return "", fmt.Errorf("mkdir thumbs: %w", err)
	}

	parsed, err := url.Parse(thumbURL)
	if err != nil {
		return "", fmt.Errorf("parse thumb URL: %w", err)
	}
	filename := filepath.Base(parsed.Path)
	filename = strings.ReplaceAll(filename, " ", "_")
	localPath := filepath.Join(authorSlug, "thumbs", filename)
	fullPath := filepath.Join(imageDir, localPath)

	if _, err := os.Stat(fullPath); err == nil {
		return localPath, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, thumbURL, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("User-Agent", wc.UserAgent)

	resp, err := wc.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("download thumb: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("thumb status %d", resp.StatusCode)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("create thumb file: %w", err)
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	return localPath, nil
}

// --- Internal helpers ---

func (wc *WikimediaClient) apiGet(ctx context.Context, params url.Values) ([]byte, error) {
	u := "https://commons.wikimedia.org/w/api.php?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", wc.UserAgent)

	resp, err := wc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wikimedia API status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (wc *WikimediaClient) getImageInfo(ctx context.Context, titles []string) ([]WikimediaImage, error) {
	// Batch titles (API limit: 50 per request)
	params := url.Values{
		"action":     {"query"},
		"format":     {"json"},
		"titles":     {strings.Join(titles, "|")},
		"prop":       {"imageinfo"},
		"iiprop":     {"url|size|mime|extmetadata"},
		"iiurlwidth": {"800"}, // Request 800px thumbnail
	}

	body, err := wc.apiGet(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("image info: %w", err)
	}

	return wc.parseImageInfoResponse(body)
}

func (wc *WikimediaClient) parseImageInfoResponse(body []byte) ([]WikimediaImage, error) {
	var resp struct {
		Query struct {
			Pages map[string]struct {
				PageID    int    `json:"pageid"`
				Title     string `json:"title"`
				ImageInfo []struct {
					URL            string `json:"url"`
					ThumbURL       string `json:"thumburl"`
					Width          int    `json:"width"`
					Height         int    `json:"height"`
					Mime           string `json:"mime"`
					DescriptionURL string `json:"descriptionurl"`
					ExtMetadata    struct {
						License struct {
							Value string `json:"value"`
						} `json:"LicenseShortName"`
						Artist struct {
							Value string `json:"value"`
						} `json:"Artist"`
					} `json:"extmetadata"`
				} `json:"imageinfo"`
			} `json:"pages"`
		} `json:"query"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parse image info: %w", err)
	}

	var images []WikimediaImage
	for _, page := range resp.Query.Pages {
		if page.PageID <= 0 || len(page.ImageInfo) == 0 {
			continue
		}
		ii := page.ImageInfo[0]

		// Filter: only actual images
		if !strings.HasPrefix(ii.Mime, "image/") {
			continue
		}

		img := WikimediaImage{
			Title:          page.Title,
			PageID:         page.PageID,
			URL:            ii.URL,
			ThumbURL:       ii.ThumbURL,
			Width:          ii.Width,
			Height:         ii.Height,
			MimeType:       ii.Mime,
			License:        ii.ExtMetadata.License.Value,
			Attribution:    stripHTML(ii.ExtMetadata.Artist.Value),
			DescriptionURL: ii.DescriptionURL,
		}
		images = append(images, img)
	}

	return images, nil
}

// stripHTML removes basic HTML tags from attribution strings.
func stripHTML(s string) string {
	var result strings.Builder
	inTag := false
	for _, r := range s {
		if r == '<' {
			inTag = true
			continue
		}
		if r == '>' {
			inTag = false
			continue
		}
		if !inTag {
			result.WriteRune(r)
		}
	}
	return strings.TrimSpace(result.String())
}

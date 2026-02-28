package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/martyria/martyria/internal/config"
	"github.com/martyria/martyria/internal/db"
	"github.com/martyria/martyria/internal/images"
	"github.com/martyria/martyria/internal/models"
)

// Handler holds dependencies for all HTTP handlers.
type Handler struct {
	DB       *db.DB
	Config   *config.Config
	ImageSvc *images.Service
}

func NewHandler(database *db.DB, cfg *config.Config, imgSvc *images.Service) *Handler {
	return &Handler{DB: database, Config: cfg, ImageSvc: imgSvc}
}

// --- Health ---

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, models.HealthResponse{
		Status:  "ok",
		Version: h.Config.Version,
	})
}

// --- Authors ---

func (h *Handler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(queryParam(r, "page", "1"))
	perPage, _ := strconv.Atoi(queryParam(r, "per_page", "20"))

	f := models.AuthorFilter{
		Era:       queryParam(r, "era", ""),
		Tradition: queryParam(r, "tradition", ""),
		Search:    queryParam(r, "search", ""),
		Page:      page,
		PerPage:   perPage,
	}

	authors, total, err := h.DB.ListAuthors(r.Context(), f)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Error:   "failed to list authors",
			Message: err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, models.PaginatedResponse{
		Data:       authors,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int64(db.TotalPages(total, perPage)),
	})
}

func (h *Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "slug required"})
		return
	}

	author, err := h.DB.GetAuthor(r.Context(), slug)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if author == nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "author not found"})
		return
	}

	// Attach primary image
	img, _ := h.DB.GetPrimaryImage(r.Context(), author.ID)
	if img != nil {
		imageURL := fmt.Sprintf("%s/data/images/%s", h.Config.BaseURL, img.LocalPath)
		author.ImageURL = &imageURL
	}

	writeJSON(w, http.StatusOK, author)
}

func (h *Handler) GetAuthorQuotes(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	page, _ := strconv.Atoi(queryParam(r, "page", "1"))
	perPage, _ := strconv.Atoi(queryParam(r, "per_page", "20"))

	f := models.QuoteFilter{
		AuthorSlug: slug,
		Page:       page,
		PerPage:    perPage,
	}

	quotes, total, err := h.DB.ListQuotes(r.Context(), f)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, models.PaginatedResponse{
		Data:       quotes,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int64(db.TotalPages(total, perPage)),
	})
}

// --- Quotes ---

func (h *Handler) ListQuotes(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(queryParam(r, "page", "1"))
	perPage, _ := strconv.Atoi(queryParam(r, "per_page", "20"))

	f := models.QuoteFilter{
		AuthorSlug: queryParam(r, "author", ""),
		TopicSlug:  queryParam(r, "topic", ""),
		Era:        queryParam(r, "era", ""),
		Tradition:  queryParam(r, "tradition", ""),
		Language:   queryParam(r, "language", ""),
		Page:       page,
		PerPage:    perPage,
	}

	// Parse verified filter
	if v := queryParam(r, "verified", ""); v != "" {
		b := v == "true"
		f.Verified = &b
	}

	quotes, total, err := h.DB.ListQuotes(r.Context(), f)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, models.PaginatedResponse{
		Data:       quotes,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int64(db.TotalPages(total, perPage)),
	})
}

func (h *Handler) GetQuote(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "invalid quote id"})
		return
	}

	quote, err := h.DB.GetQuote(r.Context(), id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if quote == nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "quote not found"})
		return
	}

	writeJSON(w, http.StatusOK, quote)
}

func (h *Handler) RandomQuote(w http.ResponseWriter, r *http.Request) {
	f := models.QuoteFilter{
		AuthorSlug: queryParam(r, "author", ""),
		TopicSlug:  queryParam(r, "topic", ""),
		Era:        queryParam(r, "era", ""),
		Tradition:  queryParam(r, "tradition", ""),
	}

	quote, err := h.DB.GetRandomQuote(r.Context(), f)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if quote == nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "no quotes found"})
		return
	}

	writeJSON(w, http.StatusOK, quote)
}

func (h *Handler) DailyQuote(w http.ResponseWriter, r *http.Request) {
	dateStr := queryParam(r, "date", "")
	date := time.Now()
	if dateStr != "" {
		parsed, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			date = parsed
		}
	}

	quote, reason, err := h.DB.GetDailyQuote(r.Context(), date)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if quote == nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "no daily quote available"})
		return
	}

	resp := models.QuoteOfTheDay{
		Quote: *quote,
		Date:  date.Format("2006-01-02"),
	}
	if reason != nil {
		resp.Reason = *reason
	}

	writeJSON(w, http.StatusOK, resp)
}

// --- Topics ---

func (h *Handler) ListTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := h.DB.ListTopics(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, topics)
}

func (h *Handler) GetTopicQuotes(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	page, _ := strconv.Atoi(queryParam(r, "page", "1"))
	perPage, _ := strconv.Atoi(queryParam(r, "per_page", "20"))

	f := models.QuoteFilter{
		TopicSlug: slug,
		Page:      page,
		PerPage:   perPage,
	}

	quotes, total, err := h.DB.ListQuotes(r.Context(), f)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, models.PaginatedResponse{
		Data:       quotes,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int64(db.TotalPages(total, perPage)),
	})
}

// --- Images ---

func (h *Handler) GetAuthorImages(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	author, err := h.DB.GetAuthor(r.Context(), slug)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if author == nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "author not found"})
		return
	}

	imgs, err := h.DB.GetAuthorImages(r.Context(), author.ID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Populate URLs
	for i := range imgs {
		if imgs[i].LocalPath != nil {
			u := fmt.Sprintf("%s/data/images/%s", h.Config.BaseURL, *imgs[i].LocalPath)
			imgs[i].FullURL = u
		}
		if imgs[i].ThumbnailPath != nil {
			u := fmt.Sprintf("%s/data/images/%s", h.Config.BaseURL, *imgs[i].ThumbnailPath)
			imgs[i].ThumbnailURL = u
		}
	}

	writeJSON(w, http.StatusOK, imgs)
}

func (h *Handler) FetchAllImages(w http.ResponseWriter, r *http.Request) {
	if h.ImageSvc == nil {
		writeJSON(w, http.StatusServiceUnavailable, models.ErrorResponse{Error: "image service not configured"})
		return
	}

	go func() {
		if err := h.ImageSvc.FetchAllAuthors(context.Background()); err != nil {
			log.Printf("Batch image fetch error: %v", err)
		}
	}()

	writeJSON(w, http.StatusAccepted, map[string]string{
		"status":  "accepted",
		"message": "Image fetch started for all authors without images",
	})
}

func (h *Handler) FetchAuthorImages(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if h.ImageSvc == nil {
		writeJSON(w, http.StatusServiceUnavailable, models.ErrorResponse{Error: "image service not configured"})
		return
	}

	author, err := h.DB.GetAuthor(r.Context(), slug)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if author == nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "author not found"})
		return
	}

	count, err := h.ImageSvc.FetchForAuthor(r.Context(), *author)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"author":       slug,
		"images_found": count,
	})
}

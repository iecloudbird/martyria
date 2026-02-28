package models

import "time"

type AuthorEra string

const (
	EraApostolic    AuthorEra = "apostolic"
	EraAnteNicene   AuthorEra = "ante_nicene"
	EraNicene       AuthorEra = "nicene"
	EraPostNicene   AuthorEra = "post_nicene"
	EraMedieval     AuthorEra = "medieval"
	EraReformation  AuthorEra = "reformation"
	EraModern       AuthorEra = "modern"
	EraContemporary AuthorEra = "contemporary"
)

type AuthorTradition string

const (
	TraditionPreSchism          AuthorTradition = "pre_schism"
	TraditionOrthodox           AuthorTradition = "orthodox"
	TraditionCatholic           AuthorTradition = "catholic"
	TraditionProtestant         AuthorTradition = "protestant"
	TraditionAnglican           AuthorTradition = "anglican"
	TraditionNonDenominational  AuthorTradition = "non_denominational"
)

type CopyrightStatus string

const (
	CopyrightPublicDomain    CopyrightStatus = "public_domain"
	CopyrightFairUse         CopyrightStatus = "short_quote_fair_use"
	CopyrightPermissionGrant CopyrightStatus = "permission_granted"
	CopyrightCCBYSA          CopyrightStatus = "cc_by_sa"
)

type Author struct {
	ID               int64           `json:"id"`
	Slug             string          `json:"slug"`
	Name             string          `json:"name"`
	NameOriginal     *string         `json:"name_original,omitempty"`
	Title            *string         `json:"title,omitempty"`
	BornYear         *int            `json:"born_year,omitempty"`
	DiedYear         *int            `json:"died_year,omitempty"`
	Era              AuthorEra       `json:"era"`
	Tradition        AuthorTradition `json:"tradition"`
	Bio              *string         `json:"bio,omitempty"`
	BioShort         *string         `json:"bio_short,omitempty"`
	Canonized        bool            `json:"canonized"`
	CanonizedDate    *string         `json:"canonized_date,omitempty"`
	CanonizedBy      *string         `json:"canonized_by,omitempty"`
	FeastDayOrthodox *string         `json:"feast_day_orthodox,omitempty"`
	FeastDayCatholic *string         `json:"feast_day_catholic,omitempty"`
	CopyrightStatus  CopyrightStatus `json:"copyright_status"`
	WikipediaURL     *string         `json:"wikipedia_url,omitempty"`
	WikimediaCategory *string        `json:"wikimedia_category,omitempty"`
	QuoteCount       int             `json:"quote_count,omitempty"`
	ImageURL         *string         `json:"image_url,omitempty"`
	PrimaryImage     *Image          `json:"primary_image,omitempty"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

type Topic struct {
	ID          int64   `json:"id"`
	Slug        string  `json:"slug"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	QuoteCount  int     `json:"quote_count,omitempty"`
}

type Quote struct {
	ID              int64   `json:"id"`
	AuthorID        int64   `json:"author_id"`
	Text            string  `json:"text"`
	TextOriginal    *string `json:"text_original,omitempty"`
	Language        string  `json:"language"`
	SourceWork      *string `json:"source_work,omitempty"`
	SourceChapter   *string `json:"source_chapter,omitempty"`
	SourcePublisher *string `json:"source_publisher,omitempty"`
	SourcePage      *string `json:"source_page,omitempty"`
	SourceURL       *string `json:"source_url,omitempty"`
	License         string  `json:"license"`
	Verified        bool    `json:"verified"`

	// Joined fields
	Author      *Author  `json:"author,omitempty"`
	Topics      []Topic  `json:"topics,omitempty"`
	Attribution *string  `json:"attribution,omitempty"` // Computed for fair-use quotes

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ImageSourceType string

const (
	ImageSourceWikimedia ImageSourceType = "wikimedia_commons"
	ImageSourceMet       ImageSourceType = "met_museum"
	ImageSourceCleveland ImageSourceType = "cleveland_museum"
	ImageSourceAI        ImageSourceType = "ai_generated"
	ImageSourceManual    ImageSourceType = "manual_upload"
)

type ImageStyle string

const (
	StyleByzantineIcon ImageStyle = "byzantine_icon"
	StyleOilPainting   ImageStyle = "oil_painting"
	StyleFresco        ImageStyle = "fresco"
	StyleMosaic        ImageStyle = "mosaic"
	StyleManuscript    ImageStyle = "manuscript"
	StylePhotograph    ImageStyle = "photograph"
	StyleEngraving     ImageStyle = "engraving"
	StyleAIPortrait    ImageStyle = "ai_portrait"
	StyleOther         ImageStyle = "other"
)

type Image struct {
	ID                int64           `json:"id"`
	AuthorID          int64           `json:"author_id"`
	SourceType        ImageSourceType `json:"source_type"`
	SourceURL         *string         `json:"source_url,omitempty"`
	SourceAttribution *string         `json:"source_attribution,omitempty"`
	SourceLicense     *string         `json:"source_license,omitempty"`
	Style             ImageStyle      `json:"style"`
	Width             *int            `json:"width,omitempty"`
	Height            *int            `json:"height,omitempty"`
	MimeType          *string         `json:"mime_type,omitempty"`
	LocalPath         *string         `json:"-"`
	ThumbnailPath     *string         `json:"-"`
	ThumbnailURL      string          `json:"thumbnail_url,omitempty"`
	FullURL           string          `json:"full_url,omitempty"`
	IsAIGenerated     bool            `json:"is_ai_generated"`
	IsPrimary         bool            `json:"is_primary"`
	QualityScore      int             `json:"quality_score"`
	CreatedAt         time.Time       `json:"created_at"`
}

// API response types

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int64       `json:"total"`
	TotalPages int64       `json:"total_pages"`
}

type QuoteOfTheDay struct {
	Date   string `json:"date"`
	Quote  Quote  `json:"quote"`
	Reason string `json:"reason,omitempty"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// Query parameters

type QuoteFilter struct {
	AuthorSlug string
	TopicSlug  string
	Era        string
	Tradition  string
	Verified   *bool
	Language   string
	Page       int
	PerPage    int
}

type AuthorFilter struct {
	Era       string
	Tradition string
	Search    string
	Page      int
	PerPage   int
}

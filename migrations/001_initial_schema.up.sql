-- Martyria: Patristic & Christian Quotes API
-- Initial schema migration

-- Enum types
CREATE TYPE author_era AS ENUM (
    'apostolic',        -- 1st-2nd century
    'ante_nicene',      -- 2nd-3rd century (before 325 AD)
    'nicene',           -- 4th century (325-451 AD)
    'post_nicene',      -- 5th-8th century
    'medieval',         -- 9th-15th century
    'reformation',      -- 16th century
    'modern',           -- 17th-20th century
    'contemporary'      -- 20th-21st century
);

CREATE TYPE author_tradition AS ENUM (
    'pre_schism',       -- Before the Great Schism (1054)
    'orthodox',
    'catholic',
    'protestant',
    'anglican',
    'non_denominational'
);

CREATE TYPE copyright_status AS ENUM (
    'public_domain',
    'short_quote_fair_use',
    'permission_granted',
    'cc_by_sa'
);

CREATE TYPE image_source_type AS ENUM (
    'wikimedia_commons',
    'met_museum',
    'cleveland_museum',
    'ai_generated',
    'manual_upload'
);

CREATE TYPE image_style AS ENUM (
    'byzantine_icon',
    'oil_painting',
    'fresco',
    'mosaic',
    'manuscript',
    'photograph',
    'engraving',
    'ai_portrait',
    'other'
);

-- Authors table
CREATE TABLE authors (
    id              BIGSERIAL PRIMARY KEY,
    slug            TEXT NOT NULL UNIQUE,
    name            TEXT NOT NULL,
    name_original   TEXT,                   -- Original language name
    title           TEXT,                   -- e.g., "Bishop of Hippo", "Elder"
    born_year       INTEGER,
    died_year       INTEGER,
    era             author_era NOT NULL,
    tradition       author_tradition NOT NULL,
    bio             TEXT,
    bio_short       TEXT,                   -- 1-2 sentence summary

    -- Canonization
    canonized           BOOLEAN NOT NULL DEFAULT false,
    canonized_date      DATE,
    canonized_by        TEXT,               -- e.g., "Ecumenical Patriarchate"

    -- Feast days
    feast_day_orthodox  TEXT,               -- e.g., "June 29 / July 12"
    feast_day_catholic  TEXT,

    -- Copyright
    copyright_status    copyright_status NOT NULL DEFAULT 'public_domain',

    -- Wikimedia/external references
    wikipedia_url       TEXT,
    wikimedia_category  TEXT,               -- For image harvesting

    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_authors_era ON authors(era);
CREATE INDEX idx_authors_tradition ON authors(tradition);
CREATE INDEX idx_authors_slug ON authors(slug);

-- Topics table
CREATE TABLE topics (
    id      BIGSERIAL PRIMARY KEY,
    slug    TEXT NOT NULL UNIQUE,
    name    TEXT NOT NULL,
    description TEXT
);

CREATE INDEX idx_topics_slug ON topics(slug);

-- Quotes table
CREATE TABLE quotes (
    id              BIGSERIAL PRIMARY KEY,
    author_id       BIGINT NOT NULL REFERENCES authors(id) ON DELETE CASCADE,
    text            TEXT NOT NULL,
    text_original   TEXT,                   -- Original language text
    language        TEXT NOT NULL DEFAULT 'en',

    -- Source citation
    source_work     TEXT,                   -- e.g., "Confessions", "Homily on Matthew"
    source_chapter  TEXT,                   -- e.g., "Book X, Chapter 27"
    source_publisher TEXT,
    source_page     TEXT,
    source_url      TEXT,

    -- Licensing
    license         TEXT NOT NULL DEFAULT 'public_domain',

    -- Verification
    verified        BOOLEAN NOT NULL DEFAULT false,
    verified_by     TEXT,
    verified_at     TIMESTAMPTZ,

    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_quotes_author ON quotes(author_id);
CREATE INDEX idx_quotes_verified ON quotes(verified);

-- Quote-Topic join table
CREATE TABLE quote_topics (
    quote_id    BIGINT NOT NULL REFERENCES quotes(id) ON DELETE CASCADE,
    topic_id    BIGINT NOT NULL REFERENCES topics(id) ON DELETE CASCADE,
    PRIMARY KEY (quote_id, topic_id)
);

-- Images table
CREATE TABLE images (
    id                  BIGSERIAL PRIMARY KEY,
    author_id           BIGINT NOT NULL REFERENCES authors(id) ON DELETE CASCADE,

    -- Source
    source_type         image_source_type NOT NULL,
    source_url          TEXT,               -- Original URL on Wikimedia/museum
    source_attribution  TEXT,               -- Required attribution text
    source_license      TEXT,               -- e.g., "CC-BY-SA-4.0", "CC0", "public_domain"

    -- Image metadata
    style               image_style NOT NULL DEFAULT 'other',
    width               INTEGER,
    height              INTEGER,
    mime_type           TEXT,

    -- Local storage
    local_path          TEXT,               -- Path in object storage
    thumbnail_path      TEXT,

    -- Flags
    is_ai_generated     BOOLEAN NOT NULL DEFAULT false,
    is_primary          BOOLEAN NOT NULL DEFAULT false,  -- Primary image for author
    quality_score       INTEGER DEFAULT 0,               -- 0-100 quality ranking

    created_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_images_author ON images(author_id);
CREATE INDEX idx_images_primary ON images(is_primary) WHERE is_primary = true;

-- Quote sources (detailed provenance tracking)
CREATE TABLE quote_sources (
    id              BIGSERIAL PRIMARY KEY,
    quote_id        BIGINT NOT NULL REFERENCES quotes(id) ON DELETE CASCADE,
    source_type     TEXT NOT NULL,          -- 'book', 'oral_teaching', 'letter', 'homily', 'sermon'
    source_title    TEXT NOT NULL,
    publisher       TEXT,
    year            INTEGER,
    page            TEXT,
    url             TEXT,
    license         TEXT
);

CREATE INDEX idx_quote_sources_quote ON quote_sources(quote_id);

-- Daily quotes (pre-scheduled rotation)
CREATE TABLE daily_quotes (
    date        DATE PRIMARY KEY,
    quote_id    BIGINT NOT NULL REFERENCES quotes(id) ON DELETE CASCADE,
    reason      TEXT                        -- e.g., "Feast of St. John Chrysostom"
);

-- API keys
CREATE TABLE api_keys (
    id          BIGSERIAL PRIMARY KEY,
    key_hash    TEXT NOT NULL UNIQUE,        -- SHA-256 hash of the API key
    name        TEXT NOT NULL,               -- User/app name
    email       TEXT,
    tier        TEXT NOT NULL DEFAULT 'free', -- 'free', 'registered', 'unlimited'
    rate_limit  INTEGER NOT NULL DEFAULT 100, -- Requests per hour
    active      BOOLEAN NOT NULL DEFAULT true,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_used   TIMESTAMPTZ
);

CREATE INDEX idx_api_keys_hash ON api_keys(key_hash);

-- Permissions tracking (for monastery outreach)
CREATE TABLE permissions (
    id              BIGSERIAL PRIMARY KEY,
    author_id       BIGINT REFERENCES authors(id),
    contact_name    TEXT,
    organization    TEXT NOT NULL,           -- e.g., "Holy Monastery of Souroti"
    status          TEXT NOT NULL DEFAULT 'pending', -- 'pending', 'granted', 'denied', 'no_response'
    date_contacted  DATE,
    date_responded  DATE,
    notes           TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER authors_updated_at BEFORE UPDATE ON authors
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER quotes_updated_at BEFORE UPDATE ON quotes
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

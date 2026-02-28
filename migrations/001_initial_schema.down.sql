DROP TRIGGER IF EXISTS quotes_updated_at ON quotes;
DROP TRIGGER IF EXISTS authors_updated_at ON authors;
DROP FUNCTION IF EXISTS update_updated_at();

DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS api_keys;
DROP TABLE IF EXISTS daily_quotes;
DROP TABLE IF EXISTS quote_sources;
DROP TABLE IF EXISTS images;
DROP TABLE IF EXISTS quote_topics;
DROP TABLE IF EXISTS quotes;
DROP TABLE IF EXISTS topics;
DROP TABLE IF EXISTS authors;

DROP TYPE IF EXISTS image_style;
DROP TYPE IF EXISTS image_source_type;
DROP TYPE IF EXISTS copyright_status;
DROP TYPE IF EXISTS author_tradition;
DROP TYPE IF EXISTS author_era;

# Martyria ✠

**μαρτυρία — Faithful Witness to the Fathers**

A RESTful API serving quotes from the Church Fathers, Saints, and Christian thinkers — from the Apostolic era to modern Orthodox elders. Paired with historical icons and sacred art sourced from public domain collections.

## Quick Start

### With Docker (recommended)

```bash
docker compose up -d
```

The API will be available at `http://localhost:8080`.

### Without Docker

1. **Prerequisites**: Go 1.24+, PostgreSQL 16+, Redis 7+

2. **Database setup**:
```bash
createdb martyria
```

3. **Run**:
```bash
cp .env.example .env
go run ./cmd/martyria
```

4. **Seed data**:
```bash
psql martyria < seeds/001_topics.sql
psql martyria < seeds/002_authors.sql
psql martyria < seeds/003_quotes.sql
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/v1/authors` | List all authors (paginated) |
| GET | `/v1/authors/{slug}` | Get author by slug |
| GET | `/v1/authors/{slug}/quotes` | Get quotes by author |
| GET | `/v1/quotes` | List all quotes (paginated) |
| GET | `/v1/quotes/random` | Get a random quote |
| GET | `/v1/quotes/daily` | Quote of the day |
| GET | `/v1/quotes/{id}` | Get specific quote |
| GET | `/v1/topics` | List all topics |
| GET | `/v1/topics/{slug}/quotes` | Get quotes by topic |

### Query Parameters

**Filtering** (on `/v1/quotes` and `/v1/authors`):
- `era` — apostolic, ante_nicene, nicene, post_nicene, medieval, modern, contemporary
- `tradition` — pre_schism, orthodox, catholic, protestant
- `topic` — topic slug (quotes only)
- `author` — author slug (quotes only)
- `verified` — true/false (quotes only)
- `language` — en, el, la, etc. (quotes only)
- `search` — full-text search (authors only)

**Pagination**:
- `page` — page number (default: 1)
- `per_page` — items per page (default: 20, max: 100)

### Example Requests

```bash
# Random quote from an Orthodox saint
curl "http://localhost:8080/v1/quotes/random?tradition=orthodox"

# All quotes by Paisios of Mount Athos
curl "http://localhost:8080/v1/authors/paisios-of-mount-athos/quotes"

# Quotes about prayer
curl "http://localhost:8080/v1/topics/prayer/quotes"

# Quote of the day
curl "http://localhost:8080/v1/quotes/daily"

# Search for authors
curl "http://localhost:8080/v1/authors?search=chrysostom"
```

## Architecture

```
cmd/martyria/main.go     — Server entrypoint
internal/
  api/                   — HTTP handlers, router, middleware
  config/                — Environment configuration
  db/                    — PostgreSQL connection & queries
  models/                — Domain types (Author, Quote, Topic, Image)
  images/                — Wikimedia/museum image fetcher (planned)
  ai/                    — AI quote extraction pipeline (planned)
  compose/               — Quote-on-image composition (planned)
migrations/              — SQL schema migrations
seeds/                   — Initial data (authors, quotes, topics)
docker/                  — Dockerfile
```

## Authors Coverage

| Tier | Era | Count | Examples |
|------|-----|-------|---------|
| 1 | Apostolic (50-150) | 4 | Clement, Ignatius, Polycarp |
| 2 | Ante-Nicene (100-325) | 6 | Irenaeus, Athanasius, Tertullian |
| 3 | Nicene & Post-Nicene (325-800) | 11 | Chrysostom, Augustine, Basil, Maximus |
| 4 | Medieval (800-1500) | 4 | Palamas, Aquinas, Francis |
| 5 | Early Modern (1700-1900) | 4 | Seraphim of Sarov, Theophan the Recluse |
| 6 | Modern Saints & Elders | 11 | Paisios, Porphyrios, Silouan, Cleopa |

## Tech Stack

- **Go 1.24** — HTTP server with stdlib `net/http` (Go 1.22+ routing)
- **PostgreSQL 16** — Primary data store
- **Redis 7** — Caching & rate limiting (planned)
- **Docker Compose** — One-command deployment

## License

MIT

---

*Glory to God for all things.* — St. John Chrysostom

# Content Getter – Frontend

Modern React single‑page interface for submitting multiple web page URLs to the Content Getter backend,
extracting matched assets (images or other targets), and downloading them in bulk.

---
## 1. Highlights
- 🚀 Fast, minimal UI (React 18 + CRA runtime)
- 🧵 Multi‑URL batch submission (one URL per line)
- 🔍 Custom regex patterns for both content targets and title folder naming
- ↩️ Graceful skipping of 404 sources (handled server‑side)
- 📥 Supports plain download or download+store modes
- 🧱 Clean component structure & centralized styles (App.css)
- 🧪 Ready for extension (routing, theming, state libs, etc.)

---
## 2. Tech Stack
| Layer      | Choice              | Notes                          |
|------------|---------------------|--------------------------------|
| UI         | React 18 (CRA)      | Simplicity & zero‑config build |
| Styling    | Plain CSS (App.css) | Easy theming / override        |
| HTTP       | fetch API           | Small footprint                |
| Env Config | `REACT_APP_API_URL` | Optional; proxy fallback       |

---
## 3. Directory Layout
```
frontend/
  public/            # Static assets (favicon.svg, logo.svg, index.html)
  src/
    App.js           # Core component & form logic
    App.css          # Global + component styles
    index.js         # React bootstrap
  package.json       # Scripts & dependencies
```

---
## 4. Quick Start
```bash
# 1. Install dependencies
npm install

# 2. (Optional) If backend runs on a different port
# Windows (cmd/PowerShell)
set REACT_APP_API_URL=http://localhost:8080/api && npm start
# Linux / macOS
REACT_APP_API_URL=http://localhost:8080/api npm start

# 3. Default (with proxy in package.json if backend is :8080)
npm start
```
Visit: http://localhost:3000

---
## 5. Environment Variables
| Variable            | Purpose                                         | Default |
|---------------------|-------------------------------------------------|---------|
| `REACT_APP_API_URL` | Absolute base URL for API (must include `/api`) | `/api`  |

If not set, requests are made relative to the frontend origin (works with a dev proxy or same host deployment).

---
## 6. Core Endpoints (Backend Contract)
| Method | Path                      | Description                                  |
|--------|---------------------------|----------------------------------------------|
| POST   | `/api/download`           | Download (no persistence) from multiple URLs |
| POST   | `/api/download-and-store` | Download + persist (server path)             |
| GET    | `/api/default-patterns`   | Returns default regex patterns               |
| GET    | `/api/health`             | Health probe                                 |

### 6.1 Request Body (POST /api/download[-and-store])
```json
{
  "urls": ["https://example.com/page1", "https://example.com/page2"],
  "contentPattern": "<custom-regex-or-empty>",
  "titlePattern": "<custom-title-regex-or-empty>"
}
```
- `urls`: Non-empty array of HTTP/HTTPS page URLs.
- `contentPattern`: Optional override (falls back server-side if empty).
- `titlePattern`: Optional override (falls back server-side if empty).

# Image-host
### About
A self-hosted image hosting service built with Go and React. Upload screenshots via paste (Ctrl+V), drag-and-drop, or file selection and get a shareable URL back instantly.
## Stack
- **Backend:** Go (stdlib only)
- **Frontend:** React + TypeScript + Vite
- **Database:** PostgreSQL
- **Reverse proxy:** Nginx (serves frontend, proxies API)
- **Infrastructure:** Docker + Docker Compose

## Running Locally (Development)
 
**Backend:**
```bash
cd backend
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/imagehost?sslmode=disable"
export STORAGE_PATH="./uploads"
export BASE_URL="http://localhost:8080"
go run ./cmd/api
```
 
**Frontend:**
```bash
cd frontend
npm install
npm run dev
```
 
The Vite dev server proxies `/upload` and `/uploads/` to the Go backend automatically.
 
## Deploy (Homelab / Self-hosted)
 
**Requirements:**
- Docker and Docker Compose
- Git
 
**Steps:**
Change the `docker-compose.yml` baseURL field first
 
```bash
git clone https://github.com/your-username/image-host /opt/services/image-host
cd /opt/services/image-host
```
 
Edit `docker-compose.yml` and set `BASE_URL` to your host's IP:
```yaml
BASE_URL: "http://YOUR-HOST-IP:8081"
```
 
```bash
docker compose up --build -d
docker compose exec -T postgres psql -U postgres -d imagehost < backend/migrations/001_create_images.sql
```
 
Access the app at `http://YOUR-HOST-IP:8081`.
## Environment Variables

| Variable       | Description                              | Example                                                       |
|----------------|------------------------------------------|---------------------------------------------------------------|
| `DATABASE_URL` | PostgreSQL connection string             | `postgres://postgres:postgres@postgres:5432/imagehost?sslmode=disable` |
| `STORAGE_PATH` | Directory where uploaded files are saved | `./uploads`                                                   |
| `BASE_URL`     | Public base URL used to build image URLs | `http://192.168.0.5:8081`                                     |

## Usage
 
Open the app in a browser and either:
- **Ctrl+V** — paste a screenshot directly from clipboard
- **Drag and drop** — drop an image file onto the page
- **File picker** — click to select a file
 
The generated URL is displayed on screen and can be used directly in Markdown:
```markdown
![](http://192.168.0.5:8081/uploads/your-image.png)
```
 
---
 

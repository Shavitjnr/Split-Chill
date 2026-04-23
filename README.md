# Split Chill AI

Split Chill AI is a full-stack personal finance application with:

- A Go backend API/server
- A Vue 3 + Vite frontend (desktop and mobile UI)
- SQLite-based local persistence by default
- Optional production packaging via executable or Docker flow

This repository is configured so you can run frontend and backend separately during development.

## Tech Stack

- Backend: Go
- Frontend: Vue 3, Vite, TypeScript, Vuetify, Framework7
- Data: SQLite (default)
- Package manager: npm

## Project Structure

- `src/` - frontend source code
- `public/` - frontend static assets
- `data/` - SQLite database files
- `log/` - backend logs
- `conf/` - backend configuration (`splitchill-ai.ini`)
- `splitchill-ai.exe` - backend executable for Windows
- `package.json` - frontend scripts and dependencies

## Prerequisites

Install the following on your machine:

- Node.js 18+ (recommended latest LTS)
- npm 9+
- Go (only needed if you want to build backend from source)
- Git

For normal local usage on Windows in this repo, you can run backend directly using `splitchill-ai.exe` (no Go build required).

## Quick Start (Current Machine)

### 1) Install frontend dependencies

```bash
npm install
```

### 2) Start backend

Windows PowerShell:

```powershell
.\splitchill-ai.exe server run
```

Backend default URL:

- `http://localhost:8080`

### 3) Start frontend dev server

In a second terminal:

```bash
npm run serve
```

Frontend dev URL:

- `http://localhost:8081`

## Run On Another PC From Scratch

Use these exact steps on a fresh machine.

### Step 1: Clone repository

```bash
git clone https://github.com/Shavitjnr/Split-Chill.git
cd Split-Chill
```

### Step 2: Install Node dependencies

```bash
npm install
```

### Step 3: Configure backend

Check config file:

- `conf/splitchill-ai.ini`

Set a strong `secret_key` (required for safe usage).  
Do not keep default/empty value in production.

### Step 4: Start backend

Windows:

```powershell
.\splitchill-ai.exe server run
```

If running from source on Linux/macOS after build:

```bash
./splitchill-ai server run
```

### Step 5: Start frontend

Open another terminal:

```bash
npm run serve
```

### Step 6: Open app

- Frontend: `http://localhost:8081`
- Backend API/base server: `http://localhost:8080`

## Access From Other Devices In Same Network (LAN)

If you want to open this app from another PC/phone on your Wi-Fi:

1. Keep backend and frontend running.
2. Find host machine local IP (example: `192.168.1.20`).
3. Open firewall ports if blocked:
   - `8080` (backend)
   - `8081` (frontend dev server)
4. Use:
   - `http://<HOST_IP>:8081` for frontend
   - `http://<HOST_IP>:8080` for backend

Notes:

- In dev mode, frontend hot reload is available.
- In strict networks, additional firewall/router rules may be required.

## Build Frontend

```bash
npm run build
```

Output is generated in `dist/`.

## Useful Scripts

- `npm run serve` - start frontend dev server
- `npm run build` - build frontend for production
- `npm run serve:dist` - preview built frontend
- `npm run lint` - type-check and lint

## Data and Logs

- Database file: `data/splitchill-ai.db`
- Runtime journal files may appear:
  - `data/splitchill-ai.db-shm`
  - `data/splitchill-ai.db-wal`
- Backend logs:
  - `log/splitchill-ai.log`

These files are environment/runtime artifacts and will change as you use the app.

## Troubleshooting

### Port already in use

- Change port in backend config (`conf/splitchill-ai.ini`) or stop the process using it.
- For frontend, ensure no other Vite instance is occupying `8081`.

### Frontend cannot talk to backend

- Confirm backend is running on `8080`.
- Check browser console/network and backend logs.
- Verify host/IP and firewall if accessing from another device.

### Security warning about `secret_key`

- Set a strong custom `secret_key` in `conf/splitchill-ai.ini`.
- Restart backend after changing config.

## Contribution Workflow

1. Create a branch
2. Make changes
3. Commit with clear message
4. Push branch and open PR

---

If you want, this README can also be extended with:

- Docker one-command setup
- systemd/PM2 service setup
- reverse proxy setup (Nginx/Caddy)
- HTTPS production deployment checklist

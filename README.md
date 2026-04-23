# Split Chill AI

[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Split Chill AI is a self-hosted personal finance app with a Go backend and a Vue 3 (Vite) frontend. It supports desktop and mobile experiences, and is designed to run locally (SQLite by default) or on a server for multi-device access on your network.

## What’s inside this repo

### Backend

- Go HTTP server (default port `8080`)
- Serves API endpoints and static assets (when built/packaged)
- Persists data to a database (SQLite by default)
- Writes logs to `log/`

### Frontend

- Vue 3 + Vite + TypeScript
- Desktop UI (Vuetify) and Mobile UI (Framework7)
- Development server runs on port `8081` by default

## Features (high level)

- **Bookkeeping**: accounts, categories, tags, templates, recurring/scheduled items
- **Analytics**: filtering, statistics, insights, charts
- **Security**: 2FA, login rate limiting, application lock
- **Data**: imports/exports (various finance formats)
- **PWA**: installable web app experience (mobile-friendly)

## Project structure

- `src/` — frontend source
- `public/` — static assets
- `conf/` — backend config (`splitchill-ai.ini`)
- `data/` — local database files (SQLite by default)
- `log/` — backend logs
- `storage/` — uploaded files / attachments (local filesystem mode)
- `dist/` — frontend production build output
- `splitchill-ai.exe` — backend executable (Windows)
- `go.mod` — Go module definition
- `package.json` — frontend scripts and dependencies

## Ports and URLs

- **Backend**: `http://localhost:8080`
- **Frontend dev (Vite)**: `http://localhost:8081`

When deployed for real usage, you typically expose only one URL (the backend “root URL”), and it serves the UI + API together.

## Prerequisites

### To run from source (recommended for development)

- Node.js (LTS recommended)
- npm
- Git

Optional:

- Go (only if you want to build the backend from source instead of using `splitchill-ai.exe`)

### To run the released binary / executable

- No Node/Go required (depending on package type)

## Run locally (development setup)

This is the best way when you’re working on frontend changes.

### Step 1 — Install frontend dependencies

```bash
npm install
```

### Step 2 — Start backend (Windows)

```powershell
.\splitchill-ai.exe server run
```

Backend will load configuration from:

- `conf\splitchill-ai.ini`

### Step 3 — Start frontend dev server

In another terminal:

```bash
npm run serve
```

Open:

- Frontend: `http://localhost:8081`

## Run on another device (step-by-step)

There are 2 common meanings of “other device”:

1) **Another PC wants to run its own copy** (fresh install on a different computer)  
2) **Another phone/PC wants to access the same running server** (multi-device access over LAN/Wi‑Fi)

### A) Fresh install on another PC (Windows) — from scratch

#### Step 1 — Install prerequisites

Install:

- Git
- Node.js (LTS) + npm

#### Step 2 — Clone the repository

```bash
git clone https://github.com/Shavitjnr/Split-Chill.git
cd Split-Chill
```

#### Step 3 — Install frontend dependencies

```bash
npm install
```

#### Step 4 — Configure backend (important)

Open:

- `conf/splitchill-ai.ini`

Set a strong `secret_key`. The backend will warn you if it’s missing.  
This is critical for protecting user data in a real deployment.

#### Step 5 — Start backend

```powershell
.\splitchill-ai.exe server run
```

#### Step 6 — Start frontend (dev mode)

In a new terminal:

```bash
npm run serve
```

#### Step 7 — Open in browser

- `http://localhost:8081`

### B) Fresh install on another PC (macOS/Linux) — from scratch

#### Step 1 — Install prerequisites

- Git
- Node.js (LTS) + npm
- Go (if you’ll build backend from source)

#### Step 2 — Clone repo

```bash
git clone https://github.com/Shavitjnr/Split-Chill.git
cd Split-Chill
```

#### Step 3 — Install frontend dependencies

```bash
npm install
```

#### Step 4 — Run backend

If you have a built `splitchill-ai` binary:

```bash
./splitchill-ai server run
```

Otherwise, build/package using the provided build scripts in this repo (see “Build & package” below).

#### Step 5 — Run frontend

```bash
npm run serve
```

### C) Access the same running server from another phone/PC (LAN)

Use this when you want **one machine to host**, and other devices open the app in a browser.

#### Step 1 — Run backend on the host machine

```powershell
.\splitchill-ai.exe server run
```

By default it binds to `0.0.0.0:8080` (accessible from your network).

#### Step 2 — Choose how you’ll serve the frontend

Option 1 (simple dev sharing): run Vite and open port `8081`

```bash
npm run serve
```

Option 2 (recommended for sharing): build frontend and serve from backend (production-style)

```bash
npm run build
```

Then ensure the backend is configured to serve the correct static root (already defaults to `public/` for packaged mode; production packaging usually wires this together).

#### Step 3 — Find host machine LAN IP

Example: `192.168.1.20`

#### Step 4 — Open firewall ports on host

Allow inbound TCP:

- `8080` (backend)
- `8081` (only if using Vite dev server)

#### Step 5 — Open from another device

- If using Vite dev server: `http://192.168.1.20:8081`
- If using backend-served UI: `http://192.168.1.20:8080`

## Configuration

Backend config file:

- `conf/splitchill-ai.ini`

Important settings to review:

- `secret_key` (must be set for real deployments)
- `http_port` / `http_addr` (ports / bind)
- database settings (SQLite by default, other DBs supported by configuration)
- storage settings (`storage/` local filesystem by default)

## Build & package

### Frontend (production build)

```bash
npm run build
```

Output: `dist/`

### Full package scripts

Windows:

- `build.bat`
- `build.ps1`

Linux/macOS:

- `build.sh`

See those scripts for packaging options (zip/tar.gz, Docker, etc.).

## Troubleshooting

### Backend warns “secret_key is not set”

Edit `conf/splitchill-ai.ini` and set a strong `secret_key`, then restart backend.

### “Port already in use”

Stop the process using the port or change the port in config.

### Git errors about “unable to unlink … db-wal / log”

On Windows, these files can be locked by a running backend process. Stop the backend before switching branches/pulling:

- stop the `splitchill-ai.exe` process

### Frontend shows proxy errors / cannot reach backend

Confirm backend is reachable at `http://localhost:8080` (or your host IP) and that firewall rules allow it.

## License

[MIT](LICENSE)

Split Chill AI is a lightweight, self-hosted personal finance app with a user-friendly interface and powerful bookkeeping features. It's easy to deploy, and you can start it with just one single Docker command. Designed to be resource-efficient and highly scalable, it can run smoothly on devices as small as a Raspberry Pi, or scale up to NAS, MicroServers, and even large cluster environments.

Split Chill AI offers tailored interfaces for both mobile and desktop devices. With support for PWA (Progressive Web Apps), you can even [add it to your mobile home screen](https://raw.githubusercontent.com/wiki/Shavitjnr/split-chill-ai/img/mobile/add_to_home_screen.gif) and use it like a native app.

Live Demo: [https://splitchill-ai.shavitjnr.net](https://splitchill-ai.shavitjnr.net)

## Features

- **Open Source & Self-Hosted**
    - Built for privacy and control
- **Lightweight & Fast**
    - Optimized for performance, runs smoothly even on low-resource environments
- **Easy Installation**
    - Docker-ready
    - Supports SQLite, MySQL, PostgreSQL
    - Cross-platform (Windows, macOS, Linux)
    - Works on x86, amd64, ARM architectures
- **User-Friendly Interface**
    - UI optimized for both mobile and desktop
    - PWA support for native-like mobile experience
    - Dark mode
- **AI-Powered Features**
    - Receipt image recognition
    - Supports MCP (Model Context Protocol) for AI integration
- **Powerful Bookkeeping**
    - Two-level accounts and categories
    - Attach images to transactions
    - Location tracking with maps
    - Recurring transactions
    - Advanced filtering, search, visualization, and analysis
- **Localization & Globalization**
    - Multi-language and multi-currency support
    - Automatic exchange rates
    - Multi-timezone awareness
    - Custom formats for dates, numbers, and currencies
- **Security**
    - Two-factor authentication (2FA)
    - Login rate limiting
    - Application lock (PIN code / WebAuthn)
- **Data Import/Export**
    - Supports CSV, OFX, QFX, QIF, IIF, Camt.053, MT940, GnuCash, Firefly III, Beancount, and more

## Screenshots

### Desktop Version

[![Split Chill AI](https://raw.githubusercontent.com/wiki/Shavitjnr/split-chill-ai/img/desktop/en.png)](https://raw.githubusercontent.com/wiki/Shavitjnr/split-chill-ai/img/desktop/en.png)

### Mobile Version

[![Split Chill AI](https://raw.githubusercontent.com/wiki/Shavitjnr/split-chill-ai/img/mobile/en.png)](https://raw.githubusercontent.com/wiki/Shavitjnr/split-chill-ai/img/mobile/en.png)

## Installation

### Run with Docker

Visit [Docker Hub](https://hub.docker.com/r/Shavitjnr/split-chill-ai) to see all images and tags.

**Latest Release:**

    $ docker run -p8080:8080 Shavitjnr/split-chill-ai

**Latest Daily Build:**

    $ docker run -p8080:8080 Shavitjnr/split-chill-ai:latest-snapshot

### Install from Binary

Download the latest release: [https://github.com/Shavitjnr/split-chill-ai](https://github.com/Shavitjnr/split-chill-ai)

**Linux / macOS**

    $ ./splitchill-ai server run

**Windows**

    > .\splitchill-ai.exe server run

By default, Split Chill AI listens on port 8080. You can then visit `http://{YOUR_HOST_ADDRESS}:8080/` .

### Build from Source

Make sure you have [Golang](https://golang.org/), [GCC](https://gcc.gnu.org/), [Node.js](https://nodejs.org/) and [NPM](https://www.npmjs.com/) installed. Then download the source code, and follow these steps:

**Linux / macOS**

    $ ./build.sh package -o splitchill-ai.tar.gz

All the files will be packaged in `splitchill-ai.tar.gz`.

**Windows**

    > .\build.bat package -o splitchill-ai.zip

or

    PS > .\build.ps1 package -Output splitchill-ai.zip

All the files will be packaged in `splitchill-ai.zip`.

You can also build a Docker image. Make sure you have [Docker](https://www.docker.com/) installed, then follow these steps:

**Linux**

    $ ./build.sh docker

## Contributing

We welcome contributions of all kinds.

Found a bug? [Submit an issue](https://github.com/Shavitjnr/split-chill-ai/issues)

Want to contribute code? Feel free to fork and send a pull request.

Contributions of all kinds — bug reports, feature suggestions, documentation improvements, or code — are highly appreciated.

Check out our [Contributor Graph](https://github.com/Shavitjnr/split-chill-ai/graphs/contributors) to see the amazing people who've already helped.

## Translating

Help make Split Chill AI accessible to users around the world. If you want to contribute a translation, please refer to our [translation guide](https://splitchill-ai.shavitjnr.net/translating).

Currently available translations:

| Tag     | Language           | Contributors                                                                                                                                                                                     |
| ------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| de      | Deutsch            | [@chrgm](https://github.com/chrgm)                                                                                                                                                               |
| en      | English            | /                                                                                                                                                                                                |
| es      | Español            | [@Miguelonlonlon](https://github.com/Miguelonlonlon), [@abrugues](https://github.com/abrugues), [@AndresTeller](https://github.com/AndresTeller), [@diegofercri](https://github.com/diegofercri) |
| fr      | Français           | [@brieucdlf](https://github.com/brieucdlf)                                                                                                                                                       |
| it      | Italiano           | [@waron97](https://github.com/waron97)                                                                                                                                                           |
| ja      | 日本語             | [@tkymmm](https://github.com/tkymmm)                                                                                                                                                             |
| kn      | ಕನ್ನಡ              | [@Darshanbm05](https://github.com/Darshanbm05)                                                                                                                                                   |
| ko      | 한국어             | [@overworks](https://github.com/overworks)                                                                                                                                                       |
| nl      | Nederlands         | [@automagics](https://github.com/automagics)                                                                                                                                                     |
| pt-BR   | Português (Brasil) | [@thecodergus](https://github.com/thecodergus)                                                                                                                                                   |
| ru      | Русский            | [@artegoser](https://github.com/artegoser)                                                                                                                                                       |
| sl      | Slovenščina        | [@thehijacker](https://github.com/thehijacker)                                                                                                                                                   |
| ta      | தமிழ்              | [@hhharsha36](https://github.com/hhharsha36)                                                                                                                                                     |
| th      | ไทย                | [@natthavat28](https://github.com/natthavat28)                                                                                                                                                   |
| tr      | Türkçe             | [@aydnykn](https://github.com/aydnykn)                                                                                                                                                           |
| uk      | Українська         | [@nktlitvinenko](https://github.com/nktlitvinenko)                                                                                                                                               |
| vi      | Tiếng Việt         | [@f97](https://github.com/f97)                                                                                                                                                                   |
| zh-Hans | 中文 (简体)        | /                                                                                                                                                                                                |
| zh-Hant | 中文 (繁體)        | /                                                                                                                                                                                                |

Don't see your language? Help us add it.

## Documentation

1. [English](https://splitchill-ai.shavitjnr.net)
1. [中文 (简体)](https://splitchill-ai.shavitjnr.net/zh_Hans)

## License

[MIT](https://github.com/Shavitjnr/split-chill-ai/blob/master/LICENSE)

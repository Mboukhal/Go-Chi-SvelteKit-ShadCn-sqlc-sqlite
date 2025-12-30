# Go-Chi + Air + SvelteKit + ShadCN + sqlc + SQLite🥰

Full-stack starter combining **Go backend (Chi + sqlc + SQLite)** and **SvelteKit frontend** with **ShadCN UI + TailwindCSS**.

## 🚀 Features

- Go API using **Chi router**
- Type-safe DB with **sqlc** (SQLite)
- Frontend with **SvelteKit as static + TailwindCSS + ShadCN UI**
- Single binary backend + frontend dev proxy
- Docker & Makefile for easy setup

## 🧱 Stack

- **Go** (API)
- **Chi** (router)
- **sqlc** (generate Go DB code from SQL)
- **SQLite** (embedded DB)
- **SvelteKit** (web UI)
- **TailwindCSS + ShadCN UI**

## 📥 Setup

### Dev

Start backend + frontend dev servers with hot reload, frontend at `localhost:1337` but proxied through backend at `localhost:3000`.

```
make install
make
```

## 🐳 Deploy With Docker

```
make build_prod
make docker_start
```

## ⚙️ Env

Copy and fill:

.env.example → .env

## 📁 Structure

/cmd # go app

/cmd/ui # SvelteKit UI

/core # business logic

/internal # handlers, services

/sqlc.yaml # sqlc config

## 🎯 Usage

- Build API in Go with type safety using sqlc
- UI in SvelteKit using ShadCN components

## 🤝 Contribute

PRs/issues welcome.

## 📜 License

MIT

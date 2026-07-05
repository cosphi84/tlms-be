# TLMS-BE Deployment Guide (Linux - Rocky Linux 9 + Nginx + PostgreSQL 18)

Dokumen ini berisi langkah lengkap deployment backend TLMS (Go + Gin) di server Linux production.

## Stack
- Rocky Linux 9.8
- PostgreSQL 18 (PGDG)
- Nginx
- Go (Gin)
- systemd

---

## 1. Install PostgreSQL 18

sudo dnf -qy module disable postgresql

sudo dnf install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-9-x86_64/pgdg-redhat-repo-latest.noarch.rpm

sudo dnf install -y postgresql18 postgresql18-server postgresql18-contrib

sudo -u postgres /usr/pgsql-18/bin/initdb \
  -D /var/lib/pgsql/18/data \
  --locale=en_US.UTF-8 \
  --encoding=UTF8

sudo systemctl enable postgresql-18
sudo systemctl start postgresql-18

---

## 2. Setup Database

sudo -u postgres psql

CREATE USER tlms WITH PASSWORD 'your_password';
CREATE DATABASE tlms OWNER tlms ENCODING 'UTF8';

\c tlms
CREATE EXTENSION IF NOT EXISTS pgcrypto;

---

## 3. Clone Project

cd /opt
mkdir -p tlms
cd tlms

git clone git@github.com:cosphi84/tlms-be.git backend
cd backend

---

## 4. Environment (.env)

APP_ENV=production
PORT=2000

DB_HOST=localhost
DB_PORT=5432
DB_USER=tlms
DB_PASSWORD=your_password
DB_NAME=tlms

GIN_MODE=release
JWT_SECRET=your_secret

---

## 5. Build & Run

go mod download
go build -o tlms-be cmd/main.go
./tlms-be

---

## 6. systemd Service

/etc/systemd/system/tlms-be.service

[Unit]
Description=TLMS Backend
After=network.target

[Service]
User=tlms
WorkingDirectory=/opt/tlms/backend
ExecStart=/opt/tlms/backend/tlms-be
Restart=always
EnvironmentFile=/opt/tlms/backend/.env

[Install]
WantedBy=multi-user.target

sudo systemctl daemon-reload
sudo systemctl enable tlms-be
sudo systemctl start tlms-be

---

## 7. Nginx Config

location /api/tlms/ {
    proxy_pass http://127.0.0.1:2000/;

    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
}

sudo nginx -t
sudo systemctl reload nginx

---

## 8. Test

curl http://127.0.0.1:2000/health
curl https://elektrodukasi.id/api/tlms/health

---

## 9. Update Deployment

cd /opt/tlms/backend
git pull
go build -o tlms-be cmd/main.go
sudo systemctl restart tlms-be

# Team Task Manager

A full-stack Team Task Management System built using Golang, React, PostgreSQL, JWT Authentication, and Role-Based Access Control (RBAC).

---

# Live Demo

## Frontend
https://team-task-manager-five-zeta.vercel.app/

## Backend
https://team-task-manager-bzwg.onrender.com

---

# Features

## Authentication
- User Registration
- User Login
- JWT Authentication
- Logout Functionality
- Protected Routes
- Persistent Login Sessions

---

# Role-Based Access Control

## Admin
- Create Projects
- Add Members To Projects
- Create Tasks
- Assign Tasks
- Filter Tasks
- View Dashboard Statistics
- Manage Project Members

## Member
- View Assigned Projects
- View Assigned Tasks
- Filter Tasks
- Update Task Status

---

# Task Features
- Task Assignment
- Task Priority
- Due Date Support
- Overdue Task Logic
- Task Status Updates
- Project-wise Task Filtering
- Priority Filtering
- Status Filtering
- Project Name Mapping In Tasks

---

# Dashboard Features
- Total Projects
- Total Tasks
- Completed Tasks
- Overdue Tasks
- Dynamic Dashboard Statistics

---

# Tech Stack

## Backend
- Golang
- Chi Router
- PostgreSQL
- SQLX
- JWT Authentication
- REST APIs

## Frontend
- React
- Vite
- TailwindCSS
- Axios
- React Router DOM

## Deployment
- Render
- PostgreSQL Render Database

---

# Project Structure

## Backend

```bash
handlers/
middleware/
router/
database/
dbhelpers/
model/
utils/
migrations/
```

## Frontend

```bash
frontend/
├── src/
│   ├── pages/
│   ├── components/
│   ├── services/
│   ├── routes/
│   └── assets/
```

---

# Installation

# Backend Setup

## Clone Repository

```bash
git clone <repository-url>
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Configure PostgreSQL

Create PostgreSQL database.

Update environment variables:

```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
JWT_SECRET=
```

---

## Run Migrations

Run migration SQL files from:

```bash
migrations/
```

---

## Start Backend Server

```bash
go run main.go
```

Backend runs on:

```bash
http://localhost:8080
```

---

# Frontend Setup

## Go To Frontend

```bash
cd frontend
```

---

## Install Dependencies

```bash
npm install
```

---

## Start Frontend

```bash
npm run dev
```

Frontend runs on:

```bash
http://localhost:5173
```

---

# API Endpoints

# Authentication

| Method | Endpoint | Description |
|---|---|---|
| POST | /register | Register User |
| POST | /login | Login User |
| POST | /logout | Logout User |

---

# Projects

| Method | Endpoint | Description |
|---|---|---|
| GET | /projects | Get Projects |
| POST | /projects | Create Project |
| POST | /projects/{id}/members | Add Member |

---

# Tasks

| Method | Endpoint | Description |
|---|---|---|
| GET | /tasks | Get Tasks |
| POST | /tasks | Create Task |
| PATCH | /tasks/{id}/status | Update Task Status |

---

# Dashboard

| Method | Endpoint | Description |
|---|---|---|
| GET | /dashboard/stats | Dashboard Statistics |

---

# Security Features
- JWT Authentication
- Password Hashing using Bcrypt
- Protected Routes
- RBAC Authorization
- Middleware-Based Authentication
- CORS Middleware

---

# Deployment

## Backend Deployment
- Render Web Service
- PostgreSQL Render Database

## Frontend Deployment
- Render Static Site

---

# Future Improvements
- File Attachments
- Notifications
- Comments On Tasks
- Real-Time Updates
- Email Notifications
- Docker Deployment
- WebSocket Integration
- Activity Logs
- Team Chat

---

# Author

Ruchi Kumari

Built using Golang, React, PostgreSQL, JWT Authentication, and TailwindCSS.

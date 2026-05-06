# Team Task Manager

A full-stack Team Task Management System built using Golang, React, PostgreSQL, JWT Authentication, and Role-Based Access Control (RBAC).

---

# Features

## Authentication
- User Registration
- User Login
- JWT Authentication
- Logout Functionality
- Protected Routes

---

# Role-Based Access Control

## Admin
- Create Projects
- Add Members To Projects
- Create Tasks
- Assign Tasks
- View Dashboard Statistics

## Member
- View Assigned Projects
- View Assigned Tasks
- Update Task Status

---

# Task Features
- Task Assignment
- Task Priority
- Due Date Support
- Overdue Task Logic
- Task Status Updates

---

# Dashboard Features
- Total Projects
- Total Tasks
- Completed Tasks
- Overdue Tasks

---

# Tech Stack

## Backend
- Golang
- Chi Router
- PostgreSQL
- SQLX
- JWT Authentication

## Frontend
- React
- Vite
- TailwindCSS
- Axios
- React Router DOM

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
```

## Frontend

```bash
frontend/
├── src/
│   ├── pages/
│   ├── components/
│   ├── services/
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

Run your SQL migration files.

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
- CORS Middleware

---

# Future Improvements
- File Attachments
- Notifications
- Comments On Tasks
- Real-Time Updates
- Email Notifications
- Docker Deployment

---

# Author

Ruchi Kumari

Built using Golang, React, PostgreSQL, and TailwindCSS.

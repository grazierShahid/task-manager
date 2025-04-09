# Task Manager API

A simple task management system built with Go and Gin framework. This project provides a RESTful API for managing tasks with user authentication.

## Features

- User authentication with JWT (signup, login, logout)
- Task management (create, read, update, delete)
- Task filtering by status and priority
- User-specific task lists

## Tech Stack

- Go programming language
- Gin web framework
- GORM for database operations
- PostgreSQL database (configurable)

## API Endpoints

### User Authentication

- POST `/signup` - Create a new user account
- POST `/login` - Authenticate and receive a token
- POST `/logout` - End the current session

### Task Management

- POST `/addtask` - Create a new task
- POST `/deletetask` - Remove an existing task
- POST `/edittask` - Update task details
- GET `/task/:id` - Get a specific task by ID
- GET `/tasks` - List all tasks with optional filtering
- GET `/user/:id/tasks` - Get all tasks for a specific user

## Getting Started

1. Clone the repository
2. Set up environment variables
3. Run the application with `go run main.go`
4. Access the API at http://localhost:3000

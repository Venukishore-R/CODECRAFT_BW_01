# Basic REST API with CRUD Operations

## Task Overview
This project involves creating a RESTful API to perform basic **CRUD** (Create, Read, Update, Delete) operations on a `users` resource. The goal is to design an API that handles user data efficiently using an in-memory data structure.

---

## Features
- CRUD operations for `users` resource:
  - **Create** a new user.
  - **Read** user details by Email.
  - **Update** existing user details.
  - **Delete** a user by ID.
- In-memory data storage using a hashmap or similar structure.
- Proper HTTP status codes and error handling:
  - `404 Not Found` for non-existing resources.
  - `400 Bad Request` for invalid inputs.
- Basic input validation, such as:
  - Checking if the email is valid.

---

## User Model
Each user will have the following fields:
- `id` (UUID): A unique identifier for each user.
- `name` (string): The name of the user.
- `email` (string): The user's email address (validated).
- `age` (integer): The age of the user.

---

## Endpoints

| Method | Endpoint        | Description               |
|--------|-----------------|---------------------------|
| POST   | `/users`        | Create a new user         |
| GET    | `/users/:email`    | Get details of a user     |
| PUT    | `/users/:email`    | Update details of a user  |
| DELETE | `/users/:email`    | Delete a user by ID       |

---

## Requirements
- Language/Framework: Any (e.g., Node.js, Python, Go, etc.).
- Use an in-memory data structure for data persistence.
- Follow REST API design principles.

---

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Venukishore-R/CODECRAFT_BW_01.git
   cd CODECRAFT_BW_01
2. Install Go dependencies (if applicable):
   ```bash
   go mod tidy
3. Run the server:
   ```bash
   go run cmd/main.go


# Project Name

This project is a backend application for managing jobs and user authentication. It provides gRPC and HTTP APIs for job management, user authentication, and job applications.

## Technologies Used

- Golang
- gRPC
- Protocol Buffers
- PostgreSQL

## Project Structure

The project is structured into several packages:

- `adapters`: Contains database adapter implementations.
- `auth`: Implements user authentication functionality.
- `domain`: Defines domain structures and interfaces.
- `job`: Implements job management functionality.
- `somniumsystem`: Implements gRPC and HTTP endpoints.

## Setup

1. Clone the repository.
2. Install dependencies using `go mod tidy`.
3. Set up your PostgreSQL database and configure the connection in the `.env` file.
4. Run the application using `go run main.go`.

## APIs

### Authentication

- `POST /auth/signin`: Sign in with email and password.
- `POST /auth/signup`: Sign up with email and password.

### Jobs

- `GET /jobs`: List all jobs.
- `GET /job/{id}`: Get details of a job by ID.
- `POST /job/create`: Create a new job.
- `PUT /job/update/{id}`: Update an existing job.
- `DELETE /job/delete/{id}`: Delete a job by ID.
- `POST /job/apply`: Apply for a job.
- `GET /user/{userID}/history`: Get job application history for a user.

### Job Status Switcher


- `GET /job/switch/{id}`: Switch the status of a job from Active to Inactive, and vice versa.

## Usage

You can use tools like `curl` or Postman to interact with the APIs. Here's an example of how to sign in and get a list of jobs using `curl`:

```bash
curl -X POST http://localhost:8080/auth/signin -d '{"email": "user@example.com", "password": "password"}'
curl -X GET http://localhost:8080/jobs -H "Authorization: Bearer <access_token>"


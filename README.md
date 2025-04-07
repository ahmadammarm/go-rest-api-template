# Go REST API Template

A RESTful API template built with Go programming language.

## Prequisites

- Go 1.24 or later
- PostgreSQL 17 or later
- Docker (optional)


## Features

- **Middleware**: Implement the JWT Auth for the middleware.
- **Service Layer**: Implement the service layer for separating business logic from application logic.
- **Repository Pattern**: Managing interactions with databases using the repository pattern.
- **Configuration**: Supports database environment configuration via .env files.
- **Deployment**: Supports build and deploy using Docker.

## REST API Design

The project provides a RESTful API for posts. The API follows standard REST conventions:

### Users API Routes

- `GET /api/v1/users` - Get all users.
- `GET /api/v1/user/:id` - Get an user by id.
- `POST /api/v1/user/register` - Register an user.
- `POST /api/v1/user/login` - Login for the registered user.


### News API Routes

- `GET /api/v1/news` - Get all news.
- `GET /api/v1/news/:id` - Get a news by id.
- `POST /api/v1/news` - Create a news.
- `PUT /api/v1/news/:id` - Edit a news by id.
- `DELETE /api/v1/news/:id` - Delete a news by id.


## Getting Started

1. Clone the repository:

```sh
git clone https://github.com/ahmadammarm/go-rest-api-template.git
```

2. Navigate to the project directory:

```sh
cd go-rest-api-template
```

3. Install the project dependencies:

```sh
go mod download
```

4. Configure Environment Variable: Copy the file `.env.example` to `.env` and adjust it to your configuration:

```sh
cp .env.example .env
```


5. Run the project:

```sh
go run cmd/main.go
```


## Getting Started with Docker

1. Clone the repository:
```sh
git clone https://github.com/ahmadammarm/go-rest-api-template.git
```

2. Navigate to the project directory:

```sh
cd go-rest-api-template
```

3. Run the Docker Compose:

```sh
docker-compose up
```

The project will be available at:

`http://localhost:8080`

### This project also available for your contributions, thank you :)


# Go REST API Template

A RESTful API template built with Go programming language.


![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) ![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens) ![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## Prequisites

- Go 1.24 or later
- PostgreSQL 17 or later
- Docker (optional)


## Features

- **Unit Testing**: Implement the unit testing in order to code maintenance.
- **Middleware**: Implement the JWT Auth for the middleware.
- **Service Layer**: Implement the service layer for separating business logic from application logic.
- **Repository Pattern**: Managing interactions with databases using the repository pattern.
- **CRUD Implementation**: Implement CRUD logic in the model entity.
- **Configuration**: Supports environment variables configuration via .env files.
- **CORS Support**: Supports Cross-Origin Resource Sharing (CORS) configuration.
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

```
POSTGRES_HOST=localhost
POSTGRES_PORT=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
JWT_SECRET_KEY=your-jwt-secret-key
CORS_ALLOW_ORIGINS=your-header-http-domain
```

- POSTGRES: Configuration for PostgreSQL database connection.

- JWT_SECRET_KEY: The secret key used for JWT token signing and verification.

- CORS_ALLOW_ORIGINS: The allowed origins for Cross-Origin Resource Sharing (CORS). This is the domain that will be able to access resources from this API. For example, if you are running the frontend on http://localhost:5173, you should set this environment variable to http://localhost:5173.


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

or run in detach mode:
```sh
docker-compose up -d
```

4. To stop the Docker Compose:
```sh
docker-compose down
```

The project will be available at:

`http://localhost:8080`

### This project also available for your contributions, thank you :)


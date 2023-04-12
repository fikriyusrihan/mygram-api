# My Gram API

My Gram API is an API that allows users to create, read, update, and delete photos.
Built with Go and PostgreSQL. This is a final project of Hacktiv8's Go Scalable Web Service Bootcamp.

## Run the Application

This application requires Go and PostgreSQL to run. Make sure you have them installed on your machine before running the
application.
Please change the application configuration in `config/config.yml` to match your environment.

```bash
# Move to the workspace directory
cd workspace

# Clone the repository
git clone https://github.com/fikriyusrihan/mygram-api.git

# Move to the application directory
cd mygram-api

# Build the application
go build -o .build/mygram-api

# Run the application
.build/mygram-api
```

## API Documentation

Access the API documentation at `http://localhost:8000/swagger/index.html` when the application is running.

## Tools

- [Go](https://golang.org/) (Programming Language)
- [Gin](https://gin-gonic.com/) (Web Framework)
- [GORM](https://gorm.io/) (ORM)
- [PostgreSQL](https://www.postgresql.org/) (Database)
- [Swagger](https://swagger.io/) (API Documentation)

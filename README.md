# Real Estate Management

![Go Version](https://img.shields.io/badge/Go-1.23.6-00ADD8?style=flat&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2.52.6-00BFB3?style=flat&logo=fiber)
![GORM](https://img.shields.io/badge/GORM-v1.25.12-5C6BC0?style=flat&logo=gorm)
![PostgreSQL Driver](https://img.shields.io/badge/PostgreSQL-v1.5.11-336791?style=flat&logo=postgresql)

## Overview
This project is focused on real estate management, providing functionality for handling property data, user management, and other core features commonly required in real estate applications. It utilizes a fast and minimalistic web framework with an efficient ORM for database interactions, making it easy to manage and interact with PostgreSQL databases.

## Dependencies

### ORM (Object-Relational Mapping)
- **[gorm.io/gorm](https://github.com/go-gorm/gorm)**: GORM is a powerful ORM for Go. It's used for handling database operations in an object-oriented manner.
  - Version: v1.25.12
- **[gorm.io/driver/postgres](https://github.com/go-gorm/postgres)**: PostgreSQL driver for GORM.
  - Version: v1.5.11

### HTTP Framework
- **[github.com/gofiber/fiber/v2](https://github.com/gofiber/fiber)**: A fast and minimalistic web framework for Go. It's used to create the HTTP server and route handlers.
  - Version: v2.52.6

### Environment Variables
- **[github.com/caarlos0/env/v11](https://github.com/caarlos0/env)**: A simple and effective way to load environment variables into Go structures.
  - Version: v11.3.1

### Configuration Management
- **[github.com/joho/godotenv](https://github.com/joho/godotenv)**: A Go package that helps you manage environment variables from `.env` files.
  - Version: v1.5.1

## Setup

1. Install dependencies:
   ```bash
   go mod tidy

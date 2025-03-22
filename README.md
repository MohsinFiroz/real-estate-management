# Teng Dragon Real Estate Management

![Go Version](https://img.shields.io/badge/Go-1.23.6-00ADD8?style=flat&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2.52.6-00BFB3?style=flat)
![GORM](https://img.shields.io/badge/GORM-v1.25.12-5C6BC0?style=flat)
![Gorm PostgreSQL Driver](https://img.shields.io/badge/PostgreSQL-v1.5.11-336791?style=flat&logo=postgresql)

## Overview
This project focuses on real estate management, offering features for property data and user management. It uses a fast web framework and an efficient ORM for seamless PostgreSQL integration.

## Dependencies

- **ORM**:  
  - [gorm.io/gorm](https://github.com/go-gorm/gorm): ORM for Go  
  - [gorm.io/driver/postgres](https://github.com/go-gorm/postgres): PostgreSQL driver for GORM

- **HTTP Framework**:  
  - [github.com/gofiber/fiber/v2](https://gofiber.io): Minimalistic web framework

- **Environment Variables**:  
  - [github.com/caarlos0/env/v11](https://github.com/caarlos0/env): Load environment variables into Go structures

- **Configuration Management**:  
  - [github.com/joho/godotenv](https://github.com/joho/godotenv): Manage environment variables from `.env` files

## Setup

1. Install dependencies:
   ```bash
   go mod tidy

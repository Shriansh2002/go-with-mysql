# Golang MySQL Starter

A simple starter template for a Golang application with MySQL database interaction.

## Table of Contents

-   [Introduction](#introduction)
-   [Prerequisites](#prerequisites)
-   [Project Structure](#project-structure)
-   [Getting Started](#getting-started)
-   [Configuration](#configuration)
-   [Contributing](#contributing)

## Introduction

This is a basic Golang application template that demonstrates how to connect to a MySQL database. It includes a modular structure with separate files for configuration, database operations, and the main application logic.

## Prerequisites

-   Go installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/)
-   MySQL server installed and running, or online MySQL database credentials

## Project Structure

```bash
.
├── cmd/
│   └── main.go
├── internal/
│   └── app/
│       └── main.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   └── database/
│       └── database.go
├── .env
├── .go.mod
└── .go.sum
```

## Getting Started

1. **Install Go:**
   Download and install Go from [https://golang.org/dl/](https://golang.org/dl/)

2. **Set Up Your Project:**
   Create a project directory and organize files based on the provided structure.

3. **Initialize Go Modules:**
   In your project directory, run:

    ```bash
    go mod init yourmodulepath
    ```

    Replace yourmodulepath with your actual module path or leave it empty.

4. **Install Dependencies:**
   In your project directory, run:

    ```bash
    go mod tidy
    ```

5. Run the application:

    ```bash
    go run cmd/main.go
    ```

## Configuration

Configure your database connection details by updating the .env file.

```env
USERNAME=root
PASSWORD=mysecretpassword
HOST=localhost
PORT=3306
DATABASE=mydatabase
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

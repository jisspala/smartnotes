## SmartNotes

This is a simple notepad application that provides a REST API to create, list, edit, and delete notes. The notes are stored in a PostgreSQL database.

## Requirements

- Go 1.20 or higher
- PostgreSQL database
- Gin Gonic
- Swagger
- gorm

## Installation

1.  Clone the repository
2.  Run the application with `docker-compose up`

## Usage

You can access the Swagger GUI to try out the API by visiting `http://localhost:8001/v1/docs/index.html#/`.

### Endpoints

- `GET /health`: get health of app
- `POST /notes`: Create a new note
- `GET /notes`: List all existing notes
- `PUT /note/:id`: Edit a note by ID
- `DELETE /note/:id`: Delete a note by ID
- `DELETE /notes`: Delete multiple notes by IDs (accepts an array of IDs in the request body)

## Technical overview

We are using **Go** with **Gin** to create APIs. T. To generate API docs and test APIs by manual we are using **Swagger**. For data-base operation and **Auto-migrations**, we are using **gorm**

## How to Test in local

1.  Install the required dependencies with `go mod tidy` `

2.  If you don't need docker db, then create a new PostgreSQL database and add the following code in main.go
    ** os.Setenv("POSTGRES_USER", "user")
    os.Setenv("POSTGRES_PASSWORD", "admin")
    os.Setenv("POSTGRES_PORT", "5432")
    os.Setenv("POSTGRES_HOST", "localhost")
    os.Setenv("POSTGRES_DB", "notepad") **

3.  Auto-migration will be take care by **gorm**
4.  Run the application with `go run .`
5.  You can access the Swagger GUI to try out the API by visiting `http://localhost:8001/v1/docs/index.html#/`.

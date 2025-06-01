<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [BRUD](#brud)
  - [Description](#description)
- [Run (local)](#run-local)
  - [Prerequisites](#prerequisites)
  - [Running](#running)
- [TASK](#task)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# BRUD

A blog CRUD API backend using Go and Fiber

## Description

BRUD is a simple blog CRUD API backend built with Go and the Fiber web framework. It allows users to create, read, update, and delete blog posts. The API supports basic operations for managing blog posts with fields such as title, description, body, created_at, and updated_at.

# Run (local)

## Prerequisites

- Go installed on your machine (version 1.18 or higher)
- A postgres database set up and running
- Environment variables configured for database connection (check `.env.example` for reference)
- Install dependencies using Go modules

```bash
go mod tidy
```

- Install Goose for database migrations (MacOS)

```bash
brew install goose
```

- Install Goose for database migrations (Linux)

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

- Run migrations

```bash
goose up
```

## Running

To run the BRUD API locally, follow these steps:

```bash
go run cmd/main.go
```

You can access server at [http://localhost:3000](http://localhost:3000).

And the Swagger documentation at [http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html).

# TASK

Problem Statement:

Have to create a blog CRUD API backend

Blog-post with the keys (title, descriptions, body, created_at, updated_at)

Expectation :

Crud Blog app with the following APIs

/api/blog-post => post => to add a post

/api/blog-post => get => to get all added posts

/api/blog-post/:id => get => get single post

/api/blog-post/:id => delete => to delete post

/api/blog-post/:id => patch => to update post

It should be deployed on any accessible hosting server like netlify, render.com, etc

The swagger link should be shared with this mail

Push it in Private Repo and invite taquartiz, nirmal-quartiz

Bonus:

Go-Fiber as API engine

API unit test with coverage

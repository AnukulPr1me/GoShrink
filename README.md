# Go Link Shortener

A simple and efficient URL shortening service built using Go, Redis, and Docker. This service allows users to generate short URLs for long links, making it easy to share and manage URLs.

## Features

- Shorten long URLs to a compact format.
- Store and retrieve short links using Redis for fast access.
- Redirect users to the original URL when the short link is accessed.
- Simple, lightweight, and easy to deploy with Docker.

## Technologies Used

- **Go**: Backend logic for handling URL shortening and redirection.
- **Redis**: In-memory database for storing and retrieving URL mappings.
- **Docker**: Containerization for easy deployment.
- **HTTP**: Go's `net/http` package for handling HTTP requests and responses.

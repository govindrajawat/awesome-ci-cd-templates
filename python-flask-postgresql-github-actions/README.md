# Flask + PostgreSQL CI/CD Template

This template provides a production-ready Flask application with PostgreSQL, Docker, and GitHub Actions CI/CD pipeline.

## Features

- Flask web application with PostgreSQL database
- Docker containerization with multi-stage build
- GitHub Actions workflow for:
  - Running tests
  - Building and pushing Docker images to GitHub Container Registry
  - Deploying to Kubernetes (optional)
- Database migrations with Flask-SQLAlchemy

## Prerequisites

- Python 3.9+
- Docker
- PostgreSQL (for local development)
- GitHub account

## Setup

1. Clone the repository
2. Create and activate a virtual environment:
   ```bash
   python -m venv venv
   source venv/bin/activate  # On Windows: venv\Scripts\activate
3. Install dependencies:
    pip install -r requirements.txt
4. Copy .env.example to .env and update with your database credentials
5. Run the application:
    flask run

## Docker Setup

1. Build and run with Docker Compose:
    docker-compose up --build
2. The application will be available at http://localhost:5000

## CI/CD Pipeline

The GitHub Actions workflow includes:
Test: Runs pytest with a PostgreSQL service container
Build and Push: Builds Docker image and pushes to GitHub Container Registry
Deploy: Deploys to Kubernetes (requires Kubernetes cluster setup)

## Customization
Update Dockerfile and docker-compose.yml with your application specifics
Modify the GitHub Actions workflow in .github/workflows/ci-cd.yml
Add your Kubernetes manifests in the k8s/ directory
# Go + PostgreSQL + GitHub Actions + Kubernetes Template

This template provides a production-ready Go application with PostgreSQL, Docker, GitHub Actions CI/CD, and Kubernetes deployment.

## Features

- Go HTTP server with PostgreSQL database
- Docker containerization with multi-stage build
- GitHub Actions workflow for:
  - Running tests with PostgreSQL service
  - Code linting and quality checks
  - Building and pushing Docker images to GitHub Container Registry
  - Deploying to Kubernetes cluster with health verification
- Kubernetes manifests for deployment with health checks
- Optional Helm chart for deployment
- Health check endpoint for monitoring
- Comprehensive test coverage

## Prerequisites

- Go 1.21+
- Docker
- PostgreSQL (for local development)
- Kubernetes cluster (for deployment)
- Helm (optional)

## Setup

1. **Clone the repository**
2. **Install dependencies:**
   ```bash
   go mod download
   ```
3. **Run the application locally:**
   ```bash
   DB_HOST=localhost DB_PORT=5432 DB_USER=postgres DB_PASSWORD=postgres DB_NAME=postgres go run cmd/app/main.go
   ```
   The API will be available at http://localhost:8080

## Docker Setup

1. **Build and run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```
2. **The API will be available at http://localhost:8080/api/messages**
3. **Health check endpoint: http://localhost:8080/health**

## Testing

1. **Run tests:**
   ```bash
   go test -v ./...
   ```
2. **Run tests with coverage:**
   ```bash
   go test -v -coverprofile=coverage.out ./...
   go tool cover -html=coverage.out -o coverage.html
   ```
3. **Run linter:**
   ```bash
   go install golang.org/x/lint/golint@latest
   golint ./...
   ```

## GitHub Actions Setup

1. **Push the repository to GitHub**
2. **Configure the following secrets in GitHub:**
   - `KUBE_CONFIG`: Base64-encoded kubeconfig for your Kubernetes cluster
3. **The pipeline will automatically run on push to main branch**

## Kubernetes Deployment

1. **Apply the Kubernetes manifests:**
   ```bash
   kubectl apply -f deployments/k8s/
   ```
2. **Or install the Helm chart:**
   ```bash
   helm install go-postgres-app ./deployments/helm
   ```

## API Endpoints

- `GET /api/messages` - Get all messages
- `POST /api/messages` - Create a new message
- `GET /health` - Health check endpoint

## CI/CD Pipeline Stages

1. **Test**: Runs linting, tests, and code coverage
2. **Build and Push**: Builds Docker image and pushes to GitHub Container Registry
3. **Deploy**: Deploys to Kubernetes with health verification

## Project Structure

```
go-postgresql-githubactions-k8s/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   └── repository/
├── pkg/
│   └── database/
├── deployments/
│   ├── k8s/
│   └── helm/
├── .github/
│   └── workflows/
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## Customization

1. **Update database connection settings** in docker-compose.yml and Kubernetes manifests
2. **Modify the GitHub Actions workflows** for your specific needs
3. **Add your deployment scripts** for production environment
4. **Update the Docker image name** in deployment manifests

## Health Monitoring

The application includes comprehensive health monitoring:

- Health check endpoint at `/health`
- Kubernetes liveness and readiness probes
- Database connection health checks
- Deployment verification in CI/CD pipeline
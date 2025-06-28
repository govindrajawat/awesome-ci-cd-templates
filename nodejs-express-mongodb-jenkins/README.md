# Node.js + Express + MongoDB + Jenkins CI/CD Template

This template provides a production-ready Node.js application with Express, MongoDB, Docker, and Jenkins CI/CD pipeline.

## Features

- Node.js + Express REST API
- MongoDB database integration
- Docker containerization with multi-stage build
- Jenkins pipeline for:
  - Running tests with MongoDB service
  - Building and pushing Docker images to Docker Hub
  - Deploying to staging environment
- Unit and integration tests with Jest
- ESLint for code quality
- Health check endpoint for monitoring
- Comprehensive test coverage reporting

## Prerequisites

- Node.js 18+
- Docker
- MongoDB (for local development)
- Jenkins server

## Setup

1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```
3. Copy `env.example` to `.env` and update with your configuration:
   ```bash
   cp env.example .env
   ```
4. Run the application:
   ```bash
   npm start
   npm run dev
   ```

## Docker Setup

1. Build and run with Docker Compose:
   ```bash
   docker-compose up --build
   ```
2. The API will be available at http://localhost:3000/api/messages
3. Health check endpoint: http://localhost:3000/health

## Jenkins Setup

1. Create a new Jenkins pipeline job
2. Point it to your repository's Jenkinsfile
3. Configure required credentials in Jenkins:
   - `docker-hub-credentials`: Docker Hub username/password
   - `staging-server-credentials`: SSH key for deployment server
4. Set environment variables in Jenkins:
   - `DOCKER_IMAGE`: Your Docker Hub image name
5. Run the pipeline

## Testing

1. Run tests with:
   ```bash
   npm test
   ```
2. Run linter with:
   ```bash
   npm run lint
   ```
3. View coverage report:
   ```bash
   open coverage/lcov-report/index.html
   ```

## API Endpoints

- `GET /api/messages` - Get all messages
- `POST /api/messages` - Create a new message
- `GET /health` - Health check endpoint

## Customization

1. Update Dockerfile and docker-compose.yml with your application specifics
2. Modify the Jenkins pipeline in Jenkinsfile
3. Update staging deployment configuration in docker-compose.staging.yml
4. Add your production deployment scripts

## CI/CD Pipeline Stages

1. **Checkout**: Clone the repository
2. **Install Dependencies**: Run `npm ci`
3. **Lint**: Run ESLint for code quality
4. **Test**: Run Jest tests with MongoDB service container
5. **Build Docker Image**: Create optimized Docker image
6. **Push to Docker Hub**: Push image to registry
7. **Deploy to Staging**: Deploy to staging environment (main branch only)
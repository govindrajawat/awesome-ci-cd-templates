# Spring Boot + MySQL + GitLab CI/CD Template

This template provides a production-ready Spring Boot application with MySQL, Docker, and GitLab CI/CD pipeline.

## Features

- Spring Boot REST API with JPA
- MySQL database integration
- Docker containerization with multi-stage build
- GitLab CI pipeline for:
  - Building and testing the application
  - Building and pushing Docker images to GitLab Container Registry
  - Deploying to staging environment
- Unit and integration tests
- Database migration with Hibernate
- Health check endpoints for monitoring
- Spring Boot Actuator for production monitoring

## Prerequisites

- Java 17+
- Maven
- Docker
- MySQL (for local development)
- GitLab account

## Setup

1. Clone the repository
2. Install dependencies:
   ```bash
   mvn install
   ```
3. Run the application:
   ```bash
   mvn spring-boot:run
   ```

## Docker Setup

1. Build and run with Docker Compose:
   ```bash
   docker-compose up --build
   ```
2. The API will be available at http://localhost:8080/api/messages
3. Health check endpoint: http://localhost:8080/health
4. Actuator health endpoint: http://localhost:8080/actuator/health

## GitLab CI Setup

1. Push the repository to GitLab
2. Configure the following environment variables in GitLab CI/CD settings:
   - `STAGING_SSH_KEY`: SSH private key for deployment server
3. The pipeline will automatically run on push

## Testing

1. Run tests with:
   ```bash
   mvn test
   ```
2. Run tests with specific profile:
   ```bash
   mvn test -Dspring.profiles.active=test
   ```

## API Endpoints

- `GET /api/messages` - Get all messages
- `POST /api/messages` - Create a new message
- `GET /health` - Custom health check endpoint
- `GET /actuator/health` - Spring Boot Actuator health endpoint

## Application Profiles

- **dev**: Development environment (default)
- **test**: Testing environment with test database
- **prod**: Production environment with environment variables

## CI/CD Pipeline Stages

1. **Build**: Compile the application
2. **Test**: Run tests with MySQL service container
3. **Package**: Create JAR file
4. **Docker Build**: Build and push Docker image
5. **Deploy**: Deploy to staging environment (main branch only)

## Customization

1. Update `pom.xml` with your dependencies
2. Modify the GitLab CI pipeline in `.gitlab-ci.yml`
3. Update staging deployment configuration in `docker-compose.staging.yml`
4. Add your production deployment scripts

## Database Migrations

For production use, consider adding Flyway or Liquibase for database migrations:

1. Add Flyway dependency to `pom.xml`:
   ```xml
   <dependency>
       <groupId>org.flywaydb</groupId>
       <artifactId>flyway-core</artifactId>
   </dependency>
   ```
2. Create SQL migration files in `src/main/resources/db/migration`

## Health Monitoring

The application includes comprehensive health monitoring:

- Custom health endpoint at `/health`
- Spring Boot Actuator health endpoint at `/actuator/health`
- Database connection health checks
- Docker health checks in containers
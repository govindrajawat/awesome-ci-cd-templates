# .NET Core + SQL Server + Azure Pipelines CI/CD Template

This template provides a production-ready .NET 7 Web API application with SQL Server, Docker, and Azure Pipelines CI/CD.

## Features

- .NET 7 Web API with EF Core
- SQL Server database integration
- Docker containerization with multi-stage build
- Azure Pipelines for:
  - Building and testing the application
  - Code coverage reporting
  - Building and pushing Docker images to Azure Container Registry
  - Deploying to Azure App Service or Kubernetes
- Unit tests with xUnit and Moq
- Health check endpoint for monitoring
- Swagger/OpenAPI for API documentation

## Prerequisites

- .NET 7 SDK
- Docker
- SQL Server (for local development)
- Azure DevOps account

## Setup

1. **Clone the repository**
2. **Restore dependencies and build:**
   ```bash
   dotnet restore
   dotnet build
   ```
3. **Run the application locally:**
   ```bash
   dotnet run --project src/dotnet-core-sqlserver-azurepipelines.csproj
   ```
   The API will be available at http://localhost:5000 (or as configured).

4. **Run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```
   The API will be available at http://localhost:8080/api/messages

## Testing

1. **Run unit tests:**
   ```bash
   dotnet test
   ```

## Health Check & API Docs

- **Health check endpoint:**  
  [http://localhost:8080/health](http://localhost:8080/health)
- **Swagger UI (development only):**  
  [http://localhost:8080/swagger](http://localhost:8080/swagger)

## Azure Pipelines Setup

1. **Create a new pipeline in Azure DevOps and point it to `azure-pipelines.yml`.**
2. **Set the following pipeline variables/secrets:**
   - `your-docker-registry-service-connection`
   - `your-azure-subscription`
   - `ACR_USERNAME`, `ACR_PASSWORD` (for Azure Container Registry)
   - Any other variables referenced in the pipeline

## Customization

- Update Dockerfile and docker-compose.yml as needed.
- Modify the Azure Pipelines YAML for your deployment target.
- Add integration tests for advanced scenarios.

## Project Structure

## Database Migrations

To create and apply migrations:
   ```bash
    dotnet ef migrations add InitialCreate
    dotnet ef database update
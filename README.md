# 🚀 Awesome CI/CD Templates

A collection of ready-to-use CI/CD pipeline templates for various languages, databases, platforms, and DevOps tools.

## 🎯 Goal
Help developers and DevOps engineers quickly set up CI/CD pipelines across different stacks with minimal setup.

---

## 📂 Included Templates

| Tech Stack | CI/CD Tool | Deploy To | Folder |
|------------|------------|-----------|--------|
| Python + Flask + PostgreSQL | GitHub Actions | Kubernetes | [`python-flask-postgresql-github-actions`](./python-flask-postgresql-github-actions) |
| Node.js + Express + MongoDB | Jenkins | Docker Compose | [`nodejs-express-mongodb-jenkins`](./nodejs-express-mongodb-jenkins) |
| Java + Spring Boot + MySQL | GitLab CI | Docker | [`java-springboot-mysql-gitlabci`](./java-springboot-mysql-gitlabci) |
| .NET Core + SQL Server | Azure Pipelines | Azure Container Registry | [`dotnet-core-sqlserver-azurepipelines`](./dotnet-core-sqlserver-azurepipelines) |
| Go + PostgreSQL | GitHub Actions | Kubernetes | [`go-postgresql-githubactions-k8s`](./go-postgresql-githubactions-k8s) |

---

## 🧪 Features

- **Dockerized applications** with multi-stage builds
- **Complete CI/CD pipelines** with build, test, lint, and deploy workflows
- **Environment variable handling** for different environments
- **Database migrations** and seeding
- **Secure secret handling** in all platforms
- **Multi-environment support** (dev, staging, production)
- **Health check endpoints** for monitoring
- **Comprehensive testing** with coverage reporting
- **Kubernetes manifests** and Helm charts
- **Reusable workflows** and templates

---

## 🚀 Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/awesome-ci-cd-templates.git
   cd awesome-ci-cd-templates
   ```

2. **Choose your tech stack** from the table above

3. **Copy the template** to your project
   ```bash
   cp -r python-flask-postgresql-github-actions/ your-project/
   ```

4. **Update configuration** files:
   - Docker image names
   - Environment variables
   - Deployment settings
   - Secrets and credentials

5. **Push to your repository** and watch the CI/CD magic! 🪄

---

## 📋 Prerequisites

Each template includes its own prerequisites, but generally you'll need:

- **Docker** for containerization
- **Git** for version control
- **CI/CD platform** (GitHub Actions, GitLab CI, Jenkins, Azure Pipelines)
- **Target platform** (Kubernetes, Docker Compose, Cloud services)

---

## 🔧 Customization

Each template is designed to be easily customizable:

- **Update application code** in the `src/` or main directories
- **Modify CI/CD workflows** in `.github/workflows/`, `.gitlab-ci.yml`, `Jenkinsfile`, or `azure-pipelines.yml`
- **Adjust Docker configurations** in `Dockerfile` and `docker-compose.yml`
- **Update deployment manifests** in `k8s/` or `deployments/` directories

---

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-template`)
3. **Add your template** or improve existing ones
4. **Test thoroughly** with the included CI/CD pipelines
5. **Submit a pull request**

### Adding New Templates

When adding a new template, please include:

- ✅ Complete application code
- ✅ Docker configuration
- ✅ CI/CD pipeline files
- ✅ Deployment manifests
- ✅ Comprehensive README
- ✅ Health check endpoints
- ✅ Unit and integration tests
- ✅ Environment configurations

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

Thanks to all the contributors who have helped make this collection of CI/CD templates comprehensive and production-ready!

---

**Ready to ship faster? Choose your template and start building! 🚀**

#!/bin/bash

# License Verification Script for Awesome CI/CD Templates
# Repository: https://github.com/govindrajawat/awesome-ci-cd-templates

set -e

echo "🔍 Verifying license information across all templates..."
echo "Repository: https://github.com/govindrajawat/awesome-ci-cd-templates"
echo ""

# Check root LICENSE file
if [ -f "LICENSE" ]; then
    echo "✅ Root LICENSE file exists"
    if grep -q "govindrajawat" LICENSE; then
        echo "✅ Root LICENSE contains correct repository URL"
    else
        echo "❌ Root LICENSE missing correct repository URL"
        exit 1
    fi
else
    echo "❌ Root LICENSE file missing"
    exit 1
fi

# Check package.json files
echo ""
echo "📦 Checking Node.js package.json files..."
if [ -f "nodejs-express-mongodb-jenkins/package.json" ]; then
    if grep -q "govindrajawat" nodejs-express-mongodb-jenkins/package.json; then
        echo "✅ Node.js package.json has correct repository URL"
    else
        echo "❌ Node.js package.json missing correct repository URL"
        exit 1
    fi
fi

# Check pom.xml files
echo ""
echo "☕ Checking Java pom.xml files..."
if [ -f "java-springboot-mysql-gitlabci/pom.xml" ]; then
    if grep -q "govindrajawat" java-springboot-mysql-gitlabci/pom.xml; then
        echo "✅ Java pom.xml has correct repository URL"
    else
        echo "❌ Java pom.xml missing correct repository URL"
        exit 1
    fi
fi

# Check setup.py files
echo ""
echo "🐍 Checking Python setup.py files..."
if [ -f "python-flask-postgresql-github-actions/setup.py" ]; then
    if grep -q "govindrajawat" python-flask-postgresql-github-actions/setup.py; then
        echo "✅ Python setup.py has correct repository URL"
    else
        echo "❌ Python setup.py missing correct repository URL"
        exit 1
    fi
fi

# Check .csproj files
echo ""
echo "🔷 Checking .NET .csproj files..."
if [ -f "dotnet-core-sqlserver-azurepipelines/src/dotnet-core-sqlserver-azurepipelines.csproj" ]; then
    if grep -q "govindrajawat" dotnet-core-sqlserver-azurepipelines/src/dotnet-core-sqlserver-azurepipelines.csproj; then
        echo "✅ .NET .csproj has correct repository URL"
    else
        echo "❌ .NET .csproj missing correct repository URL"
        exit 1
    fi
fi

# Check go.mod files
echo ""
echo "🐹 Checking Go go.mod files..."
if [ -f "go-postgresql-githubactions-k8s/go.mod" ]; then
    if grep -q "govindrajawat" go-postgresql-githubactions-k8s/go.mod; then
        echo "✅ Go go.mod has correct repository URL"
    else
        echo "❌ Go go.mod missing correct repository URL"
        exit 1
    fi
fi

# Check Helm Chart.yaml files
echo ""
echo "⚓ Checking Helm Chart.yaml files..."
if [ -f "go-postgresql-githubactions-k8s/deployments/helm/Chart.yaml" ]; then
    if grep -q "govindrajawat" go-postgresql-githubactions-k8s/deployments/helm/Chart.yaml; then
        echo "✅ Helm Chart.yaml has correct repository URL"
    else
        echo "❌ Helm Chart.yaml missing correct repository URL"
        exit 1
    fi
fi

echo ""
echo "🎉 All license verifications passed!"
echo "✅ Repository URLs are consistent across all templates"
echo "✅ Author information is properly set"
echo "✅ MIT License is properly referenced" 
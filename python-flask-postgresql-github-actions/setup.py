from setuptools import setup, find_packages

with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()

setup(
    name="flask-postgresql-app",
    version="1.0.0",
    author="Govind Rajawat",
    author_email="govindrajawat260398@gmail.com",
    description="Flask + PostgreSQL CI/CD Template",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/govindrajawat/awesome-ci-cd-templates",
    packages=find_packages(),
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.9",
    ],
    python_requires=">=3.9",
    install_requires=[
        "Flask>=2.0.1",
        "Flask-SQLAlchemy>=2.5.1",
        "psycopg2-binary>=2.9.1",
        "python-dotenv>=0.19.0",
        "gunicorn>=20.1.0",
    ],
    extras_require={
        "dev": [
            "pytest>=6.2.5",
        ],
    },
    license="MIT",
) 
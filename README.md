
# Check Assets API

This project uses the clean architecture concepts with Golang and MongoDB to build a API to storage and monitor price's financial assets.

**REQUIREMENTS**: Docker

## Getting started

To run the project, you need to execute `docker-compose up --build -d`. The default port is **8080** (and you can change in the docker-compose). The tests can be executed using `./test.sh` shell script found in the folder root.

The API documentation is running on **:8081** (and you can also change in `docker-compose.yml`).

The free version of Yahoo Finance API is limited by 100 requests per day. Because of that, there are five API Keys available in `docker-compose.yml`. You can change if it is necessary.
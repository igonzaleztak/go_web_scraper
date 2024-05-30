#!/bin/bash

cd ..

# Build the Docker image
docker build -t intelygenz_scraper:latest .

# Run the Docker container
docker run --rm --name intelygenz_scraper intelygenz_scraper:latest
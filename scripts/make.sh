#!/bin/bash

cd ..

# Run tests
make test

# Generate coverage
make coverage

# Build the binary file
make build

# Run the binary file
./intelygenz_scraper
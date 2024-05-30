# Intelygenz scraper
- [Intelygenz scraper](#intelygenz-scraper)
  - [Guidelines](#guidelines)
  - [Overview](#overview)
  - [Thought process](#thought-process)
  - [Installation and usage](#installation-and-usage)
    - [Local](#local)
    - [Docker](#docker)

## Guidelines

A golang web scraper designed for the Hacker News webpage. This scraper will take the first 30 entries from the Hacker News frontpage (API: https://github.com/HackerNews/API) and then it will sort them in the following manner:

1. First come entries with more than 5 words in their title
2. Long title entries will then be sorted by number of comments
3. Short title entries will be sorted by number of points

Then sorted entries will be printed out by console output.
Candidates are required to fill a README in ENGLISH explaining their thought process while completing the assignment, discussing key technical points, blocking points, etc. They should also explain test coverage and how to test the code.

- PLUS: Try scrapping different new sources (Hint: Use goroutines)
- PLUS: Automation (Testing, running, publishing) with scripting, Make, Taskfile,...


## Overview

This repository contains a golang web scraper. This scraper is a command-line tool
that takes the first 30 entries from the Hacker News API and then sorts them as it is specified in the guidelines.

The user can specify the number of words that a title must have to be considered long and the number of news that will 
be fetched from the sources. By default, the tool will fetch 30 news from each source and will consider a title long if
it has more than 5 words.

At the moment, the tool supports the Hacker News API and the Space Flight News API. Therefore, when the user runs the
tool, it will scrap information from both sources. The second source has been added to achieve the first 
PLUS requirement.

The second PLUS requirement has been achieved by generating the Makefile and the Dockerfile. This way, the user can
test, run and publish the tool easily.

The tool accepts the following arguments:

| Argument              | Description                                                                | Type      | Default Value | Required |
|-----------------------|----------------------------------------------------------------------------|-----------|---------------|----------|
| `-v`, `--verbose`     | Enable verbose mode. Supported modes Debug: 0, Info: 1                     | `integer` | `0`           | No       |
| `-n`, `--max-stories` | Defines the number of news that will be fetched from the sources           | `integer` | `30`          | No       |
| `-w`, `--num words`   | Indicates the number of words that a title must have to be considered long | `integer` | `5`           | No       |
| `-h`, `--help`        | Display help information                                                   | `boolean` | `false`       | No       |


**Output**: The tool will print to console the scraped news sorted as it is indicated in the guidelines. You can easily 
identify the data source at the beginning of the printed line. You can see below how the tool displays the news.

```sh
# data scraped from Space Flight API
[spaceflight.scraper] story ID: 23668

# data scraped from Hacker News API
[hackernews.scraper] story ID: 40523582
```

## Thought process

A detailed description of the thought process that I have followed to complete this skill test can be found in the the 
[process.md](docs/process.md) file.

## Installation and usage

The tool can either be run locally or using Docker. The following sections describe how to run the tool in both ways.

### Local
A makefile has been provided to facilitate the execution of the tool. The following commands can be used to run the tool:

```bash
# test the tool
make test

# Generate coverage
make coverage

# Build the tool
make build

# Run the tool
./intelygenz_scraper
```

You can also run the [make.sh](scripts/make.sh) script to execute the previous commands. 

The tool will fetch 30 news from each source and will consider a title long if it has more than 5 words. However, the
user can specify the number of news that will be fetched from the sources and the number of words that a title must have
to be considered long. The following commands show how to run the tool with different configurations:

```bash
# Run the tool with 10 news from each source
./intelygenz_scraper -n 10

# Run the tool with 10 news from each source and consider a title long if it has more than 10 words
./intelygenz_scraper -n 10 -w 10
```

User can use the `-h` flag to display the help information.

```bash
# Display help information
./intelygenz_scraper -h

# Output
Intelygenz Scraper aims to obtain several news from the Hacker News website

Usage:
  scrap [flags]

Flags:
  -h, --help              help for scrap
  -n, --max-stories int   Defines the number of news that will be fetched from the sources
  -w, --num-words int     Indicates the number of words that a title must have to be considered long
  -v, --verbose verbose   Enable verbose mode. Supported modes Debug: 0, Info: 1 (default log)
```


### Docker
The tool can also be run using Docker. The following commands can be used to run the tool:

```bash
# Build the Docker image
docker build -t intelygenz_scraper:latest .

# Run the Docker container
docker run -rm intelygenz_scraper
```

You can also run the [docker.sh](scripts/make.sh) script to execute the previous commands.

To modify the dockerized tool's behavior, the user can pass the same arguments as in the local execution. The following commands
show how to run the tool with different configurations:

```bash
# Run the tool with 10 news from each source
docker run -rm intelygenz_scraper -n 10

# Run the tool with 10 news from each source and consider a title long if it has more than 10 words
docker run -rm intelygenz_scraper -n 10 -w 10
```

# Line of thought

This document describes the process that I have followed to complete this skill test. The first thing that has popped
into my mind is that this application should be a CLI tool that takes several input arguments to modify its
functionality. Here, I had doubts whether this was the best option to design the tool, but I decided to go with this 
approach since it does not make sense to create an API for this tool.

The arguments that the tool will accept are the following:

| Argument              | Description                                                                                                                                                                                                                                                                                  | Type      | Default Value | Required |
|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------|---------------|----------|
| `-m`, `--mode`        | Mode in which the tool should work. It accepts two modes: `api` and `web`.If the `api` mode is sset the tool will query directly the source API to obtain the desired data. Otherwise, if this input is set to `web` the tool will scrap the website HTML to obtain the desired information. | `string`  | `api`         | No       |
| `-v`, `--verbose`     | Enable verbose mode                                                                                                                                                                                                                                                                          | `boolean` | `false`       | No       |
| `-n`, `--max-stories` | Defines the number of news that will be fetched from Hackers News                                                                                                                                                                                                                            | `integer` | `30`          | No       |
| `-w`, `--num words`   | Indicates the number of words that a title must have to be considered long                                                                                                                                                                                                                   | `integer` | `5`           | No       |
| `-h`, `--help`        | Display help information                                                                                                                                                                                                                                                                     | `boolean` | `false`       | No       |

After defining what type of application should I design, I proceeded to create the project structure. The project 
structure is as follows:

```
.
├── config
│   └──  config.yaml    
├── docs
│   └── process.md
├── logs
├── pkg 
│   └── cmd 
│   └── config
│   └── enums
│   └── logs
│   └── schemas
│   └── scraper
│   └── utils 
├── go.mod
├── go.sum
├── main.go
```

The folder config located on the root of the project contains the configuration file that the tool will use to define
the app behavior. Here, users can define the URLs of the Web page to scrap, the API URL, and the path to the log file.

While the README.md will contain the instructions to run the application, the process.md file will contain the line of
process that I have follow to complete this skill test.

Logs will contain the log file that the tool will generate to store the information of the execution.

The pkg folder contains the different packages that I have created to build the tool.

- cmd: Contains the different commands that the tool will accept.
- config: Contains the structs that stores the configuration of the tool and the flags that the tool will accept.
- enums: Contains the different enums that the tool will use.
- logs: Contains the logger that the tool will use to log the information.
- schemas: Contains the structs that the tool accepts when querying the Hacker News API.
- scraper: Contains the different functions that the tool will use to scrap the website.

The main.go file contains the entry point of the application.

I have tried to design the tool in a way that it is easy to extend and maintain. Thus, I have used libraries
that are de facto standard for this purpose, including:
    
- cobra: A CLI library that allows the creation of powerful modern CLI interfaces. I have used this library for the
creation of the cli command to interact with the tool. You can find the library's documentation [here](github.com/spf13/viper).
- viper: A library that allows the reading of configuration files. I have used this library to read the configuration
file that the tool uses to define its behavior. You can find the library's documentation [here](github.com/spf13/viper).
- zap: A library that allows the logging of information. I have used this library to log the information of the 
tool, specifically I have defined two levels of verbosity: info and debug. You can find the library's documentation [here](github.com/uber-go/zap).




# Test coverage

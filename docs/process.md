# Line of thought

This document describes the process that I have followed to complete this skill test. The first thing that has popped
into my mind is that this application should be a CLI tool that takes several input arguments to modify its
functionality. The arguments that the tool will accept are the following:

| Argument              | Description                                                                | Type      | Default Value | Required |
|-----------------------|----------------------------------------------------------------------------|-----------|---------------|----------|
| `-v`, `--verbose`     | Enable verbose mode. Supported modes Debug: 0, Info: 1                     | `integer` | `0`           | No       |
| `-n`, `--max-stories` | Defines the number of news that will be fetched from the sources           | `integer` | `30`          | No       |
| `-w`, `--num words`   | Indicates the number of words that a title must have to be considered long | `integer` | `5`           | No       |
| `-h`, `--help`        | Display help information                                                   | `boolean` | `false`       | No       |

After defining what type of application should I design, I proceeded to create the project structure as it can be seen
below.

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
the app behavior. Here, users can define the URLs of the APIs that are going to be scraped and the path to the log
folder.

While the README.md contains a general overview of the tool and the instructions to run the application,
the [process.md](process.md) file contains the line of process that I have followed to complete this skill test.

The folder logs contains all the logs that the tool has generated during its execution. I wanted to include this log
system to ease the visualization of news that the tool has fetched. This way, the user can check the logs to see the
news that the tool is working properly. Depending on the level of verbosity, the logs file will be more or less
detailed.

The pkg folder contains tool's source code.

- cmd: Contains the initialization of the CLI command.
- config: Contains the structs that stores the configuration of the tool and the flags that the tool will accept.
- enums: Contains the different enums that the tool will use. In this case, there is only one enum that defines the
  verbosity level.
- logs: Contains the logger that the tool will use to log the information. This logger outputs the logs to the console
  and to the log file.
- schemas: Contains the structs that the tool accepts when querying the APIs. These structs are used to unmarshal
  the JSON response from the APIs.
- scraper: Contains the different functions that the tool will use to scrap the website. I will describe this package
  in more detailed below.

The main.go file contains the entry point of the application.

# Technical key points

I have tried to design the tool in a way that it is easy to extend and maintain. Thus, I have used tools that are
commonly used.

- cobra: A CLI library that allows the creation of powerful modern CLI interfaces. I have used this library for the
  creation of the cli command to interact with the tool. You can find the library's
  documentation [here](https://github.com/spf13/cobra).
- viper: A library that allows the reading of configuration files. I have used this library to read the configuration
  file. You can find the library's documentation [here](https://github.com/spf13/viper).
- zap: A logging's library. I have used this library to log the information of the tool, specifically I have defined two
  levels of verbosity: info and debug. You can find the library's
  documentation [here](https://github.com/uber-go/zap).
- automaxprocs: Library that sets automatically the number of GOMAXPROCS to the host's CPU quota. I have used this
  library due to the requirement of scraping different data sources using goroutines. You can find the library's
  documentation [here](https://github.com/uber-go/automaxprocs).

Like I have mentioned earlier, I wanted to create the tool as generic as possible, so it can be easy to maintain. Thus,
the scraper functionality has been created using a simple generic interface with two methods: `Scrap()` and `Print()`.
The definition of this interface can be found on the [pkg/scraper/interface.go](../pkg/scraper/interface.go) file.

```go
type Scraper interface {
  Scrap() error
  Print()
}
```

By using this interface, we can create as many scrapers as required. Therefore, to create a new scraper, we would need
to follow the next steps:

1. Write the URL of the scraper on the app's config file. This is optional, we can always hardcode the API's URL
   directly in the struct.
2. Create a scraper struct that implements these two methods.
3. Start the scraper in the [pkg/cmd/root.go](../pkg/cmd/root.go) file. In this file, we would need to initialize the
   scraper that we want to use and start the scraping process with this new scraper as it can be seen below.

```go
// initialize the scraper we want to use
spaceFlightScraper := spaceFlight.NewScraper(config.AppConfig.SpaceFlightNewsAPI)

// add the function to the errgroup subroutines manager. This will start the scraper in a new goroutine.
errs.Go(func () error { return scraper.StartScraperProcess(spaceFlightScraper) })
```

In the previous code snippet, you'll notice that the initialized API is the Space Flight News API, not the Hacker News
API. This choice aligns with one of the bonus points of this task, which was to scrape different news sources. The Space
Flight News API is open and does not require an API key for queries. You can find the documentation for this
API[here](https://api.spaceflightnewsapi.net/v4/docs/).

Additionally, I have used the errgroup package to manage the goroutines that the tool uses to scrape the data. This
package was chosen to effectively handle errors that the goroutines might return.

# Test coverage

The tool contains unit tests and integration tests. The unit tests are located in the same package that the function
that is being tested. For example, the unit test for the `Scrap()` function of the `HackerNewsScraper` struct is located
in the [pkg/scraper/hackerNews/scraper_test.go](../pkg/scraper/hackerNews/scraper_test.go) file. The same applies to
other functions.

The integration test is located in the root of the project as [main_test.go](../main_test.go). This test will test the
whole application.

To create the test files, I have used the strechr/testify library. This library provides a set of functions that makes
the testing process easier. You can find the library's documentation [here](https://github.com/stretchr/testify). I
always try to encapsulate the test logic in testify suites to avoid collision between tests.

Additionally, I have used the go test -coverprofile flag to generate the test coverage. You can visualize the test
coverage by running the following command:

```bash
make coverage
```

This command will generate a `coverage.out` file with all the test coverage information.

# Running the tool

To build the tool you can use the Makefile that I have created. By executing the command `make build` the script will
build the tool and create an executable file called `intelygenz_scraper`. This file will be located in the
root of the project. Then, you can execute the tool by following the instructions described in
the [README.md](../README.md) file.

Furthermore, I have created a Dockerfile that will build the tool and create a Docker image. You can build the Docker
image and execute the tool by following the instructions described in the [README.md](../README.md) file. I have
decided to create a Docker image to make sure that the tool can be executed in any environment.

# Line of thought

 This document describes the process that I have followed to complete this skill test. The first thing that has popped
 into my mind is that this application should be a CLI tool that takes several input arguments to modify its 
 functionality. These are described below.

| Argument          | Description                                                                                                                                                                                                                                                                                            | Type      | Default Value | Required |
|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------|---------------|----------|
| `-m`, `--mode`    | Mode in which the tool should work. It accepts two modes: `api` and `web`.<br/>If the `api` mode is sset the tool will query directly the source API to obtain the desired data. <br/>Otherwise, if this input is set to `web` the tool will scrap the website HTML to obtain the desired information. | `string`  | `api`         | No       |
| `-v`, `--verbose` | Enable verbose mode                                                                                                                                                                                                                                                                                    | `boolean` | `false`       | No       |
| `-h`, `--help`    | Display help information                                                                                                                                                                                                                                                                               | `boolean` | `false`       | No       |


After defining what type of should I design, the next step consists on defining the folder structure of the 
project.

-- CLI library -> github.com/spf13/cobra@latest
-- read default config  -> github.com/spf13/viper

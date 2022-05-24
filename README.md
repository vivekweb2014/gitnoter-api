## GitNoter API Module

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/git-noter/gitnoter-api/Test/main?color=forestgreen)](https://github.com/git-noter/gitnoter-api/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/git-noter/gitnoter-api/branch/main/graph/badge.svg?token=pWRurWucMC)](https://codecov.io/gh/git-noter/gitnoter-api)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/6b8593ef77a041a89a2dfa9aa9eea154)](https://www.codacy.com/gh/git-noter/gitnoter-api/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=git-noter/gitnoter-api&amp;utm_campaign=Badge_Grade)

This is the api service of gitnoter application built with golang.

It exposes ReST endpoints to access and manage markdown notes from user's git repository. Currently it only supports github repository for storing & managing notes.

### Local Development Setup

#### Prerequisites
*   Go version `1.18` or above
*   Docker Desktop

#### Start postgres database container
Below commands use docker to start the database container.
```shell
make network
make postgres
make createdb
```

#### Create configuration file from template
The `config.yaml` is the configuration template file containing default config values.
```shell
cp config.yaml .config.yaml
```
Application uses `.config.yaml` file to get the config values. Please update placeholder-values from this config file to the actual ones. 

#### Start the server
Make sure that the `.config.yaml` file is configured correctly & database container is up & running.
Then run the below commands to setup the db schema and start the web server.
```shell
go run main.go migrateup
go run main.go serve
```
This will start the server on port specified in `.config.yaml` file. You can now access the api endpoints.

#### Run tests
```shell
go test -v -cover ./...
```
This will execute all the tests and also prints the code coverage percentage.

### Contribution Guidelines
> Every Contribution Makes a Difference

Read the [Contribution Guidelines](CONTRIBUTING.md) before you contribute.

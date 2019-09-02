[![Build Status](https://travis-ci.org/YAWAL/TodayInHistory.svg?branch=master)](https://travis-ci.org/YAWAL/TodayInHistory)

# Day In History

Web server is using API https://history.muffinlabs.com/#api .

## Prerequisites:
1. [Golang](https://golang.org/dl/) (1.12)
2. [Postgres](https://www.postgresql.org/download/) 
3. [Dependency management for Go](https://golang.github.io/dep/docs/installation.html) Dep

## How to start a server:
1. Clone or download repository to your local machine
2. Open root folder and run command ```dep ensure```
3. Set up config file - provide port and Postgres credential, f.e.:
```
{
    "host": ":8888",
    "database": {
        "dialect": "postgres",
        "user": "yourPostgresUser",
        "db_name": "yourDatabaseName",
        "ssl_mode": "disable",
        "password": "yourPostgresPassword"
    }
}
```
4. From root folder run ```go run main.go```
4. Open browser on port ```:8888``` or port provided by you in config file and use following endpoints

## Endpoints:

- http://localhost:8888/history/date
- http://localhost:8888/history/date/{month}/{day}, where month (1-12) and day (1-31)
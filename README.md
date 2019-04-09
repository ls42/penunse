# Penunse

Expense Tracker Web Application written in Go, HTML, CSS and SQL. Ready 
to be deployed as a single binary.

The descriptions in this README might not be up-to-date.

## Architecture

This project is made of two parts:

1. An API, that handles all data storage tasks and makes the data 
available via a REST interface. The API is available under the 
`/api/*`-route.
1. A reference client, that is written with Go-HTML-Templates.

### API Server

There is an API server that's written in Go. It's external dependencies are listed in [go.mod](penunse/blob/master/go.mod)


### Reference Client

When accessing the route `/` the server will hand out a reference 
client. I initially designed this as a JavaScript frontend for the API, 
but then I got turned off by all the JavaScript ecosystem overhead (with 
Vue.js, NPM, vue-cli and suddenly I had around 200 external 
dependencies) and so I decided to switch to a frontend that's made off 
Go and it's `template/html` stdlib.

The goal of the client is to have as few dependencies as possible 
(that's the goal for the whole project in fact).  


## Tasks

### Update data/penunse.sql

This might need to be updated once in a while, when the gorm schema changes, or when I want to add more sample data (or hundred other reasons).

```bash
sqlite3 data/penunse.db
sqlite> .output data/penunse.sql
sqlite> .dump
sqlite> .quit
# easy peasy
```


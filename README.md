# Penunse Personal Expense Tracker

Expense Tracker Web Application

## Architecture

This project is made of two parts:

1. An API server, that handles all data storage tasks and makes the data available via a REST interface.
1. A reference client, that is written in JavaScript, CSS and HTML.

### API Server

There is an API server that's written in Go. It's external dependencies are listed in [go.mod](penunse/tree/master/go.mod)


### JavaScript reference client

When accessing the route `/` the server will hand out a reference API client.


Goals are:

- Stick with vanilla ES6 source. I try to keep external dependencies minimal and include the source directly in this repo, no CDN or external requirements during build time.

So far these are the dependecies I think I need (sure we can strip those too, somehow):

* [ToasterJS](https://github.com/ZitRos/toaster-js)

The client talks to the server via the JavaScript Request API. It's an REST API.

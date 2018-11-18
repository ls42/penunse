# Penunse Personal Expense Tracker

Learn where your money is going. Maybe I'll bake in some useful stats later.

## Architecture

### API Server

There is an API server that's written in Go. It's external dependencies are listed in [go.mod](~stephan/penunse/tree/master/go.mod)


### JavaScript reference client

When accessing the route `/` the server will hand out a reference API client.


Goals are:

- Stick with vanilla ES6 source. I try to keep external dependencies minimal and include the source directly in this repo, no CDN or external requirements during build time.

So far these are the dependecies I think I need (sure we can strip those too, somehow):

* [ToasterJS](https://github.com/ZitRos/toaster-js)

The client talks to the server via the JavaScript Request API. It's an REST API.

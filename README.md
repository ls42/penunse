# Penunse Personal Expense Tracker

Learn where your money is going. Maybe I'll bake in some useful stats later.

# Architecture

## API Server

There is an API server that's written in Go. It's external dependencies are:

* [Bolt](https://github.com/boltdb/bolt)


## JavaScript reference client

Stick with vanilla ES6 source. I try to keep external dependencies minimal and include the source directly in this repo, no CDN or external requirements during build time:

* [ToasterJS](https://github.com/ZitRos/toaster-js)

The client talks to the API server via the JavaScript Request API. Via REST.

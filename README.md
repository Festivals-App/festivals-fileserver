# Eventus File Server

[![License](https://img.shields.io/github/license/phisto/eventusfileserver.svg)](https://github.com/Phisto/eventusfileserver)

Eventus file server, a live and lightweight go server app.

## Overview

A simple RESTful API with Go using go-chi/chi.

## Requirements

-  go 1.11

## Installation

```bash
go get github.com/Phisto/eventusfileserver
```

# Build and Run
```bash
cd $GOPATH/github.com/Phisto/eventusfileserver
go build main.go
./main
# API Endpoint : http://localhost:1910
```

## Structure
```
├── server
│   ├── server.go               // server logic
│   │
│   ├── handler                 // API handlers
│   │   ├── common.go           // Common response functions
│   │   ├── image.go            // APIs for handling images
│   │   └── status.go           // APIs getting server status information
│   │
│   └── manipulate
│       ├── resize.go            // APIs for resizing images
│       └── toolbox.go           // Misc funktions
│
├── config
│   └── config.go               // Configuration
│
└── main.go               
```


## Eventus API

#### /images/upload
* `POST`    : Upload an image

#### /images/{imageIdentifier}
* `GET`     : Get an image
* `PATCH`   : Replace an image

#### /status
* `GET`     : Get the server status

#### /status/files
* `GET`     : Get a list of all files

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 

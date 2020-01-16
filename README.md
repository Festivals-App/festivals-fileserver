# Eventus Image Server

[![License](https://img.shields.io/github/license/phisto/eventusimageserver.svg)](https://github.com/Phisto/eventusimageserver)

Eventus image server, a live and lightweight go server app.

## Overview

A simple RESTful API with Go using go-chi/chi.

## Requirements

-  go 1.11

## Installation

```bash
go get github.com/Phisto/eventusimageserver
```

# Build and Run
```bash
go build
./server

# API Endpoint : http://localhost:1910
```

## Structure
```
├── server
│   ├── server.go               // server logic
│   │
│   ├── handler                 // API handlers
│   │   ├── common.go           // Common response functions
│   │   └── festival.go         // APIs for the Festival model
│   │
│   └── model
│       └── model.go            // The data models
│
├── config
│   └── config.go               // Configuration
│
└── main.go               
```


## Eventus API

#### /festivals
* `GET`     : Get all festivals
* `POST`    : Create a new festival

#### /festivals/{objectID}
* `GET`     : Get a festival
* `PATCH`   : Update a festival
* `DELETE`  : Delete a festival

#### /festivals/{objectID}/{image|links|place|tags}
* `GET`     : Get the given associated objects

#### /festivals/{objectID}/{image|links|place|tags}/{resourceID}
* `POST`     : Associates the object with the given ID with the festival with the given ID

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 

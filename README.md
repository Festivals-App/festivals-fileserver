<h1 align="center">
    Festivals App File Server
</h1>

<p align="center">
   <a href="https://github.com/festivals-app/festivals-fileserver/commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/festivals-app/festivals-fileserver?style=flat"></a>
   <a href="https://github.com/festivals-app/festivals-fileserver/issues" title="Open Issues"><img src="https://img.shields.io/github/issues/festivals-app/festivals-fileserver?style=flat"></a>
   <a href="./LICENSE" title="License"><img src="https://img.shields.io/github/license/festivals-app/festivals-fileserver.svg"></a>
</p>

<p align="center">
  <a href="#development">Development</a> •
  <a href="#usage">Usage</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#engage">Engage</a> •
  <a href="#licensing">Licensing</a>
</p>

A live and lightweight go server app providing a simple RESTful API using [go-chi/chi](https://github.com/go-chi/chi) and [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql).

## Development

TBA

### Requirements

-  go 1.11

### Setup development

TBA

## Usage

TBA

### API

#### /images/upload
* `POST`    : Upload an image

#### /images/{imageIdentifier}
* `GET`     : Get an image
* `PATCH`   : Replace an image

#### /status
* `GET`     : Get the server status

#### /status/files
* `GET`     : Get a list of all files

### Structure
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
### Documentation

The full documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. The documentation repository contains technical documents, architecture information, UI/UX specifications, and whitepapers related to this implementation.

## Deployment

Before running the API server, you should set the database config with your values in config/config.go

```bash
go get github.com/Festivals-App/festivals-fileserver
```

### Build and Run
```bash
cd $GOPATH/github.com/Festivals-App/festivals-fileserver
go build main.go
./main
# API Endpoint : http://localhost:1910
```

## Engage

TBA

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Concept Feedback**    | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="Open Concept Feedback"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/architecture.svg?style=flat-square"></a>  |
| **Other Requests**    | <a href="mailto:phisto05@gmail.com" title="Email Festivals Team"><img src="https://img.shields.io/badge/email-Festivals%20team-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

## Licensing

Copyright (c) 2020 Simon Gaus.

Licensed under the **GNU Lesser General Public License v3.0** (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License at https://www.gnu.org/licenses/lgpl-3.0.html.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the [LICENSE](./LICENSE) for the specific language governing permissions and limitations under the License.
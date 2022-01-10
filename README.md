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
  <a href="#deployment">Deployment</a> •
  <a href="#festivalsfilesapi">FestivalsFilesAPI</a> •
  <a href="#architecture">Architecture</a> •
  <a href="#engage">Engage</a> •
  <a href="#licensing">Licensing</a>
</p>

A live and lightweight go server app providing a simple RESTful API called FestivalsFilesAPI using [go-chi/chi](https://github.com/go-chi/chi).

## Development

The developement of the [FestivalsFilesAPI](./DOCUMENTATION.md) and the festivals-fileserver is quite forward and does *not* dependend on the [festivals-api-ios](https://github.com/Festivals-App/festivals-api-ios) client library.

### Requirements

- [Golang](https://go.dev/) Version 1.17+
- [Visual Studio Code](https://code.visualstudio.com/download) 1.63.2+
    * Plugin recommendations are managed via [workspace recommendations](https://code.visualstudio.com/docs/editor/extension-marketplace#_recommended-extensions).
- [Bash script](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) friendly environment

### Structure
```
├── server
│   ├── server.go               // server logic
│   │
│   ├── config
│   │   └── config.go           // Server configuration
|   |
│   ├── handler                 // API handlers
│   │   ├── common.go           // Common response functions
│   │   ├── image.go            // APIs for handling images
│   │   ├── pdf.go              // APIs for handling pdfs
│   │   └── status.go           // Handling status server status requests
│   │
│   └── manipulate
│   │   ├── resize.go            // APIs for resizing images
│   │   └── toolbox.go           // Misc funktions
│   │
│   └── status
│       └── status.go           // API for the status of the application
│
└── main.go               
```

## Deployment

Running the festivals-fileserver is pretty easy because Go binaries are able to run without system dependencies 
on the target for which they are compiled. The only dependency is that the festivals-fileserver expects either a config file at `/etc/festivals-fileserver.conf`,
the environment variables set or the template config file present in the directory it runs from.

### Build and Run manually

```bash
cd /path/to/repository/festivals-fileserver
make build
(make install)
make run

# Default API Endpoint : http://localhost:1910
```

### VM deployment

The install, update and uninstall scripts should work with any system that uses *systemd* and *firewalld*.
Additionally the scripts will somewhat work under macOS but won't configure the firewall or launch service.

Installing
```bash
curl -o install.sh https://raw.githubusercontent.com/Festivals-App/festivals-fileserver/master/operation/install.sh
chmod +x install.sh
sudo ./install.sh
```
Updating
```bash
curl -o update.sh https://raw.githubusercontent.com/Festivals-App/festivals-fileserver/master/operation/update.sh
chmod +x update.sh
sudo ./update.sh
```
Uninstalling
```bash
curl -o uninstall.sh https://raw.githubusercontent.com/Festivals-App/festivals-fileserver/master/operation/uninstall.sh
chmod +x uninstall.sh
sudo ./uninstall.sh
```

To see if the server is running use:
```bash
systemctl status festivals-fileserver
```

### Container deployment

```bash
TBA
```

### FestivalsFilesAPI

The FestivalFilesAPI is documented in detail [here](./DOCUMENTATION.md).

## Architecture

![Figure 1: Architecture Overview Highlighted](https://github.com/Festivals-App/festivals-documentation/blob/main/images/architecture/overview_fileserver.png "Figure 1: Architecture Overview Highlighted")

The general documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. 
The documentation repository contains architecture information, general deployment documentation, templates and other helpful documents.

## Engage

I welcome every contribution, whether it is a pull request or a fixed typo. The best place to discuss questions and suggestions regarding the festivals-server is the [issues](https://github.com/festivals-app/festivals-fileserver/issues/) section. More general information and a good starting point if you want to get involved is the [festival-documentation](https://github.com/Festivals-App/festivals-documentation) repository.

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Other Requests**    | <a href="mailto:simon.cay.gaus@gmail.com" title="Email me"><img src="https://img.shields.io/badge/email-Simon-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

## Licensing

Copyright (c) 2017-2022 Simon Gaus.

Licensed under the **GNU Lesser General Public License v3.0** (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License at https://www.gnu.org/licenses/lgpl-3.0.html.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the [LICENSE](./LICENSE) for the specific language governing permissions and limitations under the License.
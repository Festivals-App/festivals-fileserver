<p align="center">
   <a href="https://github.com/festivals-app/festivals-fileserver/commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/festivals-app/festivals-fileserver?style=flat"></a>
   <a href="https://github.com/festivals-app/festivals-fileserver/issues" title="Open Issues"><img src="https://img.shields.io/github/issues/festivals-app/festivals-fileserver?style=flat"></a>
   <a href="./LICENSE" title="License"><img src="https://img.shields.io/github/license/festivals-app/festivals-fileserver.svg"></a>
</p>

<h1 align="center">
    <br/><br/>
    Festivals App File Server
    <br/><br/>
</h1>

<p align="center">
  <a href="#development">Development</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#festivalsfilesapi">FestivalsFilesAPI</a> •
  <a href="#architecture">Architecture</a> •
  <a href="#engage">Engage</a> •
  <a href="#licensing">Licensing</a>
</p>

A lightweight go server app providing a simple RESTful API using [go-chi/chi](https://github.com/go-chi/chi). The FestivalsFilesAPI
exposes file storage and file manipulation functions to be used for all file assets needed by the FestivalsApp.

![Figure 1: Architecture Overview Highlighted](https://github.com/Festivals-App/festivals-documentation/blob/main/images/architecture/overview_files.png "Figure 1: Architecture Overview Highlighted")

## Development

The developement of the [FestivalsFilesAPI](./DOCUMENTATION.md) and the festivals-fileserver is quite forward and does *not* dependend on the [festivals-api-ios](https://github.com/Festivals-App/festivals-api-ios) client library directly.

### Requirements

- [Golang](https://go.dev/) Version 1.20+
- [Visual Studio Code](https://code.visualstudio.com/download) 1.83.1+
    * Plugin recommendations are managed via [workspace recommendations](https://code.visualstudio.com/docs/editor/extension-marketplace#_recommended-extensions).
- [Bash script](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) friendly environment

## Deployment

The festivals-fileserver expects either a config file at `/etc/festivals-fileserver.conf` or the template config file present in the directory it runs from, 
it also expects the root certificate and the server certificate and key to establish secure inter-service communication as well as the client-service communication.

### Build and Run manually

```bash
cd /path/to/repository/festivals-fileserver
make build
(make install)
make run

# Default API Endpoint : https://localhost:1910
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
sudo systemctl status festivals-fileserver
```

### FestivalsFilesAPI

The FestivalFilesAPI is documented in detail [here](./DOCUMENTATION.md).

## Architecture

The general documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. 
The documentation repository contains architecture information, general deployment documentation, templates and other helpful documents.

## Engage

I welcome every contribution, whether it is a pull request or a fixed typo. The best place to discuss questions and suggestions regarding the festivals-fileserver is the [issues](https://github.com/festivals-app/festivals-fileserver/issues/) section. More general information and a good starting point if you want to get involved is the [festival-documentation](https://github.com/Festivals-App/festivals-documentation) repository.

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Other Requests**    | <a href="mailto:simon.cay.gaus@gmail.com" title="Email me"><img src="https://img.shields.io/badge/email-Simon-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

## Licensing

Copyright (c) 2017-2023 Simon Gaus.

Licensed under the **GNU Lesser General Public License v3.0** (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License at https://www.gnu.org/licenses/lgpl-3.0.html.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the [LICENSE](./LICENSE) for the specific language governing permissions and limitations under the License.
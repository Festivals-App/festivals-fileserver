<p align="center">
   <a href="https://github.com/festivals-app/festivals-fileserver/commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/festivals-app/festivals-fileserver?style=flat"></a>
   <a href="https://github.com/festivals-app/festivals-fileserver/issues" title="Open Issues"><img src="https://img.shields.io/github/issues/festivals-app/festivals-fileserver?style=flat"></a>
   <a href="./LICENSE" title="License"><img src="https://img.shields.io/github/license/festivals-app/festivals-fileserver.svg"></a>
</p>

<h1 align="center">
    <br/><br/>
    FestivalsApp File Server
    <br/><br/>
</h1>

A lightweight service providing a RESTful API called **FestivalsFilesAPI**. The FestivalsFilesAPI
exposes file storage and file manipulation functions to be used for all file assets needed by the FestivalsApp.

![Figure 1: Architecture Overview Highlighted](https://github.com/Festivals-App/festivals-documentation/blob/main/images/architecture/architecture_overview_file_server.svg "Figure 1: Architecture Overview Highlighted")

<hr />
<p align="center">
  <a href="#development">Development</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#engage">Engage</a>
</p>
<hr />

## Development

The developement of the FestivalsFilesAPI [(see documentation)](./DOCUMENTATION.md) and the festivals-fileserver is quite forward and does *not* dependend on the [festivals-api-ios](https://github.com/Festivals-App/festivals-api-ios) client library directly.

To find out more about the architecture and technical information see the [ARCHITECTURE](./ARCHITECTURE.md) document. The general documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. The documentation repository contains architecture information, general deployment documentation, templates and other helpful documents.

- [Golang](https://go.dev/) Version 1.21.5+
- [Visual Studio Code](https://code.visualstudio.com/download) 1.85.2+
  - Plugin recommendations are managed via [workspace recommendations](https://code.visualstudio.com/docs/editor/extension-marketplace#_recommended-extensions).
- [Bash script](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) friendly environment

## Deployment

The Go binaries are able to run without system dependencies so there are not many requirements for the system to run the festivals-fileserver binary.
The config file needs to be placed at `/etc/festivals-fileserver.conf` or the template config file needs to be present in the directory the binary runs in.

You also need to provide certificates in the right format and location:

- The default path to the root CA certificate is          `/usr/local/festivals-fileserver/ca.crt`
- The default path to the server certificate is           `/usr/local/festivals-fileserver/server.crt`
- The default path to the corresponding key is            `/usr/local/festivals-fileserver/server.key`

Where the root CA certificate is required to validate incoming requests and the server certificate and key is required to make outgoing connections.
For instructions on how to manage and create the certificates see the [festivals-pki](https://github.com/Festivals-App/festivals-pki) repository.

### VM

The install, update and uninstall scripts should work with any system that uses *systemd* and *firewalld*.
Additionally the scripts will somewhat work under macOS but won't configure the firewall or launch service.

```bash
#Installing
curl -o install.sh https://raw.githubusercontent.com/Festivals-App/festivals-fileserver/master/operation/install.sh
chmod +x install.sh
sudo ./install.sh

#Updating
curl -o update.sh https://raw.githubusercontent.com/Festivals-App/festivals-fileserver/master/operation/update.sh
chmod +x update.sh
sudo ./update.sh

#To see if the server is running use:
sudo systemctl status festivals-fileserver
```

#### Build and run using make

```bash
make build
make run
# Default API Endpoint : https://localhost:1910
```

## Engage

I welcome every contribution, whether it is a pull request or a fixed typo. The best place to discuss questions and suggestions regarding the festivals-fileserver is the [issues](https://github.com/festivals-app/festivals-fileserver/issues/) section. More general information and a good starting point if you want to get involved is the [festival-documentation](https://github.com/Festivals-App/festivals-documentation) repository.

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Other Requests**    | <a href="mailto:simon.cay.gaus@gmail.com" title="Email me"><img src="https://img.shields.io/badge/email-Simon-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

### Licensing

Copyright (c) 2017-2024 Simon Gaus. Licensed under the [**GNU Lesser General Public License v3.0**](./LICENSE)

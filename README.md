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

![Figure 1: Architecture Overview Highlighted](https://github.com/Festivals-App/festivals-documentation/blob/main/images/architecture/export/architecture_overview_file_server.svg "Figure 1: Architecture Overview Highlighted")

<hr />
<p align="center">
  <a href="#development">Development</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#engage">Engage</a>
</p>
<hr/>

## Development

The developement of the FestivalsFilesAPI [(see documentation)](./DOCUMENTATION.md) and the festivals-fileserver is quite forward and does *not* dependend on the [festivals-api-ios](https://github.com/Festivals-App/festivals-api-ios) client library directly.

The general documentation for the Festivals App is in the [festivals-documentation](https://github.com/festivals-app/festivals-documentation) repository. The documentation repository contains architecture information, general deployment documentation, templates and other helpful documents.

### Requirements

- [Golang](https://go.dev/) Version 1.24.1+
- [Visual Studio Code](https://code.visualstudio.com/download) 1.98.2+
  - Plugin recommendations are managed via [workspace recommendations](https://code.visualstudio.com/docs/editor/extension-marketplace#_recommended-extensions).
- [Bash script](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) friendly environment

## Deployment

The Go binaries are able to run without system dependencies so there are not many requirements for the system to run the festivals-fileserver binary,
just follow the [**deployment guide**](./operation/DEPLOYMENT.md) for deploying it inside a virtual machine or the [**local deployment guide**](./operation/local/README.md)
for running it on your macOS developer machine.

## Engage

I welcome every contribution, whether it is a pull request or a fixed typo. The best place to discuss questions and suggestions regarding the festivals-fileserver is the [issues](https://github.com/festivals-app/festivals-fileserver/issues/) section. More general information and a good starting point if you want to get involved is the [festival-documentation](https://github.com/Festivals-App/festivals-documentation) repository.

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Other Requests**    | <a href="mailto:simon@festivalsapp.org" title="Email me"><img src="https://img.shields.io/badge/email-Simon-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

### Licensing

Copyright (c) 2017-2025 Simon Gaus. Licensed under the [**GNU Lesser General Public License v3.0**](./LICENSE)

<!--
-->

# Apache APISIX Dashboard

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/apache/apisix-dashboard/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/apache/apisix-dashboard)](https://goreportcard.com/report/github.com/apache/apisix-dashboard)
[![DockerHub](https://img.shields.io/docker/pulls/apache/apisix-dashboard.svg)](https://hub.docker.com/r/apache/apisix-dashboard)
[![Cypress.io](https://img.shields.io/badge/tested%20with-Cypress-04C38E.svg)](https://www.cypress.io/)
[![Slack](https://badgen.net/badge/Slack/Join%20Apache%20APISIX?icon=slack)](https://apisix.apache.org/slack)

<p align="center">
  <a href="https://apisix.apache.org/">Website</a> •
  <a href="https://github.com/apache/apisix/tree/master/docs">Docs</a> •
  <a href="https://twitter.com/apacheapisix">Twitter</a>
</p>

- The master version should be used with Apache APISIX master version.

- The latest released version is [3.0.1](https://apisix.apache.org/downloads/) and is compatible with [Apache APISIX 3.0.x](https://apisix.apache.org/downloads/).

## What's Apache APISIX Dashboard

The Apache APISIX Dashboard is designed to make it as easy as possible for users to operate [Apache APISIX](https://github.com/apache/apisix) through a frontend interface.

The Dashboard is the control plane and performs all parameter checks; Apache APISIX mixes data and control planes and will evolve to a pure data plane.

Note: Currently the Dashboard does not have complete coverage of Apache APISIX features, [visit here](https://github.com/apache/apisix-dashboard/milestones) to view the milestones.

![architecture](./docs/assets/images/architecture.png)

## Demo

[Online Playground](https://apisix-dashboard.apiseven.com/)

```text
Username: admin
Password: admin
```

## Works with APISIX Ingress Controller

Currently, APISIX Ingress Controller automatically manipulates some APISIX resources, which is not very compatible with APISIX Dashboard. In addition, users should not modify resources labeled `managed-by: apisix-ingress-controllers` via APISIX Dashboard.

## Project structure

```text
.
├── CHANGELOG.md
├── CODE_OF_CONDUCT.md
├── CONTRIBUTING.md
├── Dockerfile
├── LICENSE
├── Makefile
├── NOTICE
├── README.md
├── api
├── docs
├── licenses
└── web
```

1. The `api` directory is used to store the `Manager API` source codes, which is used to manage `etcd` and provide APIs to the frontend interface.
2. The `web` directory is used to store the frontend source codes.

## Build then launch

Support the following ways currently.

- [Docker, RPM, Source Codes](./docs/en/latest/install.md)
- [Rebuild docker image](./docs/en/latest/deploy-with-docker.md)

## Development

Pull requests are encouraged and always welcome. [Pick an issue](https://github.com/apache/apisix-dashboard/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22) and help us out!

Please refer to the [Development Guide](./docs/en/latest/develop.md).

## User Guide

Please refer to the [User Guide](./docs/en/latest/USER_GUIDE.md).

## Contributing

Please refer to the [Contribution Guide](./CONTRIBUTING.md) for a more detailed information.

## FAQ

Please refer to the [FAQ](./docs/en/latest/FAQ.md) for more known issues.

## License

[Apache License 2.0](./LICENSE)

# rhttp
rhttp allows you to periodically clone a Git Repository and expose its contents to the web

## Setup
Setting up rhttp is fairly easy. Either use the Docker image or build it yourself. Configuration is done via
environment variables which can also be set using a .env file in the applications directory.

## Environment variables
| Variable              | Default   | Description                                                    |
|-----------------------|-----------|----------------------------------------------------------------|
| `RHTTP_REPO_URL`      | `<empty>` | The URL of the repository to use                               |
| `RHTTP_UPDATE_PERIOD` | `1h`      | The period to use to update the repository                     |
| `RHTTP_BLACKLIST`     | `.git/`   | A set of blacklisted files and/or directories, separated by `` |
| `RHTTP_WEB_ADDRESS`   | `:8080`   | The web address rhttp should listen to                         |

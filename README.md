# gotoggl

[![release](https://github.com/akira393/gotoggl/actions/workflows/release.yml/badge.svg)](https://github.com/akira393/gotoggl/actions/workflows/release.yml)

gotoggl is command line interface tool for toggl api client.

## Setup

```bash
$ gotoggl configure

toggl api token: xxxxxxxxxxxxxxx
```

check install

```bash
$ gotoggl version
gotoggl version: 1.0.0, revision: a1d30dd, buildTag: v1.0.0
```

set toggl api token

```bash
$ gotoggl configure
toggl api token: xxxxxxxxxxxxxx
$ cat ~/.gotoggl.yaml
toggl_api_token: xxxxxxxxxxxxxx
```

## Usage

```bash
$ gotoggl help
Using config file: /Users/YOURHOMENAME/.gotoggl.yaml

The  gotoggl Command  Line  Interface is a unified tool to manage your toggl services.

Usage:
  gotoggl [flags]
  gotoggl [command]

Available Commands:
  client      Manage client resources
  completion  generate the autocompletion script for the specified shell
  configure   Create config
  help        Help about any command
  project     Manage project resources
  report      Manage report resources
  tag         Manage tag resources
  version     Print version

Flags:
      --config string   config file (default is $HOME/.gotoggl.yaml)
  -h, --help            help for gotoggl
  -t, --toggle          Help message for toggle

Use "gotoggl [command] --help" for more information about a command.
```

## LICENSE

check this file. =>[LICENSE](./LICENSE)
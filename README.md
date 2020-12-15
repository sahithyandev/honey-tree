honey-tree
==========

A tool to create &amp; use project templates

[![oclif](https://img.shields.io/badge/cli-oclif-brightgreen.svg)](https://oclif.io)
[![Version](https://img.shields.io/npm/v/honey-tree.svg)](https://npmjs.org/package/honey-tree)
[![Downloads/week](https://img.shields.io/npm/dw/honey-tree.svg)](https://npmjs.org/package/honey-tree)
[![License](https://img.shields.io/npm/l/honey-tree.svg)](https://github.com/sahithyandev/honey-tree/blob/master/package.json)

<!-- toc -->
* [Usage](#usage)
* [Commands](#commands)
<!-- tocstop -->
# Usage
<!-- usage -->
```sh-session
$ npm install -g honey-tree
$ honey-tree COMMAND
running command...
$ honey-tree (-v|--version|version)
honey-tree/0.0.1 linux-x64 node-v14.15.1
$ honey-tree --help [COMMAND]
USAGE
  $ honey-tree COMMAND
...
```
<!-- usagestop -->
# Commands
<!-- commands -->
* [`honey-tree hello [FILE]`](#honey-tree-hello-file)
* [`honey-tree help [COMMAND]`](#honey-tree-help-command)

## `honey-tree hello [FILE]`

describe the command here

```
USAGE
  $ honey-tree hello [FILE]

OPTIONS
  -f, --force
  -h, --help       show CLI help
  -n, --name=name  name to print

EXAMPLE
  $ honey-tree hello
  hello world from ./src/hello.ts!
```

_See code: [src/commands/hello.ts](https://github.com/sahithyandev/honey-tree/blob/v0.0.1/src/commands/hello.ts)_

## `honey-tree help [COMMAND]`

display help for honey-tree

```
USAGE
  $ honey-tree help [COMMAND]

ARGUMENTS
  COMMAND  command to show help for

OPTIONS
  --all  see all commands in CLI
```

_See code: [@oclif/plugin-help](https://github.com/oclif/plugin-help/blob/v3.2.1/src/commands/help.ts)_
<!-- commandsstop -->

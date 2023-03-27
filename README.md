# Neko Project Generator

[![Go Report Card](https://goreportcard.com/badge/github.com/KGRC199913/neko-projectGenerator)](https://goreportcard.com/report/github.com/KGRC199913/neko-projectGenerator)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/KGRC199913/neko-projectGenerator/blob/main/LICENSE)

## Introduction

Neko Project Generator is a CLI tool that generates project templates for various programming languages. Currently, it supports the following languages:

- Go
- Node.js

This tool can be helpful for developers who want to start new projects quickly without worrying about the initial setup.

## Installation

To install the tool, you need to have Go installed on your system. Then, run the following command:

```bash
go get github.com/KGRC199913/neko-projectGenerator
```

## Usage

After installing the tool, you can use it to generate project templates by running the following command:

```bash
neko-projectGenerator generate <project-name> <language>
```

Replace <project-name> with the name of your project, and <language> with the programming language you want to use (either go or node).

For example, to generate a Go project template named my-project, run the following command:

```bash

neko-projectGenerator generate my-project go
```


## Contributing

Contributions to this project are welcome. If you find a bug or have a feature request, please open an issue on GitHub.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
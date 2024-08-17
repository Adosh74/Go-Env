# Go-Env

1. [Description](#description)
2. [Project Structure](#project-structure)
3. [Packages](#packages)
4. [Installation](#installation)
5. [Usage](#usage)

## Description

This my base project for Go web applications. It includes a basic structure and clean architecture for web applications, configuration management, and CLI.

## Project Structure

```
go-env
├───.git
├───assets
├───cmd
├───config
├───internal
│   ├───modules
│   │   └───home
│   │       ├───html
│   │       └───routes
│   └───providers
│       ├───routes
│       └───view
└───pkg
    ├───bootstrap
    ├───config
    ├───html
    ├───routing
    └───static
```

## Packages

- Gin: Web framework
- Viper: Configuration management
- Cobra: CLI

## Installation

```bash
git clone https://github.com/Adosh74/Go-Env.git

cd Go-Env

go mod download
```

## Usage

```bash
go run main.go serve
```

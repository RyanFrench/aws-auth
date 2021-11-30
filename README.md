# aws-auth

Small CLI utility to assume a role in AWS

Using this utility, a command can be run using an assumed role, as opposed to your default profile. The role assumption is not persisted to your shell, so subsequent requests must also use this utility.


## Installation

### From source

Currently, the only way to install the utility is from source. It requires you to have [go](https://golang.org) installed on your system.

`go get github.com/ryanfrench/aws-auth`

## Usage

### Assume a role

`aws-auth --role-arn=role [--duration=seconds] [command]`

This will assume `role` and then run `command` using that role.

### Parameters

| Parameter | Description| Required | Default |
|-----------|------------|----------|---------|
| role-arn  | The role to assume | Yes |
| duration  | The duration, in seconds, before the token will expire | No | 3600

### Help

`aws-auth --help`

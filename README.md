# aws-role

Small CLI utility to assume a role in AWS

Using this utility, a command can be run using an assumed role, as opposed to your default profile. The role assumption is not persisted to your shell, so subsequent requests must also use this utility.


## Installation

### From source

Currently, the only way to install the utility is from source. It requires you to have [go](https://golang.org) installed on your system.

`go get github.com/ryanfrench/aws-role`

## Usage

### Assume a role

`aws-role --role-arn=[role] [command]`

This will assume `[role]` and then run `[command]` using that role.

### Help

`aws-role --help`

# Microsample Service

This is the Microsample service

Generated with

```
micro new mysamples/microsample --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.microsample
- Type: srv
- Alias: microsample

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./bin/microsample-srv
```

Build a docker image
```
make docker
```
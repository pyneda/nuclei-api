# Nuclei API

The Nuclei API package provides a simple gRPC API wrapper around the Nuclei scanner, allowing easy integration with other tools.


## Installation

```bash
go install -v github.com/pyneda/nuclei-api@latest
```

or

```
git clone https://github.com/pyneda/nuclei-api
cd nuclei-api
go build
```


## Usage

```bash
nuclei-api start -h
```

## Example client

An example client is provided in the `cmd/example` directory. 

It can be built with:

```bash
go build cmd/client/main.go
```

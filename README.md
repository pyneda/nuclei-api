# Nuclei API

The Nuclei API package provides a simple gRPC wrapper around the [Nuclei scanner](https://github.com/projectdiscovery/nuclei), allowing easy integration with other tools.


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
go build cmd/example/client.go
```

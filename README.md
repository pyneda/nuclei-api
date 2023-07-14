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

Start the server:

```bash
nuclei-api start -h
```

### Configuration

Most of the options available in Nuclei to choose and filter templates are available as a part of the scan request.

Other options such as rate limits, interactsh configuration, etc are available as a part of the server configuration which can be configured using a YAML file. The following command is available to dump the default configuration:

```bash
nuclei-api dumpconfig
```

## Example client

An example client is provided in the `cmd/example` directory. 

It can be built with:

```bash
go build cmd/example/client.go
```

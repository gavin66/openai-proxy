# OPENAI Proxy Server

## Building

- `git clone https://github.com/gavin66/openai-proxy`
- `cd openai-proxy`
- `go build`

The output should be a binary `openai-proxy`

## Running

To run the binary, simply execute it with `./openai-proxy`

## Deployment

There are multiple ways to deploy this proxy: `docker` or simply running it on a server.

### Docker

- Install [Docker](https://docs.docker.com/get-docker/)
- Build the docker image with `docker build -t openai-proxy .`
- Remember to build the binary first

### Running on a server

- Just run the binary on a server of the same architecture as the machine you built it on

### Configuration

The only configuration is the port to run the proxy on. This can be set with the `PORT` environment variable. The
default port is `8890`.

You can use `nginx` to reverse proxy the port to a domain and/or Cloudflare to proxy the port to a domain.
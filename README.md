# Go Proxy Server

This project aims to create a simple proxy server using Go, providing an opportunity to understand how proxy servers work and gain hands-on experience with Go programming language.

## Features

- [ ] Forward HTTP requests from clients to target servers
- [ ] Forward responses from target servers back to clients
- [ ] Handle multiple concurrent connections efficiently
- [ ] Logging of requests and responses
- [ ] Basic caching mechanism for improved performance

## Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/): Ensure Go is properly installed on your system.

## Getting Started

1. **Clone the repository:**

    ```sh
    git clone <repository_url>
    cd go-proxy-server
    ```

2. **Build the project:**

    ```sh
    go build
    ```

3. **Run the server:**

    ```sh
    ./go-proxy-server
    ```

## Usage

To use the proxy server, configure your client to use `localhost` at the port the proxy server is running on (default is 8080).

## Configuration

You can configure the proxy server by modifying the `config.json` file. Options include:

- Port: The port on which the proxy server listens for incoming connections.
- TargetHost: The host to which the proxy server forwards incoming requests.

## Example

Here's an example of how you can configure the proxy server in `config.json`:

```json
{
  "port": 8080,
  "targetHost": "example.com"
}

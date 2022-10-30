# @block-api/block-node-example

<p align="center" width="100%">
<img src="./docs/images/golang-gopher.png" width="250" alt="golang gopher"/>
</p>

In this repository you can find dockerized two simple microservices built with **block-node**:

- `hello-world-service`
- `ping-pong-service`

## Run it and test yourself

In order to run it and test you need to have installed [Docker](https://www.docker.com/) and [Go language](https://go.dev/) v1.18.

Once you have these two installed, run command below in root directory of the project to build docker images and start services:

```shell
docker-compose up
```

When running it for the first time it might take a bit longer to start as this example require [Redis](https://redis.io/) which is used as a communication transporter - Docker will need to download it as well as Go language.

### hello-world-service

This small microservice has simple implementation of HTTP server exposing two endpoints at `localhost:8090`:

#### `http://localhost:8090/ping`

This endpoint will call `ping-pong-service` which is not exposing any HTTP endpoints and can be reached out only through other microservice in the network.

#### `http://localhost:8090/hello?name=Jhon`

This endpoint will call itself (locally) to invoke proper action and return data.

### ping-pong-service

This microservice contains one action `ping` which should return `pong` as a result. It does not expose any HTTP endpoints like `hello-world-service` and can be reached out by another microservice in the network.

## Developement

If you would like to make any changes inside examples and run it locally for testing without Docker you can go into `hello-world-service` directory or `ping-pong-service` and run command below to start it in developement mode:

```shell
make dev
```

<small>**Note:** To Build binary run</small>

```shell
make build
```

<small>This command will compile files into one executable file which can be found in `build` directory of the service.</small>

For information about configuration options head to **block-node** [repository](https://github.com/block-api/block-node).

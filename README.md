# @block-api/block-node-example

<p align="center" width="100%">
<img src="./docs/images/golang-gopher.png" width="250" alt="golang gopher"/>
</p>

In this repository you can find dockerized two simple microservices built with **block-node**:

- `hello-world-service`
- `ping-pong-service`
- `sqlite-service`

## Run it and test yourself

In order to run it and test you need to have installed [Docker](https://www.docker.com/) and [Go language](https://go.dev/) v1.18.

Once you have these two installed, run command below in root directory of the project to build docker images and start services:

```shell
docker-compose up --scale ping-pong-service=2
```
<small>**Note:** Please keep in mind that in this example implementation if you scale `sqlite-service` each of them will have separate database and there is no data synchronization between those.</small>

When running it for the first time it might take a bit longer to start as this example require [Redis](https://redis.io/) which is used as a communication transporter - Docker will need to download it as well as Go language.

### hello-world-service

This small microservice has simple implementation of HTTP server exposing two endpoints at `localhost:8090`:

#### `http://localhost:8090/ping`

This endpoint will call `ping-pong-service` which is not exposing any HTTP endpoints and can be reached out only through other microservice in the network.

#### `http://localhost:8090/hello?name=Jhon`

This endpoint will call itself (locally) to invoke proper action and return data.

#### `http://localhost:8090/user/add?name=Jhon`

This endpoint will call `sqlite-service` and will add new entry to SQLite database

#### `http://localhost:8090/user/get?name=Jhon`

This endpoint will call `sqlite-service` and will get ID of row in SQLite database for given user name

#### `http://localhost:8090/user/delete?name=Jhon`

This endpoint will call `sqlite-service` and will delete entry from SQLite database

### ping-pong-service

This microservice contains one action `ping` which should return `pong` as a result. It does not expose any HTTP endpoints like `hello-world-service` and can be reached out by another microservice in the network.

### sqlite-service

This microservice has three very simple (please keep in mind there is no sanity checks etc - this is not how you should do it on production environment) actions to `add`, `get` and `delete` entry in SQLite database.

## Development

If you would like to make any changes inside examples and run it locally for testing without Docker you can go into `hello-world-service` directory, `ping-pong-service` or `sqlite-service` and run command below to start it in developement mode:

```shell
make dev
```

<small>**Note:** To Build binary run</small>

```shell
make build
```

<small>This command will compile files into one executable file which can be found in `build` directory of the service.</small>

For information about configuration options head to **block-node** [repository](https://github.com/block-api/block-node).

## Benchmark
Below you can find benchmark results done with [wrk](https://github.com/wg/wrk) benchmarking tool.

Tests were performed on MacOS Monetery *MacBook Pro, 2 GHz Quad-Core Intel Core i5, 32 GB 3733 MHz LPDDR4X* 
running both services and Redis instance.

*Local method invocation*<br>
`wrk -t12 -c500 -d30s http://localhost:8090/hello?name=Jhon`

```text
Running 30s test @ http://localhost:8090/hello?name=Jhon
  12 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.37ms    2.28ms  41.90ms   90.72%
    Req/Sec     8.61k     3.80k   52.65k    55.81%
  3087124 requests in 30.10s, 379.79MB read
  Socket errors: connect 253, read 121, write 0, timeout 0
Requests/sec: 102568.82
Transfer/sec:     12.62MB
```

*Remote node method invocation*<br>
`wrk -t12 -c500 -d30s http://localhost:8090/ping`

```text
Running 30s test @ http://localhost:8090/ping
  12 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    45.69ms    6.38ms  88.01ms   79.38%
    Req/Sec   433.66    226.21     1.17k    61.94%
  155489 requests in 30.02s, 18.24MB read
  Socket errors: connect 253, read 128, write 0, timeout 0
Requests/sec:   5178.84
Transfer/sec:    622.07KB
```
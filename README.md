## Total Block Amount

###### Overview

The service provides single HTTP endpoint: `GET /api/block/<block_number>/total`

where `<block_number>` is a requested block number in decimal format.

Request example: /api/block/ 11508993 /total

###### Configuration

Edit the `conf.json` configuration file to pick/change the following options:

- *api_key:*  the etherscan.io API secret value. If not specified, the default key is provided
- *port:* the http port on which the service will be listening
- *enable_caching:* a bool value indicating whether the service will be caching request results with Redis

###### Usage

To start the service, execute 

```
go run .
```

command in the main directory. To use Redis, the following commands must be executed to download it`s image and run it:

```
$ docker pull redis
$ docker run --name redis-test-instance -p 6379:6379 -d redis
```
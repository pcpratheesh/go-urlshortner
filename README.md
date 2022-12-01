# go-urlshortner
A golang api application which can be used to shorten any url.

## URL Shortening
```
URL shortening is a technique on the World Wide Web in which a Uniform Resource Locator may be made substantially shorter and still direct to the required page. This is achieved by using a redirect which links to the web page that has a long URL
```

## Tech Stack
- [Golang](https://go.dev/)
- [Fiber](https://docs.gofiber.io/) - Express inspired web framework
- [Docker](https://www.docker.com/) 
- [Docker compose](https://docs.docker.com/compose/)


### Golang
Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency

credits:google

### Fiber
Fiber is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind

### Docker
Docker is a set of platform as a service products that use OS-level virtualization to deliver software in packages called containers. The service has both free and premium tiers. The software that hosts the containers is called Docker Engine. It was first started in 2013 and is developed by Docker, Inc

credits:google


### Docker Compose 
Docker Compose is a tool for running multi-container applications on Docker defined using the Compose file format. A Compose file is used to define how the one or more containers that make up your application are configured.

credits:google


## Installation
- install Docker
    - [Linux](https://docs.docker.com/desktop/install/linux-install/)
    - [Windows](https://docs.docker.com/desktop/install/windows-install/)
    - [Mac OS](https://docs.docker.com/desktop/install/mac-install/)
- Install Docker compose
    - https://docs.docker.com/compose/install/other/

- After successfull installation, run the following command to build and up the docker container

        docker-compose up 
        or 
        docker-compose up -d : to run the container in background
- Server will be running now

## API KEY
`encode` is secured with an api key. You can use `3cbc5291f1e04ebe5ea24bfdba6763c49c597cea` as default.

## Try API
curl --request POST \
  --url http://127.0.0.1:8080/encode \
  --header 'Content-Type: application/json' \
  --header 'x-api-key: 3cbc5291f1e04ebe5ea24bfdba6763c49c597cea' \
  --data '{
	"url":"https://google.com"
}'



## How to implement new storage
At the moment, the system is developed with easy implementation of new storages, such as databases, caches, etc.

Follow the below steps to implement the redis cache
- Install redis
- Create `redis.go` and `redis_test.go` files under `pkg/storage`
- Create new structure `RedisStorage` within redis.go 
- Implement `CheckURLExists`, `SaveURL`, `RetrieveURL` methods in `RedisStorage` struct, and a `NewRedisStorage` function to create new object of `RedisStorage`
- Create new global variable `STORE_TYPE_REDIS` with value `Redis`
- Add a new case within the `NewStorage` function to switch to `RedisStorage`
- And the last set teh env config as `Redis` to run the application with `Redis Cache`



## Useful Docker commands

    docker build -t pcpratheesh/go_url_shortner .

    docker images

    docker container ps
    
    docker logs <container_id> --follow

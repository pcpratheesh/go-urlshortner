# go-urlshortner
A golang api application which can be used to shorten any url.

## URL Shortening
```
URL shortening is a technique on the World Wide Web in which a Uniform Resource Locator may be made substantially shorter and still direct to the required page. This is achieved by using a redirect which links to the web page that has a long URL
```

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


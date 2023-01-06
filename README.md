
# Erply remote server with Cache system

## Overall description
This server reads and writes customer data through Erply remote server. 
It also uses Redis for cache to get data without accessing Erply remote 
server if the data was fetched within specific time span

 <br />

## Architecture description
It uses [Gin framework](https://github.com/gin-gonic/gin) and [Redis](https://github.com/go-redis/redis) for local storage caching and 
its architecture focuses on [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). <br />

### Structure details 
- Entity/Model (Storing structs) is in [entity](https://github.com/canigetyourwhatwhat/cloud_base_customer_management/tree/main/entity) folder.
- Usecase/Service (Storing business logic) is in [service](https://github.com/canigetyourwhatwhat/cloud_base_customer_management/tree/main/service) folder.
- Controller (First entrypoint for receiving API) is in [controller](https://github.com/canigetyourwhatwhat/cloud_base_customer_management/tree/main/controllers) folder.
- DB (handling database) is in [database](https://github.com/canigetyourwhatwhat/cloud_base_customer_management/tree/main/infra/database)
  - [Infra](https://github.com/canigetyourwhatwhat/cloud_base_customer_management/tree/main/infra) should contain all the external contents such as handling DB, Payment, send email, Cloud service functions.
- All the API endpoints/paths are in [router.go](https://github.com/canigetyourwhatwhat/cloud_base_customer_management/blob/main/middlewares/router.go).


## How to run
> Note: Docker should be already installed

First, run the container that has Redis inside
```sh
docker run -p 6379:6379 redis
```

Next, run the Go server
```go
go run main.go
```

## Swagger 
> Note: Server and the container must be running 

[here](http://localhost:9000/swagger/index.html) is the Swagger documentation



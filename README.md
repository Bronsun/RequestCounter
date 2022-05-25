# RequestCounter
Request Counter is a simple solution for counting http requests on the multiple server instances.

### Definitions

- Cluster - contains all instances of HTTPserver
- Instance - one instance of HTTPserver

It counts the number of requests to the single server instance and also counts all requests to the cluster. Total number of requests to the cluster is saved to Redis. To protect app from race condition, saving to redis is based on transactions. This solution is great for accuracy in number of requests, however, HTTP Handler function is slower due to waiting for transaction to be closed.

 Docker-compose is used as a conterization solution, but it also simulates 3 instances using Replica method. NGINX is used as the main reverse proxy solution.

![](showcase.gif)


## Requirements

You need to install the following tools:

- Docker with docker-compose -> https://docs.docker.com/
- Make:
    
    For Linux Distro:  
    ```
    sudo apt-get install build-essential
    ```

    For Windows: 
    
    1. Install chocolatey -> https://chocolatey.org/install 
    
    2. Run command
    ```
    choco install make
    ``` 

    For MacOS: 
    
    1. Xcode
    ```
    xcode-select --install or brew install make
    ```

    2. Brew -> https://brew.sh/index_pl
    ```
    brew install make
    ```

## RUN & BUILD

Build project in docker:
```
make build-docker
```

Run project in docker
```
make run-docker
```

Clear docker containers
```
make clean-docker
```

## Tools used

#### NGINX

It is used as a reverse proxy for our web server. 

#### REDIS

It is used as key-value store to saving total number of requests from all instances

## Packages

In project was used Redis, as key-value store for saving total numbers of requests on all server instances. For easier development Redis package was used as a solution for connecting to database.

https://github.com/go-redis/redis









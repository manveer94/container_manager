# Container Manager

## Prerequisite

1. Go should be setup in the machine and present in path.
2. go-sqlite3 dependency has to be installed. Instruction to install can be found [here](https://pkg.go.dev/github.com/mattn/go-sqlite3#section-readme:~:text=Author-,Installation,-This%20package%20can)
        

## Usage
To run the application, use the following command in the terminal.

`go run ./main.go [PORT]`

Server will start on whichever port is provided.

## Project
The project have following packages
1. **server** - It contains code to initialize and start the api server.
2. **database** - Contains code to intilize the connection to database.
3. **model** - Contains all the datatypes.
4. **utils** - Package for utility functions.
5. **container_service** - This package have all the endpoints and services related to containers.
6. **host_service** - This package have all the endpoints and services related to hosts.

container_service and host_service are divided into routes, controller and service.

## API Reference

### Host
1. Lists of all hosts
    
    `GET /host/all`
    
2. Get host by id

    `GET /host/:host_id`


### Container
1. List all containers
    
    `GET /container/all`

2. Get container by id

    `GET /container/:container_id`

3. List containers by host id

    `GET /container/all/host/:host_id`

4. Create new container

    ``` 
    PUT /container

    {
        "ImageName" : "image name",
        "HostId" : 1
    }

    ```
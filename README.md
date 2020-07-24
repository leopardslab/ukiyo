# ukiyo

## Idea incubation

- ukiyo will act as a watcher for docker containers. It will run alongside with the other running containers and will be responsible for automatic updates. Updates will be based on push based model compared to existing solutions such as [watchtower](https://github.com/containrrr/watchtower) and [ouroboros](https://github.com/pyouroboros/ouroboros)

- Push events will be recived from ukiyo via webhooks. Docker registries provide webhooks to subscribe and listen to image changes. Locally running images will change only after such an event is received by ukiyo.

- Pull based model can be implemented as an optional way of updating the running containers.

## Components

- Container manager
- Push manager (webhooks configuration) 
- Scheduler 
- OPTIONAL - Pull based update implementation

## Execution modes

Two modes of execution

- As a container running alongside other containers (Should mount docker.sock to run docker commands inside the ukiyo docker container)
- As a standalone executable

## Language

- Go lang

### Go style guide

- https://github.com/rajikaimal/go-styleguide

## Dev Setup Guide

Setup docker
```sh
$ go mod init ukiyo
$ go mod tidy
```

Create docker binary file
```sh	
$ set GOARCH=amd64
$ set GOOS=linux
$ go build -ldflags="-s -w" -o ukiyo main.go
```

Docker build command
```sh
$ docker build -f Dockerfile -t agentukiyo/ukiyo .
$ docker push agentukiyo/ukiyo
$ docker run -p 8080:8080 \
     -v /var/run/docker.sock:/var/run/docker.sock \
     -v /home/reinventor/dbs:/dbs \
     agentukiyo/ukiyo
```

Add webhook to your dockerhub repository
```
http:{serverIP}:8080/ukiyo-web-hook
```

Add your own docker registy details
```
http:{serverIP}:8080/add-container
{
     "username":"agentukiyo",
     "repoName":"Ukiyo Docker registry",
     "accessToken":"f44e334e-1440-4166-a16f-d8fc9d0eb188",
     "email":"hansika.16@itfac.mrt.ac.lk",
     "serverAddress":"http://docker.io/v1"
}

http:{serverIP}:8080/edit-container-token
{
     "username":"agentukiyo",
     "accessToken":"f44e334e-1440-4166-a16f-d8fc9d0eb188"
}
```
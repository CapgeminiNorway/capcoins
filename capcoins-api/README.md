 CapCoins-API crafted with Golang  
==================================  

CapCoins API server.     

### Prerequisites

In order to build and run this app you need to have a couple of things installed:  

- The [Go SDK](https://golang.org)              
- An IDE for the development, like [Atom](https://atom.io) or IntelliJ/Goland      
- The Docker Toolbox or native Docker, whatever you prefer. See [Docker](https://docs.docker.com) and [Docker-Compose](https://docs.docker.com/compose/)       


#### Clone this repo     

```bash
  # clone this repo  
$ git clone GIT_REPO_URL  
$ cd path/to/capcoins-api  

```   

#### Prepare env-vars  

**WARNING** This APP requires _a valid security token_ for connecting to **MS Teams API**!!!          
You can get it by following [this guide](https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/outgoingwebhook) from Microsoft Teams.      

By default, app checks for env-var to find _vimeoToken_, if not found then it asks user for input.            
```bash
# set vimeoToken as env-variable    
$ export TEAMS_SECRET=<your-security-token>     

# OR just enter it when app asks you for it     

```

### Building & Running this App    
```bash

# Clone this repo and verify compile  
$ git clone GIT_REPO_URL   
$ cd path/to/capcoins-api    

# Run as App from source  
$ go run capcoins-api.go

# build & run it as executable 
$ go build capcoins-api.go
$ ./capcoins-api  

# by default, 'go build' generates executable per env it runs.  
# see https://golang.org/pkg/go/build/    

# if you want to specify per OS compatibility, see each below      
## - linux   
$ GOOS=linux GOARCH=amd64 go build capcoins-api.go     

## - mac   
$ GOOS=darwin GOARCH=amd64 go build capcoins-api.go  

## - windows   
$ GOOS=windows GOARCH=amd64 go build capcoins-api.go  

```
see also [make-build.sh](make-build.sh) which builds for all common platforms   
When you run, it should become available via [http://localhost:8088](http://localhost:8088)      
  

### Containerization with Docker  

Building, publishing and running via _Docker_ :       
```bash
# set env vars for ease-of-use
# NB! please just replace 'zeusbaba' with your own dockerhub username    
$ export dockerhubUser=zeusbaba \
  export appName=capcoins-api \
  export appVersion=1.0.0
$ export dockerImage=${dockerhubUser}/${appName}:${appVersion}

  # required for compatibility
$ GOOS=linux GOARCH=amd64 go build capcoins-api.go

## using Docker!!!       
# build a docker image  
$ docker build \
  -t ${dockerImage} \
  --rm --no-cache .    
$ docker images  	
# (optional) publish the image to docker hub  
$ docker push ${dockerImage}  
```
Alternatively, you can use also use [make-docker.sh](make-docker.sh) script for easy dockerization. 
`sh make-docker.sh` should yield the same result with automation.   

*(optional) you can also run this docker image locally*      
```    
$ docker run \
	-p 8088:8088 \
	${dockerImage}  
```
Now it should be available via [http://localhost:8088](http://localhost:8088)  

Building and running via _Docker-Compose_:         
```bash   
$ docker-compose up --build   

  # NOTE: in linux-env you might have permission problems with 'docker-data-*' folders      
  # to fix; stop docker-compose, set permissions as below, then run docker-compose again.    
$ sudo chown 1001:1001 -R docker-data-*  

  # to shut it down, ctrl+c and   
$ docker-compose down   
```
`docker-compose up` gets the instace up and running together with MongoDB  

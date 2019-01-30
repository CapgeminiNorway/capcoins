#!/usr/bin/env sh

echo "`go version`"

echo "making build... linux"
GOOS=linux GOARCH=amd64 go build capcoins-api.go

echo "prepare env"
# NB! replace this with your own docker hub username
export dockerhubUser=zeusbaba
export appName=capcoins-api
export appVersion=1.0.0
export dockerImage=${dockerhubUser}/${appName}:${appVersion}

echo "docker build for dockerImage: ${dockerImage}"
docker build -t ${dockerImage} --rm --no-cache .

# NB! uncomment this if you want the image to be published on docker
#echo "docker push for dockerImage: ${dockerImage}"
#docker push ${dockerImage}

# to avoid name conflicts of further builds for other envs
mv capcoins-api capcoins-api_linux

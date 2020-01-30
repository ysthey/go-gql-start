#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
cfgFile="config.yml"
outputPath="build"
app="gql-server"
output="$outputPath/$app"
src="$srcPath/$app/$pkgFile"
bcmd="time go build -o $output $src"

printf "\nBuilding: $app\n"
docker run -it --rm -v $GOPATH:/go -v `pwd`:/app -w /app golang:1.13-buster bash -c "$bcmd"
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\npopulate config file"
cp  ./conf/$cfgFile $outputPath
printf "\nclient files"
cp -r ./ui $outputPath/static
printf "\nDone building: $app\n\n"

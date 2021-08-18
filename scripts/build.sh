#!/bin/sh
cfgFile="config.yml"
outputPath="build"
output="$outputPath/"
src="./..."

rm -rf ${outputPath}/*
printf "\nBuilding: $app\n"
time go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\npopulate config file"
cp  ./conf/$cfgFile $outputPath
printf "\nclient files"
cp -r ./static $outputPath/static
printf "\nDone building: $app\n\n"

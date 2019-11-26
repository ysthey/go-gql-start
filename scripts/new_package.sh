#!/bin/sh
#TODO: parameterize this script
grep -rl "github.com/ysthey/go-gql-start" ./* | xargs sed -i 's/github\.com\/ysthey\/go-gql-start/gitlab\.com\/ysthey\/newpackage/g'


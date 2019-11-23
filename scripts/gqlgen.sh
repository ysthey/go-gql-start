#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/gql/exec-gen.go \
    internal/gql/models/models-gen.go \
    internal/gql/resolvers/generated/resolvers-gen.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"
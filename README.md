# Golang GraphQl Dev Environment [Project Template]

Quick way to create / run Graphql server dev environment with Golang. 

## Golang packages:
        github.com/99designs/gqlgen
        github.com/gin-gonic/gin
        github.com/google/uuid
        github.com/jinzhu/gorm
        github.com/spf13/viper
        github.com/vektah/gqlparser
        gopkg.in/gormigrate.v1

## Usage:
1. Configuration files are ".env" and "./conf/*.yml"

2. To build:

        $ ./scripts/build.sh

3. To run:

        $ ./scripts/run.sh

4. To run tests:

        $ ./scripts/run_tests.sh

5. To update schema:

        $ vim ./internal/gql/schemas/schema.graphql
        $ ./scripts/gqlgen.sh


## Inspired by :
1. Creating an opinionated GraphQL server with Go

* https://dev.to/cmelgarejo/creating-an-opinionated-graphql-server-with-go-part-1-3g3l
* https://dev.to/cmelgarejo/creating-an-opinionated-graphql-server-with-go-part-2-46io
* https://dev.to/cmelgarejo/creating-an-opinionated-graphql-server-with-go-part-3-3aoi
	
2. UUID or GUID as Primary Keys? Be Careful!

* https://tomharrisonjr.com/uuid-or-guid-as-primary-keys-be-careful-7b2aa3dcb439

3. Golang configuration in 12 factor applications

* https://blog.container-solutions.com/golang-configuration-in-12-factor-applications

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

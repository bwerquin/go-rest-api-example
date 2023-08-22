# go-rest-api-example
Go rest API example with Gorillaz/Mux

You have to specify your config file in "/home/appuser/config.yml" (cf. helpers/app_config.go)

## Get started

`go run main.go .`

### Build Docker image (from goland)

Step 1 : build Docker image (with go build)

`docker build -t example -f Dockerfile .`

Step 2 : run it

`docker run -it -p 8080:8080 example`

### Build Light Docker image (from scratch)

Step 1 : build go app et Docker image (multistage)

`docker build -t example-scratch -f scratch.Dockerfile .`

Step 2 : run it

`docker run -it -p 8080:8080 example-scratch`




# go-rest-api-example
Go rest API example with Gorillaz/Mux

## Get started

`go run main.go .`

### Build Docker image (from goland)

Step 1 : build Docker image (with go build)
`docker build -t example -f Dockerfile .`
Step 2 : run it
`docker run -it -p 8080:8080 example`

### Build Light Docker image (from scratch)

Step 1 : build go app
`CGO_ENABLED=0 GOOS=linux go build -o main .`
Step 2 : build Docker image
`docker build -t example-scratch -f DockerfileWithoutGoland .`
Step 3 : run it
`docker run -it -p 8080:8080 example-scratch`




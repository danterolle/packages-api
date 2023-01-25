# ParrotOS Packages API

Displaying JSON files is easily done using a web server, such as NGINX. This API, once synchronized with the [packages-filter](https://github.com/danterolle/packages-filter) output, 
can instead be useful to show a single package, choosing the branch, the architecture and of course the package name.

## Run it locally

You can try this program by running it in its interpreter mode via `go run main.go` or build it with `go build -o packages-api`

Then just make a GET request to the following address, as in the example: http://localhost:8080/packages?branch=main&arch=amd64&package=nginx, or just use cURL or Postman.

## Options 

The options for branch and arch (architecture) are:

`branch = main, contrib, non-free`

`arch = amd64, arm64, armhf, i386`



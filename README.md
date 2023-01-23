# ParrotOS Packages API

Expose the contents of JSON files via an API using Go's net/http library.
The `json` folder with the relative branches and architectures of the packages must be present inside.
Once this is done, the program runs in interpreter mode via `go run main.go` or can be built with the appropriate command
`go build -o packages-api`

If started, just make a GET request to the following address, as in the example: http://localhost:8080/packages?branch=main&arch=amd64

Where the options for branch and arch (architecture) are:

`branch = main, contrib, non-free`

`arch = amd64, arm64, armhf, i386`



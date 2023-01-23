BINARY_NAME=packages-api

all: build

build:
	go build -o $(BINARY_NAME)

run:
	./$(BINARY_NAME)

deploy:
	docker build -t $(BINARY_NAME) .
	docker run -p 8080:8080 $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
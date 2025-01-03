all: build 

build:
	go build -o scrollbacktamer .

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -f scrollbacktamer coverage.out coverage.html

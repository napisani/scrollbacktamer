all: build 

build:
	go build -o scrollbacktamer .

test:
	go test -v ./...

clean:
	rm -f scrollbacktamer 

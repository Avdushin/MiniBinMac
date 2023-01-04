Default:
	go run src/cmd/main.go
start:
	./bin/minibinmac
build:
	go build -o bin/minibinmac src/cmd/main.go
install:
	./install.sh
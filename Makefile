build:
	go build -o contactBook

run:
	go run main.go

setup:
	go mod tidy

clean:
	rm -f contactBook

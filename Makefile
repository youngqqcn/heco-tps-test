all:
	go build -o build/main ./cmd/

run:
	./build/main
build:
	go build -o bin/tracks

run:
	@ templ generate
	go run .

test:
	go test -v ./...

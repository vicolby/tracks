build:
	go build -o bin/tracks
	@ templ generate

run:
	@ templ generate
	go run .

gen:
	@ templ generate

test:
	go test -v ./...

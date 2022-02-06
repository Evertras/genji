.PHONY: default
default:
	go run ./cmd/genji/main.go

.PHONY: build-all
build-all: ./build/pi/genji

./build/pi/genji: ./cmd/genji/main.go ./pkg/screens/disk/*.go
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./build/pi/genji ./cmd/genji/main.go


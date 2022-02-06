default:
	go run ./cmd/genji/main.go

./build/pi/genji:
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./build/pi/genji ./cmd/genji/main.go


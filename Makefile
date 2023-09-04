APP_ID = tech.innoai.example

dev:
	go run ./cmd/example

test:
	go test -v ./pkg/...

build.windows:
	GOOS=windows go build -o ./build/example.exe ./cmd/example

fmt:
	goimports -w -l .
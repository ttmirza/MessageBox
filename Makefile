build-go:
	go mod download && go mod verify
	go build -v -o ./bin/messagebox ./cmd/main.go

run-go:
	./bin/messagebox

go: build-go run-go

build-docker:
	docker build -t messagebox .

run-docker:
	docker run --name messagebox -d -p 3001:3001 messagebox

docker: build-docker run-docker

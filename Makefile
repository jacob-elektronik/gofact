all: build

.SILENT:
	
build: test
	cd example/ && go build -o gofact

test:
	go test ./... -covermode=count -coverprofile=testcover.out

show_testcover:
	go tool cover -html=testcover.out

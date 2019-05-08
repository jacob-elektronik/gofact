all: build

.SILENT:
	
build: test
	go build -o gofact

test:
	go test ./... -covermode=count -coverprofile=testcover.out

test_html:
	go tool cover -html=testcover.out

clean: 
	if [ -a ./gofact ]; then rm ./gofact; fi;

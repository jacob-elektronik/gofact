all: build

.SILENT:
	
build: test
	go build -o gofact

test:
	go test ./... -covermode=count -coverprofile=testcover.out

test_html:
	go tool cover -html=testcover.out

profile:
	./gofact -message edi_messages/huge_file2.edi  -cpuprofile cpu.prof

clean: 
	if [ -a ./gofact ]; then rm ./gofact; fi;

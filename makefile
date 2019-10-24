linux: main.go
	docker run --rm -v $(shell pwd):/go/src/spa-template -w /go/src/spa-template golang:1.13 /bin/bash -c "go get && go build -v"
	# We should also rebuild the main container.
	docker build -t="spa-template" .

run: main.go
	docker run -p 8080:8080 --rm spa-template
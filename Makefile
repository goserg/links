.PHONY:build
build:
	go build -tags=postgres -o bin/links ./cmd/.

.PHONY:build-inmemory
build-inmemory:
	go build -tags=inmemory -o bin/links ./cmd/.

.PHONY:run
run:
	./bin/links

.PHONY:test
test:
	go test -tags=postgres -v ./...

test:
	go test -v -cover -covermode=atomic ./internal/usecase/...

unittest:
	go test -short  ./internal/usecase/...

lint:
	golangci-lint run


https://gitlab.sefthost/group/subgroup/package?go-get=1
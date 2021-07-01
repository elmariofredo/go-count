.PHONY: test

test:
	go test ./... -v -covermode=count -coverprofile coverage.out
	#go tool cover -func coverage.out

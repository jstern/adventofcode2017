export INPUTS := $(CURDIR)/inputs

run:
	go run main.go $(e)

test:
	go test ./...

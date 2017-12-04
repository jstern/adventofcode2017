run:
	go run main.go $(e)

test:
	go test github.com/jstern/adventofcode2017/ex$(e)

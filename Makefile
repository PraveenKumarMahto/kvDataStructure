build: main.go main_test.go
	go build -o bin/main main.go

run: main.go main_test.go build
	go test
	./bin/main

clean: ./bin/main
	rm -r ./bin



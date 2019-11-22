build:
	go build -o bin/pps-client src/main.go

lint:
	golint src/main.go ;\
  mdl *.md -r ~MD013

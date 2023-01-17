.PHONY: default
default: build

all: clean get-deps build test

version := "0.0.1"

unit-tests:
	gocov test ./... -p 1 | gocov-xml > bin/coverage.xml

build:
	mkdir -p bin
	go build -o bin/service-sonar main.go

test: build
	go test ./... -coverprofile bin/coverage.out
	go tool cover -func=bin/coverage.out

clean:
	rm -rf ./bin

sonar: test
	sonar-scanner -Dsonar.projectVersion="$(version)"

start-sonar:
	docker run --name sonarqube -p 9000:9000 sonarqube
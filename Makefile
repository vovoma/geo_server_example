appname := geo

build:
	mkdir build/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(appname)

docker:
	docker build -t gusga/geolocation .

clean:
	\rm -rf build/


.PHONY: build docker clean

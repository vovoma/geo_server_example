appname := geo
image_name := gusga/geolocation
compose_file := docker-compose-mesos.yml

build:
	mkdir build/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(appname)

compose_up:
	docker-compose -f $(compose_file) up

docker:
	docker build -t $(image_name) .

push:
	docker push $(image_name)

clean:
	\rm -rf build/


.PHONY: build docker clean

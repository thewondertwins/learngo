
build:
	docker run -v `pwd`:/srv/gitbook -v `pwd`/_book:/srv/html bketelsen/gitbook

docker:
	docker build -t bketelsen/gitbook .

docker-push:
	docker push bketelsen/gitbook

# Change this value to the name of your project
PROJECT=client

APPUSER=$(shell echo `id -u`:`id -g`)

IMG_TAG=ngdev
CONTAINER=${IMG_TAG}-serve


all: new


image:
	docker run -it --name ngcontainer --network host  node:12-alpine sh -c "npm install npm -g && npm install  -g @angular/cli@^8"
	docker commit ngcontainer ngdev:latest
	docker container rm ngcontainer

new: npm-install
	docker run -u ${APPUSER} --network=host --rm -it -v "`pwd`:/app" -w /app ${IMG_TAG} sh -c "ng new ${PROJECT} --directory=."

npm-install: image
	docker run -u ${APPUSER} --network=host --rm -it -v "`pwd`:/app" -w /app ${IMG_TAG} sh -c "npm install"

serve:
	docker run -u ${APPUSER} --network=host --rm -d --name=${CONTAINER} -v "`pwd`:/app" -w /app ${IMG_TAG} sh -c "ng serve --host 0.0.0.0 --port 4200"

stop-serve:
	docker container stop ${CONTAINER}

logs:
	docker container logs -f ${CONTAINER}

sh:
	docker run -u ${APPUSER} --network=host --rm -it -v "`pwd`:/app" -w /app ${IMG_TAG} sh

prod:
	docker run -u ${APPUSER} --rm -it -v "`pwd`:/app" -w /app ${IMG_TAG} sh -c "ng build"

clean:
	git clean -fxd

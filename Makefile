docker/develop/build:
	docker build -t obada/fullcore:develop -f docker/Dockerfile .

docker/develop/publish:
	docker push obada/fullcore:develop

docker/develop: docker/develop/build docker/develop/publish

src/protogen:
	cd src && starport generate proto-go

src/run:
	cd src && go run ./src/cmd/fullcored/main.go

export GOPRIVATE=github.com/obada-foundation
src/vendor:
	cd src && go mod tidy && go mod vendor


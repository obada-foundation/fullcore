build:
	docker build -t obada/fullcore -f docker/Dockerfile .

protogen:
	cd src && starport generate proto-go

export GOPRIVATE=github.com/obada-foundation
vendor:
	cd src && go mod tidy && go mod vendor


.PHONY: all
all: build
FORCE: ;

.PHONY: build

build: build-aperta build-curat build-simplex build-metrik

build-aperta:
	cd aperta; go build -o bin/aperta main.go

build-curat:
	cd curat; go build -o bin/curat main.go

build-simplex:
	cd simplex; go build -o bin/simplex main.go

build-test-aperta:
	cd aperta; go test ./...

build-test-curat:
	cd curat; go test ./...

build-test-simplex:
	cd simplex; go test ./...


build-metrik:
	cd metrik; go build -o bin/metrik main.go

build-aperta-linux:
	cd aperta; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "netgo" -installsuffix netgo -o bin/aperta main.go

build-curat-linux:
	cd curat; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "netgo" -installsuffix netgo -o bin/curat main.go

build-simplex-linux:
	cd simplex; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "netgo" -installsuffix netgo -o bin/simplex main.go

build-metrik-linux:
	cd metrik; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "netgo" -installsuffix netgo -o bin/metrik main.go

build-linux: build-aperta-linux build-curat-linux build-simplex-linux build-metrik-linux

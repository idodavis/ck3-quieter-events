BINARY_NAME=vanilla_mod_merger
SRC_DIR=.

build:
	go build -o ${BINARY_NAME} main.go

clean:
	rm -rf ${BINARY_NAME}

run: build
	./${BINARY_NAME} merge-config.yaml
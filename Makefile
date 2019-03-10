# This is how we want to name the binary output
OUTPUT=user-api-passport

# These are the values we want to pass for Version and BuildTime
GITTAG=`git describe --tags`
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
# LDFLAGS=-ldflags "-X main.GitTag=${GITTAG} -X main.BuildTime=${BUILD_TIME}"
LDFLAGS=-ldflags "-X main.BuildTime=${BUILD_TIME}"

.PHONY:all clean release docker
all:clean release

clean:
	rm -f ${OUTPUT}

release:
	go build ${LDFLAGS} -o ${OUTPUT} main.go

docker:
	docker build --build-arg HOST=$HOST -t platform-${OUTPUT}:latest



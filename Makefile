

SRCS=main.go
BIN=static_httpd
TWISTEDDOCS=~/source/Twisted-19.10.0/docs/_build/html

static_httpd: ${SRCS}
	go build -o ${BIN} ${SRCS}

twisted:
	./static_httpd -hostport=:8080 -rootdir=${TWISTEDDOCS}

clean: ${BIN}
	rm ${BIN}

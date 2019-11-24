package static_httpd

import (
	"net/http"
)

func StaticServer(rootDir http.Dir) http.Handler {
	return http.FileServer(rootDir)
}

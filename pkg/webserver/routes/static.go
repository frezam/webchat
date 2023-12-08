package routes

import "net/http"

func StaticPublicDir(prefix string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir("./public/")))
}

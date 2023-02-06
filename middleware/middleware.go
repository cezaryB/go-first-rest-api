package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func JSONcontentMidleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		next(w, r, ps)
	}
}

func RouterWithLog(router http.Handler, port string) http.Handler {
	log.Printf("Server listenning on port %v ", port)
	return router
}
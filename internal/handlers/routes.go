package handlers

import "net/http"

func (h *Handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", http.FileServer(http.Dir("./ui/static"))))

	mux.HandleFunc("/register", h.CreateClient)

	return h.AllHandler(mux)
}

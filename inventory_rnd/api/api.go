package api

import (
	"net/http"
	"strings"
)

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /api
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")

	// ------------------ ADD THIS CODE ------------------
	if strings.HasPrefix(r.URL.Path, "/todolist") {
		TodoListHandler(w, r)
		return
	}
	// ---------------------------------------------------

	if strings.HasPrefix(r.URL.Path, "/custom_list") {
		CustomListHandler(w, r)
		return
	}

	// --------------------- ADD THIS CODE ---------------------
	if strings.HasPrefix(r.URL.Path, "/add_friend") {
		AddFriendHandler(w, r)
		return
	}
	// ---------------------------------------------------------

}

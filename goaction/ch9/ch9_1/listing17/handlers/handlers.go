package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes 路由
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON 响应json文档
func SendJSON(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

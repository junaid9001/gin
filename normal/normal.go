package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "wrong request", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"method": "post",
		})
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Bad req", http.StatusBadRequest)
		}
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "something wrong", http.StatusBadRequest)
		}
		fmt.Fprintln(w, "username is "+user.Name)

		w.Header().Set("content-type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status":   "success",
			"username": user.Name,
		})
	})

	http.ListenAndServe(":8080", nil)

}

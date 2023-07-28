package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	db "github.com/12ain13owz/project-borntodev/database"
)

func HandlerUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := db.GetUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write(result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func HandlerUserByID(w http.ResponseWriter, r *http.Request) {
	urlPathSegmentx := strings.Split(r.URL.Path, "user/")
	if len(urlPathSegmentx[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(urlPathSegmentx[len(urlPathSegmentx)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := db.GetUserByID(userID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write(result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

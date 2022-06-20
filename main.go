package main

import (
	"encoding/json"
	"net/http"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// 4
func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	// 들어오는 요청의 Response Header에 Content-Type을 추가
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		// 전달 받은 http.Handler를 호출한다.
		next.ServeHTTP(rw, r)
	})
}

func main() {
	// 1
	mux := http.NewServeMux()

	// 2
	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // 조회
			json.NewEncoder(wr).Encode(users)
		case http.MethodPost: // 등록
			var user User
			json.NewDecoder(r.Body).Decode(&user)

			users[user.Email] = &user

			json.NewEncoder(wr).Encode(user)
		}
	})

	// 3
	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
	http.ListenAndServe(":8080", mux)
}

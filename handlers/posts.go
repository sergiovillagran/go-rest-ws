package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"github.com/sergiovillagran/rest-ws/models"
	"github.com/sergiovillagran/rest-ws/repository"
	"github.com/sergiovillagran/rest-ws/server"
)

type InsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type InsertPostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

func InserPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			user, err := repository.GetUser(r.Context(), claims.UserId)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var postRequest = InsertPostRequest{}

			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			id, err := ksuid.NewRandom()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			post := models.Post{
				Id:          id.String(),
				PostContent: postRequest.PostContent,
				UserId:      user.Id,
			}

			err = repository.InsertPost(r.Context(), &post)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(InsertPostResponse{
				Id:          post.Id,
				PostContent: post.PostContent,
			})
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		post, err := repository.GetPostById(r.Context(), params["id"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

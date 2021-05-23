package main

import (
	"net/http"
	"os"
	"regexp"

	"github.com/0x4c6565/secret.lee.io/pkg/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()

	storage := storage.NewRedisStorage(
		viper.GetString("storage_redis_host"),
		viper.GetInt("storage_redis_port"),
		viper.GetString("storage_redis_password"),
		viper.GetInt("storage_redis_db"),
	)
	handler := NewHandler(storage)

	r := mux.NewRouter()
	r.HandleFunc(`/secret/{uuid:[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}}`, handler.HandleGet).Methods("GET")
	r.HandleFunc("/secret", handler.HandlePost).Methods("POST")
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok, _ := regexp.MatchString(`^\/[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}$`, r.URL.Path); ok {
			http.ServeFile(w, r, "./static/index.html")
			return
		}

		fs.ServeHTTP(w, r)
	})).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8080", loggedRouter)
}

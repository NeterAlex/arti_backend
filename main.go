package main

import (
	"C"
	"arti_backend/models"
	"arti_backend/pkg/setting"
	"arti_backend/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	models.InitDB()
	err := server.ListenAndServe()
	if err != nil {
		println(err)
	}
}

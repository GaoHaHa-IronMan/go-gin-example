package main

import (
	"fmt"
	"github.com/GaoHaHa-IronMan/go-gin-example/pkg/setting"
	"github.com/GaoHaHa-IronMan/go-gin-example/routers"
	"net/http"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
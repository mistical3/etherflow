package main

import (
	"app.go/internal/user"
	"app.go/pkg/db"
	"app.go/pkg/logs"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logs.GetLogger()
	logger.Info("create router %s")
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	handler := user.NewHandler()
	logger.Info("register user handler")
	handler.Register(router)
	db.Init()
	start(router)
}

func start(router *httprouter.Router) {
	logger := logs.GetLogger()
	logger.Info("create router %s")
	list, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	serv := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info("localhost:1234")
	logger.Fatal(serv.Serve(list))

}

package main

import (
	// "encoding/json"
	"net/http"
	"fmt"
	"github.com/yuanyu90221/go-start/todo_server/controller"
	"log"
	"./model"
	_ "github.com/lib/pq"
)

func main() {
	mux := controller.Register()
	fmt.Println("test");
	db := model.Connect()
	defer db.Close()
	fmt.Println("Servering...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

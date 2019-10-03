package main

import (
	"context"
	c "controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	clientop := options.Client().ApplyURI(viper.GetString("db.connect"))
	client, err := mongo.Connect(context.TODO(), clientop)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/bootstrap/").Handler(http.StripPrefix("/bootstrap/", http.FileServer(http.Dir("bootstrap/"))))
	r.HandleFunc("/", c.Index).Methods("GET")
	r.HandleFunc("/insert", c.Insert).Methods("POST")
	r.HandleFunc("/new", c.New)
	r.HandleFunc("/edit/{id}", c.Edit).Methods("GET")
	r.HandleFunc("/update", c.Update).Methods("POST")
	r.HandleFunc("/delete/{id}", c.Del).Methods("GET")
	r.HandleFunc("/view/{id}", c.View).Methods("GET")
	r.HandleFunc("/rc", c.Rc).Methods("POST")
	fmt.Println("ready")
	err = http.ListenAndServeTLS(viper.GetString("port_https"), "server.crt", "server.key", r)
	if err != nil {
		log.Fatal(err)
	}
}

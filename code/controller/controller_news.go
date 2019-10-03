package controller

import (
	"context"
	"log"
	m "model"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var view = template.Must(template.ParseGlob("view/*"))

func config() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()
}

func conn() (*mongo.Database, error) {
	config()
	clientop := options.Client()
	clientop.ApplyURI(viper.GetString("db.connect"))
	client, err := mongo.NewClient(clientop)
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	return client.Database(viper.GetString("db.dbname")), nil
}

//Index a
func Index(w http.ResponseWriter, r *http.Request) {
	config()
	db, err := conn()
	collection := db.Collection(viper.GetString("db.collection"))
	querry, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	var result []*m.News
	for querry.Next(context.TODO()) {
		var news m.News
		err := querry.Decode(&news)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &news)
	}
	if err := querry.Err(); err != nil {
		log.Fatal(err)
	}
	querry.Close(context.TODO())
	view.ExecuteTemplate(w, "index", result)
}

//Insert a
func Insert(w http.ResponseWriter, r *http.Request) {
	config()
	db, err := conn()
	if err != nil {
		log.Fatal(err)
	}
	collection := db.Collection(viper.GetString("db.collection"))
	if r.Method == "POST" {
		ireact, err := strconv.ParseInt(r.FormValue("react"), 10, 0)
		icount, err := strconv.ParseInt(r.FormValue("count"), 10, 0)
		news := m.News{Title: r.FormValue("title"), Body: r.FormValue("body"), Author: r.FormValue("author"),
			React: ireact, Count: icount, Date: time.Now().String()[0:10]}
		_, err = collection.InsertOne(context.TODO(), news)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.Redirect(w, r, "/", 301)
}

//New a
func New(w http.ResponseWriter, r *http.Request) {
	view.ExecuteTemplate(w, "new", nil)
}

//Edit a
func Edit(w http.ResponseWriter, r *http.Request) {
	config()
	db, err := conn()
	if err != nil {
		log.Fatal(err)
	}
	collection := db.Collection(viper.GetString("db.collection"))
	var news m.News
	url := mux.Vars(r)
	rid := url["id"]
	rid = rid[10:34]
	brid, _ := primitive.ObjectIDFromHex(rid)
	err = collection.FindOne(context.TODO(), bson.M{"_id": brid}).Decode(&news)
	if err != nil {
		log.Fatal(err)
	}
	view.ExecuteTemplate(w, "edit", news)
}

//Update a
func Update(w http.ResponseWriter, r *http.Request) {
	config()
	db, err := conn()
	collection := db.Collection(viper.GetString("db.collection"))
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "POST" {
		rid := r.FormValue("id")
		rid = rid[10:34]
		brid, _ := primitive.ObjectIDFromHex(rid)
		ireact, err := strconv.ParseInt(r.FormValue("react"), 10, 0)
		icount, err := strconv.ParseInt(r.FormValue("count"), 10, 0)
		unews := bson.M{
			"$set": bson.M{
				"title":  r.FormValue("title"),
				"body":   r.FormValue("body"),
				"author": r.FormValue("author"),
				"react":  ireact,
				"count":  icount,
			},
		}
		_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": brid}, unews)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.Redirect(w, r, "/", 301)
}

//Del a
func Del(w http.ResponseWriter, r *http.Request) {
	config()
	db, err := conn()
	collection := db.Collection(viper.GetString("db.collection"))
	if err != nil {
		log.Fatal(err)
	}
	url := mux.Vars(r)
	rid := url["id"]
	rid = rid[10:34]
	brid, _ := primitive.ObjectIDFromHex(rid)
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": brid})
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", 301)
}

//View a
func View(w http.ResponseWriter, r *http.Request) {
	config()
	db, err := conn()
	collection := db.Collection(viper.GetString("db.collection"))
	if err != nil {
		log.Fatal(err)
	}
	url := mux.Vars(r)
	rid := url["id"]
	rid = rid[10:34]
	brid, _ := primitive.ObjectIDFromHex(rid)
	var result m.News
	err = collection.FindOne(context.TODO(), bson.M{"_id": brid}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	view.ExecuteTemplate(w, "view", result)
}

//Rc a
func Rc(w http.ResponseWriter, r *http.Request) {
	urid := "/view/"
	config()
	db, err := conn()
	collection := db.Collection(viper.GetString("db.collection"))
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "POST" {
		rid := r.FormValue("id")
		urid = urid + r.FormValue("id")
		rid = rid[10:34]
		brid, _ := primitive.ObjectIDFromHex(rid)
		var ireact, icount int64
		if r.FormValue("positive") != "" {
			ireact = 1
			icount = 1
		} else if r.FormValue("neutral") != "" {
			ireact = 0
			icount = 1
		} else {
			ireact = -1
			icount = 1
		}
		unews := bson.M{
			"$inc": bson.M{
				"react": ireact,
				"count": icount,
			},
		}
		_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": brid}, unews)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.Redirect(w, r, urid, 301)
}

package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// MOngoCN es el objeto de conexion a la BD
var MongoCN = ConectarBD()

// add url by .env
var clientOptions = options.Client().ApplyURI("mongodb+srv://rodriG0:MkzEeEBMNwORFZm5@go.3jazpeb.mongodb.net/twittor")

/* ConectarBD es la funcion que me permite conectar a la BD*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa a la DB")
	return client
}

// ChequeoConnection es el Ping a la BD
func ChequeoConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}

package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb+srv://jmc:root@cluster0.wnnlh.mongodb.net/test"

/* MongoCN es el objeto de conexión */
var MongoCN = conectarBD()

var clienteOptions = options.Client().ApplyURI(uri)

/* ConectarBD es la función que me permite conectar la BD */
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		return logError(err, client)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return logError(err, client)
	}
	log.Println("Conexión Exitosa con la DB")
	return client
}

/* ChequeoConexion es el Ping a la BD*/
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

func logError(err error, client *mongo.Client) *mongo.Client {
	log.Fatal(err.Error())
	return client
}

package bd

import (
	"context"
	"time"

	"github.com/ingjeffer/twittor2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodo(ID string, page int64, search string, searchType string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	limit := int64(20)
	findOptions := options.Find()
	findOptions.SetSkip(limit * (page - 1))
	findOptions.SetLimit(limit)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false
		encontrado, err = ConsultoRelacion(r)
		if searchType == "new" && !encontrado {
			incluir = true
		}

		if searchType == "follow" && encontrado {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}

	return results, true

}

package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ingjeffer/twittor2/bd"
	"github.com/ingjeffer/twittor2/models"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("jmc_4321*")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}

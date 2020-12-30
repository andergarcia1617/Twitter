package bd

import (
	"context"
	"time"

	"github.com/andergarcia1617/Twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro para insertar nuestro registro*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitterander")
	col := db.Collection("users")
	u.Password, _ = encriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

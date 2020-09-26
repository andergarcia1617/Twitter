package bd

import (
	"context"
	"time"
	"github.com/andergarcia1617/Twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func InsertoRegistro(u models, Usuario)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Twitter")
	col := db.Collection("usuarios")
	u.Password, _= EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err !=nil {
		return "", false, err
	}
	ObjID, _ := result.InsertID.(primitive.ObjectID)
	return objID.String(), true, nil

}

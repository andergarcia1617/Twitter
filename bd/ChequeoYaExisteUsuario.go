package bd

import(
	"context"
	"time"
	"github.com/andergarcia1617/Twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)
/*ChequeoYaexisteUsuario recibe u email de parametro y chequea si existe*/
func ChequeoYaexisteUsuario(email string)(model.Usuario, bool, string) 
ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()
condicion := bson.M{"email":email}
var resultado models.Usuario
err := col.FindOne(ctx, condicion).Decode(&resultado)
ID :=resultado.ID.Hex()
if err{
	return resultado, false, ID
	}
	return resultado, true , ID
}


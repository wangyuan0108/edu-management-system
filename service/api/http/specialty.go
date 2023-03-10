package http

import (
	"context"
	"edu-management-system/db"
	"edu-management-system/schema"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func GetSpecialtyList(filter bson.D) ([]schema.Specialty, error) {

	var specialtyList []schema.Specialty
	cursor, err := db.Mongo().
		Database(os.Getenv("MONGODB_DB_EDU")).
		Collection(schema.Specialty{}.Collection()).
		Find(context.Background(), filter)
	if err = cursor.All(context.TODO(), &specialtyList); err != nil {
		return nil, err
	}

	return specialtyList, nil
}

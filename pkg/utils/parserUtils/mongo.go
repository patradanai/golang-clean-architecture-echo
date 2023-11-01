package parserUtils

import "go.mongodb.org/mongo-driver/bson"

func MappingMongo(query map[string]interface{}) bson.D {
	result := bson.D{}

	for k, v := range query {
		result = append(result, bson.D{{Key: k, Value: v}}...)
	}

	return result
}

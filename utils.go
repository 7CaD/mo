package mo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func kv(k string, v interface{}) bson.E {
	return bson.E{Key: k, Value: v}
}

func kvNew(pairs ...bson.E) bson.D {
	n := bson.D{}
	for _, pair := range pairs {
		n = append(n, pair)
	}
	return n
}

func processTuples(tuples []interface{}) ([]bsonTuple, error) {
	var results []bsonTuple
	for i := 0; i < len(tuples); i += 2 {
		if i+1 < len(tuples) {
			key, ok1 := tuples[i].(string)
			value := tuples[i+1]
			if ok1 {
				results = append(results, bsonTuple{
					Key:   key,
					Value: value,
				})
			} else {
				return results, fmt.Errorf("processTuples error: expected a string as the key")
			}
		} else {
			return results, fmt.Errorf(
				"processTuples error: unmatched key-value pair at idx %d, value is %v",
				i,
				tuples[i],
			)
		}
	}
	return results, nil
}

type bsonTuple struct {
	Key   string
	Value interface{}
}

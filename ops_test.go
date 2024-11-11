package mo

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func TestInitQueryPredicate(t *testing.T) {
	tests := []struct {
		name   string
		invoke func() *QueryPredicate
		want   bson.D
	}{
		{
			name: "test eq",
			invoke: func() *QueryPredicate {
				return Eq("test", 1)
			},
			want: bson.D{
				{Key: "$eq", Value: bson.D{
					{
						Key:   "test",
						Value: 1,
					},
				}},
			},
		},
		{
			name: "test ne",
			invoke: func() *QueryPredicate {
				return Ne("foo", "bar")
			},
			want: bson.D{
				{Key: "$ne", Value: bson.D{
					{
						Key:   "foo",
						Value: "bar",
					},
				}},
			},
		},
		{
			name: "test gt",
			invoke: func() *QueryPredicate {
				return Gt("foo", 1)
			},
			want: bson.D{
				{Key: "$gt", Value: bson.D{
					{
						Key:   "foo",
						Value: 1,
					},
				}},
			},
		},
		{
			name: "test gte",
			invoke: func() *QueryPredicate {
				return Gte("foo", 1)
			},
			want: bson.D{
				{Key: "$gte", Value: bson.D{
					{
						Key:   "foo",
						Value: 1,
					},
				}},
			},
		},
		{
			name: "test lt",
			invoke: func() *QueryPredicate {
				return Lt("foo", 1)
			},
			want: bson.D{
				{Key: "$lt", Value: bson.D{
					{
						Key:   "foo",
						Value: 1,
					},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.invoke(); !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("%s: got = %v, want %v", tt.name, got.Value, tt.want)
			}
		})
	}
}

func TestComplex(t *testing.T) {
	t.Run("test complex", func(t *testing.T) {
		predicate := And(
			Eq("test", 1),
			Ne("foo", "bar"),
			Gt("foo", 1),
			Gte("foo", 1),
			Lt("foo", 1),
		)
		expected := bson.D{{"$and", bson.A{
			bson.D{{"$eq", bson.D{{"test", 1}}}},
			bson.D{{"$ne", bson.D{{"foo", "bar"}}}},
			bson.D{{"$gt", bson.D{{"foo", 1}}}},
			bson.D{{"$gte", bson.D{{"foo", 1}}}},
			bson.D{{"$lt", bson.D{{"foo", 1}}}},
		}}}
		if !reflect.DeepEqual(predicate.Value, expected) {
			t.Errorf("got = %v, want = %v", predicate.Value, expected)
		}
	})
}

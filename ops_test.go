package mo

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func TestInitQueryPredicate(t *testing.T) {
	logicalQueryChildren := []*QueryPredicate{
		Eq("test", 1),
		Ne("foo", "bar"),
		Gt("foo", 1),
		Gte("foo", 1),
		Lt("foo", 1),
	}

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
				{Key: "test", Value: bson.D{
					{
						Key:   "$eq",
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
				{Key: "foo", Value: bson.D{
					{
						Key:   "$ne",
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
		{
			name: "test lte",
			invoke: func() *QueryPredicate {
				return Lte("foo", 1)
			},
			want: bson.D{
				{Key: "$lte", Value: bson.D{
					{
						Key:   "foo",
						Value: 1,
					},
				}},
			},
		},
		{
			name: "test in",
			invoke: func() *QueryPredicate {
				return In("foo", []int{1, 2, 3})
			},
			want: bson.D{
				{Key: "$in", Value: bson.D{
					{
						Key:   "foo",
						Value: []int{1, 2, 3},
					},
				}},
			},
		},
		{
			name: "test nin",
			invoke: func() *QueryPredicate {
				return Nin("foo", []int{1, 2, 3})
			},
			want: bson.D{
				{Key: "$nin", Value: bson.D{
					{
						Key:   "foo",
						Value: []int{1, 2, 3},
					},
				}},
			},
		},
		{
			name: "test and",
			invoke: func() *QueryPredicate {
				return And(logicalQueryChildren...)
			},
			want: bson.D{
				{Key: "$and", Value: bson.A{
					bson.D{{"$eq", bson.D{{"test", 1}}}},
					bson.D{{"$ne", bson.D{{"foo", "bar"}}}},
					bson.D{{"$gt", bson.D{{"foo", 1}}}},
					bson.D{{"$gte", bson.D{{"foo", 1}}}},
					bson.D{{"$lt", bson.D{{"foo", 1}}}},
				}},
			},
		},
		{
			name: "test or",
			invoke: func() *QueryPredicate {
				return Or(logicalQueryChildren...)
			},
			want: bson.D{
				{Key: "$or", Value: bson.A{
					bson.D{{"$eq", bson.D{{"test", 1}}}},
					bson.D{{"$ne", bson.D{{"foo", "bar"}}}},
					bson.D{{"$gt", bson.D{{"foo", 1}}}},
					bson.D{{"$gte", bson.D{{"foo", 1}}}},
					bson.D{{"$lt", bson.D{{"foo", 1}}}},
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
			bson.D{{"test", bson.D{{"$eq", 1}}}},
			bson.D{{"foo", bson.D{{"$ne", "bar"}}}},
			bson.D{{"foo", bson.D{{"$gt", 1}}}},
			bson.D{{"foo", bson.D{{"$gte", 1}}}},
			bson.D{{"foo", bson.D{{"$lt", 1}}}},
		}}}
		if !reflect.DeepEqual(predicate.Value, expected) {
			t.Errorf("got = %v, want = %v", predicate.Value, expected)
		}
	})
}

# mo

Pipeline query builder designed for [mongo-go-driver](https://github.com/mongodb/mongo-go-driver). It provides a fluent and intuitive interface to construct MongoDB aggregation pipelines in Go.

## Usage

```sh
go get github.com/7CaD/mo
```

Example: 
```go
package main

import (
    "github.com/7CaD/mo"
    "go.mongodb.org/mongo-driver/bson"

    "context"
)

func main() {
    pipeline := mo.NewPipeline().
        Match(bson.M{"status": "A"}).
        Group(bson.M{"_id": "$cust_id", "total": bson.M{"$sum": "$amount"}}).
        Sort(bson.M{"total": -1}).
        Build()

    // Use 'pipeline' with your mongo-go-driver collection
    collection := client.Database("your_database").Collection("your_collection")
    cursor, err := collection.Aggregate(context.TODO(), pipeline)

    // ...
}
```

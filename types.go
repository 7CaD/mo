package mo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IPipelineBuilder interface {
	Set(kvs ...interface{}) PipelineBuilder
	Sort(kvs ...interface{}) PipelineBuilder
	Limit(limit int) PipelineBuilder
	Skip(limit int) PipelineBuilder
	Match(kvs ...interface{}) PipelineBuilder
	Group(kvs ...interface{}) PipelineBuilder
	Unwind(path string) PipelineBuilder
	AddFields(kvs ...interface{}) PipelineBuilder
	Count(countKey string) PipelineBuilder
	ReplaceRoot(newRootKey string) PipelineBuilder
	Value() mongo.Pipeline
}

type IsQueryPredicate interface {
	IsQp()
}

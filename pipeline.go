package mo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PipelineBuilder underlying struct. do not instantiate this manually. Use NewPipeline().
type PipelineBuilder struct {
	stages mongo.Pipeline
	errors []error
}

func (p *PipelineBuilder) Set(kvs ...interface{}) *PipelineBuilder {
	p.genericAddToStage("$set", kvs)
	return p
}

func (p *PipelineBuilder) Sort(kvs ...interface{}) *PipelineBuilder {
	p.genericAddToStage("$sort", kvs)
	return p
}

func (p *PipelineBuilder) Limit(limit int) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$limit", Value: limit}})
	return p
}

func (p *PipelineBuilder) Skip(limit int) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$skip", Value: limit}})
	return p
}

func (p *PipelineBuilder) Match(query *QueryPredicate) *PipelineBuilder {
	p.stages = append(p.stages, query.Value)
	return p
}

func (p *PipelineBuilder) Group(kvs ...interface{}) *PipelineBuilder {
	p.genericAddToStage("$group", kvs)
	return p
}

func (p *PipelineBuilder) Unwind(path string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$unwind", Value: path}})
	return p
}

type PipelineStage struct {
	Name  string
	Value bson.D
}

func AddFields() *PipelineStage {
	return &PipelineStage{
		Name:  "$addFields",
		Value: bson.D{},
	}
}

func (p *PipelineBuilder) Count(countKey string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$count", Value: countKey}})
	return p
}

func (p *PipelineBuilder) ReplaceRoot(newRootKey string) *PipelineBuilder {
	p.stages = append(p.stages, bson.D{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: newRootKey}}}})
	return p
}

func (p *PipelineBuilder) genericAddToStage(rootKey string, kvs []interface{}) {
	if len(kvs) == 0 {
		return
	}
	if len(kvs) == 1 {
		p.stages = append(p.stages, bson.D{{Key: rootKey, Value: kvs[0]}})
		return
	}
	tuples, err := processTuples(kvs)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("syntax error in pipeline.%s: %v", rootKey, err))
	}
	match := bson.D{}
	for _, tuple := range tuples {
		match = append(match, bson.E{Key: tuple.Key, Value: tuple.Value})
	}
	p.stages = append(p.stages, bson.D{{Key: rootKey, Value: match}})
	return
}

func (p *PipelineBuilder) Value() mongo.Pipeline {
	return p.stages
}

func NewPipeline() *PipelineBuilder {
	return &PipelineBuilder{
		stages: mongo.Pipeline{},
	}
}

func (p *PipelineBuilder) Project(kvs ...interface{}) *PipelineBuilder {
	p.genericAddToStage("$project", kvs)
	return p
}

type OperatorBuildable interface {
	Value() bson.D
}

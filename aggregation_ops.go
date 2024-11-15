package mo

import "go.mongodb.org/mongo-driver/bson"

type AggExpressionOperator struct {
	Value bson.D
}

func exprToVal(expr interface{}) interface{} {
	switch expr.(type) {
	case *AggExpressionOperator:
		return expr.(*AggExpressionOperator).Value
	default:
		return expr
	}
}

func (p *PipelineStage) Abs(expr interface{}) *AggExpressionOperator {
	return &AggExpressionOperator{
		Value: obj(
			kv("$abs", exprToVal(expr)),
		),
	}
}

func Acos(expr interface{}) *AggExpressionOperator {
	return &AggExpressionOperator{
		Value: obj(
			kv("$acos", exprToVal(expr)),
		),
	}
}

func Acosh(expr *interface{}) *AggExpressionOperator {
	return &AggExpressionOperator{
		Value: obj(
			kv("$acosh", exprToVal(expr)),
		),
	}
}

func Add(exprs ...*interface{}) *AggExpressionOperator {
	vals := make([]interface{}, len(exprs))
	for i := range exprs {
		vals[i] = exprToVal(exprs[i])
	}
	return &AggExpressionOperator{
		Value: obj(
			kv("$add", vals),
		),
	}
}

func AddToSet(expr *interface{}) *AggExpressionOperator {
	return &AggExpressionOperator{
		Value: obj(
			kv("$addToSet", exprToVal(expr)),
		),
	}
}

func AllElementsTrue(expr *interface{}) *AggExpressionOperator {
	return &AggExpressionOperator{
		Value: obj(
			kv("$allElementsTrue", exprToVal(expr)),
		),
	}
}

func AnyElementTrue(expr *interface{}) *AggExpressionOperator {
	return &AggExpressionOperator
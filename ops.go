package mo

import "go.mongodb.org/mongo-driver/bson"

type QueryPredicate struct {
	Value bson.D ``
}

// region Comparison Queries

func (p *QueryPredicate) Eq(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$eq", val))))
	return p
}

func Eq(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Eq(path, val)
}

func (p *QueryPredicate) Ne(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$ne", val))))
	return p
}

func Ne(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Ne(path, val)
}

func (p *QueryPredicate) Gt(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$gt", val))))
	return p
}

func Gt(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Gt(path, val)
}

func (p *QueryPredicate) Gte(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$gte", val))))
	return p
}

func Gte(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Gte(path, val)
}

func (p *QueryPredicate) Lt(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$lt", val))))
	return p
}

func Lt(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Lt(path, val)
}

func (p *QueryPredicate) Lte(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$lte", val))))
	return p
}

func Lte(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Lte(path, val)
}

func (p *QueryPredicate) In(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$in", val))))
	return p
}

func In(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).In(path, val)
}

func (p *QueryPredicate) Nin(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$nin", val))))
	return p
}

func Nin(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Nin(path, val)
}

//endregion

// region Logical Queries

func (p *QueryPredicate) And(val ...*QueryPredicate) *QueryPredicate {
	p.Value = append(p.Value, kv("$and", list(val)))
	return p
}

func And(val ...*QueryPredicate) *QueryPredicate {
	return (&QueryPredicate{}).And(val...)
}

func (p *QueryPredicate) Or(val ...*QueryPredicate) *QueryPredicate {
	p.Value = append(p.Value, kv("$or", list(val)))
	return p
}

func Or(val ...*QueryPredicate) *QueryPredicate {
	return (&QueryPredicate{}).Or(val...)
}

func (p *QueryPredicate) Not(val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$not", val))
	return p
}

func Nor(val ...*QueryPredicate) *QueryPredicate {
	return (&QueryPredicate{}).Nor(val...)
}

func (p *QueryPredicate) Nor(val ...*QueryPredicate) *QueryPredicate {
	p.Value = append(p.Value, kv("$nor", val))
	return p
}

//endregion

// region Element Queries
const (
	BSONTypeDouble = "double"
	BSONTypeString = "string"
	BSONTypeObject = "object"
	BSONTypeArray  = "array"
	BSONTypeBinary = "binData"
	// BSONTypeUndefined deprecated
	BSONTypeUndefined = "undefined"
	BSONTypeObjectID  = "objectId"
	BSONTypeBool      = "bool"
	BSONTypeDate      = "date"
	BSONTypeNull      = "null"
	BSONTypeRegex     = "regex"
	// BSONTypeDBPointer deprecated
	BSONTypeDBPointer  = "dbPointer"
	BSONTypeJavaScript = "javascript"
	// BSONTypeSymbol deprecated
	BSONTypeSymbol     = "symbol"
	BSONTypeInt32      = "int"
	BSONTypeTimestamp  = "timestamp"
	BSONTypeInt64Long  = "long"
	BSONTypeDecimal128 = "decimal"
	BSONTypeMinKey     = "minKey"
	BSONTypeMaxKey     = "maxKey"
	BSONTypeNumber     = "number"
)

func (p *QueryPredicate) Exists(path string, exists bool) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$exists", exists))))
	return p
}

func Exists(path string, exists bool) *QueryPredicate {
	return (&QueryPredicate{}).Exists(path, exists)
}

type BSONType int

func (p *QueryPredicate) Type(path string, t ...BSONType) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$type", t))))
	return p
}

func Type(path string, t ...BSONType) *QueryPredicate {
	return (&QueryPredicate{}).Type(path, t...)
}

// endregion

// region Evaluation Queries

func (p *QueryPredicate) Expr(expr bson.D) *QueryPredicate {
	p.Value = append(p.Value, kv("$expr", expr))
	return p
}

func Expr(expr bson.D) *QueryPredicate {
	return (&QueryPredicate{}).Expr(expr)
}

//endregion

func All(path string, val ...any) *QueryPredicate {
	return (&QueryPredicate{}).All(path, val...)
}

func (p *QueryPredicate) All(path string, val ...any) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$all", val))))
	return p
}

func Size(path string, size int) *QueryPredicate {
	return (&QueryPredicate{}).Size(path, size)
}

func (p *QueryPredicate) Size(path string, size int) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$size", size))))
	return p
}

func ElemMatch(path string, query *QueryPredicate) *QueryPredicate {
	return (&QueryPredicate{}).ElemMatch(path, query)
}

func (p *QueryPredicate) ElemMatch(path string, query *QueryPredicate) *QueryPredicate {
	p.Value = append(p.Value, kv(path, obj(kv("$elemMatch", query))))
	return p
}

package mo

import "go.mongodb.org/mongo-driver/bson"

type QueryPredicate struct {
	Value bson.D
}

func (p *QueryPredicate) Eq(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$eq", kvNew(kv(path, val))))
	return p
}

func Eq(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Eq(path, val)
}

func (p *QueryPredicate) Ne(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$ne", kvNew(kv(path, val))))
	return p
}

func Ne(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Ne(path, val)
}

func (p *QueryPredicate) Gt(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$gt", kvNew(kv(path, val))))
	return p
}

func Gt(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Gt(path, val)
}

func (p *QueryPredicate) Gte(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$gte", kvNew(kv(path, val))))
	return p
}

func Gte(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Gte(path, val)
}

func (p *QueryPredicate) Lt(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$lt", kvNew(kv(path, val))))
	return p
}

func Lt(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Lt(path, val)
}

func (p *QueryPredicate) Lte(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$lte", kvNew(kv(path, val))))
	return p
}

func Lte(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Lte(path, val)
}

func (p *QueryPredicate) In(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$in", kvNew(kv(path, val))))
	return p
}

func In(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).In(path, val)
}

func (p *QueryPredicate) Nin(path string, val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$nin", kvNew(kv(path, val))))
	return p
}

func Nin(path string, val any) *QueryPredicate {
	return (&QueryPredicate{}).Nin(path, val)
}

func (p *QueryPredicate) And(val ...*QueryPredicate) *QueryPredicate {
	var tVal bson.A
	for _, predicate := range val {
		tVal = append(tVal, predicate.Value)
	}
	p.Value = append(p.Value, kv("$and", tVal))
	return p
}

func And(val ...*QueryPredicate) *QueryPredicate {
	return (&QueryPredicate{}).And(val...)
}

func (p *QueryPredicate) Or(val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$or", val))
	return p
}

func Or(val any) *QueryPredicate {
	return (&QueryPredicate{}).Or(val)
}

func (p *QueryPredicate) Not(val any) *QueryPredicate {
	p.Value = append(p.Value, kv("$not", val))
	return p
}

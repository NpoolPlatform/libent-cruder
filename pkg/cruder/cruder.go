package cruder

type Any interface{}

const (
	EQ   = "eq"
	NEQ  = "neq"
	GT   = "gt"
	GTE  = "gte"
	LT   = "lt"
	LTE  = "lte"
	IN   = "in"
	LIKE = "like"
)

type Cond struct {
	Op  string
	Val Any
}

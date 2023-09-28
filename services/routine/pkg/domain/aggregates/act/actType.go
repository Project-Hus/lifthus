package act

const (
	WEIGHT = "weight"
	TIME   = "time"
	SIMPLE = "simple"
)

var WeightType = ActType{WEIGHT}
var TimeType = ActType{TIME}
var SimpleType = ActType{SIMPLE}

type ActType struct {
	code string
}

func (pt ActType) Type() string {
	return pt.code
}

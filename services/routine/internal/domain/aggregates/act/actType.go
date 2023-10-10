package act

const (
	WEIGHT = "weight"
	TIME   = "time"
	SIMPLE = "simple"
)

func MapActType(code string) (*ActType, error) {
	switch code {
	case WEIGHT:
		return &WeightType, nil
	case TIME:
		return &TimeType, nil
	case SIMPLE:
		return &SimpleType, nil
	default:
		return nil, ErrInvalidActType
	}
}

var WeightType = ActType{WEIGHT}
var TimeType = ActType{TIME}
var SimpleType = ActType{SIMPLE}

type ActType struct {
	code string
}

func (pt ActType) Type() string {
	return pt.code
}

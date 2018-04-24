package exchange

type Exchange struct {
    name string
    valueAdded float64
}

func New(name string, valueAdded float64) Exchange {
    e := Exchange {name, valueAdded}
    return e
}

func (e *Exchange) SetName(name string) string {
    e.name = name
}

func (e Exchange) Name() string {
    return e.name
}

func (e Exchange) ValueAdded() float64 {
    return e.valueAdded
}

func (e *Exchange) SetValueAdded(valueAdded float64) {
    e.valueAdded = valueAdded
}

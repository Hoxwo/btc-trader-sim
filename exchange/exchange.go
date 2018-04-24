package exchange

type Exchange struct {
    name string
    volume float64
}

func New(name string, volume float64) Exchange {
    e := Exchange {name, volume}
    return e
}

func (e Exchange) Name() string {
    return e.name
}

func (e Exchange) Volume() float64 {
    return e.volume
}

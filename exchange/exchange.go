package exchange

type Exchange struct {
    name string
    volume float32
}

func New(name string, volume float32) Exchange {
    e := Exchange {name, volume}
    return e
}

func (e Exchange) Name() string {
    return e.name
}

func (e Exchange) Volume() float32 {
    return e.volume
}

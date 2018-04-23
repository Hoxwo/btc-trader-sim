package exchange

type exchange struct {
    name string
    volume float32
}

func New(name string, volume float32) exchange {
    e := exchange {name, volume}
    return e
}

func (e exchange) Name() string {
    return e.name
}

func (e exchange) Volume() float32 {
    return e.volume
}

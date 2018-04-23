package coin

type coin struct {
    name string
    price float32
    industry string
    trend string
}

func New(name string, price float32, industry string, trend string) coin {
    c := coin {name, price, industry, trend}
    return c
}

func (c coin) Name() string {
    return c.name
}

func (c coin) Price() float32 {
    return c.price
}

func (c coin) DailyPriceAdjustment() {
    c.price = c.price + 1.00
}

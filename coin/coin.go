package coin

type Coin struct {
    name string
    price float32
    industry string
    trend string
}

func New(name string, price float32, industry string, trend string) Coin {
    c := Coin {name, price, industry, trend}
    return c
}

func (c Coin) Name() string {
    return c.name
}

func (c Coin) Price() float32 {
    return c.price
}

func (c Coin) DailyPriceAdjustment() {
    c.price = c.price + 1.00
}

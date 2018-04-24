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

func (c *Coin) SetName(name string) {
    c.name = name
}

func (c Coin) Name() string {
    return c.name
}

func (c *Coin) SetPrice(price float32) {
    c.price = price
}

func (c Coin) Price() float32 {
    return c.price
}

func (c Coin) DailyPriceAdjustment() float32 {
    c.SetPrice(c.Price() + 1.77)
    return c.Price()
}

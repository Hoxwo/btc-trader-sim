package coin

import percent "github.com/dariubs/percent"
import "math/rand"
import "time"

type Coin struct {
    name string
    price float64
    industry string
    trend int
}

//trend 1: VERY BULLISH
//trend 2: BULLISH
//trend 3: SIDEWAYS
//trend 4: BEARISH
//trend 5: VERY BEARISH
func New(name string, price float64, industry string, trend int) Coin {
    c := Coin {name, price, industry, trend}
    return c
}

func (c *Coin) SetName(name string) {
    c.name = name
}

func (c Coin) Name() string {
    return c.name
}

func (c *Coin) SetPrice(price float64) {
    c.price = price
}

func (c Coin) Price() float64 {
    return c.price
}

func (c Coin) Trend() int {
    return c.trend
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func (c *Coin) DailyPriceAdjustment() float64 {

    if (c.Trend() == 1) {
	dailyGains := random(17, 35)
	c.SetPrice(c.Price() + float64(percent.Percent(dailyGains, int(c.Price()))))
    } else if (c.Trend() == 2) {
	dailyGains := random(4, 17)
	c.SetPrice(c.Price() + float64(percent.Percent(dailyGains, int(c.Price()))))
    } else if (c.Trend() == 3) {
	dailyGains := random(1, 4)
	c.SetPrice(c.Price() + float64(percent.Percent(dailyGains, int(c.Price()))))
    } else if (c.Trend() == 4) {
	dailyGains := random(4, 17) 
	c.SetPrice(c.Price() - float64(percent.Percent(dailyGains, int(c.Price()))))
    } else {
	dailyGains := random(17, 35) 
	c.SetPrice(c.Price() - float64(percent.Percent(dailyGains, int(c.Price()))))
    }

    return c.Price()
}

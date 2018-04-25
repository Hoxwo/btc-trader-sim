package coin

import "math/rand"
import "time"

type Coin struct {
    name string
    symbol string
    price float64
    supply int         //total supply of this coin available 
    launchDay int      //days into game this coin should become available (approx, maybe modify in main)
}

func New(name string, symbol string, price float64, supply int, launchDay int) Coin {
    c := Coin {name, symbol, price, supply, launchDay}
    return c
}

func (c *Coin) SetName(name string) {
    c.name = name
}

func (c Coin) Name() string {
    return c.name
}

func (c *Coin) SetSymbol(symbol string) {
    c.symbol = symbol
}

func (c Coin) Symbol() string {
    return c.symbol
}

func (c *Coin) SetPrice(price float64) {
    c.price = price
}

func (c Coin) Price() float64 {
    return c.price
}

func (c *Coin) SetLaunchDay(launchDay int) {
    c.launchDay = launchDay
}

func (c Coin) LaunchDay() int {
    return c.launchDay
}

func (c Coin) Supply() int {
    return c.supply
}

func (c *Coin) SetSupply(supply int) {
    c.supply = supply
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func (c *Coin) DailyPriceAdjustment(price float64) float64 {
    //set price calculated in main
    //add more here later maybe
    c.SetPrice(price)
    return c.Price()
}

func (c *Coin) DailyLaunchAdjustment(marketTrend int) int {
    //depending on overall market trend, this coin's launch will slide forward or backward
    if (marketTrend == 1) {
	dailySlide := random(1, 7)
	c.SetLaunchDay(c.LaunchDay() - dailySlide) 
    } else if (marketTrend == 2) {
	dailySlide := random(1, 3)
	c.SetLaunchDay(c.LaunchDay() - dailySlide) 
    } else if (marketTrend == 3) {
	dailySlide := random(1, 3)
	c.SetLaunchDay(c.LaunchDay() + dailySlide) 
    } else {
	dailySlide := random(1, 7)
	c.SetLaunchDay(c.LaunchDay() + dailySlide) 
    }

    return c.LaunchDay()
}

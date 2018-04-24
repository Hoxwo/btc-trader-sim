package coin

import percent "github.com/dariubs/percent"
import "math/rand"
import "time"

type Coin struct {
    name string
    symbol string
    price float64
    trend int          //BULLISH with HYPE 1, BULLISH 2, BEARISH 3, or BEARISH with DOUBT 4, affects market share modifier
    minShare int       //min total market cap share this coin can have, out of 100
    maxShare int       //max total market cap share this coin can have, out of 100
    currentShare int   //current total market cap share this coin has
    supply uint64         //total supply of this coin available 
    launchDay int      //days into game this coin should become available (approx, maybe modify in main)
}

func New(name string, symbol string, price float64, trend int, minShare int, maxShare int, currentShare int, supply uint64, launchDay int) Coin {
    c := Coin {name, symbol, price, trend, minShare, maxShare, currentShare, supply, launchDay}
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

func (c *Coin) SetTrend(trend int) {
    c.trend = trend
}

func (c Coin) MinShare() int {
    return c.minShare
}

func (c *Coin) SetMinShare(minShare int) {
    c.minShare = minShare
}

func (c Coin) MaxShare() int {
    return c.maxShare
}

func (c *Coin) SetMaxShare(maxShare int) {
    c.maxShare = maxShare
}

func (c *Coin) SetLaunchDay(launchDay int) {
    c.launchDay = launchDay
}

func (c Coin) LaunchDay() int {
    return c.launchDay
}

func (c Coin) CurrentShare() int {
    return c.trend
}

func (c *Coin) SetCurrentShare(currentShare int) {
    c.currentShare = currentShare
}

func (c Coin) Supply() uint64 {
    return c.supply
}

func (c *Coin) SetSupply(supply uint64) {
    c.supply = supply
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func (c *Coin) DailyPriceAdjustment(totalMarketCap float64) float64 {
    //add or remove total market share, find the total marketcap price / this coin's share, divide by this coin's supply for price
    //set price
    if (c.Trend() == 1) {
	dailyGains := random(17, 35)
	c.SetPrice(c.Price() + float64(percent.Percent(dailyGains, int(c.Price())))) 
    } else {
	dailyGains := random(4, 17)
	c.SetPrice(c.Price() - float64(percent.Percent(dailyGains, int(c.Price()))))
    }

    return c.Price()
}

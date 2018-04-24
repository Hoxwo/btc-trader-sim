package coin

import percent "github.com/dariubs/percent"
import "math/rand"
import "time"

type Coin struct {
    name string
    symbol string
    price float64
    trend int          //BULLISH with HYPE 1, BULLISH 2, BEARISH 3, or BEARISH with DOUBT 4, affects market share modifier and launch times 
    minShare int       //min total market cap share this coin can have, out of 100
    maxShare int       //max total market cap share this coin can have, out of 100
    currentShare int   //current total market cap share this coin has
    supply int         //total supply of this coin available 
    launchDay int      //days into game this coin should become available (approx, maybe modify in main)
}

func New(name string, symbol string, price float64, trend int, minShare int, maxShare int, currentShare int, supply int, launchDay int) Coin {
    c := Coin {name, symbol, price, trend, minShare, maxShare, currentShare, supply, launchDay}
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

func (c *Coin) DailyPriceAdjustment(totalMarketCap int) float64 {
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

func (c *Coin) DailyLaunchAdjustment(marketTrend int) int {
    //depending on overall market trend, this coin's launch will slide forward or backward
    if (c.Trend() == 1) {
	dailySlide := random(1, 7)
	c.SetLaunchDay(c.LaunchDay() - dailySlide) 
    } else if (c.Trend() == 2) {
	dailySlide := random(1, 3)
	c.SetLaunchDay(c.LaunchDay() - dailySlide) 
    } else if (c.Trend() == 3) {
	dailySlide := random(1, 3)
	c.SetLaunchDay(c.LaunchDay() + dailySlide) 
    } else {
	dailySlide := random(1, 7)
	c.SetLaunchDay(c.LaunchDay() + dailySlide) 
    }

    return c.LaunchDay()
}

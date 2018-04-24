package main

import "fmt"
import "time"
import percent "github.com/dariubs/percent"
import coin "btc-trader-sim/coin"
import trader "btc-trader-sim/trader"
import exchange "btc-trader-sim/exchange"
import termui "github.com/gizak/termui"

func main() {
	// set the time to Jan 1st, 2010
	currentTime := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	//to track the number of days	
	dayCounter := 0
        //track the total market cap
        var totalMarketCap int = 0
	// the current market trend
	var marketTrend int = 2 //BULLISH with HYPE 1, BULLISH 2, BEARISH 3, or BEARISH with DOUBT 4, affects market share modifier and launches
	// array of coins 
	coins := make([]*coin.Coin, 0)
	// array of coin prices
	coinPrices := make(map[string]float64)
	// array of maps for storing coin price history	
	coinPriceHistory := make(map[string][]float64)
	// array of exchanges
	exchanges := make([]*exchange.Exchange, 0)
	// array of exchange valueAdded
	exchangeValues := make(map[string]int)
	// array of maps for storing exchange value history	
	exchangeValueHistory := make(map[string][]int)

	// set up coins
	// 	{name, symbol, price, trend, minShare, maxShare, currentShare, supply, launchDay}
	c0 := coin.New("Bitcoin",           "BTC", 0.00, 1, 15, 100, 100,     20,       1)
	c1 := coin.New("LightCoin",         "LGC", 0.00, 1, 2,  30,  2,       55,     125)
	c2 := coin.New("Nethereum", 	    "NTH", 0.00, 1, 5,  75,  5,      100,    490)
	c3 := coin.New("Nethereum Vintage", "NTV", 0.00, 1, 3,  25,  3,      100,    855)	
	c4 := coin.New("Riddle",            "XRD", 0.00, 1, 5,  50,  5,    50000, 1220)
	c5 := coin.New("ZEO",               "ZEO", 0.00, 1, 2,  50,  2,       70,    1585)
	c6 := coin.New("YCash",             "YEC", 0.00, 1, 10, 25,  10,       4,     1850)
	c7 := coin.New("Interstellar",      "ILM", 0.00, 1, 5,  35,  5,    18000, 2215)
	c8 := coin.New("Bitbeets",          "BBT", 0.00, 1, 1,  15,  1,     2000,  2580)
	c9 := coin.New("TRAM",              "TRM", 0.00, 1, 1,  20,  1,    70000, 2945)
	
	//add em to our master list
	coins = append(coins, &c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9)
	
	//start the price and history tracking	
	coinPrices[c0.Name()] = c0.Price()
	coinPrices[c1.Name()] = c1.Price()
	coinPrices[c2.Name()] = c2.Price()
	coinPrices[c3.Name()] = c3.Price()
	coinPrices[c4.Name()] = c4.Price()
	coinPrices[c5.Name()] = c5.Price()
	coinPrices[c6.Name()] = c6.Price()	
	coinPrices[c7.Name()] = c7.Price()
	coinPrices[c8.Name()] = c8.Price()
	coinPrices[c9.Name()] = c9.Price()
	arr1 := make([]float64, 0)
	coinPriceHistory[c0.Name()] = append(arr1, c0.Price())
	coinPriceHistory[c1.Name()] = append(arr1, c1.Price())
	coinPriceHistory[c2.Name()] = append(arr1, c2.Price())
	coinPriceHistory[c3.Name()] = append(arr1, c3.Price())
	coinPriceHistory[c4.Name()] = append(arr1, c4.Price())
	coinPriceHistory[c5.Name()] = append(arr1, c5.Price())
	coinPriceHistory[c6.Name()] = append(arr1, c6.Price())
	coinPriceHistory[c7.Name()] = append(arr1, c7.Price())
	coinPriceHistory[c8.Name()] = append(arr1, c8.Price())
	coinPriceHistory[c9.Name()] = append(arr1, c9.Price())

	//set up exchanges 
	e0 := exchange.New("Mt Ganx",   0,  250, 1)
	e1 := exchange.New("GDOX",      0,  500, 366)
	e2 := exchange.New("BitSaurus", 0, 1000, 730)
	e3 := exchange.New("CoinHQ",    0, 1500, 1095)
	e4 := exchange.New("Czinance",  0, 1000, 1460)
	e5 := exchange.New("Napoleox",  0,  750, 1825)
	e6 := exchange.New("YoCoin",    0,  250, 2190)
	e7 := exchange.New("CoinHawk",  0,  250, 2555)	

	// add to master list
	exchanges = append(exchanges, &e0, &e1, &e2, &e3, &e4, &e5, &e6, &e7)

	//start the exchange and history tracking	
	exchangeValues[e0.Name()] = e0.ValueAdded()
	exchangeValues[e1.Name()] = e1.ValueAdded()
	exchangeValues[e2.Name()] = e2.ValueAdded()
	exchangeValues[e3.Name()] = e3.ValueAdded()
	exchangeValues[e4.Name()] = e4.ValueAdded()
	exchangeValues[e5.Name()] = e5.ValueAdded()
	exchangeValues[e6.Name()] = e6.ValueAdded()	
	exchangeValues[e7.Name()] = e7.ValueAdded()
	arr2 := make([]int, 0)
	exchangeValueHistory[e0.Name()] = append(arr2, e0.ValueAdded())
	exchangeValueHistory[e1.Name()] = append(arr2, e1.ValueAdded())
	exchangeValueHistory[e2.Name()] = append(arr2, e2.ValueAdded())
	exchangeValueHistory[e3.Name()] = append(arr2, e3.ValueAdded())
	exchangeValueHistory[e4.Name()] = append(arr2, e4.ValueAdded())
	exchangeValueHistory[e5.Name()] = append(arr2, e5.ValueAdded())
	exchangeValueHistory[e6.Name()] = append(arr2, e6.ValueAdded())
	exchangeValueHistory[e7.Name()] = append(arr2, e7.ValueAdded())

  	//t := trader.New("kc", 100.00)
        
    
	//coinPrice := fmt.Sprintf("%s | %.6f \n", c.Name(), c.Price())
    	//playerSavings := fmt.Sprintf("%s has savings balance %.6f \n", t.Name(), t.SavingsBalance())
        //exchangeVolume := fmt.Sprintf("%s has volume %.6f \n", e.Name(), e.Volume())

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	//termui.Render(spls1, singlespl0, par0)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/i", func(termui.Event) {
		currentTime = currentTime.Add(time.Hour * 24 * 1)
		dayCounter++
		AdvanceOneDay(coins, exchanges, coinPrices, exchangeValues, coinPriceHistory, exchangeValueHistory, dayCounter, totalMarketCap, 						marketTrend)
	
		// Short term dollar amounts, or estimate of day until launch
		shorttermhisttitle0 := ShortTermCoinTitle(coins[0], dayCounter)
		shorttermhist0 := termui.NewSparkline()
		shorttermhist0.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitcoin", coinPriceHistory))
		shorttermhist0.Title = shorttermhisttitle0
		shorttermhist0.LineColor = termui.ColorGreen

		shorttermhisttitle1 := ShortTermCoinTitle(coins[1], dayCounter)
		shorttermhist1 := termui.NewSparkline()
		shorttermhist1.Data = FloatToInts(GetHistoricPriceDataForCoin("Lightcoin", coinPriceHistory))
		shorttermhist1.Title = shorttermhisttitle1
		shorttermhist1.LineColor = termui.ColorCyan

		shorttermhisttitle2 := ShortTermCoinTitle(coins[2], dayCounter)
		shorttermhist2 := termui.NewSparkline()
		shorttermhist2.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum", coinPriceHistory))
		shorttermhist2.Title = shorttermhisttitle2
		shorttermhist2.LineColor = termui.ColorMagenta

		shorttermhisttitle3 := ShortTermCoinTitle(coins[3], dayCounter)
		shorttermhist3 := termui.NewSparkline()	
		shorttermhist3.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum Vintage", coinPriceHistory))
		shorttermhist3.Title = shorttermhisttitle3
		shorttermhist3.LineColor = termui.ColorGreen

		shorttermhisttitle4 := ShortTermCoinTitle(coins[4], dayCounter)
		shorttermhist4 := termui.NewSparkline()
		shorttermhist4.Data = FloatToInts(GetHistoricPriceDataForCoin("Riddle", coinPriceHistory))
		shorttermhist4.Title = shorttermhisttitle4
		shorttermhist4.LineColor = termui.ColorCyan
	
		shorttermhisttitle5 := ShortTermCoinTitle(coins[5], dayCounter)
		shorttermhist5 := termui.NewSparkline()
		shorttermhist5.Data = FloatToInts(GetHistoricPriceDataForCoin("ZEO", coinPriceHistory))
		shorttermhist5.Title = shorttermhisttitle5
		shorttermhist5.LineColor = termui.ColorMagenta
	
		shorttermhisttitle6 := ShortTermCoinTitle(coins[6], dayCounter)
		shorttermhist6 := termui.NewSparkline()
		shorttermhist6.Data = FloatToInts(GetHistoricPriceDataForCoin("YCash", coinPriceHistory))
		shorttermhist6.Title = shorttermhisttitle6
		shorttermhist6.LineColor = termui.ColorGreen
	
		shorttermhisttitle7 := ShortTermCoinTitle(coins[7], dayCounter)
		shorttermhist7 := termui.NewSparkline()
		shorttermhist7.Data = FloatToInts(GetHistoricPriceDataForCoin("Intersteller", coinPriceHistory))
		shorttermhist7.Title = shorttermhisttitle7
		shorttermhist7.LineColor = termui.ColorCyan

		shorttermhisttitle8 := ShortTermCoinTitle(coins[8], dayCounter)
		shorttermhist8 := termui.NewSparkline()
		shorttermhist8.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitbeets", coinPriceHistory))
		shorttermhist8.Title = shorttermhisttitle8
		shorttermhist8.LineColor = termui.ColorMagenta

		shorttermhisttitle9 := ShortTermCoinTitle(coins[9], dayCounter)
		shorttermhist9 := termui.NewSparkline()
		shorttermhist9.Data = FloatToInts(GetHistoricPriceDataForCoin("TRAM", coinPriceHistory))
		shorttermhist9.Title = shorttermhisttitle9
		shorttermhist9.LineColor = termui.ColorGreen

		// put them together
		shorttermhistograms := termui.NewSparklines(shorttermhist0, shorttermhist1, shorttermhist2, 
					shorttermhist3, shorttermhist4, shorttermhist5,
					shorttermhist6, shorttermhist7, shorttermhist8,
					shorttermhist9)
		shorttermhistograms.Height = 20
		shorttermhistograms.Width = 32
		shorttermhistograms.Y = 12
		shorttermhistograms.X = 64
		shorttermhistograms.BorderLabel = "Short Term $ History"
		
		//List of exchanges - presented as gauges of total cap
		exchangeGauge0 := termui.NewGauge()
		exchangeGauge0.Percent = ExchangeInfoPercent(exchanges[0])
		exchangeGauge0.Width = 28
		exchangeGauge0.Height = 3
		exchangeGauge0.Y = 0
		exchangeGauge0.X = 0
		exchangeGauge0.BorderLabel = ExchangeInfoString(exchanges[0], dayCounter)
		exchangeGauge0.Label = ExchangeInfoLabel(exchanges[0])
		exchangeGauge0.BarColor = termui.ColorMagenta
		exchangeGauge0.BorderFg = termui.ColorWhite
		exchangeGauge0.LabelAlign = termui.AlignRight

		exchangeGauge1 := termui.NewGauge()
		exchangeGauge1.Percent = ExchangeInfoPercent(exchanges[1])
		exchangeGauge1.Width = 28
		exchangeGauge1.Height = 3
		exchangeGauge1.Y = 3
		exchangeGauge1.X = 0
		exchangeGauge1.BorderLabel = ExchangeInfoString(exchanges[1], dayCounter)
		exchangeGauge1.Label = ExchangeInfoLabel(exchanges[1])
		exchangeGauge1.BarColor = termui.ColorMagenta
		exchangeGauge1.BorderFg = termui.ColorWhite
		exchangeGauge1.LabelAlign = termui.AlignRight
		
		exchangeGauge2 := termui.NewGauge()
		exchangeGauge2.Percent = ExchangeInfoPercent(exchanges[2])
		exchangeGauge2.Width = 28
		exchangeGauge2.Height = 3
		exchangeGauge2.Y = 6
		exchangeGauge2.X = 0
		exchangeGauge2.BorderLabel = ExchangeInfoString(exchanges[2], dayCounter)
		exchangeGauge2.Label = ExchangeInfoLabel(exchanges[2])
		exchangeGauge2.BarColor = termui.ColorMagenta
		exchangeGauge2.BorderFg = termui.ColorWhite
		exchangeGauge2.LabelAlign = termui.AlignRight

		exchangeGauge3 := termui.NewGauge()
		exchangeGauge3.Percent = ExchangeInfoPercent(exchanges[3])
		exchangeGauge3.Width = 28
		exchangeGauge3.Height = 3
		exchangeGauge3.Y = 9
		exchangeGauge3.X = 0
		exchangeGauge3.BorderLabel = ExchangeInfoString(exchanges[3], dayCounter)
		exchangeGauge3.Label = ExchangeInfoLabel(exchanges[3])
		exchangeGauge3.BarColor = termui.ColorMagenta
		exchangeGauge3.BorderFg = termui.ColorWhite
		exchangeGauge3.LabelAlign = termui.AlignRight

		exchangeGauge4 := termui.NewGauge()
		exchangeGauge4.Percent = ExchangeInfoPercent(exchanges[4])
		exchangeGauge4.Width = 28
		exchangeGauge4.Height = 3
		exchangeGauge4.Y = 12
		exchangeGauge4.X = 0
		exchangeGauge4.BorderLabel = ExchangeInfoString(exchanges[4], dayCounter)
		exchangeGauge4.Label = ExchangeInfoLabel(exchanges[4])
		exchangeGauge4.BarColor = termui.ColorMagenta
		exchangeGauge4.BorderFg = termui.ColorWhite
		exchangeGauge4.LabelAlign = termui.AlignRight

		exchangeGauge5 := termui.NewGauge()
		exchangeGauge5.Percent = ExchangeInfoPercent(exchanges[5])
		exchangeGauge5.Width = 28
		exchangeGauge5.Height = 3
		exchangeGauge5.Y = 15
		exchangeGauge5.X = 0
		exchangeGauge5.BorderLabel = ExchangeInfoString(exchanges[5], dayCounter)
		exchangeGauge5.Label = ExchangeInfoLabel(exchanges[5])
		exchangeGauge5.BarColor = termui.ColorMagenta
		exchangeGauge5.BorderFg = termui.ColorWhite
		exchangeGauge5.LabelAlign = termui.AlignRight

		exchangeGauge6 := termui.NewGauge()
		exchangeGauge6.Percent = ExchangeInfoPercent(exchanges[6])
		exchangeGauge6.Width = 28
		exchangeGauge6.Height = 3
		exchangeGauge6.Y = 18
		exchangeGauge6.X = 0
		exchangeGauge6.BorderLabel = ExchangeInfoString(exchanges[6], dayCounter)
		exchangeGauge6.Label = ExchangeInfoLabel(exchanges[6])
		exchangeGauge6.BarColor = termui.ColorMagenta
		exchangeGauge6.BorderFg = termui.ColorWhite
		exchangeGauge6.LabelAlign = termui.AlignRight

		exchangeGauge7 := termui.NewGauge()
		exchangeGauge7.Percent = ExchangeInfoPercent(exchanges[7])
		exchangeGauge7.Width = 28
		exchangeGauge7.Height = 3
		exchangeGauge7.Y = 21
		exchangeGauge7.X = 0
		exchangeGauge7.BorderLabel = ExchangeInfoString(exchanges[7], dayCounter)
		exchangeGauge7.Label = ExchangeInfoLabel(exchanges[7])
		exchangeGauge7.BarColor = termui.ColorMagenta
		exchangeGauge7.BorderFg = termui.ColorWhite
		exchangeGauge7.LabelAlign = termui.AlignRight
	
	par1 := termui.NewPar(currentTime.Format("01-02-2006"))
	par1.Height = 1
	par1.Width = 20
	par1.X = 16
	par1.Y = 12
	par1.Border = false

	par3 := termui.NewPar(fmt.Sprintf("days %d", dayCounter))
	par3.Height = 1
	par3.Width = 20
	par3.X = 20
	par3.Y = 10
	par3.Border = false		
	termui.Render( shorttermhistograms, par1, par3, exchangeGauge0, exchangeGauge1, exchangeGauge2, exchangeGauge3,
				exchangeGauge4, exchangeGauge5, exchangeGauge6, exchangeGauge7)
	})

	termui.Loop()

}

func AdvanceOneDay(coins []*coin.Coin, exchanges []*exchange.Exchange, coinPrices map[string]float64, exchangeValues map[string]int, 			coinPriceHistory map[string][]float64, exchangeValueHistory map[string][]int, dayCounter int, totalMarketCap int, marketTrend int) {
	//save today's exchange valueAdded for all exchanges
	//find new exchange valueAdded for all exchanges
	for _, e := range exchanges {
	    if(e.LaunchDay() > dayCounter) {
	    	e.DailyLaunchAdjustment(marketTrend)
	    } else {
	        currentValueHistory := exchangeValueHistory[e.Name()]
	        delete(exchangeValueHistory, e.Name())
	        exchangeValueHistory[e.Name()] = append(currentValueHistory, exchangeValues[e.Name()])
	        exchangeValues[e.Name()] = e.DailyValueAdjustment(totalMarketCap, marketTrend)
	    }
	}			

	// save coin price history for all coins
	// and find next day's value
	for _, c := range coins {
	    currentPriceHistory := coinPriceHistory[c.Name()]
	    delete(coinPriceHistory, c.Name())
	    coinPriceHistory[c.Name()] = append(currentPriceHistory, coinPrices[c.Name()])
	    coinPrices[c.Name()] = c.DailyPriceAdjustment(totalMarketCap)
	    if(c.LaunchDay() > dayCounter) {
	    	c.DailyLaunchAdjustment(marketTrend)
	    }
	}		

}

func ExchangeInfoString(exchange *exchange.Exchange, dayCounter int) string {
	exchangeInfo := ""

	if(exchange.LaunchDay() > dayCounter) {
		exchangeInfo = fmt.Sprintf("%s ETA %d days", exchange.Name(), exchange.LaunchDay())
	} else {
		exchangeInfo = fmt.Sprintf("%s value: $%d B", exchange.Name(), exchange.ValueAdded())
	}
	
	return exchangeInfo
}

func ExchangeInfoPercent(exchange *exchange.Exchange) int {
	return int(percent.PercentOf(exchange.ValueAdded(), exchange.MaxValueAdded()))
}

func ExchangeInfoLabel(exchange *exchange.Exchange) string {
	exchangeLabel := ""
	exchangeLabel = fmt.Sprintf(" of $%dB", exchange.MaxValueAdded())
	return exchangeLabel
}

func ShortTermCoinTitle(coin *coin.Coin, dayCounter int) string {
	title := ""
	
	if(coin.LaunchDay() > dayCounter) {
		title = fmt.Sprintf("%s - ETA %d days", coin.Symbol(), coin.LaunchDay())
	} else {
		title = fmt.Sprintf("%s - $%.2f ", coin.Symbol(), coin.Price())
	}
	
	return title
}

func GetHistoricPriceDataForCoin(coin string, coinPriceHistory map[string][]float64) []float64 {
	return coinPriceHistory[coin]
}

func GetTraderDollarValueForCoin(t trader.Trader, coin string, coinPriceHistory map[string][]float64) []float64 {
	traderBalance := t.HistoricBalanceForCoin(coin)
	traderDollarValue := make([]float64, len(traderBalance))
	for i, _ := range traderBalance {
		traderDollarValue[i] = (traderBalance[i]*coinPriceHistory[coin][i])
	}

	return traderDollarValue
}

func GetTraderDollarValueForAllCoins(t trader.Trader, coinPriceHistory map[string][]float64) []float64 {
	ownedCoins := t.OwnedCoins()
	sumAllCoins := make([]float64, len(GetTraderDollarValueForCoin(t, "bitcoin", coinPriceHistory)))
	for _, c := range ownedCoins {
		oneCoinHistory := GetTraderDollarValueForCoin(t, c, coinPriceHistory)
		for i, _ := range oneCoinHistory {
			sumAllCoins[i] = sumAllCoins[i] + oneCoinHistory[i]
		}
	}

	return sumAllCoins
}

func FloatToInts(floatArray []float64) []int {
	intArray := make([]int, len(floatArray))
	for i, _ := range floatArray {
		intArray[i] = int(floatArray[i])
	}

	return intArray
}

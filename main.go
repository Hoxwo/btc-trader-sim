package main

import "fmt"
import "time"
import coin "btc-trader-sim/coin"
import trader "btc-trader-sim/trader"
//import exchange "btc-trader-sim/exchange"
import termui "github.com/gizak/termui"

func main() {
	// set the time to Jan 1st, 2010
	currentTime := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	//to track the number of days	
	dayCounter := 0
        //track the total market cap
        var totalMarketCap float64 = 0.00
	// array of coins 
	coins := make([]*coin.Coin, 0)
	// array of coin prices
	coinPrices := make(map[string]float64)
	// array of maps for storing coin price history	
	coinPriceHistory := make(map[string][]float64)
	// array of exchanges
	//exchanges := make([]exchange, 20)

	// set up coins
	// 	{name, symbol, price, trend, minShare, maxShare, currentShare, supply, launchDay}
	c0 := coin.New("Bitcoin",           "BTC", 0.00, 1, 15, 100, 100,  20000000,       1)
	c1 := coin.New("LightCoin",         "LGC", 0.00, 1, 2,  30,  2,    55000000,     125)
	c2 := coin.New("Nethereum", 	    "NTH", 0.00, 1, 5,  75,  5,    100000000,    490)
	c3 := coin.New("Nethereum Vintage", "NTV", 0.00, 1, 3,  25,  3,    100000000,    855)	
	c4 := coin.New("Riddle",            "XRD", 0.00, 1, 5,  50,  5,    50000000000, 1220)
	c5 := coin.New("ZEO",               "ZEO", 0.00, 1, 2,  50,  2,    70000000,    1585)
	c6 := coin.New("YCash",             "YEC", 0.00, 1, 10, 25,  10,   4000000,     1850)
	c7 := coin.New("Interstellar",      "ILM", 0.00, 1, 5,  35,  5,    18000000000, 2215)
	c8 := coin.New("Bitbeets",          "BBT", 0.00, 1, 1,  15,  1,    2000000000,  2580)
	c9 := coin.New("TRAM",              "TRM", 0.00, 1, 1,  20,  1,    70000000000, 2945)
	
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


  	//t := trader.New("kc", 100.00)
        //e := exchange.New("gdax", 100.00)
    
	//coinPrice := fmt.Sprintf("%s | %.6f \n", c.Name(), c.Price())
    	//playerSavings := fmt.Sprintf("%s has savings balance %.6f \n", t.Name(), t.SavingsBalance())
        //exchangeVolume := fmt.Sprintf("%s has volume %.6f \n", e.Name(), e.Volume())

    //ColorRed
    //ColorGreen
    //ColorYellow
    //ColorBlue
    //ColorMagenta
    //ColorCyan
    //ColorWhite

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
		currentTime.AddDate(0, 0, 1)
		dayCounter++
		AdvanceOneDay(coins, coinPrices, coinPriceHistory, dayCounter, totalMarketCap)
	
		// Short term dollar amounts, or estimate of day until ICO
		shorttermhist0 := termui.NewSparkline()
		shorttermhist0.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitcoin", coinPriceHistory))
		shorttermhist0.Title = "BTC"
		shorttermhist0.LineColor = termui.ColorGreen

		shorttermhist1 := termui.NewSparkline()
		shorttermhist1.Data = FloatToInts(GetHistoricPriceDataForCoin("Lightcoin", coinPriceHistory))
		shorttermhist1.Title = "LGC"
		shorttermhist1.LineColor = termui.ColorCyan

		shorttermhist2 := termui.NewSparkline()
		shorttermhist2.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum", coinPriceHistory))
		shorttermhist2.Title = "NTH"
		shorttermhist2.LineColor = termui.ColorMagenta

		shorttermhist3 := termui.NewSparkline()	
		shorttermhist3.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum Vintage", coinPriceHistory))
		shorttermhist3.Title = "NTV"
		shorttermhist3.LineColor = termui.ColorGreen

		shorttermhist4 := termui.NewSparkline()
		shorttermhist4.Data = FloatToInts(GetHistoricPriceDataForCoin("Riddle", coinPriceHistory))
		shorttermhist4.Title = "XRD"
		shorttermhist4.LineColor = termui.ColorCyan
	
		shorttermhist5 := termui.NewSparkline()
		shorttermhist5.Data = FloatToInts(GetHistoricPriceDataForCoin("ZEO", coinPriceHistory))
		shorttermhist5.Title = "ZEO"
		shorttermhist5.LineColor = termui.ColorMagenta
	
		shorttermhist6 := termui.NewSparkline()
		shorttermhist6.Data = FloatToInts(GetHistoricPriceDataForCoin("YCash", coinPriceHistory))
		shorttermhist6.Title = "YEC"
		shorttermhist6.LineColor = termui.ColorGreen
	
		shorttermhist7 := termui.NewSparkline()
		shorttermhist7.Data = FloatToInts(GetHistoricPriceDataForCoin("Intersteller", coinPriceHistory))
		shorttermhist7.Title = "ILM"
		shorttermhist7.LineColor = termui.ColorCyan
	
		shorttermhist8 := termui.NewSparkline()
		shorttermhist8.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitbeets", coinPriceHistory))
		shorttermhist8.Title = "BBT"
		shorttermhist8.LineColor = termui.ColorMagenta

		shorttermhist9 := termui.NewSparkline()
		shorttermhist9.Data = FloatToInts(GetHistoricPriceDataForCoin("TRAM", coinPriceHistory))
		shorttermhist9.Title = "TRM"
		shorttermhist9.LineColor = termui.ColorGreen

		// put them together
		shorttermhistograms := termui.NewSparklines(shorttermhist0, shorttermhist1, shorttermhist2, 
					shorttermhist3, shorttermhist4, shorttermhist5,
					shorttermhist6, shorttermhist7, shorttermhist8,
					shorttermhist9)
		shorttermhistograms.Height = 20
		shorttermhistograms.Width = 24
		shorttermhistograms.Y = 12
		shorttermhistograms.X = 50
		shorttermhistograms.BorderLabel = "Short Term $ History"

	// single
        singledata := FloatToInts(GetHistoricPriceDataForCoin("bitcoin", coinPriceHistory))
	single0 := termui.NewSparkline()
	single0.Data = singledata
	single0.Title = "Test"
	single0.LineColor = termui.ColorCyan
		
	singlespl0 := termui.NewSparklines(single0)
	singlespl0.Height = 6
	singlespl0.Width = 20
	singlespl0.Border = false
	
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
	termui.Render(shorttermhistograms, singlespl0, par1, par3)
	})

	termui.Loop()

}

func AdvanceOneDay(coins []*coin.Coin, coinPrices map[string]float64, coinPriceHistory map[string][]float64, dayCounter int, totalMarketCap float64) {
	// save coin price history for all coins
	// and find next day's value
	for _, c := range coins {
	    currentPriceHistory := coinPriceHistory[c.Name()]
	    delete(coinPriceHistory, c.Name())
	    coinPriceHistory[c.Name()] = append(currentPriceHistory, coinPrices[c.Name()])
	    coinPrices[c.Name()] = c.DailyPriceAdjustment(totalMarketCap)
	}		
	//find new exchange volume for all exchanges

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

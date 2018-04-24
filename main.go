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
	dayCounter := 0
	// array of coins 
	coins := make([]coin.Coin, 1)
	// array of coin prices
	coinPrices := make(map[string]float32)
	// array of maps for storing coin price history	
	coinPriceHistory := make(map[string][]float32)
	// array of exchanges
	//exchanges := make([]exchange, 20)

	// set up Bitcoin, player Trader, and an exchange
	c := coin.New("bitcoin", 10.01, "FINANCE", "BULLISH")
	c2 := coin.New("eth", 3.01, "FINANCE", "BULLISH")
	c3 := coin.New("ltc", 1.01, "FINANCE", "BULLISH")
	coins = append(coins, c, c2, c3)
	coinPrices[c.Name()] = c.Price()
	coinPrices[c2.Name()] = c2.Price()
	coinPrices[c3.Name()] = c3.Price()
	arr1 := make([]float32, 1)
	coinPriceHistory[c.Name()] = append(arr1, c.Price())
	coinPriceHistory[c2.Name()] = append(arr1, c2.Price())
	coinPriceHistory[c3.Name()] = append(arr1, c3.Price())
  	//t := trader.New("kc", 100.00)
        //e := exchange.New("gdax", 100.00)
    
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
		dayCounter++
		AdvanceOneDay(coins, coinPrices, coinPriceHistory, dayCounter)
	//////
	spl0 := termui.NewSparkline()
	spl0.Data = FloatToInts(GetHistoricPriceDataForCoin("bitcoin", coinPriceHistory))
	spl0.Title = "BTC"
	spl0.LineColor = termui.ColorGreen

	spl1 := termui.NewSparkline()
	spl1.Data = FloatToInts(GetHistoricPriceDataForCoin("eth", coinPriceHistory))
	spl1.Title = "ETH"
	spl1.LineColor = termui.ColorRed

	spl2 := termui.NewSparkline()
	spl2.Data = FloatToInts(GetHistoricPriceDataForCoin("ltc", coinPriceHistory))
	spl2.Title = "LTC"
	spl2.LineColor = termui.ColorMagenta

	// group
	spls1 := termui.NewSparklines(spl0, spl1, spl2)
	spls1.Height = 10
	spls1.Width = 20
	spls1.Y = 16
	spls1.BorderLabel = "Dollar Values"

	// single
        singledata := FloatToInts(GetHistoricPriceDataForCoin("bitcoin", coinPriceHistory))
	single0 := termui.NewSparkline()
	single0.Data = singledata
	single0.Title = "Sparkline 0"
	single0.LineColor = termui.ColorCyan
		
	singlespl0 := termui.NewSparklines(single0)
	singlespl0.Height = 6
	singlespl0.Width = 20
	singlespl0.Border = false

	par0 := termui.NewPar(currentTime.Format("01-02-2006"))
	par0.Height = 1
	par0.Width = 20
	par0.Y = 1
	par0.Border = false
	
	par1 := termui.NewPar(fmt.Sprintf("coins size %d", len(coins)))
	par1.Height = 1
	par1.Width = 20
	par1.X = 20
	par1.Y = 4
	par1.Border = false

	par2 := termui.NewPar(fmt.Sprintf("btc price size %d", len(GetHistoricPriceDataForCoin("bitcoin", coinPriceHistory))))
	par2.Height = 1
	par2.Width = 20
	par2.X = 20
	par2.Y = 6
	par2.Border = false

	par3 := termui.NewPar(fmt.Sprintf("days %d", dayCounter))
	par3.Height = 1
	par3.Width = 20
	par3.X = 20
	par3.Y = 8
	par3.Border = false		
	/////
		termui.Render(spls1, singlespl0, par0, par1, par2, par3)
	})

	termui.Loop()

}

func AdvanceOneDay(coins []coin.Coin, coinPrices map[string]float32, coinPriceHistory map[string][]float32, dayCounter int) {
	// save coin price history for all coins
	// and find next day's value
	for _, c := range coins {
	    currentPriceHistory := coinPriceHistory[c.Name()]
	    coinPriceHistory[c.Name()] = append(currentPriceHistory, coinPrices[c.Name()])
    	    c.DailyPriceAdjustment()
	}		
	//find new exchange volume for all exchanges

	dayCounter++
}

func GetHistoricPriceDataForCoin(coin string, coinPriceHistory map[string][]float32) []float32 {
	return coinPriceHistory[coin]
}

func GetTraderDollarValueForCoin(t trader.Trader, coin string, coinPriceHistory map[string][]float32) []float32 {
	traderBalance := t.HistoricBalanceForCoin(coin)
	traderDollarValue := make([]float32, len(traderBalance))
	for i, _ := range traderBalance {
		traderDollarValue[i] = (traderBalance[i]*coinPriceHistory[coin][i])
	}

	return traderDollarValue
}

func GetTraderDollarValueForAllCoins(t trader.Trader, coinPriceHistory map[string][]float32) []float32 {
	ownedCoins := t.OwnedCoins()
	sumAllCoins := make([]float32, len(GetTraderDollarValueForCoin(t, "bitcoin", coinPriceHistory)))
	for _, c := range ownedCoins {
		oneCoinHistory := GetTraderDollarValueForCoin(t, c, coinPriceHistory)
		for i, _ := range oneCoinHistory {
			sumAllCoins[i] = sumAllCoins[i] + oneCoinHistory[i]
		}
	}

	return sumAllCoins
}

func FloatToInts(floatArray []float32) []int {
	intArray := make([]int, len(floatArray))
	for i, _ := range floatArray {
		intArray[i] = int(floatArray[i])
	}

	return intArray
}

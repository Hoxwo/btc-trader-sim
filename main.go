package main

//import "fmt"
//import "time"
import termui "github.com/gizak/termui"
//import "btc-trader-sim/coin"
//import "btc-trader-sim/trader"
//import "btc-trader-sim/exchange"

func main() {
	// set the time to Jan 1st, 2010
	//currentTime := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	//dayCounter := 0
	// array of coins
	// coinPrices = make(map[string]float32)
	// array of maps for storing coin price history	
	// coinPriceHistory := make([]map[string]float32, 5900)	

	// set up Bitcoin, player Trader, and an exchange
	//c := coin.New("bitcoin", 1.01, "FINANCE", "BULLISH")
  //  	t := trader.New("kc", 0.00)
      //  e := exchange.New("gdax", 100.00)
    
	//coinPrice := fmt.Sprintf("%s has price %.6f \n", c.Name(), c.Price())
    	//playerSavings := fmt.Sprintf("%s has savings balance %.6f \n", t.Name(), t.SavingsBalance())
        //exchangeVolume := fmt.Sprintf("%s has volume %.6f \n", e.Name(), e.Volume())

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	data := []int{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
	spl0 := termui.NewSparkline()
	spl0.Data = data[3:]
	spl0.Title = "Sparkline 0"
	spl0.LineColor = termui.ColorGreen

	spl1 := termui.NewSparkline()
	spl1.Data = data
	spl1.Title = "Sparkline 1"
	spl1.LineColor = termui.ColorRed

	spl2 := termui.NewSparkline()
	spl2.Data = data[5:]
	spl2.Title = "Sparkline 2"
	spl2.LineColor = termui.ColorMagenta

	// group
	spls1 := termui.NewSparklines(spl0, spl1, spl2)
	spls1.Height = 8
	spls1.Width = 20
	spls1.Y = 3
	spls1.BorderLabel = "Group Sparklines"

	// single
        singledata := []int{13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6, 4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10}
	single0 := termui.NewSparkline()
	single0.Data = singledata
	single0.Title = "Sparkline 0"
	single0.LineColor = termui.ColorCyan
		
	singlespl0 := termui.NewSparklines(single0)
	singlespl0.Height = 2
	singlespl0.Width = 20
	singlespl0.Border = false

	termui.Render(spls1, singlespl0)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/i", func(termui.Event) {
		//advance one day
	})

	termui.Loop()

}

func AdvanceOneDay(daysSinceStart int, coinPriceHistory []map[string]float32) {
	//save coin price history for all coins

	//find new price for all coins

	//find new exchange volume for all exchanges

}

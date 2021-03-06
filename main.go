package main

import "fmt"
import "time"
import "math/rand"
import "strings"
import percent "github.com/dariubs/percent"
import coin "btc-trader-sim/coin"
import trader "btc-trader-sim/trader"
import exchange "btc-trader-sim/exchange"
import termui "github.com/gizak/termui"

func main() {
	// set the time to Jan 1st, 2010
	currentTime := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	//for randos
	rand.Seed(time.Now().UTC().UnixNano())
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
	exchangeValueHistory := make(map[string][]int, 8)
	// coin market shares
	coinMarketShares := make(map[string]int)
	//news for the day
	news := ""
	//news History
	newsHistory := make([]string,0)
	//message History
	messageHistory := make([]string,0)
	//index of Selected Coin
	selected := 0
	// current state
	state := 1 // 1 - watching, 2 - buying, 3 - selling
	//last player quantity input
	lastPlayerQuantityInput := 0

	// set up coins
	// 	       {name,              symbol, price, supply,     launchDay}
	c0 := coin.New("Bitcoin",           "BTC",  0.00,      21,         0)
	c1 := coin.New("LightCoin",         "LGC",  0.00,      55,        60)
	c2 := coin.New("Nethereum", 	    "NTH",  0.00,     100,        90)
	c3 := coin.New("Nethereum Vintage", "NTV",  0.00,     100,       120)	
	c4 := coin.New("Riddle",            "XRD",  0.00,   50000,       180)
	c5 := coin.New("Ledge",            "XLG",  0.00,    14000,       270)				
	c6 := coin.New("Bancem",            "BNC",  0.00,     850,       300)
	c7 := coin.New("ZEO",               "ZEO",  0.00,      70,       400)
	c8 := coin.New("YCash",             "YEC",  0.00,       4,       500)
	c9 := coin.New("Interstellar",      "ILM",  0.00,   18000,       550)
	c10 := coin.New("Bitbeets",          "BBT",  0.00,   2000,       800)
	c11 := coin.New("TRAM",             "TRM",  0.00,   70000,       920)
	c12 := coin.New("DigiLink",         "DGL", 0.00,      350,       950)
	c13 := coin.New("XTRAbits",         "XBI",  0.00,     650,       1020)
	
	//add em to our master list
	coins = append(coins, &c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13)
	//select BTC
	selected = 0
	
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
	coinPrices[c10.Name()] = c10.Price()
	coinPrices[c11.Name()] = c11.Price()
	coinPrices[c12.Name()] = c12.Price()
	coinPrices[c13.Name()] = c13.Price()
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
	coinPriceHistory[c10.Name()] = append(arr1, c10.Price())
	coinPriceHistory[c11.Name()] = append(arr1, c11.Price())
	coinPriceHistory[c12.Name()] = append(arr1, c12.Price())
	coinPriceHistory[c13.Name()] = append(arr1, c13.Price())

	// initialize shares of the market - only BTC at first	
	coinMarketShares[c0.Name()] = 100
	coinMarketShares[c1.Name()] = 0
	coinMarketShares[c2.Name()] = 0
	coinMarketShares[c3.Name()] = 0
	coinMarketShares[c4.Name()] = 0
	coinMarketShares[c5.Name()] = 0
	coinMarketShares[c6.Name()] = 0	
	coinMarketShares[c7.Name()] = 0
	coinMarketShares[c8.Name()] = 0
	coinMarketShares[c9.Name()] = 0
	coinMarketShares[c10.Name()] = 0
	coinMarketShares[c11.Name()] = 0
	coinMarketShares[c12.Name()] = 0
	coinMarketShares[c13.Name()] = 0

	//set up exchanges  
	e0 := exchange.New("Mt Ganx",   4,    25,   0)
	e1 := exchange.New("GDOX",      0,   125,  75)
	e2 := exchange.New("BitSaurus", 0,   500, 120)
	e3 := exchange.New("CoinHQ",    0,  1000, 350)
	e4 := exchange.New("Czinance",  0,   900, 500)
	e5 := exchange.New("Napoleox",  0,   750, 600)
	e6 := exchange.New("YoCoin",    0,   350, 660)
	e7 := exchange.New("CoinHawk",  0,   300, 800)	

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

	initialCoinBalancesMap := make(map[string]int,14)
	initialCoinBalancesMap[c0.Name()] = 0
	initialCoinBalancesMap[c1.Name()] = 0
	initialCoinBalancesMap[c2.Name()] = 0
	initialCoinBalancesMap[c3.Name()] = 0
	initialCoinBalancesMap[c4.Name()] = 0
	initialCoinBalancesMap[c5.Name()] = 0
	initialCoinBalancesMap[c6.Name()] = 0
	initialCoinBalancesMap[c7.Name()] = 0
	initialCoinBalancesMap[c8.Name()] = 0
	initialCoinBalancesMap[c9.Name()] = 0
	initialCoinBalancesMap[c10.Name()] = 0
	initialCoinBalancesMap[c11.Name()] = 0
	initialCoinBalancesMap[c12.Name()] = 0
	initialCoinBalancesMap[c13.Name()] = 0
  	t := trader.New("Trader", 100.00, initialCoinBalancesMap)
    
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

	//if in buying state, player gains 1 of selected coin
	//	player loses savings equal to value of 1 coin
	//if in selling state, player loses 1 of selected coin
	//	player gains savings equal to value of 1 coin
	termui.Handle("/sys/kbd/1", func(termui.Event) {
		lastPlayerQuantityInput = 1		
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/2", func(termui.Event) {
		lastPlayerQuantityInput = 2
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/3", func(termui.Event) {
		lastPlayerQuantityInput = 3
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/4", func(termui.Event) {
		lastPlayerQuantityInput = 4
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/5", func(termui.Event) {
		lastPlayerQuantityInput = 5
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/6", func(termui.Event) {
		lastPlayerQuantityInput = 6
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/7", func(termui.Event) {
		lastPlayerQuantityInput = 7
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/8", func(termui.Event) {
		lastPlayerQuantityInput = 8
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})

	termui.Handle("/sys/kbd/9", func(termui.Event) {
		lastPlayerQuantityInput = 9
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})		

	termui.Handle("/sys/kbd/b", func(termui.Event) {
		//player buy
		state = 2
	})
	
	termui.Handle("/sys/kbd/s", func(termui.Event) {
		//player sell
		state = 3
	})

	termui.Handle("/sys/kbd/x", func(termui.Event) {
		//player back to normal
		state = 1
	})

	termui.Handle("/sys/kbd/k", func(termui.Event) {
		if(selected > 0) {		
			selected--
		}
	})

	termui.Handle("/sys/kbd/m", func(termui.Event) {
		if(selected < 13) {		
			selected++
		}
	})

	termui.Handle("/sys/kbd/,", func(termui.Event) {
		if(selected > 0) {		
			selected--
		}
	})

	termui.Handle("/sys/kbd/.", func(termui.Event) {
		if(selected < 13) {		
			selected++
		}
	})
	
	//"/sys/kbd/g"
	termui.Handle("/timer/1s", func(termui.Event) {
		currentTime = currentTime.Add(time.Hour * 24 * 1)
		dayCounter++

		sixSidedDie := random(1,7)
		m := int(currentTime.Month())		
		if( m == 8 || m == 9 || m == 10 || m == 2 || m == 3 || m == 4) {
			// make trend sticky
			if( marketTrend == 3 || marketTrend == 4) {
				sixSidedDie = sixSidedDie - 1
				if (strings.Compare(news, "") != 0) {
					sixSidedDie = sixSidedDie + 1
				}
			}	
	
			if (sixSidedDie <= 2) {
				marketTrend = 4
			} else if (sixSidedDie == 3 || sixSidedDie == 4) {
				marketTrend = 3
			} else if (sixSidedDie == 5) {
				marketTrend = 2
			} else {
				marketTrend = 1
			}	
		} else { 
			// make trend sticky
			if( marketTrend == 1 || marketTrend == 2) {
				sixSidedDie = sixSidedDie - 1
				if (strings.Compare(news, "") != 0) {
					sixSidedDie = sixSidedDie - 1
				}		
			}	

			if (sixSidedDie <= 2) {
					marketTrend = 1
			} else if (sixSidedDie == 3 || sixSidedDie == 4) {
				marketTrend = 2
			} else if (sixSidedDie == 5) {
				marketTrend = 3
			} else {
				marketTrend = 4
			}	
		} 
		
		totalMarketCap = AdvanceOneDay(coins, exchanges, coinPrices, exchangeValues, coinPriceHistory, exchangeValueHistory, 							    coinMarketShares, dayCounter, marketTrend, news, &newsHistory, &t)
		
		//selected coin info	
		selectedInfo := termui.NewPar("")
		if(state == 1) {		
			selectedInfo = termui.NewPar(SelectedCoinTextState1(coins, selected))
		} else if(state == 2) {
			selectedInfo = termui.NewPar(SelectedCoinTextState2(coins, selected))
		} else {
			selectedInfo = termui.NewPar(SelectedCoinTextState3(coins, selected))
		}
		selectedInfo.Height = 8
		selectedInfo.Width = 21
		selectedInfo.X = 30
		selectedInfo.Y = 16
		selectedInfo.BorderLabel = ""
		selectedInfo.BorderFg = termui.ColorCyan
		selectedInfo.TextFgColor = termui.ColorGreen

		//trader info
		traderInfo := termui.NewPar("")		
		traderInfo = termui.NewPar(TraderInfoText(t, coins))
		traderInfo.Height = 18
		traderInfo.Width = 27
		traderInfo.X = 51
		traderInfo.Y = 10
		traderInfo.BorderLabel = fmt.Sprintf("Balances -- %s", currentTime.Format("01-02-2006"))
		traderInfo.BorderFg = termui.ColorCyan
		traderInfo.TextFgColor = termui.ColorWhite			
	
		//savings history
		savingsHistory := termui.NewLineChart()
		savingsHistory.BorderLabel = "Savings"
		savingsHistory.Mode = "dot"
		if(dayCounter < 27) {
			savingsHistory.Data = t.SavingsBalanceHistory()[:dayCounter]
		} else {
			savingsHistory.Data = t.SavingsBalanceHistory()[dayCounter-27:dayCounter]
		}
		savingsHistory.Width = 27	
		savingsHistory.Height = 6
		savingsHistory.X = 51
		savingsHistory.Y = 28
		savingsHistory.DotStyle = '.'
		savingsHistory.AxesColor = termui.ColorCyan
		savingsHistory.LineColor = termui.ColorBlue | termui.AttrBold

		//Coin Worth $
		coinWorth := termui.NewLineChart()
		coinWorth.BorderLabel = "Crypto Net Worth $"
		if(dayCounter < 21) {
			coinWorth.Data = GetHistoricTotalMarketCapAsFloatArray(exchangeValueHistory)[:dayCounter]
		} else {
			coinWorth.Data = GetTraderDollarValueForAllCoins(t, coinPriceHistory, dayCounter)[dayCounter-21:dayCounter]			
		}
		coinWorth.Width = 21
		coinWorth.Height = 10
		coinWorth.X = 30
		coinWorth.Y = 24
		coinWorth.AxesColor = termui.ColorWhite
		coinWorth.LineColor = termui.ColorYellow

		// Short term dollar amounts, or estimate of day until launch
		shorttermhisttitle0 := ShortTermCoinTitle(coins[0], dayCounter, 0, selected)
		shorttermhist0 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist0.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitcoin", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist0.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitcoin", coinPriceHistory)[dayCounter-30:dayCounter])
		}			
		shorttermhist0.Title = shorttermhisttitle0
		shorttermhist0.LineColor = termui.ColorCyan

		shorttermhisttitle1 := ShortTermCoinTitle(coins[1], dayCounter, 1, selected)
		shorttermhist1 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist1.Data = FloatToInts(GetHistoricPriceDataForCoin("LightCoin", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist1.Data = FloatToInts(GetHistoricPriceDataForCoin("LightCoin", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist1.Title = shorttermhisttitle1
		shorttermhist1.LineColor = termui.ColorGreen

		shorttermhisttitle2 := ShortTermCoinTitle(coins[2], dayCounter, 2, selected)
		shorttermhist2 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist2.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist2.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist2.Title = shorttermhisttitle2
		shorttermhist2.LineColor = termui.ColorMagenta

		shorttermhisttitle3 := ShortTermCoinTitle(coins[3], dayCounter, 3, selected)
		shorttermhist3 := termui.NewSparkline()	
		if(dayCounter < 30) {
			shorttermhist3.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum Vintage", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist3.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum Vintage", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist3.Title = shorttermhisttitle3
		shorttermhist3.LineColor = termui.ColorCyan

		shorttermhisttitle4 := ShortTermCoinTitle(coins[4], dayCounter, 4, selected)
		shorttermhist4 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist4.Data = FloatToInts(GetHistoricPriceDataForCoin("Riddle", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist4.Data = FloatToInts(GetHistoricPriceDataForCoin("Riddle", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist4.Title = shorttermhisttitle4
		shorttermhist4.LineColor = termui.ColorGreen
	
		shorttermhisttitle5 := ShortTermCoinTitle(coins[5], dayCounter, 5, selected)
		shorttermhist5 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist5.Data = FloatToInts(GetHistoricPriceDataForCoin("Ledge", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist5.Data = FloatToInts(GetHistoricPriceDataForCoin("Ledge", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist5.Title = shorttermhisttitle5
		shorttermhist5.LineColor = termui.ColorMagenta
	
		shorttermhisttitle6 := ShortTermCoinTitle(coins[6], dayCounter, 6, selected)
		shorttermhist6 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist6.Data = FloatToInts(GetHistoricPriceDataForCoin("Bancem", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist6.Data = FloatToInts(GetHistoricPriceDataForCoin("Bancem", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist6.Title = shorttermhisttitle6
		shorttermhist6.LineColor = termui.ColorCyan
	
		shorttermhisttitle7 := ShortTermCoinTitle(coins[7], dayCounter, 7, selected)
		shorttermhist7 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist7.Data = FloatToInts(GetHistoricPriceDataForCoin("ZEO", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist7.Data = FloatToInts(GetHistoricPriceDataForCoin("ZEO", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist7.Title = shorttermhisttitle7
		shorttermhist7.LineColor = termui.ColorGreen

		shorttermhisttitle8 := ShortTermCoinTitle(coins[8], dayCounter, 8, selected)
		shorttermhist8 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist8.Data = FloatToInts(GetHistoricPriceDataForCoin("YCash", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist8.Data = FloatToInts(GetHistoricPriceDataForCoin("YCash", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist8.Title = shorttermhisttitle8
		shorttermhist8.LineColor = termui.ColorMagenta

		shorttermhisttitle9 := ShortTermCoinTitle(coins[9], dayCounter, 9, selected)
		shorttermhist9 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist9.Data = FloatToInts(GetHistoricPriceDataForCoin("Interstellar", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist9.Data = FloatToInts(GetHistoricPriceDataForCoin("Interstellar", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist9.Title = shorttermhisttitle9
		shorttermhist9.LineColor = termui.ColorCyan

		shorttermhisttitle10 := ShortTermCoinTitle(coins[10], dayCounter, 10, selected)
		shorttermhist10 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist10.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitbeets", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist10.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitbeets", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist10.Title = shorttermhisttitle10
		shorttermhist10.LineColor = termui.ColorGreen

		shorttermhisttitle11 := ShortTermCoinTitle(coins[11], dayCounter, 11, selected)
		shorttermhist11 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist11.Data = FloatToInts(GetHistoricPriceDataForCoin("TRAM", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist11.Data = FloatToInts(GetHistoricPriceDataForCoin("TRAM", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist11.Title = shorttermhisttitle11
		shorttermhist11.LineColor = termui.ColorMagenta
		
		shorttermhisttitle12 := ShortTermCoinTitle(coins[12], dayCounter, 12, selected)
		shorttermhist12 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist12.Data = FloatToInts(GetHistoricPriceDataForCoin("DigiLink", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist12.Data = FloatToInts(GetHistoricPriceDataForCoin("DigiLink", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist12.Title = shorttermhisttitle12
		shorttermhist12.LineColor = termui.ColorCyan

		shorttermhisttitle13 := ShortTermCoinTitle(coins[13], dayCounter, 13, selected)
		shorttermhist13 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist13.Data = FloatToInts(GetHistoricPriceDataForCoin("XTRAbits", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist13.Data = FloatToInts(GetHistoricPriceDataForCoin("XTRAbits", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist13.Title = shorttermhisttitle13
		shorttermhist13.LineColor = termui.ColorGreen

		// put them together
		shorttermhistograms := termui.NewSparklines(shorttermhist0, shorttermhist1, shorttermhist2, 
					shorttermhist3, shorttermhist4, shorttermhist5,
					shorttermhist6, shorttermhist7, shorttermhist8,
					shorttermhist9, shorttermhist10, shorttermhist11,
					shorttermhist12, shorttermhist13)
		shorttermhistograms.Height = 30
		shorttermhistograms.Width = 30
		shorttermhistograms.Y = 4
		shorttermhistograms.X = 0
		shorttermhistograms.BorderLabel = "Coin - supply - price"
		
		//List of exchanges - presented as gauges of total cap
		exchangeGauge0 := termui.NewGauge()
		exchangeGauge0.Percent = ExchangeInfoPercent(exchanges[0])
		exchangeGauge0.Width = 28
		exchangeGauge0.Height = 3
		exchangeGauge0.Y = 10
		exchangeGauge0.X = 78
		exchangeGauge0.BorderLabel = ExchangeInfoString(exchanges[0], dayCounter)
		exchangeGauge0.Label = ExchangeInfoLabel(exchanges[0])
		exchangeGauge0.BarColor = termui.ColorMagenta
		exchangeGauge0.BorderFg = termui.ColorWhite
		exchangeGauge0.LabelAlign = termui.AlignRight

		exchangeGauge1 := termui.NewGauge()
		exchangeGauge1.Percent = ExchangeInfoPercent(exchanges[1])
		exchangeGauge1.Width = 28
		exchangeGauge1.Height = 3
		exchangeGauge1.Y = 13
		exchangeGauge1.X = 78
		exchangeGauge1.BorderLabel = ExchangeInfoString(exchanges[1], dayCounter)
		exchangeGauge1.Label = ExchangeInfoLabel(exchanges[1])
		exchangeGauge1.BarColor = termui.ColorMagenta
		exchangeGauge1.BorderFg = termui.ColorWhite
		exchangeGauge1.LabelAlign = termui.AlignRight
		
		exchangeGauge2 := termui.NewGauge()
		exchangeGauge2.Percent = ExchangeInfoPercent(exchanges[2])
		exchangeGauge2.Width = 28
		exchangeGauge2.Height = 3
		exchangeGauge2.Y = 16
		exchangeGauge2.X = 78
		exchangeGauge2.BorderLabel = ExchangeInfoString(exchanges[2], dayCounter)
		exchangeGauge2.Label = ExchangeInfoLabel(exchanges[2])
		exchangeGauge2.BarColor = termui.ColorMagenta
		exchangeGauge2.BorderFg = termui.ColorWhite
		exchangeGauge2.LabelAlign = termui.AlignRight

		exchangeGauge3 := termui.NewGauge()
		exchangeGauge3.Percent = ExchangeInfoPercent(exchanges[3])
		exchangeGauge3.Width = 28
		exchangeGauge3.Height = 3
		exchangeGauge3.Y = 19
		exchangeGauge3.X = 78
		exchangeGauge3.BorderLabel = ExchangeInfoString(exchanges[3], dayCounter)
		exchangeGauge3.Label = ExchangeInfoLabel(exchanges[3])
		exchangeGauge3.BarColor = termui.ColorMagenta
		exchangeGauge3.BorderFg = termui.ColorWhite
		exchangeGauge3.LabelAlign = termui.AlignRight

		exchangeGauge4 := termui.NewGauge()
		exchangeGauge4.Percent = ExchangeInfoPercent(exchanges[4])
		exchangeGauge4.Width = 28
		exchangeGauge4.Height = 3
		exchangeGauge4.Y = 22
		exchangeGauge4.X = 78
		exchangeGauge4.BorderLabel = ExchangeInfoString(exchanges[4], dayCounter)
		exchangeGauge4.Label = ExchangeInfoLabel(exchanges[4])
		exchangeGauge4.BarColor = termui.ColorMagenta
		exchangeGauge4.BorderFg = termui.ColorWhite
		exchangeGauge4.LabelAlign = termui.AlignRight

		exchangeGauge5 := termui.NewGauge()
		exchangeGauge5.Percent = ExchangeInfoPercent(exchanges[5])
		exchangeGauge5.Width = 28
		exchangeGauge5.Height = 3
		exchangeGauge5.Y = 25
		exchangeGauge5.X = 78
		exchangeGauge5.BorderLabel = ExchangeInfoString(exchanges[5], dayCounter)
		exchangeGauge5.Label = ExchangeInfoLabel(exchanges[5])
		exchangeGauge5.BarColor = termui.ColorMagenta
		exchangeGauge5.BorderFg = termui.ColorWhite
		exchangeGauge5.LabelAlign = termui.AlignRight

		exchangeGauge6 := termui.NewGauge()
		exchangeGauge6.Percent = ExchangeInfoPercent(exchanges[6])
		exchangeGauge6.Width = 28
		exchangeGauge6.Height = 3
		exchangeGauge6.Y = 28
		exchangeGauge6.X = 78
		exchangeGauge6.BorderLabel = ExchangeInfoString(exchanges[6], dayCounter)
		exchangeGauge6.Label = ExchangeInfoLabel(exchanges[6])
		exchangeGauge6.BarColor = termui.ColorMagenta
		exchangeGauge6.BorderFg = termui.ColorWhite
		exchangeGauge6.LabelAlign = termui.AlignRight

		exchangeGauge7 := termui.NewGauge()
		exchangeGauge7.Percent = ExchangeInfoPercent(exchanges[7])
		exchangeGauge7.Width = 28
		exchangeGauge7.Height = 3
		exchangeGauge7.Y = 31
		exchangeGauge7.X = 78
		exchangeGauge7.BorderLabel = ExchangeInfoString(exchanges[7], dayCounter)
		exchangeGauge7.Label = ExchangeInfoLabel(exchanges[7])
		exchangeGauge7.BarColor = termui.ColorMagenta
		exchangeGauge7.BorderFg = termui.ColorWhite
		exchangeGauge7.LabelAlign = termui.AlignRight

		marketCap := termui.NewLineChart()
		marketCap.BorderLabel = MarketCapInfoString(totalMarketCap)
		marketCap.Mode = "dot"
		marketCapWindow := make([]float64, 30)
		if(dayCounter < 28) {
			marketCapWindow = GetHistoricTotalMarketCapAsFloatArray(exchangeValueHistory)[:dayCounter]
		} else {
			marketCapWindow = GetHistoricTotalMarketCapAsFloatArray(exchangeValueHistory)[dayCounter-28:dayCounter]			
		}
		marketCap.Data = marketCapWindow	
		marketCap.Width = 28
		marketCap.Height = 10
		marketCap.X = 78
		marketCap.Y = 0
		marketCap.DotStyle = '.'
		marketCap.AxesColor = termui.ColorWhite
		marketCap.LineColor = termui.ColorGreen | termui.AttrBold

		marketShares := termui.NewBarChart()
		data := []int{MarketShareForCoin(coinMarketShares,coins[0]), MarketShareForCoin(coinMarketShares,coins[1]), 
				MarketShareForCoin(coinMarketShares,coins[2]),MarketShareForCoin(coinMarketShares,coins[3]),
				 MarketShareForCoin(coinMarketShares,coins[4]),MarketShareForCoin(coinMarketShares,coins[5]),
					MarketShareForCoin(coinMarketShares,coins[6]), MarketShareForCoin(coinMarketShares,coins[7]), 						MarketShareForCoin(coinMarketShares,coins[8]),MarketShareForCoin(coinMarketShares,coins[9]),
					 MarketShareForCoin(coinMarketShares,coins[10]), MarketShareForCoin(coinMarketShares,coins[11]),
					MarketShareForCoin(coinMarketShares,coins[12]), MarketShareForCoin(coinMarketShares,coins[13])}
		labels := []string{coins[0].Symbol(), coins[1].Symbol(), coins[2].Symbol(), coins[3].Symbol(),
					coins[4].Symbol(), coins[5].Symbol(), coins[6].Symbol(), coins[7].Symbol(),
					coins[8].Symbol(), coins[9].Symbol(), coins[10].Symbol(), coins[11].Symbol(),
					coins[12].Symbol(), coins[13].Symbol()}
		marketShares.BorderLabel = "Market Share"
		marketShares.Data = data
		marketShares.Width = 64
		marketShares.Height = 4
		marketShares.X=0
		marketShares.Y=0
		marketShares.DataLabels = labels
		marketShares.TextColor = termui.ColorWhite
		marketShares.BarColor = termui.ColorBlue
		marketShares.NumColor = termui.ColorWhite
		
		//news
		recentNews := termui.NewList()	
		if(len(newsHistory) == 0) {		
			recentNews.Items = make([]string,0)
		} else if(len(newsHistory) < 4) {
			recentNews.Items = newsHistory[:len(newsHistory)]
		} else {
			recentNews.Items = newsHistory[len(newsHistory)-4:len(newsHistory)]
		}
		recentNews.ItemFgColor = termui.ColorWhite
		recentNews.BorderLabel = fmt.Sprintf("Latest News")
		recentNews.Height = 6
		recentNews.Width = 48
		recentNews.Y = 4
		recentNews.X = 30

		//messages for player
		messages := termui.NewList()	
		if(len(messageHistory) == 0) {		
			recentNews.Items = make([]string,0)
		} else if(len(messageHistory) < 4) {
			messages.Items = messageHistory[:len(messageHistory)]
		} else {
			messages.Items = messageHistory[len(messageHistory)-4:len(messageHistory)]
		}
		messages.ItemFgColor = termui.ColorCyan
		messages.BorderLabel = fmt.Sprintf("Messages")
		messages.Height = 6
		messages.Width = 21
		messages.Y = 10
		messages.X = 30

		//market sentiment
		sentiment := termui.NewPar(SentimentString(marketTrend))
		sentiment.Height = 4
		sentiment.Width = 14
		sentiment.X = 64
		sentiment.Y = 0
		sentiment.BorderLabel = "Sentiment"
		sentiment.BorderFg = termui.ColorCyan
		if(marketTrend == 1 || marketTrend == 2) {
			sentiment.TextFgColor = termui.ColorGreen
		} else {
			sentiment.TextFgColor = termui.ColorRed
		}


		debug := termui.NewPar(fmt.Sprintf(" day %d", dayCounter))
		debug.Height = 1
		debug.Width = 20
		debug.X = 34
		debug.Y = 30
		debug.Border = false		
	
	termui.Render( shorttermhistograms, debug, exchangeGauge0, exchangeGauge1, exchangeGauge2, exchangeGauge3,
				exchangeGauge4, exchangeGauge5, exchangeGauge6, exchangeGauge7, marketCap, marketShares, 
				recentNews, sentiment, selectedInfo, traderInfo, messages, savingsHistory, coinWorth)
	})

	termui.Loop()

}

func TraderInfoText(t trader.Trader, coins []*coin.Coin) string {
	traderInfo := ""
	traderInfo = traderInfo + fmt.Sprintf("Cash: $%.2f", t.SavingsBalance())
	traderInfo = traderInfo + "\n---"
	for _, c := range coins {
		traderInfo = traderInfo + fmt.Sprintf("\n%s: %d",c.Name(),t.BalanceForCoin(c.Name()))
	}
	return traderInfo
}

func SelectedCoinTextState1(coins []*coin.Coin, selected int) string {
	//show selected coin name, b to buy, s to sell
	return fmt.Sprintf("%s (%s)\n[b] to buy\n[s] to sell\n\n[k,<]select up\n[m,>]select down",coins[selected].Name(), coins[selected].Symbol())
}

func SelectedCoinTextState2(coins []*coin.Coin, selected int) string {
	//show name, quantity, enter quantity to buy, 1-9
	return fmt.Sprintf("%s (%s)\nBUY\nEnter [1-9]\n\n[x] to exit\n[<,>] to scroll coins",coins[selected].Name(), coins[selected].Symbol())
}

func SelectedCoinTextState3(coins []*coin.Coin, selected int) string {
	//show name, quantity, enter quantity to sell, 1-9
	return fmt.Sprintf("%s (%s)\nSELL\nEnter [1-9]\n\n[x] to exit\n[<,>] to scroll coins",coins[selected].Name(), coins[selected].Symbol())
}

func GenerateTweets(coins []*coin.Coin) map[string]int {
	coinDailyTweets := make(map[string]int)	
	
	for i, _ := range coins {
		numberOfTweets := random(1,8)
		coinDailyTweets[coins[i].Name()] = numberOfTweets 
	}
	
	return coinDailyTweets
}

func GenerateNews(coins []*coin.Coin, dayCounter int) string {
	chanceOfNews := random(1,10) // 1/3 chance of news for a random coin
	randomIdx := (random(1,15)-1)
	news := ""	
	
	if(chanceOfNews % 2 == 0 || randomIdx == 0) { //BTC always gets its news
		if(coins[randomIdx].LaunchDay() < dayCounter) {
			eventStringRandomizer := random(1,7)
			if(eventStringRandomizer == 1) {
				news = fmt.Sprintf("%s had a wallet release today", coins[randomIdx].Name())
			} else if(eventStringRandomizer == 2) {
				news = fmt.Sprintf("%s team spoke at a conference this week", coins[randomIdx].Name())
			} else if(eventStringRandomizer == 3) {
				news = fmt.Sprintf("%s announced a new partnership", coins[randomIdx].Name())
			} else if(eventStringRandomizer == 4) {
				news = fmt.Sprintf("%s released a developer update", coins[randomIdx].Name())
			} else if(eventStringRandomizer == 5) {
				news = fmt.Sprintf("%s founder appeared on CABC", coins[randomIdx].Name())
			} else if(eventStringRandomizer == 6) {
				news = fmt.Sprintf("%s was tweeted by an e-celeb", coins[randomIdx].Name())
			}
		}
	}

	return news
}

func AdvanceOneDay(coins []*coin.Coin, exchanges []*exchange.Exchange, coinPrices map[string]float64, exchangeValues map[string]int, 			coinPriceHistory map[string][]float64, exchangeValueHistory map[string][]int, coinMarketShares map[string]int, dayCounter int, 			marketTrend int, news string, newsHistory *[]string, t *trader.Trader) int {
	//compute totalMarketCap
	totalCap := 0
	for _, e := range exchanges {
		totalCap = totalCap + e.ValueAdded()
	}
	
	//save today's exchange valueAdded for all exchanges
	//find new exchange valueAdded for all exchanges
	for _, e := range exchanges {
	    if(e.LaunchDay() > dayCounter) {
		if(e.LaunchDay() >= dayCounter+10) {
	    		e.DailyLaunchAdjustment(marketTrend)
		}
	    } else {
	        currentValueHistory := exchangeValueHistory[e.Name()]
	        delete(exchangeValueHistory, e.Name())
	        exchangeValueHistory[e.Name()] = append(currentValueHistory, exchangeValues[e.Name()])
	        exchangeValues[e.Name()] = e.DailyValueAdjustment(totalCap, marketTrend)
	    }
	}			

	// save coin price history for all coins
	// and find next day's value
	for _, c := range coins {
	    if(c.LaunchDay() > dayCounter) {
		if(c.LaunchDay() >= dayCounter+10) {
	    		c.DailyLaunchAdjustment(marketTrend)
		}
		currentPriceHistory := coinPriceHistory[c.Name()]
	    	delete(coinPriceHistory, c.Name())
	    	coinPriceHistory[c.Name()] = append(currentPriceHistory, coinPrices[c.Name()])
	    } else if(c.LaunchDay() == dayCounter) {	
		icoShare := random(1,5)
		if(coinMarketShares["Bitcoin"] > 60) {
			//take icoShare from BTC
			newBtcShares := coinMarketShares["Bitcoin"] - icoShare
			delete(coinMarketShares, "Bitcoin")
	        	coinMarketShares["Bitcoin"] = newBtcShares
			delete(coinMarketShares, c.Name())
			coinMarketShares[c.Name()] = icoShare
			
			//adjust price and records
			var capShare float64 =  float64(float64(totalCap)*float64(1000)) /* how many millions */ / 
					float64(MarketShareForCoin(coinMarketShares, c)/100)
			price := capShare / float64(c.Supply())
	    		currentPriceHistory := coinPriceHistory[c.Name()]
	    		delete(coinPriceHistory, c.Name())
	    		coinPriceHistory[c.Name()] = append(currentPriceHistory, coinPrices[c.Name()])
	    		coinPrices[c.Name()] = c.DailyPriceAdjustment(price)
		} else {
			sharesToGive := 0
			for _, c2 := range coins {
				sharesToTake := random(1,5)
				rollForDodge := random(1,4)
				if(c.LaunchDay() < dayCounter && rollForDodge != 1) {	
					if(coinMarketShares[c2.Name()] > 5) {		
						previousShares := coinMarketShares[c2.Name()]
						delete(coinMarketShares, c2.Name())
						coinMarketShares[c2.Name()] = previousShares - sharesToTake		
						sharesToGive = sharesToGive + sharesToTake
					}
				}

			}

			// if we came out of that with 0 shares to give just take 10 or 1 from btc its buffed enuff
			if(sharesToGive == 0) {
				if(coinMarketShares["Bitcoin"] > 24) {
					newBtcShares := coinMarketShares["Bitcoin"] - 10
					delete(coinMarketShares, "Bitcoin")
	        			coinMarketShares["Bitcoin"] = newBtcShares
					sharesToGive = 10
				} else {
					newBtcShares := coinMarketShares["Bitcoin"] - 1
					delete(coinMarketShares, "Bitcoin")
	        			coinMarketShares["Bitcoin"] = newBtcShares
					sharesToGive = 1
				}
			}
		
			prevShares := coinMarketShares[c.Name()]
			delete(coinMarketShares, c.Name())		
			coinMarketShares[c.Name()] = prevShares + sharesToGive
	 	}
	    } else {
		//adjust price and records
		var capShare float64 =  float64(float64(totalCap)*float64(1000)) /* how many millions */ * 
					float64(float64(MarketShareForCoin(coinMarketShares, c))/float64(100))
		price := capShare / float64(c.Supply())
		if(price == 0.0) {
			price = 0.01
		}
	    	currentPriceHistory := coinPriceHistory[c.Name()]
	    	delete(coinPriceHistory, c.Name())
	    	coinPriceHistory[c.Name()] = append(currentPriceHistory, coinPrices[c.Name()])
	    	coinPrices[c.Name()] = c.DailyPriceAdjustment(price)
	    }
	}	

	//make and save today's news
	news = GenerateNews(coins, dayCounter)
	if(strings.Compare(news,"") != 0) {
		*newsHistory = append(*newsHistory, news)
	}

	//shuffle shares		
	coinMarketShares = ShuffleMarketShare(coinMarketShares, coins, news, GenerateTweets(coins), dayCounter, marketTrend)
	
	//record trader balances
	t.RecordBalances(dayCounter)

	return totalCap
}

func ShuffleMarketShare(coinMarketShares map[string]int, coins []*coin.Coin, news string, coinDailyTweets map[string]int, dayCounter int, marketTrend int) map[string]int {
	//take 1 from every coin's share, leaves us with n share to distribute, n is number of coins that are past their launch day	
	// also find which coin got news
	sharesToGive := 0	
	newsworthy := ""
	
	for _, c := range coins {
		if(strings.Contains(news, c.Name())) {
			newsworthy = c.Name()			
		}
	}

	//!! If the market is very bearish, bitcoin will get the boost, but the news will still be reported for the other coin
	if(marketTrend == 4 && strings.Compare(newsworthy,"") != 0) {
		newsworthy = "Bitcoin"
	}

	// if a coin got news take 1-3 shares from every coin with shares over 4 and give it to the coin with news
	if(strings.Compare(newsworthy,"") != 0) {	
		// take two off btc and give them away if btc is still packin
		if(coinMarketShares["Bitcoin"] > 20) {
			newBtcShares := coinMarketShares["Bitcoin"] - 2
			delete(coinMarketShares, "Bitcoin")
		        coinMarketShares["Bitcoin"] = newBtcShares
			sharesToGive = sharesToGive + 2 
		}
	
		for _, c := range coins {
			sharesToTake := random(1,4)
			rollForDodge := random(1,4)
			if(c.LaunchDay() < dayCounter && rollForDodge != 1) {	
				if(coinMarketShares[c.Name()] > 4) {		
					previousShares := coinMarketShares[c.Name()]
					delete(coinMarketShares, c.Name())
					coinMarketShares[c.Name()] = previousShares - sharesToTake		
					sharesToGive = sharesToGive + sharesToTake
				}
			}
		}
	}

	//for now, give it to the coin who got news
	prevShares := coinMarketShares[newsworthy]
	delete(coinMarketShares, newsworthy)		
	coinMarketShares[newsworthy] = prevShares + sharesToGive

	return coinMarketShares
}



func MarketCapInfoString(totalMarketCap int) string {
	maxCap := 4000
	marketCapInfo := fmt.Sprintf("$%dB/$%dB", totalMarketCap, maxCap)
	return marketCapInfo
}

func MarketShareForCoin(coinMarketShares map[string]int, coin *coin.Coin) int {
	return coinMarketShares[coin.Name()]
}

func SentimentString(marketTrend int) string {
	if(marketTrend == 1) { 
		return "BULLISH+"
	} else if(marketTrend == 2) {
		return "BULLISH"
	} else if(marketTrend == 3) {
		return "BEARISH"
	} else {
		return "BEARISH-"
	}
}

func ExchangeInfoString(exchange *exchange.Exchange, dayCounter int) string {
	exchangeInfo := ""

	if(exchange.LaunchDay() > dayCounter) {
		exchangeInfo = fmt.Sprintf("%s ETA %d days", exchange.Name(), (exchange.LaunchDay() - dayCounter))
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

func ShortTermCoinTitle(coin *coin.Coin, dayCounter int, coinsIdx int, selected int) string {
	title := ""
	cursor := ""
	
	if(coinsIdx == selected) {
		cursor = ">> "
	}

	if(coin.LaunchDay() > dayCounter) {
		title = fmt.Sprintf("%s%s - ETA %d days", cursor, coin.Symbol(), (coin.LaunchDay() - dayCounter))
		
	} else {
		if(coin.Supply() > 1000) {
			title = fmt.Sprintf("%s%s - %d Bil - $%.2f", cursor, coin.Symbol(), (coin.Supply()/1000), coin.Price())
		} else {
			title = fmt.Sprintf("%s%s - %d Mil - $%.2f", cursor, coin.Symbol(), coin.Supply(), coin.Price())
		}
	}
	
	return title
}

func GetHistoricTotalMarketCapAsFloatArray(exchangeValueHistory map[string][]int) []float64 {
	totalMarketCap := make([]float64, 5900)

	keys := make([]string, 0, len(exchangeValueHistory))
	for k := range exchangeValueHistory {
    	        keys = append(keys, k)
    	}

    	for _, k := range keys {
       		oneExchangeHistoricValue := exchangeValueHistory[k]
		for i, _ := range oneExchangeHistoricValue {
			totalMarketCap[i] = totalMarketCap[i] + float64(oneExchangeHistoricValue[i])
		}
    	}
	
	return totalMarketCap
}

func GetHistoricTotalMarketCapAsIntArray(exchangeValueHistory map[string][]int) []int {
	totalMarketCap := make([]int, 5900)

	keys := make([]string, 0, len(exchangeValueHistory))
	for k := range exchangeValueHistory {
    	        keys = append(keys, k)
    	}

    	for _, k := range keys {
       		oneExchangeHistoricValue := exchangeValueHistory[k]
		for i, _ := range oneExchangeHistoricValue {
			totalMarketCap[i] = totalMarketCap[i] + oneExchangeHistoricValue[i]
		}
    	}
	
	return totalMarketCap
}

func GetHistoricPriceDataForCoin(coin string, coinPriceHistory map[string][]float64) []float64 {
	return coinPriceHistory[coin]
}

func GetTraderDollarValueForCoin(t trader.Trader, coin string, coinPriceHistory map[string][]float64, dayCounter int) []float64 {
	traderBalance := t.HistoricBalancesForCoin(coin)
	traderDollarValue := make([]float64,dayCounter)
	for i := 0; i < dayCounter-1; i++ {
		traderDollarValue[i] = (float64(traderBalance[i]))*(coinPriceHistory[coin][i])
	}

	return traderDollarValue
}

func GetTraderDollarValueForAllCoins(t trader.Trader, coinPriceHistory map[string][]float64, dayCounter int) []float64 {
	sumAllCoins := make([]float64,dayCounter)
	for k, _ := range coinPriceHistory {
		oneCoinHistory := GetTraderDollarValueForCoin(t, k, coinPriceHistory, dayCounter)
		for i := 0; i < dayCounter-1; i++ {
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

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

package main

import "fmt"
import "time"
import "btc-trader-sim/coin"
import "btc-trader-sim/trader"
import "btc-trader-sim/exchange"

func main() {
    currentTime := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
    c := coin.New("bitcoin", 1.00, "FINANCE", "BULLISH")
    t := trader.New("kc", 0.00, 0.00, nil)
    e := exchange.New("gdax", 100.00)
    
    fmt.Printf("%s has price %s \n", c.Name() ,c.Price())
    fmt.Printf("%s has savings balance %s \n", t.Name() ,t.SavingsBalance())
    fmt.Printf("%s has volume %s \n", e.Name() ,e.Volume())
    fmt.Printf("Date : ", currentTime.Format("01-02-2006"))
}

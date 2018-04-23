package main

import "fmt"
import "btc-trader-sim/coin"
import "btc-trader-sim/trader"
import "btc-trader-sim/exchange"

func main() {
    c := coin.New("bitcoin", 1.00, "FINANCE", "BULLISH")
    t := trader.New("kc", 0.00, 0.00, nil)
    e := exchange.New("gdax", 100.00)
    
    fmt.Printf("%s has price %s", c.Name() ,c.Price())
    fmt.Printf("%s has savings balance %s", t.Name() ,t.SavingsBalance())
    fmt.Printf("%s has volume %s", e.Name() ,e.Volume())
}

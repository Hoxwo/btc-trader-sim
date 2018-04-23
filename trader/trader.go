package trader

import "fmt"

type trader struct {
    name string
    savingsBalance float32
    savingsBalanceHistory []float32
    coinBalancesMap map[string]float32
    coinBalancesHistoryMap []map[string]float32
}

func New(name string, savingsBalance float32) trader {
    t := trader {name, savingsBalance, make([]float32, 5900), 
		 make(map[string]float32), make([]map[string]float32, 5900)}
    return t
}

func (t trader) Name() string {
    return t.name
}

func (t trader) SavingsBalance() float32 {
    return t.savingsBalance
}

func (t trader) SavingsBalanceOnDay(daysSinceStart int) float32 {
    return t.savingsBalanceHistory[daysSinceStart]
}

func (t trader) RecordBalances(daysSinceStart int) {
    // record savings balance
    t.savingsBalanceHistory[daysSinceStart] = t.savingsBalance
    
    // record coin balances
    t.coinBalancesHistoryMap[daysSinceStart] = t.coinBalancesMap
}

// coin: coin name
// coinAmount: amount to buy or sell
// dollarAmount: cost in fiat
// op: 1 - BUY, 2 - SELL
func (t trader) ModifyCoinAndSavingsBalance(coin string, coinAmount float32, dollarAmount float32, op int ) string {
  if(op == 1) {
	current := t.coinBalancesMap[coin]
	currentSavings := t.savingsBalance
	if((currentSavings - dollarAmount) < 0.0 || currentSavings == 0.0) {
            return "Can not buy that much - no cash"
	} else {
	    t.coinBalancesMap[coin] = current + coinAmount
	    t.savingsBalance = currentSavings - dollarAmount
            ret := fmt.Sprintf("Buy of %.6f %s successful", coinAmount, coin)
	    return ret
        }
  } else if (op == 2) {
	current := t.coinBalancesMap[coin]
	if((current - coinAmount) < 0.0 || current == 0.0) {
            return "Can not sell that much"
	} else {
	     t.coinBalancesMap[coin] = current - coinAmount
             t.savingsBalance = t.savingsBalance + dollarAmount
	     ret := fmt.Sprintf("Sell of %.6f %s successful", coinAmount, coin)
	     return ret
        }
    } else {
         return "Invalid Operation"
    }
}




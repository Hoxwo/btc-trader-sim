package trader

import "fmt"

type Trader struct {
    name string
    savingsBalance float64
    savingsBalanceHistory []float64
    coinBalancesMap map[string]int
    coinBalancesHistoryMap map[string][]int
}

func New(name string, savingsBalance float64, initialCoinBalancesMap map[string]int) Trader {
    t := Trader {name, savingsBalance, make([]float64,0), 
		 initialCoinBalancesMap, make(map[string][]int,0)}
    return t
}

func (t Trader) Name() string {
    return t.name
}

func (t Trader) SavingsBalance() float64 {
    return t.savingsBalance
}

func (t Trader) SavingsBalanceOnDay(daysSinceStart int) float64 {
    return t.savingsBalanceHistory[daysSinceStart]
}

func (t *Trader) RecordBalances(dayCounter int) {
    // record savings balance
    t.savingsBalanceHistory[dayCounter] = t.savingsBalance
    
    // record coin balances
    for k, _ := range t.coinBalancesMap {
	currentBalancesHistory := t.coinBalancesHistoryMap[k]
	t.coinBalancesHistoryMap[k] = append(currentBalancesHistory, t.coinBalancesMap[k])
    }	
}

func (t Trader) OwnedCoins() []string {
    coins := make([]string, 0)
    for k := range t.coinBalancesMap {
	if(t.coinBalancesMap[k] != 0) {
        	coins = append(coins, k)
	}
    }

    return coins
}

func (t Trader) BalanceForCoin(coin string) int {
	return t.coinBalancesMap[coin]
}


func (t Trader) HistoricBalancesForCoin(coin string) []int {
	return t.coinBalancesHistoryMap[coin]
}

// coin: coin name
// coinAmount: amount to buy or sell
// dollarAmount: cost in fiat
// op: 1 - BUY, 2 - SELL
func (t *Trader) ModifyCoinAndSavingsBalance(coin string, coinAmount int, dollarAmount float64, op int ) string {
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




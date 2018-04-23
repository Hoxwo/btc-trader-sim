package trader

type trader struct {
    name string
    savingsBalance float32
    fiatBalance float32
    coinBalancesMap map[string]float32
}

func New(name string, savingsBalance float32, fiatBalance float32, coinBalancesMap map[string]float32) trader {
    t := trader {name, savingsBalance, fiatBalance, make(map[string]float32)}
    return t
}

func (t trader) Name() string {
    return t.name
}

func (t trader) SavingsBalance() float32 {
    return t.savingsBalance
}

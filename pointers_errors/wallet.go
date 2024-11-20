package pointerserrors

type Wallet struct{}

func (w *Wallet) Deposit(amt int) int {
	return 0
}

func (w *Wallet) Balance() int {
	return 0
}

package pointers

import (
	"errors"
	"fmt"
)

var insufficientFundsError = errors.New("Cannot withdraw insufficient funds")

type Bitcoin int

/*
Implement fmt's interface Stringer by implementing the method String()
type Stringer interface {
        String() string
}
*/

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(bitcoins Bitcoin) {
	w.balance += bitcoins
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(bitcoin Bitcoin) error {
	if bitcoin > w.balance {
		return insufficientFundsError
	}

	w.balance -= bitcoin

	return nil
}
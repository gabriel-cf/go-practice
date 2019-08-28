package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	
	want := Bitcoin(10)

	AssertBalance(t, wallet, want)
}

func TestWithdraw(t *testing.T) {
	t.Run("With enough funds", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(5)

		AssertBalance(t, wallet, want)
		AssertNoError(t, err)
	})

	t.Run("Without enough funds", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(5)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(5)

		AssertBalance(t, wallet, want)	
		AssertError(t, err, insufficientFundsError)
	})
}

func AssertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func AssertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but got none")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got %q but expected none", got)
	}
}
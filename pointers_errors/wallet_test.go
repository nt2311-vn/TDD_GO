package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want Bitcoint) {
		t.Helper()

		got := w.Balance()
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Error("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoint(10))

		want := Bitcoint(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(20)}

		wallet.Withdraw(Bitcoint(10))
		want := Bitcoint(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoint(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoint(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

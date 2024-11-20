package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10
	checkWallet(t, got, want)
}

func checkWallet(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

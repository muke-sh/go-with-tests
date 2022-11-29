package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{10}

		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t, err)

		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {

		startBalance := Bitcoin(10)

		wallet := Wallet{startBalance}

		err := wallet.Withdraw(Bitcoin(20))

		assertBalance(t, wallet, startBalance)

		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {

	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {

	t.Helper()

	if got == nil {
		t.Fatal("want an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {

	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

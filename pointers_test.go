package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Solana(10))

		assertBalance(t, wallet, Solana(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Solana(20)}

		err := wallet.Withdraw(Solana(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Solana(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Solana(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Solana(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Solana) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

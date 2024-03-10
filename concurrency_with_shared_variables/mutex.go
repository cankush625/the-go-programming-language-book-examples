package main

import "sync"

var mu sync.Mutex
var bankBalance int

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if bankBalance < 0 {
		deposit(amount)
		return false // insufficient balance
	}
	return true
}

func DepositMoney(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func GetBalance() int {
	mu.Lock()
	defer mu.Unlock()
	return bankBalance
}

// deposit function requires that the lock be held
func deposit(amount int) {
	bankBalance += amount
}

func main() {

}

package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex //protects concurrent access to the account balance
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()         //lock the account before modifying the shared balance
	defer b.mutex.Unlock() //always release the lock when leaving the function
	b.balance += amount

	fmt.Println("Deposit", amount)
}

func (b *BankAccount) Withdraw(amount int) {
	b.mutex.Lock() //prevent multiple goroutines from updating the balance simultaneously
	defer b.mutex.Unlock()

	if b.balance < amount {
		fmt.Println("cannot withdraw that amount: ", amount)
		return
	}

	b.balance -= amount
	fmt.Println("Withdraw", amount)
}

func (b *BankAccount) Balance() int {
	b.mutex.Lock() //lock even for reads to avoid race conditions
	defer b.mutex.Unlock()

	return b.balance
}

func main() {
	var wg sync.WaitGroup
	var account = &BankAccount{balance: 150} //shared bank account accessed by multiple goroutines

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(amount int) { //launch concurrent deposit operations
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Deposit(amount)
		}(i + 1)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(amount int) { //launch concurrent withdrawl operations
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Withdraw(amount * 10)
		}(i + 1)
	}

	wg.Wait() //wait until every transaction completes
	fmt.Println(account.Balance())
}

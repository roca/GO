package application

import (
	"fmt"
	"grpc-go-server/data"
	"time"

	"github.com/google/uuid"
)

type BankService struct {
	Models data.Models
}

func (b *BankService) Save(data data.BankAccount) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (b *BankService) FindCurrentBalance(uuid string) float64 {
	// log.Println(acct)

	account, _ := b.Models.BankAccounts.Get(uuid)

	return account.CurrentBalance
}

// Create a function to inject dummy exchange rates to table bank_exchange_rates. The function should run and inject new exchange rate data every X seconds.
func (b *BankService) InsertExchangeRatesAtInterval(exit chan bool, interval time.Duration) {
	//er := &data.BankExchangeRate{}
	insertFunc := func() {
		//b.Models.BankExchangeRates.Insert(*er)
		fmt.Println("# b.Models.BankExchangeRates.Insert(m BankExchangeRate)")
	}
	stop := runFuncAtInterval(insertFunc, interval)
	<-exit
	fmt.Println("Finished InsertExchangeRatesAtInterval")
	stop <- true
}

func runFuncAtInterval(f func(), seconds time.Duration) chan bool {
	ticker := time.NewTicker(seconds)
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			case <-stop:
				ticker.Stop()
				fmt.Println("Finished runFuncAtInterval")
				return
			}
		}
	}()
	return stop
}

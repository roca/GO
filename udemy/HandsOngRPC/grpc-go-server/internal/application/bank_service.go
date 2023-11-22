package application

import (
	"encoding/json"
	"fmt"
	"grpc-go-server/data"
	"math/rand"
	"time"

	"github.com/google/uuid"
	up "github.com/upper/db/v4"
)

type BankService struct {
	Models   data.Models
	ExitChan chan bool
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
func (b *BankService) InsertExchangeRatesAtInterval(exit chan bool, fromCurrency, toCurrency string, interval time.Duration) {
	er := &data.BankExchangeRate{}
	insertFunc := func() {
		now := time.Now()
		validFrom := now.Truncate(time.Second).Add(3 * time.Second)
		validTo := validFrom.Add(interval).Add(-1 * time.Millisecond)
		er.FromCurrency = fromCurrency
		er.ToCurrency = toCurrency
		er.ValidFromTimestamp = validFrom
		er.ValidToTimestamp = validTo
		er.Rate = 2000 + float64(rand.Intn(3000))
		//b.Models.BankExchangeRates.Insert(*er)
		bytes, _ := json.Marshal(er)
		fmt.Println("# b.Models.BankExchangeRates.Insert(m BankExchangeRate):", string(bytes))
		b.Models.BankExchangeRates.Insert(*er)
	}
	stop := runFuncAtInterval(insertFunc, interval)
	<-exit
	fmt.Println("Finished InsertExchangeRatesAtInterval")
	stop <- true
}
func (b *BankService) StopExchangeRatesAtInterval() {
	b.ExitChan <- true
}

func (b *BankService) GetExchangeRateAtTimestamp(fromCurrency, toCurrency string, timestamp time.Time) (*data.BankExchangeRate, error) {
	cond := up.Cond{"from_currency": fromCurrency, "to_currency": toCurrency}
	rates, err := b.Models.BankExchangeRates.GetAll(cond)
	if err != nil {
		return nil, err
	}
	return rates[0], nil
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

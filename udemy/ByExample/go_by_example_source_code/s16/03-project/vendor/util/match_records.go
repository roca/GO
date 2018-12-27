package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// MatchCustomersHouses ...
func MatchCustomersHouses() {

	start := time.Now()

	var err error

	RowsCustomer, err = DB.Query("SELECT * FROM " + TableName1)
	CheckErr(err)

	// Number of rows in 'utl.TableName1'
	_ = DB.QueryRow(`SELECT COUNT(*) FROM ` +
		TableName1).Scan(&NumOfRowsCustomer)

	for RowsCustomer.Next() {
		var mrtg Mortgage

		err = RowsCustomer.Scan(&mrtg.CID, &mrtg.FirstName, &mrtg.LastName,
			&mrtg.CreditScore, &mrtg.Salary, &mrtg.Downpayment, &mrtg.HouseID)

		CheckErr(err)

		query1 := "SELECT * FROM " + TableName2 +
			" WHERE ID=" + strconv.FormatUint(uint64(mrtg.HouseID), 10)

		rowHouse := DB.QueryRow(query1)

		err = rowHouse.Scan(&mrtg.HID, &mrtg.Price,
			&mrtg.MinDownpayment, &mrtg.PropertyTax, &mrtg.MaintenanceCost)
		CheckErr(err)

		Mrtgs = append(Mrtgs, mrtg)
	}

	// for i, m := range Mrtgs {
	// 	fmt.Println(i, m)
	// }

	WaitG.Add(NumOfRowsCustomer)
	for i, mrtg := range Mrtgs {
		go checkMortgageStatus(mrtg, i)
	}
	WaitG.Wait()

	elapsed := time.Since(start)
	fmt.Printf("= = = = = = = = = = = = = = = = = = = = = = = = = = =\n")
	fmt.Printf("= MatchCustomersHouses() - time elapsed: %s \n", elapsed)
	fmt.Printf("= = = = = = = = = = = = = = = = = = = = = = = = = = =\n\n")
}

func checkMortgageStatus(mrtg Mortgage, idx int) {

	Mu.Lock()
	delay := time.Duration(rand.Intn(700) + 300) // rand.Intn(max-min) + min
	time.Sleep(delay * time.Millisecond)

	approved := false

	if (mrtg.Downpayment < mrtg.MinDownpayment) || (mrtg.CreditScore < 650) {
		approved = false
	}
	yearlyMortgage := (mrtg.Price - mrtg.Downpayment) * 0.0612
	yearlyExpense := yearlyMortgage + mrtg.PropertyTax + mrtg.MaintenanceCost

	if yearlyExpense <= (mrtg.Salary * 0.32) {
		approved = true
	}

	fmt.Printf("%2d  %-10s %-10s %4d %12.2f %12.2f %3d\n",
		mrtg.CID, mrtg.FirstName,
		mrtg.LastName, mrtg.CreditScore,
		mrtg.Salary, mrtg.Downpayment, mrtg.HouseID)

	fmt.Printf("%2d %12.2f %12.2f %12.2f %12.2f \n", mrtg.HID, mrtg.Price,
		mrtg.MinDownpayment, mrtg.PropertyTax, mrtg.MaintenanceCost)

	if approved {
		fmt.Printf(" **Mortgage Application Approved! [process time=%v]\n\n", delay)
		AppprovedIdx = append(AppprovedIdx, idx)
	} else {
		fmt.Printf(" **Mortgage Application Rejected! [process time=%v]\n\n", delay)
		RejectedIdx = append(RejectedIdx, idx)
	}

	Mu.Unlock()

	WaitG.Done()
}

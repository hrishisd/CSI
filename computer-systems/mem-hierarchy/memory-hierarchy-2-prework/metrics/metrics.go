package metrics

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type UserId int
type UserMap map[UserId]*User

type Address struct {
	fullAddress string
	zip         int
}

type DollarAmount struct {
	dollars, cents uint64
}

type Payment struct {
	amount DollarAmount
	time   time.Time
}

type User struct {
	id       UserId
	name     string
	age      int
	address  Address
	payments []Payment
}

type UserData struct {
	ages            []uint32
	paymentsInCents []uint64
}

func AverageAge(users UserData) float64 {
	ages := users.ages
	n := len(users.ages)
	nLess1 := n - 3
	// var total uint32 = 0
	var acc1 uint32 = 0
	var acc2 uint32 = 0
	var acc3 uint32 = 0
	var acc4 uint32 = 0
	var i int = 0
	for ; i < nLess1; i += 4 {
		acc1 += ages[i]
		acc2 += ages[i+1]
		acc3 += ages[i+2]
		acc4 += ages[i+3]
	}
	for ; i < n; i++ {
		acc1 += ages[i]
	}
	return float64(acc1+acc2+acc3+acc4) / float64(n)
}

func AveragePaymentAmount(users UserData) float64 {
	paymentsInCents := users.paymentsInCents
	var total uint64
	for _, cents := range paymentsInCents {
		total += cents
	}
	return float64(total) / 100.0 / float64(len(paymentsInCents))
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount(users UserData) float64 {
	paymentsInCents := users.paymentsInCents
	n := float64(len(paymentsInCents))
	var totalCents uint64 = 0
	var totalCentsSq float64 = 0
	for _, cents := range users.paymentsInCents {
		totalCents += cents
		totalCentsSq += float64(cents * cents)
	}
	totalDollars := float64(totalCents) * .01
	meanDollars := totalDollars / n
	totalDollarsSq := totalCentsSq * .0001
	variance := totalDollarsSq/n - (meanDollars * meanDollars)
	return math.Sqrt(variance)
}

func LoadData() UserData {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	users := make(UserMap, len(userLines))
	for _, line := range userLines {
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		age, _ := strconv.Atoi(line[2])
		address := line[3]
		zip, _ := strconv.Atoi(line[3])
		users[UserId(id)] = &User{UserId(id), name, age, Address{address, zip}, []Payment{}}
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	for _, line := range paymentLines {
		userId, _ := strconv.Atoi(line[2])
		paymentCents, _ := strconv.Atoi(line[0])
		datetime, _ := time.Parse(time.RFC3339, line[1])
		users[UserId(userId)].payments = append(users[UserId(userId)].payments, Payment{
			DollarAmount{uint64(paymentCents / 100), uint64(paymentCents % 100)},
			datetime,
		})
	}

	ages := make([]uint32, 0)
	paymentsInCents := make([]uint64, 0)
	for _, user := range users {
		// fmt.Printf("%d\n", user.age)
		ages = append(ages, uint32(user.age))
		for _, payment := range user.payments {
			cents := payment.amount.dollars*100 + payment.amount.cents
			paymentsInCents = append(paymentsInCents, cents)
		}
	}
	// return users
	// fmt.Printf("%v\n", ages)
	fmt.Println(len(ages))
	fmt.Println(len(users))
	return UserData{ages, paymentsInCents}
}

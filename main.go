package main

import (
	"context"
	"log"
	"os"
	"time"

	"codeid.revamptwo/config"
	"codeid.revamptwo/repositories"
	"codeid.revamptwo/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting revamp REST API!")

	log.Println("Initializing Configuration!")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing DataBase!")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	// Test insert to payment.bank, using goroutine
	// ctx := context.Background()
	// queries := repositories.New(dbHandler)

	// newPaymentBank, err := queries.CreatePaymentBank(ctx,
	// 	repositories.CreatePaymentBankParams{
	// 		BankEntityID:     5,
	// 		BankCode:         "BNI-S",
	// 		BankName:         "BNI Syariah",
	// 		BankModifiedDate: time.Now(),
	// 	},
	// )

	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	// log.Println(newPaymentBank)

	// Test insert to payment.fintech, using goroutine
	// ctx2 := context.Background()
	// queries2 := repositories.New(dbHandler)

	// newPaymentFintech, err := queries2.CreatePaymentFintech(ctx2,
	// 	repositories.CreatePaymentFintechParams{
	// 		FintEntityID:     6,
	// 		FintCode:         "GO-P",
	// 		FintName:         "Gopay",
	// 		FintModifiedDate: time.Now(),
	// 	},
	// )

	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	// log.Println(newPaymentFintech)

	// Test insert to payment.transaction_payment, using goroutine
	// ctx3 := context.Background()
	// queries3 := repositories.New(dbHandler)

	// newPaymentTransactionPayment, err := queries3.CreatePaymentTransaction_payment(ctx3,
	// 	repositories.CreatePaymentTransaction_paymentParams{
	// 		TrpaID:           4,
	// 		TrpaCodeNumber:   "20220727#000004",
	// 		TrpaOrderNumber:  "22-12-1997",
	// 		TrpaDebit:        "4444",
	// 		TrpaCredit:       "0",
	// 		TrpaType:         "refund",
	// 		TrpaNote:         "Refund Ticket Pesawat",
	// 		TrpaModifiedDate: time.Now(),
	// 		TrpaSourceID:     "131899555",
	// 		TrpaTargetID:     "99999",
	// 		TrpaUserEntityID: 3,
	// 	},
	// )

	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	// log.Println(newPaymentTransactionPayment)

	// Test insert to payment.transaction_payment, using goroutine
	ctx4 := context.Background()
	queries4 := repositories.New(dbHandler)

	newPaymentUserAccount, err := queries4.CreatePaymentUsers_account(ctx4,
		repositories.CreatePaymentUsers_accountParams{
			UsacBankEntityID:  7,
			UsacUserEntityID:  3,
			UsacAccountNumber: "087654321",
			UsacSaldo:         "55555",
			UsacType:          "payment",
			UsacStartDate:     time.Now(),
			UsacEndDate:       time.Now(),
			UsacModifiedDate:  time.Now(),
			UsacStatus:        "blocked",
		},
	)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println(newPaymentUserAccount)
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "revamp" + env
	}

	return "revamp"
}

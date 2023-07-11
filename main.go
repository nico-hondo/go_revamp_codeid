package main

import (
	"context"
	"log"
	"os"

	"codeid.revamptwo/config"
	repositories "codeid.revamptwo/repositories"
	"codeid.revamptwo/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("starting revamp restart")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	// test insert to category, using goroutine
	ctx := context.Background()
	queries := repositories.New(dbHandler)

	// newCartItems, err := queries.CreateCartItems(ctx,
	// 	repositories.CreateCartItemsParams{
	// 		CaitID:           10,
	// 		CaitQuantity:     10,
	// 		CaitUnitPrice:    1000000,
	// 		CaitModifiedDate: time.Date(2024, time.December, 15, 15, 0, 0, 0, time.Local),
	// 		CaitUserEntityID: 10,
	// 		CaitProgEntityID: 2,
	// 	},
	// )

	// newSales_order_detail, err := queries.CreateSales_order_detail(ctx,
	// 	repositories.CreateSales_order_detailParams{
	// 		SodeID:           3,
	// 		SodeQty:          3,
	// 		SodeUnitPrice:    200000,
	// 		SodeUnitDiscount: 50,
	// 		SodeLineTotal:    70000,
	// 		SodeModifiedDate: time.Date(2023, time.May, 25, 15, 0, 0, 0, time.Local),
	// 		SodeSoheID:       2,
	// 		SodeProgEntityID: 3,
	// 	},
	// )

	// newSales_order_header, err := queries.CreateSales_order_header(ctx,
	// 	repositories.CreateSales_order_headerParams{
	// 		SoheID:             3,
	// 		SoheOrderDate:      time.Date(2023, time.December, 22, 15, 0, 0, 0, time.Local),
	// 		SoheDueDate:        time.Date(2023, time.June, 22, 15, 0, 0, 0, time.Local),
	// 		SoheShipDate:       time.Date(2023, time.January, 22, 15, 0, 0, 0, time.Local),
	// 		SoheOrderNumber:    "ORD#20220727-0000004",
	// 		SoheAccountNumber:  "151899557",
	// 		SoheTrpaCodeNumber: "",
	// 		SoheSubtotal:       "20",
	// 		SoheTax:            "70000",
	// 		SoheTotalDue:       3,
	// 		SoheLicenseCode:    "3",
	// 		SoheModifiedDate:   time.Date(2023, time.March, 22, 15, 0, 0, 0, time.Local),
	// 		SoheUserEntityID:   4,
	// 		SoheStatus:         "Open",
	// 	},
	// )

	// newSales_special_offer, err := queries.CreateSales_special_offer(ctx,
	// 	repositories.CreateSales_special_offerParams{
	// 		SpofID:           3,
	// 		SpofDescription:  "Dapatkan discount 25%, untuk 3 orang",
	// 		SpofDiscount:     50,
	// 		SpofType:         "",
	// 		SpofStartDate:    time.Date(2023, time.August, 22, 15, 0, 0, 0, time.Local),
	// 		SpofEndDate:      time.Date(2023, time.September, 22, 15, 0, 0, 0, time.Local),
	// 		SpofMinQty:       4,
	// 		SpofMaxQty:       3,
	// 		SpofModifiedDate: time.Date(2023, time.October, 22, 15, 0, 0, 0, time.Local),
	// 		SpofCateID:       3,
	// 	},
	// )

	// newSpecial_offer_programs, err := queries.CreateSpecial_offer_programs(ctx,
	// 	repositories.CreateSpecial_offer_programsParams{
	// 		SocoID:           4,
	// 		SocoSpofID:       2,
	// 		SocoProgEntityID: 2,
	// 		SocoStatus:       "closed",
	// 		SocoModifiedDate: time.Date(2023, time.June, 22, 15, 0, 0, 0, time.Local),
	// 	},
	// )
	//create
	// if err != nil {
	// 	log.Fatal("Error : ", err)
	// }
	// log.Println(newSpecial_offer_programs)

	//delete
	err := queries.DeleteSpecial_offer_programs(ctx, 4)
	if err != nil {
		log.Fatal("Error : ", err)
	}
	log.Println("Delete Selesai")

}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "revamp" + env
	}

	return ("revamp")
}

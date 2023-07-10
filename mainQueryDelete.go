package main

import (
	"context"
	"log"

	"codeid.revamptwo/repositories"
)

func DeleteData(ctx context.Context, queries *repositories.Queries) {
	var nilai int32
	// log.Print("Masukkan id yang ingin anda Hapus : ")
	// fmt.Scan(&nilai)

	nilai = 6

	err := queries.DeleteTalentApplyProgress(ctx, nilai)

	if err != nil {
		log.Fatal("Error : ", err)
	}
	log.Fatal("Delete Selesai")

}

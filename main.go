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
	log.Println("starting revamp-two restapi")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	// test insert,using goroutine
	ctx := context.Background()
	queries := repositories.New(dbHandler)

	// newBatch, err := queries.CreateBatch(ctx,
	// 	repositories.CreateBatchParams{
	// 		BatchID:           2,
	// 		BatchEntityID:     2,
	// 		BatchName:         "Batch#8",
	// 		BatchDescription:  "",
	// 		BatchStartDate:    time.Now(),
	// 		BatchEndDate:      time.Date(2023, time.May, 15, 15, 0, 0, 0, time.Local),
	// 		BatchReason:       "",
	// 		BatchType:         "offline",
	// 		BatchModifiedDate: time.Now(),
	// 		BatchStatus:       "Open",
	// 		BatchPicID:        1,
	// 	},
	// )

	// newBatchTrainee, err := queries.CreateBatchTrainee(ctx,
	// 	repositories.CreateBatchTraineeParams{
	// 		BatrID:               5,
	// 		BatrStatus:           "running",
	// 		BatrCertificated:     "1",
	// 		BatreCertificateLink: "asal",
	// 		BatrAccessToken:      "as01",
	// 		BatrAccessGrant:      "1",
	// 		BatrReview:           "canti",
	// 		BatrTotalScore:       20,
	// 		BatrModifiedDate:     "2023-01-21",
	// 		BatrTraineeEntityID:  2,
	// 		BatrBatchID:          3,
	// 	},
	// )

	// newBatchTraineeEvaluation, err := queries.CreateBatchTraineeEvaluation(ctx,
	// 	repositories.CreateBatchTraineeEvaluationParams{
	// 		BtevID:              5,
	// 		BtevType:            "hardskill",
	// 		BtevHeader:          "Fundamental",
	// 		BtevSection:         "bebas",
	// 		BtevSkill:           "Golang",
	// 		BtevWeek:            3,
	// 		BtevSkor:            70,
	// 		BtevNote:            "belajar lagiya",
	// 		BtevModifiedDate:    time.Now(),
	// 		BtevBatchID:         2,
	// 		BtevTraineeEntityID: 3,
	// 	},
	// )

	// newInstructorProgram, err := queries.CreateInstructorProgram(ctx,
	// 	repositories.CreateInstructorProgramParams{
	// 		BatchID:           1,
	// 		InproEntityID:     1,
	// 		InproEmpEntityID:  4,
	// 		InproModifiedDate: time.Now(),
	// 	},
	// )

	// newProgramApply, err := queries.CreateProgramApplyProgress(ctx,
	// 	repositories.CreateProgramApplyProgressParams{
	// 		PrapUserEntityID: 3,
	// 		PrapProgEntityID: 2,
	// 		PrapTestScore:    50,
	// 		PrapGpa:          2,
	// 		PrapIqTest:       89,
	// 		PrapReview:       "saluud",
	// 		PrapModifiedDate: time.Now(),
	// 		PrapStatus:       "Failed",
	// 	},
	// )

	newProgramApplyProgress, err := queries.CreateProgramApplyProgress(ctx,
		repositories.CreateProgramApplyProgressParams{
			ParogID:           1,
			ParogUserEntityID: 2,
			ParogProgEntityID: 2,
			ParogActionDate:   time.Now(),
			ParogModifiedDate: time.Now(),
			ParogComment:      "Nice",
			ParogProgressName: "done",
			ParogEmpEntityID:  1,
			ParogStatus:       "Open",
		},
	)

	if err != nil {
		log.Fatal("Error : ", err)
	}

	log.Println(newProgramApplyProgress)
}

// 	// Menghapus data menggunakan fungsi DeleteProgramApplyProgress
// 	parogIDToDelete := 1 // Masukkan nilai yang ingin dihapus di sini

// 	err = queries.DeleteProgramApplyProgress(ctx, int32(parogIDToDelete))
// 	if err != nil {
// 		log.Fatal("Error deleting program apply progress: ", err)
// 	}

// 	log.Println("Program Apply Progress with ID", parogIDToDelete, "deleted successfully")
// }

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "revamp-two" + env
	}
	return "revamp-two"
}

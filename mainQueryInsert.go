package main

import (
	"context"
	"database/sql"

	// "log"
	"time"

	repositories "codeid.revamptwo/repositories"
)

func InsertData(ctx context.Context, queries *repositories.Queries) {
	//pembuatan data tanggal
	var nullableTime sql.NullTime
	// var makeDistanceTime sql.NullTime
	// var makeDistanceHour sql.NullTime
	// var makeDistanceHour1 sql.NullTime
	t := time.Now()
	nullableTime.Time = t
	nullableTime.Valid = true

	//membuat jarak waktu awal dengan tambahan menit * 5
	// time.Sleep(5 * time.Minute)
	// makeDistanceTime.Time = t
	// makeDistanceTime.Valid = true

	// //membuat jarak waktu awal dengan tambahan day
	// time.Sleep(72 * time.Hour)
	// makeDistanceHour.Time = t
	// makeDistanceHour.Valid = true
	// //membuat jarak waktu awal dengan tambahan day
	// time.Sleep(144 * time.Hour)
	// makeDistanceHour1.Time = t
	// makeDistanceHour1.Valid = true

	//insert Data Job Category
	// newCategory, err := queries.CreateCategory(ctx,
	// 	repositories.CreateCategoryParams{
	// 		JocaID:           6,
	// 		JocaName:         "Rails Developer",
	// 		JocaModifiedDate: nullableTime,
	// 	})

	//insert Data Job Employee Range
	// newEmployeeRange, err := queries.CreateEmployeeRange(ctx,
	// 	repositories.CreateEmployeeRangeParams{
	// 		EmraID:           1,
	// 		EmraRangeMin:     1,
	// 		EmraRangeMax:     3,
	// 		EmraModifiedDate: nullableTime,
	// 	})

	//insert DAta Job Client
	// newJobClient, err := queries.CreateClient(ctx, repositories.CreateClientParams{
	// 	ClitID:           4,
	// 	ClitName:         "BTPN BANK",
	// 	ClitAbout:        "PT Bank BTPN Tbk adalah anak usaha Sumitomo Mitsui Banking Corporation yang berkantor pusat di Jakarta.",
	// 	ClitModifiedDate: nullableTime,
	// 	ClitAddrID:       1,
	// 	ClitEmraID:       1,
	// })

	//insert Data Job Post
	// newJobPost, err := queries.CreateJobPost(ctx, repositories.CreateJobPostParams{
	// 	JopoEntityID:       2,
	// 	JopoNumber:         "JOB20220728-0002",
	// 	JopoTitle:          "Data Engineer",
	// 	JopoStartDate:      nullableTime,
	// 	JopoEndDate:        nullableTime,
	// 	JopoMinSalary:      8000000,
	// 	JopoMaxSalary:      14000000,
	// 	JopoMinExperience:  1,
	// 	JopoMaxExperience:  3,
	// 	JopoPrimarySkill:   "Can Using Python",
	// 	JopoSecondarySkill: "Make Maintenance Device is must",
	// 	JopoPublishDate:    nullableTime,
	// 	JopoModifiedDate:   nullableTime,
	// 	JopoEmpEntityID:    4,
	// 	JopoClitID:         4,
	// 	JopoJoroID:         2,
	// 	JopoJotyID:         4,
	// 	JopoJocaID:         6,
	// 	JopoAddrID:         2,
	// 	JopoStatus:         "Waiting",
	// })

	//insert data jobpost desc
	// newJobPostDesc, err := queries.CreateJobPostDesc(ctx, repositories.CreateJobPostDescParams{
	// 	JopoEntityID:       2,
	// 	JopoDescription:    sql.NullString{String: `{"item" : "Developing, implementing, and maintaining Java based components and interfaces. Coordinate with the rest of the team working on different layers of the infrastructure. Delivering code with well tested"}`},
	// 	JopoResponsibility: sql.NullString{String: `{"items":"Proficient in Core Java, with a solid understanding of object-oriented programming Have minimum 1 year of working experience in Java Programming Skill Set: Spring framework, JPA / Hibernate Experience with database Oracle, My SQL, Postgre SQL or MS SQL Server Able to working with Agile methodology Can join immediately is a plus"}`},
	// 	JopoTarget:         sql.NullString{String: `{"item": "Sertifikat Professional, D3 (Diploma), D4 (Diploma), Sarjana (S1), Diploma Pascasarjana, Gelar Professional, Magister (S2)"}`},
	// 	JopoBenefit:        sql.NullString{String: `{"item":"Career growth in software development Positive working environment Great opportunity to get experiences in several industry sectors"}`},
	// })

	//insert data talent apply
	// newTalentApply, err := queries.CreateTalentApply(ctx, repositories.CreateTalentApplyParams{
	// 	TaapUserEntityID: 4,
	// 	TaapEntityID:     2,
	// 	TaapIntro:        "Perkenalkan nama saya Nofal, saya seorang Data Engineer",
	// 	TaapScoring:      9,
	// 	TaapModifiedDate: nullableTime,
	// 	TaapStatus:       "Done",
	// })

	//insert data talent apply progress
	// newTalentApplyProgress, err := queries.CreateTalentApplyProgress(ctx, repositories.CreateTalentApplyProgressParams{
	// 	TaprID:           6,
	// 	TaapUserEntityID: 4,
	// 	TaapEntityID:     2,
	// 	TaprModifiedDate: nullableTime,
	// 	TaprStatus:       "Cancel",
	// })

	// if err != nil {
	// 	log.Fatal("Error : ", err)
	// }
	// log.Println(newTalentApplyProgress)
}

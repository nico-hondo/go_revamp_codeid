package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"codeid.revamptwo/config"
	repositories "codeid.revamptwo/repositories/curriculum"
	"codeid.revamptwo/server"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("starting northwind restapi")

	log.Println("initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	// test insert to table
	ctx := context.Background()
	queries := repositories.New(dbHandler)

	//insertDataProgramEntity(ctx, queries)

	insertDataProgramEntityDescription(ctx, queries)

	//insertDataProgramReviews(ctx, queries)

	//insertDataSections(ctx, queries)

	//insertDataSectionDetail(ctx, queries)

	//insertDataSectionDetailMaterial(ctx, queries)

}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "miniproject" + env
	}
	return "miniproject"
}

// insert data ke program_entity
func insertDataProgramEntity(ctx context.Context, queries *repositories.Queries) {
	newprogram_entity, err := queries.CreateProgramEntity(ctx,
		repositories.CreateProgramEntityParams{
			ProgEntityID:     4,
			ProgTitle:        "C#",
			ProgHeadline:     "Develop application with C#",
			ProgType:         "bootcamp",
			ProgLearningType: "offline",
			ProgRating:       9,
			ProgTotalTrainee: 5,
			ProgModifiedDate: time.Now(),
			ProgImage:        "C#.png",
			ProgBestSeller:   "1",
			ProgPrice:        550000,
			ProgLanguage:     "bahasa",
			ProgDuration:     3,
			ProgDurationType: "month",
			ProgTagSkill:     "golang",
			ProgCityID:       1,
			ProgCateID:       3,
			ProgCreatedBy:    1,
			ProgStatus:       "publish",
		},
	)
	//cek error
	if err != nil {
		log.Fatal(err)
	}

	//tampilkan
	log.Println(newprogram_entity)
}

func insertDataProgramEntityDescription(ctx context.Context, queries *repositories.Queries) {
	newprogram_entity_description, err := queries.CreateProgramEntityDescription(ctx,
		repositories.CreateProgramEntityDescriptionParams{
			PredProgEntityID: 2,
			PredItemLearning: sql.NullString{String: `{"items" : "test"}`, Valid: true},
			PredItemInclude:  sql.NullString{String: `{"items" : "test"}`, Valid: true},
			PredRequirement:  sql.NullString{String: `{"items" : "test"}`, Valid: true},
			PredDescription:  sql.NullString{String: `{"items" : "test"}`, Valid: true},
			PredTargetLevel:  sql.NullString{String: `{"items" : "test"}`, Valid: true},
		},
	)
	/* newprogram_entity_description, err := queries.CreateProgramEntityDescription(ctx,
		repositories.CreateProgramEntityDescriptionParams{
			PredProgEntityID: 2,
			PredItemLearning: sql.NullString{String: `{ "items" : ["Become an advanced, confident, and modern JavaScript developer from scratch","Become job-ready by understanding how JavaScript really works behind the scenes","JavaScript fundamentals: variables, if/else, operators, boolean logic, functions, arrays, objects, loops, strings, etc."]}`, Valid: true},
			PredItemInclude:  sql.NullString{String: `{"items":[{"value" : "69 hours on-demand video","icon" : "video"},{"value" : "22 article","icon" : "list"},{"value" : "18 downloadable resources","icon" : "download"},{"value" : "Full lifetime access","icon" : "infinite"},{"value" : "Access on mobile and TV","icon" : "phone"},{"value" : "Certificate of completion","icon" : "reward"}]}`, Valid: true},
			PredRequirement:  sql.NullString{String: `{"items" : "No coding experience is necessary to take this course! I take you from beginner to expert!","Any computer and OS will work — Windows, macOS or Linux. We will set up your text editor the course.","A basic understanding of HTML and CSS is a plus, but not a must! The course includes an HTML and CSS crash course."}`, Valid: true},
			PredDescription:  sql.NullString{String: `{"items" : "JavaScript is the most popular programming language in the world. It powers the entire modern web. It provides millions of high-paying jobs all over the world."}`, Valid: true},
			PredTargetLevel:  sql.NullString{String: `{"items" : "Take this course if you want to gain a true and deep understanding of JavaScript"}`, Valid: true},
		},
	) */
	//cek error
	if err != nil {
		log.Fatal(err)
	}

	//tampilkan
	log.Println(newprogram_entity_description)
}

// insert data ke program_reviews
func insertDataProgramReviews(ctx context.Context, queries *repositories.Queries) {
	newprogram_reviews, err := queries.CreateProgramReviews(ctx,
		repositories.CreateProgramReviewsParams{
			ProwUserEntityID: 3,
			ProwProgEntityID: 4,
			ProwReview:       "Course javascript sangat comprehensif, mudah diikutin, trainer sangat menguasai materi",
			ProwRating:       4,
			ProwModifiedDate: time.Now(),
		},
	)
	//cek error
	if err != nil {
		log.Fatal(err)
	}

	//tampilkan
	log.Println(newprogram_reviews)
}

// insert data ke sections
func insertDataSections(ctx context.Context, queries *repositories.Queries) {
	newsections, err := queries.CreateSections(ctx,
		repositories.CreateSectionsParams{
			SectID:           4,
			SectProgEntityID: 1,
			SectTitle:        "Javascript Fundamental",
			SectDescription:  "",
			SectTotalSection: 12,
			SectTotalLecture: 5,
			SectTotalMinute:  180,
			SectModifiedDate: time.Now(),
		},
	)
	//cek error
	if err != nil {
		log.Fatal(err)
	}

	//tampilkan
	log.Println(newsections)
}

// insert data ke section_detail
func insertDataSectionDetail(ctx context.Context, queries *repositories.Queries) {
	newsections_detail, err := queries.CreateSectionDetail(ctx,
		repositories.CreateSectionDetailParams{
			SecdID:           3,
			SecdTitle:        "Course Structure",
			SecdPreview:      "1",
			SecdScore:        9,
			SecdNote:         "you will learning how to create new project",
			SecdMinute:       125,
			SecdModifiedDate: time.Now(),
			SecdSectID:       1,
		},
	)
	//cek error
	if err != nil {
		log.Fatal(err)
	}

	//tampilkan
	log.Println(newsections_detail)
}

// insert data ke section_detail_material
func insertDataSectionDetailMaterial(ctx context.Context, queries *repositories.Queries) {
	newsections_detail_material, err := queries.CreateSectionDetailMaterial(ctx,
		repositories.CreateSectionDetailMaterialParams{
			SedmID:           3,
			SedmFilename:     "video1.mp4",
			SedmFilesize:     12345,
			SedmFiletype:     "MP4",
			SedmFilelink:     "https://codeacademy/assets",
			SedmModifiedDate: time.Now(),
			SedmSecdID:       2,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	//tampilkan
	log.Println(newsections_detail_material)
}

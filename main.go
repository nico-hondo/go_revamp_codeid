package main

import (
	"context"
	"log"
	"os"
	"time"

	"codeid.revamptwo/config"
	users "codeid.revamptwo/repositories/users/create"
	"codeid.revamptwo/server"
	_ "github.com/lib/pq"
)

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "db_revamp" + env
	}
	// == file db_revamp.toml
	return "db_revamp"
}

func main() {
	log.Println("Starting db_revamp restapi")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	//test insert to category
	ctx := context.Background()   // create goroutine
	queries := users.New(dbHandler)


	/* |||----------------- INSERT DATA USERS --------------------||| */


			/* >> Insert Roles << */

	// newRoles, err := queries.CreateRoles(ctx, 
	// 	users.CreateRolesParams{
	// 		RoleID: 19,
	// 		RoleName: "Calo",
	// 		RoleType: "Eksternal",
	// 		RoleModifiedDate: time.Now(),
	// 	},
	// )

	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newRoles)
	

			/* >> Insert Business Entity << */

	// newEntityId, err := queries.CreateBusinessEntity(ctx, 11)


			/* >> Insert Data Users << */

	// newUsers, err := queries.CreateUsers(ctx, 
	// 	users.CreateUsersParams{
	// 		UserEntityID:       11,
	// 		UserName:           "Purwo",
	// 		UserPassword:       "12345",
	// 		UserFirstName:      "pur",
	// 		UserLastName:       "wo",
	// 		UserBirthDate:      time.Date(1998, time.August, 20, 12, 0, 0, 0, time.Local),
	// 		UserEmailPromotion: 0,
	// 		UserDemographic:    sql.NullString{String: `{ "latitude" : 12.90, "longitude" : -99.989}`, Valid: true},
	// 		UserModifiedDate:   time.Now(),
	// 		UserPhoto:          "n/a",
	// 		UserCurrentRole:    2,
	// 	},
	// )
	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newUsers)
	

		/* >> Insert Data Address << */

	// newAddress, err := queries.CreateAddrees(ctx, 
	// 	users.CreateAddreesParams{
	// 		EtadAddrID:       3,
	// 		EtadModifiedDate: time.Now(),
	// 		EtadEntityID:     11,
	// 		EtadAdtyID:       5,
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newAddress)


		/* Insert Data Education */

	// newEducation, err := queries.CreateEducation(ctx, 
	// 	users.CreateEducationParams{
	// 		UsduID:           2,
	// 		UsduEntityID:     11,
	// 		UsduSchool:       "Poltek",
	// 		UsduDegree:       "Diploma",
	// 		UsduFieldStudy:   "Information System",
	// 		UsduGraduateYear: "2022",
	// 		UsduStartDate:    time.Date(2019, time.September, 24, 10, 0, 0, 0, time.Local),
	// 		UsduEndDate:      time.Date(2022, time.September, 25, 13, 0, 0, 0, time.Local),
	// 		UsduGrade:        "3.70",
	// 		UsduActivities:   "Im Diploma With Cumlaude",
	// 		UsduDescription:  "Never Give Up to raise the future",
	// 		UsduModifiedDate: time.Now(),
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newEducation)


		/* Insert Data Email */

	// newEmail, err := queries.CreateEmail(ctx, 
	// 	users.CreateEmailParams{
	// 		PmailEntityID:     11,
	// 		PmailID:           4,
	// 		PmailAddress:      "purwo@gmail.com",
	// 		PmailModifiedDate: time.Now(),
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newEmail)


		/* Insert Data User Experience */

	// newUsex, err := queries.CreateExperience(ctx, 
	// 	users.CreateExperienceParams{
	// 		UsexID:              3,
	// 		UsexEntityID:        11,
	// 		UsexTitle:           "Programmer",
	// 		UsexProfileHeadline: "Backend Developer",
	// 		UsexEmploymentType:  "Fulltime",
	// 		UsexCompanyName:     "PT Future Tech",
	// 		UsexIsCurrent:       "1",
	// 		UsexStartDate:       time.Date(2022, time.October, 27, 17, 0, 0, 0, time.Local),
	// 		UsexEndDate:         time.Time{},
	// 		UsexIndustry:        "IT",
	// 		UsexDescription:     "Solve problem backend programmer",
	// 		UsexExperienceType:  "Company",
	// 		UsexCityID:          3,
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newUsex)



	/* Insert Data Experience Skill */

	// newUsexSkill, err := queries.CreateExperienceSkill(ctx, 
	// 	users.CreateExperienceSkillParams{
	// 		UeskUsexID: 3,
	// 		UeskUskiID: 3,
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newUsexSkill)


	/* Insert Data Users License */

	// newLicense, err := queries.CreateLicense(ctx, 
	// 	users.CreateLicenseParams{
	// 		UsliID:           1,
	// 		UsliLicenseCode:  "AA0001",
	// 		UsliModifiedDate: time.Now(),
	// 		UsliStatus:       "Active",
	// 		UsliEntityID:     11,
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newLicense)


	/*  Insert Users Media */

	// newMedia, err := queries.CreateMedia(ctx, 
	// 	users.CreateMediaParams{
	// 		UsmeID:           3,
	// 		UsmeEntityID:     11,
	// 		UsmeFileLink:     "http://purwo",
	// 		UsmeFilename:     "cv.docx",
	// 		UsmeFilesize:     2000,
	// 		UsmeFiletype:     "word",
	// 		UsmeNote:         "cv",
	// 		UsmeModifiedDate: time.Now(),
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newMedia)


	/* Insert Data Users Phones */

	// newPhone, err := queries.CreatePhones(ctx, 
	// 	users.CreatePhonesParams{
	// 		UspoEntityID:     11,
	// 		UspoNumber:       "081625661775",
	// 		UspoModifiedDate: time.Now(),
	// 		UspoPontyCode:    "Cell",
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newPhone)


	/* Insert Data Users Skill */

	// newSkill, err := queries.CreateSkill(ctx, 
	// 	users.CreateSkillParams{
	// 		UskiID:           4,
	// 		UskiEntityID:     11,
	// 		UskiModifiedDate: time.Now(),
	// 		UskiSktyName:     "hardskill",
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newSkill)



	/* Insert Data Users Phone Number Type */

	// newPhoneNumber, err := queries.CreateUsersPhoneNumberType(ctx, 
	// 	users.CreateUsersPhoneNumberTypeParams{
	// 		PontyCode:         "Office",
	// 		PontyModifiedDate: time.Now(),
	// 	},
	// )


	// if err != nil{
	// 	log.Fatal("Eror : ", err)
	// }
	// log.Println(newPhoneNumber)


	/* Insert Data Users Roles */

	newusersRoles, err := queries.CreateUsersRoles(ctx, 
		users.CreateUsersRolesParams{
			UsroEntityID:     11,
			UsroRoleID:       2,
			UsroModifiedDate: time.Now(),
		},
	)


	if err != nil{
		log.Fatal("Eror : ", err)
	}
	log.Println(newusersRoles)
}
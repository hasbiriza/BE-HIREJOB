package helper

import (
	"be_hiring_app/src/config"
	experiencemodel "be_hiring_app/src/models/ExperienceModel"
	projectmodel "be_hiring_app/src/models/ProjectModel"
	recruitermodel "be_hiring_app/src/models/RecruiterModel"
	usermodel "be_hiring_app/src/models/UserModel"
	workermodel "be_hiring_app/src/models/WorkerModel"
)

func Migration() {
	config.DB.AutoMigrate(&usermodel.User{})
	config.DB.AutoMigrate(&workermodel.Worker{})
	config.DB.AutoMigrate(&projectmodel.Project{})
	config.DB.AutoMigrate(&recruitermodel.Recruiter{})
	config.DB.AutoMigrate(&experiencemodel.Experience{})
}

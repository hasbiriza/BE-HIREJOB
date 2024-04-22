package routes

import (
	"github.com/gofiber/fiber/v2"
	experiencecontroller "be_hiring_app/src/controllers/ExperienceController"
	projectcontroller "be_hiring_app/src/controllers/ProjectController"
	recruitercontroller "be_hiring_app/src/controllers/RecruiterController"
	usercontroller "be_hiring_app/src/controllers/UserController"
	workercontroller "be_hiring_app/src/controllers/WorkerController"
	skillcontroller "be_hiring_app/src/controllers/SkillController"
)

func Router(c *fiber.App) {
	//helmet
	// c.Use(helmet.Default())

	v1 := c.Group("/api/v1")

	v1.Post("/login", usercontroller.Login)
	v1.Post("/register-worker", usercontroller.RegisterWorker)
	v1.Post("/register-recruiter", usercontroller.RegisterRecruiter)

	user := v1.Group("/user")
	{
		user.Get("/data", usercontroller.GetAllUsers)
		user.Get("/:id", usercontroller.GetUserById)
		user.Put("/update-worker/:id", usercontroller.UpdateWorker)
		user.Put("	", usercontroller.UpdateRecruiter)
		user.Delete("/delete/:id", usercontroller.DeleteUser)
	}

	worker := v1.Group("/user/worker")
	{
		worker.Get("/data", workercontroller.GetAllWorkers)
		worker.Get("/:id", workercontroller.GetWorkerById)
		worker.Post("/create", workercontroller.PostWorker)
		// worker.Get("/paginated-data", workercontroller)
		worker.Put("/update/:id", workercontroller.UpdateWorker)
		worker.Delete("/delete/:id", workercontroller.DeleteWorker)
	}

	recruiter := v1.Group("/user/recruiter")
	{
		recruiter.Get("/data", recruitercontroller.GetAllRecruiters)
		recruiter.Get("/:id", recruitercontroller.GetRecruiterById)
		recruiter.Post("/create", recruitercontroller.PostRecruiter)
		// recruiter.Get("/paginated-data", recruitercontroller)
		recruiter.Put("/update/:id", recruitercontroller.UpdateRecruiter)
		recruiter.Delete("/delete/:id", recruitercontroller.DeleteRecruiter)
	}

	experience := v1.Group("/experience")
	{
		experience.Get("/data", experiencecontroller.GetAllExperiences)
		experience.Get("/:id", experiencecontroller.GetExperienceById)
		experience.Post("/create", experiencecontroller.PostExperience)
		// experience.Post("/post-file", experiencecontroller.FileUpload)
		experience.Post("/remote", experiencecontroller.RemoteUpload)
		// experience.Get("/paginated-data", experiencecontroller)
		experience.Put("/update/:id", experiencecontroller.UpdateExperience)
		experience.Delete("/delete/:id", experiencecontroller.DeleteExperience)
	}

	project := v1.Group("/project")
	{
		project.Get("/data", projectcontroller.GetAllProjects)
		project.Get("/:id", projectcontroller.GetProjectById)
		project.Post("/create", projectcontroller.PostProject)
		// project.Get("/paginated-data", projectcontroller)
		project.Put("/update/:id", projectcontroller.UpdateProject)
		project.Delete("/delete/:id", projectcontroller.DeleteProject)
	}

	skill := v1.Group("/skill")
	{
		skill.Get("/data", skillcontroller.GetAllSkills)
		skill.Get("/:id", skillcontroller.GetSkillById)
		skill.Post("/create", skillcontroller.PostSkill)
		// skill.Get("/paginated-data", skillcontroller)
		skill.Put("/update/:id", skillcontroller.UpdateSkill)
		skill.Delete("/delete/:id", skillcontroller.DeleteSkill)
	}
}

package rest

import (
	"os"

	"github.com/atreugo/cors"
	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/librerialeo/oklever-api/pkg/websocket"
	"github.com/savsgio/atreugo/v11"
)

// SendResponse send parsed json structure response
func SendResponse(ctx *atreugo.RequestCtx, data ...interface{}) error {
	json := atreugo.JSON{
		"token":  ctx.UserValue("token"),
		"logout": ctx.UserValue("logout"),
		"error":  ctx.UserValue("error"),
		"data":   "",
	}
	if len(data) > 0 {
		if len(data) == 1 {
			json["data"] = data[0]
		} else {
			json["data"] = data
		}
	}
	return ctx.JSONResponse(json)
}

// InitRouterHandler initialize al routes
func InitRouterHandler(r *atreugo.Atreugo, conn *pgx.Conn) {
	s := service.InitService(conn)
	io := websocket.NewIO(s)
	io.InitActions()
	cors := cors.New(cors.Config{
		AllowedOrigins:   []string{os.Getenv("STUDENTS_ROUTE"), os.Getenv("TEACHERS_ROUTE"), os.Getenv("ACADEMY_ROUTE"), os.Getenv("SUPPORT_ROUTE"), "null"},
		AllowedHeaders:   []string{"Content-Type", "X-Custom", "authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
	})
	r.Static("/static", "static")
	r.UseAfter(cors)
	r.UseBefore(func(ctx *atreugo.RequestCtx) error {
		tokenString := string(ctx.Request.Header.Peek("Authorization"))
		if tokenString != "" {
			credential, token, valid := s.CheckToken(tokenString[7:])
			if valid {
				if credential != nil {
					if token != "" {
						ctx.SetUserValue("token", token)
					}
					ctx.SetUserValue("user", credential.ID)
					ctx.SetUserValue("rol", credential.Rol)
					ctx.SetUserValue("status", credential.Status)
				} else {
					ctx.SetUserValue("logout", true)
				}
			} else {
				ctx.SetUserValue("logout", true)
				return SendResponse(ctx)
			}
		}
		return ctx.Next()
	})
	r.GET("/", func(ctx *atreugo.RequestCtx) error {
		ctx.SendFile("index.html")
		return nil
	})
	r.GET("/ws", func(ctx *atreugo.RequestCtx) error {
		websocket.SocketInit(ctx, io)
		return nil
	})
	chatsRouter := r.NewGroupPath("/chats")
	InitChatsHandler(chatsRouter, s)
	ClassesCommentsRouter := r.NewGroupPath("/classes_comments")
	InitClassesCommentsHandler(ClassesCommentsRouter, s)
	classesRouter := r.NewGroupPath("/classes")
	InitClassesHandler(classesRouter, s)
	ClassesResourcesRouter := r.NewGroupPath("/classes_resources")
	InitClassesResourcesHandler(ClassesResourcesRouter, s)
	countriesRouter := r.NewGroupPath("/countries")
	InitCountriesHandler(countriesRouter, s)
	coursesDatasheetsRouter := r.NewGroupPath("/courses_datasheets")
	InitCoursesDatasheetsHandler(coursesDatasheetsRouter, s)
	coursesForumsCommentsRouter := r.NewGroupPath("/courses_forums_comments")
	InitCoursesForumsCommentsHandler(coursesForumsCommentsRouter, s)
	coursesForumsRouter := r.NewGroupPath("/courses_forums")
	InitCoursesForumsHandler(coursesForumsRouter, s)
	coursesRouter := r.NewGroupPath("/courses")
	InitCoursesHandler(coursesRouter, s)
	coursesLevelsRouter := r.NewGroupPath("/courses_levels")
	InitCoursesLevelsHandler(coursesLevelsRouter, s)
	coursesProjectsRouter := r.NewGroupPath("/courses_projects")
	InitCoursesProjectsHandler(coursesProjectsRouter, s)
	coursesReviewsRouter := r.NewGroupPath("/courses_reviews")
	InitCoursesReviewsHandler(coursesReviewsRouter, s)
	degreesRouter := r.NewGroupPath("/degrees")
	InitDegreesHandler(degreesRouter, s)
	insigniasRouter := r.NewGroupPath("/insignias")
	InitInsigniasHandler(insigniasRouter, s)
	languagesRouter := r.NewGroupPath("/languages")
	InitLanguagesHandler(languagesRouter, s)
	modulesFeedbackRouter := r.NewGroupPath("/modules_feedback")
	InitModulesFeedbackHandler(modulesFeedbackRouter, s)
	modulesRouter := r.NewGroupPath("/modules")
	InitModulesHandler(modulesRouter, s)
	modulesResourcesRouter := r.NewGroupPath("/modules_resources")
	InitModulesResourcesHandler(modulesResourcesRouter, s)
	purchasesRouter := r.NewGroupPath("/purchases")
	InitPurchasesHandler(purchasesRouter, s)
	purchasesProductsRouter := r.NewGroupPath("/purchases_products")
	InitPurchasesProductsHandler(purchasesProductsRouter, s)
	questionsRouter := r.NewGroupPath("/questions")
	InitQuestionsHandler(questionsRouter, s)
	questionsOptionsRouter := r.NewGroupPath("/questions_options")
	InitQuestionsOptionsHandler(questionsOptionsRouter, s)
	questionsTypesRouter := r.NewGroupPath("/questions_types")
	InitQuestionsTypesHandler(questionsTypesRouter, s)
	quizzesRouter := r.NewGroupPath("/quizzes")
	InitQuizzesHandler(quizzesRouter, s)
	resourcesTypesRouter := r.NewGroupPath("/resources_types")
	InitResourcesTypesHandler(resourcesTypesRouter, s)
	rolesRouter := r.NewGroupPath("/roles")
	InitRolesHandler(rolesRouter, s)
	studentsAnswersRouter := r.NewGroupPath("/students_answers")
	InitStudentsAnswersHandler(studentsAnswersRouter, s)
	studentsCoursesRouter := r.NewGroupPath("/students_courses")
	InitStudentsCoursesHandler(studentsCoursesRouter, s)
	studentsRouter := r.NewGroupPath("/students")
	InitStudentsHandler(studentsRouter, s)
	studentsProjectsRouter := r.NewGroupPath("/students_projects")
	InitStudentsProjectsHandler(studentsProjectsRouter, s)
	studentsProjectsHistoryRouter := r.NewGroupPath("/students_projects_history")
	InitStudentsProjectsHistoryHandler(studentsProjectsHistoryRouter, s)
	studentsQuizzesRouter := r.NewGroupPath("/students_quizzes")
	InitStudentsQuizzesHandler(studentsQuizzesRouter, s)
	studentsSubscriptionsRouter := r.NewGroupPath("/students_subscriptions")
	InitStudentsSubscriptionsHandler(studentsSubscriptionsRouter, s)
	subscriptionsRouter := r.NewGroupPath("/subscriptions")
	InitSubscriptionsHandler(subscriptionsRouter, s)
	synchronousClassesRouter := r.NewGroupPath("/synchronous_classes")
	InitSynchronousClassesHandler(synchronousClassesRouter, s)
	synchronousClassesResourcesRouter := r.NewGroupPath("/synchronous_classes_resources")
	InitSynchronousClassesResourcesHandler(synchronousClassesResourcesRouter, s)
	teachersRouter := r.NewGroupPath("/teachers")
	InitTeachersHandler(teachersRouter, s)
	usersLanguagesRouter := r.NewGroupPath("/users_languages")
	InitUsersLanguagesHandler(usersLanguagesRouter, s)
	testClassesFeedbackRouter := r.NewGroupPath("/test_classes_feedback")
	InitTestClassesFeedbackHandler(testClassesFeedbackRouter, s)
	testClassesRouter := r.NewGroupPath("/test_classes")
	InitTestClassesHandler(testClassesRouter, s)
	topicsRouter := r.NewGroupPath("/topics")
	InitTopicsHandler(topicsRouter, s)
	topicsResourcesRouter := r.NewGroupPath("/topics_resources")
	InitTopicsResourcesHandler(topicsResourcesRouter, s)
	usersChatsRouter := r.NewGroupPath("/users_chats")
	InitUsersChatsHandler(usersChatsRouter, s)
	usersDegreesRouter := r.NewGroupPath("/users_degrees")
	InitUsersDegreesHandler(usersDegreesRouter, s)
	usersRouter := r.NewGroupPath("/users")
	InitUsersHandler(usersRouter, s)
	usersInsigniasRouter := r.NewGroupPath("/users_insignias")
	InitUsersInsigniasHandler(usersInsigniasRouter, s)
	usersAcademyRouter := r.NewGroupPath("/users_academy")
	InitUsersAcademyHandler(usersAcademyRouter, s)
}

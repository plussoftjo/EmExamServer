// Package routes (Setup Routes Group)
package routes

import (
	"server/config"
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Setup >>>
func Setup() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "authorization", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))

	// gin.SetMode(gin.ReleaseMode)

	r.Use(static.Serve("/public", static.LocalFile(config.ServerInfo.PublicPath+"public", true)))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})

	// -------- Auth Groups ----------//

	// ~~~ Auth Group ~~~ //
	auth := r.Group("/auth")
	// Auth Routes
	auth.POST("/login", controllers.LoginController)
	auth.POST("/register", controllers.RegisterController)
	auth.GET("/auth", controllers.Auth)
	auth.POST("/update", controllers.UpdateUser)
	auth.POST("/store/results", controllers.StoreUserResults)

	// ~~~ Main Controller ~~~ //
	main := r.Group("/main")
	main.GET("/index", controllers.Index)
	main.GET("/index_all_questions", controllers.IndexAllQuestions)
	main.GET("/index_with_auth/:user_id", controllers.IndexWithAuth)
	main.GET("/index/questions/:categories_id", controllers.Indexquestions)
	main.POST("/store_notification_token", controllers.StoreNotificationToken)

	// ~~~ Admin Controller ~~~ //
	admin := r.Group("/admin")
	admin.POST("/store/categories", controllers.StoreCategories)
	admin.POST("/store/questions", controllers.StoreQuestions)
	admin.POST("/store/answers", controllers.StoreAnswers)

	admin.POST("/store/QuestionWithAnswers", controllers.StoreQuestionsWithAnswers)
	admin.POST("/question/remove/:id", controllers.RemoveQuestion)
	r.Run(":8082")

}

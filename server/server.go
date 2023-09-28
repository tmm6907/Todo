package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/tmm6907/Todo/constants"
	"github.com/tmm6907/Todo/db"
	"github.com/tmm6907/Todo/handlers"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
	port   string
}

func (s *Server) RegisterRoutes() error {
	h := &handlers.Handler{
		DB: s.db,
	}
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "base", gin.H{
			"msg": "TodoItems",
		})
	})
	// partialsRoutes := router.Group("/partials")
	// partialsRoutes.GET("/")
	apiRoutes := s.router.Group("/api")
	apiRoutes.GET("/todos", h.GetTodoItems)
	return nil
}

func (s *Server) ParseTemplates() error {
	files := []string{}
	for _, dir := range constants.TEMPLATE_DIRS {
		ff, err := filepath.Glob(dir)
		if err != nil {
			return err
		}
		files = append(files, ff...)
	}
	log.Println("Templates loaded:", files)
	t, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}
	s.router.SetHTMLTemplate(t)
	return nil
}

func (s *Server) Run() {
	db, err := db.InitDB(constants.DB_PATH)
	if err != nil {
		log.Fatalln(err)
	}
	s.db = db
	s.ParseTemplates()
	s.RegisterRoutes()
	log.Fatal(s.router.Run(s.port))
}

func New(port string) *Server {
	return &Server{
		router: gin.Default(),
		port:   ":" + port}
}

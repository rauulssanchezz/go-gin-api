package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rauulssanchezz/go-gin-api/internal/task"
	"github.com/rauulssanchezz/go-gin-api/internal/user"
	"github.com/rauulssanchezz/go-gin-api/pkg/router"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error al cargar .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatalf("Faltan variables de entorno criticas.")
	}

	dbConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbConnString)

	if err != nil {
		log.Fatalf("Error al abrir la conexión a PostgreSQL: %v", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a PostgreSQL: %v", err)
	}
	log.Println("Conexión a PostgreSQL exitosa.")

	taskRepo := task.NewTaskRepositoryStruct(db)
	taskService := task.NewTaskService(taskRepo)
	taskHandler := task.NewTaskHandler(taskService)

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	r := gin.Default()
	router.SetUpRoutes(r, taskHandler, userHandler)

	apiPort := "8080"

	log.Printf("Servidor Gin iniciando en el puerto %s", apiPort)
	if err := r.Run(":" + apiPort); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

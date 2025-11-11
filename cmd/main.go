package main

import (
	"fmt"
	"go-petshop/internal/controllers"
	"go-petshop/internal/db"
	"go-petshop/internal/repositories"
	"go-petshop/internal/router"
	"go-petshop/internal/services"
	"log"
	"net/http"
)

func printBanner() {
	banner := `    
   ____                 ____       _       
  / ___| ___           |  _ \ ___ | |_     
 | |  _ / _ \   _____  | |_) / _ \| __|    
 | |_| | (_) | |_____| |  __/  __/| |_     
  \____|\___/          |_|   \___| \__|    
                                            
 :: Go-PetShop ::                (v1.0.0)    
`
	fmt.Println(banner)
	fmt.Println("🐾 Sistema de Gerenciamento de PetShop")
	fmt.Println("📦 Inicializando aplicação...")
	fmt.Println()
}

func main() {
	printBanner()

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	usuarioRepo := repositories.NewUsuarioRepository(db)
	usuarioService := services.NewUsuarioService(usuarioRepo)
	usuarioController := controllers.NewUsuarioController(usuarioService)

	authService := services.NewAuthService(usuarioRepo)
	authController := controllers.NewAuthController(authService)

	router := router.NewRouter(&usuarioController, &authController)
	r := router.Setup()

	fmt.Println("✅ Banco de dados conectado")
	fmt.Println("✅ Rotas configuradas")
	log.Println("🚀 Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

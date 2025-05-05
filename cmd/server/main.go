package main

import (
	"log"
	"net/http"

	"registration-app/internal/db"
	"registration-app/internal/handlers"
)

func main() {
	db.InitDB("data.db")
	defer db.Close()
	http.HandleFunc("/", handlers.Registration)
	http.HandleFunc("/registration", handlers.Registration)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/shop", handlers.ShopHandler)
	http.HandleFunc("/add-to-cart", handlers.AddToCartHandler)
	http.HandleFunc("/remove-from-cart", handlers.RemoveItemFromCartHandler)
	http.HandleFunc("/cart", handlers.CartHandler)
	http.HandleFunc("/process-order", handlers.ProcessOrderHandler)
	http.HandleFunc("/profile", handlers.GetProfile)

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

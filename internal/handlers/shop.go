package handlers

import (
	"log"
	"net/http"
	"registration-app/internal/db"
	"registration-app/internal/models"
	"strconv"
)

func ShopHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	username := cookie.Value

	rows, err := db.DB.Query("SELECT id, Name, Description, price FROM items")
	if err != nil {
		http.Error(w, "Failed to query items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price); err != nil {
			log.Println("Db scan error:", err)
		}
		items = append(items, item)
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	err = db.DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
	if err != nil {
		http.Error(w, "Failed to get username", http.StatusInternalServerError)
		return
	}

	var cartCount int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM carts WHERE user_id = ?", userID).Scan(&cartCount)
	if err != nil {
		http.Error(w, "Failed to get cart count", http.StatusInternalServerError)
		return
	}

	data := models.ShopPageData{
		Username:  username,
		Items:     items,
		CartCount: cartCount,
	}

	err = templates.ExecuteTemplate(w, "shop.html", data)
	if err != nil {
		log.Println("Template execution error:", err)
	}
}

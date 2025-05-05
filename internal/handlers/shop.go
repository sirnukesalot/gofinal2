package handlers

import (
	"net/http"
	"registration-app/internal/db"
	"registration-app/internal/models"
)

func ShopHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, Name, Description, price FROM items")
	if err != nil {
		http.Error(w, "Failed to query items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Name, &item.Description, &item.Price); err == nil {
			items = append(items, item)
		}
	}

	templates.ExecuteTemplate(w, "shop.html", map[string]interface{}{
		"Items": items,
	})
}

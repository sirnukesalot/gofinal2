package handlers

import (
	"html/template"
	"net/http"
	"registration-app/internal/db"
	"registration-app/internal/models"
	"strconv"
)

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromSession(r)
	r.ParseForm()
	itemID, _ := strconv.Atoi(r.FormValue("item"))

	_, err := db.DB.Exec("INSERT INTO carts (user_id, item_id) VALUES (?, ?)", userID, itemID)
	if err != nil {
		http.Error(w, "Failed to add item to cart", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromSession(r)

	rows, err := db.DB.Query(`
		SELECT items.id, items.name, items.description, items.price
		FROM carts
		JOIN items ON items.id = carts.item_id
		WHERE carts.user_id = ?`, userID)
	if err != nil {
		http.Error(w, "Failed to load cart", http.StatusInternalServerError)
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

	tmpl := template.Must(template.ParseFiles("templates/cart.html"))
	tmpl.Execute(w, map[string]interface{}{"CartItems": items})
}

func getUserIDFromSession(r *http.Request) int {
	return 1
}

func RemoveFromCartHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	itemID, _ := strconv.Atoi(r.FormValue("item_id"))
	userID := getUserIDFromSession(r)

	_, err := db.DB.Exec("DELETE FROM carts WHERE user_id = ? AND item_id = ? LIMIT 1", userID, itemID)
	if err != nil {
		http.Error(w, "Failed to remove item", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

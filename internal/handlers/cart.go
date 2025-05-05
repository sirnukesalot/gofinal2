package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"registration-app/internal/db"
	"registration-app/internal/models"
	"strconv"
)

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {

	userID := getUserIDFromSession(r)
	if userID == 0 {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}

	r.ParseForm()
	itemID, err := strconv.Atoi(r.FormValue("item"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("INSERT OR REPLACE INTO carts (user_id, item_id) VALUES (?, ?)", userID, itemID)
	if err != nil {
		http.Error(w, "Failed to add item to cart", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/shop", http.StatusSeeOther)
}

func getCartFromDatabase(r *http.Request) ([]models.Item, error) {
	userID := getUserIDFromSession(r)
	if userID == 0 {
		return nil, fmt.Errorf("user not logged in")
	}

	rows, err := db.DB.Query("SELECT i.id, i.name, i.description, i.price FROM items i JOIN carts c ON i.id = c.item_id WHERE c.user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	cartItems, err := getCartFromDatabase(r)
	if err != nil {
		http.Error(w, "Failed to load cart", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/cart.html"))
	tmpl.Execute(w, map[string]interface{}{"CartItems": cartItems})
}

func getUserIDFromSession(r *http.Request) int {
	cookie, err := r.Cookie("session_user")
	if err != nil {
		return 0
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return 0
	}

	return userID
}

func ProcessOrderHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromSession(r)

	rows, err := db.DB.Query("SELECT item_id FROM carts WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Failed to load cart", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var itemIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err == nil {
			itemIDs = append(itemIDs, id)
		}
	}

	for _, id := range itemIDs {
		_, err := db.DB.Exec("DELETE FROM items WHERE id = ?", id)
		if err != nil {
			log.Println("Failed to delete item:", id, err)
		}
	}

	_, err = db.DB.Exec("DELETE FROM carts WHERE user_id = ?", userID)
	if err != nil {
		log.Println("Failed to clear cart for user:", userID, err)
		http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "flash",
		Value: "Order placed successfully!",
		Path:  "/",
	})
	http.Redirect(w, r, "/shop", http.StatusSeeOther)
}

func RemoveItemFromCartHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromSession(r)
	if userID == 0 {
		log.Println("User is not logged in")
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}

	r.ParseForm()
	itemID, err := strconv.Atoi(r.FormValue("item_id"))
	if err != nil {
		log.Println("Invalid item ID:", err)
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("DELETE FROM carts WHERE user_id = ? AND item_id = ?", userID, itemID)
	if err != nil {
		log.Println("Error removing item:", err)
		http.Error(w, "Failed to remove item from cart", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

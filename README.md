# 🛒 Online Shop Web Application

This is a Go-based web application for an online shop, allowing users to register, log in, browse products, manage a cart, and place orders. It features user session handling, SQLite-backed data persistence, and clean routing for a smooth shopping experience.

## ✨ Features

- User registration and login with session cookies
- Product listing with detailed descriptions and prices
- Add-to-cart functionality with live item count
- Cart management with item removal
- Order processing with automatic stock updates

## 🧱 Technologies

- Go (Golang)
- SQLite
- HTML/CSS
- PlantUML (for architecture diagramming)

## 🗺️ System Flow Diagram

```plantuml
@startuml Online Shop
actor User as Foo
participant Website as Foo2
participant DB_Server as Foo3
participant Cart as Foo4
participant Checkout as Foo5

Activate Foo
Foo -> Foo2: Registration
Activate Foo2
Foo2 -> Foo3: Add user to db
Deactivate Foo2
Activate Foo3
Foo3 --> Foo2: Successfully added
Deactivate Foo3
Foo2 --> Foo: Redirect to login page

Foo -> Foo2: Login
Activate Foo2
Foo2 -> Foo3: Authentication
Deactivate Foo2
Foo3 --> Foo2: Success authentication
Foo2 --> Foo: Redirect to home page

Foo -> Foo2: Searches for product
Activate Foo2
Deactivate Foo2

Foo -> Foo2: Clicks on add button
Activate Foo2
Foo2 -> Foo4: Add item to cart
Deactivate Foo2
Activate Foo4

Foo4 -> Foo3: Check for stock
Activate Foo3
Foo3 --> Foo4:
Deactivate Foo3

Foo4 -> Foo5: Authorize purchase
Deactivate Foo4
Activate Foo5
Foo5 --> Foo3: Update db
Activate Foo3
Deactivate Foo3
Foo5 --> Foo: Immediate notification

Deactivate Foo5
Deactivate Foo4
Deactivate Foo
@enduml

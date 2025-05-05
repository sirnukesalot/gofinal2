BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username TEXT NOT NULL,
            email TEXT NOT NULL,
            password TEXT NOT NULL
        );
CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		name TEXT,
		description TEXT,
		price REAL
        );  
CREATE TABLE IF NOT EXISTS carts (
		user_id INTEGER,
		item_id INTEGER
        );

INSERT INTO items (name, description, price) VALUES ('Book', 'A thrilling novel', 12.99);
INSERT INTO items (name, description, price) VALUES ('Headphones', 'Noise-canceling over-ear headphones', 79.99);
INSERT INTO items (name, description, price) VALUES ('Laptop', 'Lightweight 13-inch laptop with 16GB RAM', 999.00);
INSERT INTO items (name, description, price) VALUES ('Backpack', 'Waterproof backpack with padded straps', 45.50);

COMMIT;
BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            username TEXT NOT NULL,
            email TEXT NOT NULL,
            password TEXT NOT NULL
        );
CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY,
		name TEXT,
		description TEXT,
		price REAL
        );  
CREATE TABLE IF NOT EXISTS carts (
    user_id INT,
    item_id INT,
    PRIMARY KEY (user_id, item_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (item_id) REFERENCES items(id)
);
-- INSERT INTO items (name, description, price) VALUES ('Book', 'A thrilling novel', 12.99);
-- INSERT INTO items (name, description, price) VALUES ('Headphones', 'Noise-canceling over-ear headphones', 79.99);
-- INSERT INTO items (name, description, price) VALUES ('Laptop', 'Lightweight 13-inch laptop with 16GB RAM', 999.00);
-- INSERT INTO items (name, description, price) VALUES ('Backpack', 'Waterproof backpack with padded straps', 45.50);
-- INSERT INTO items (name, description, price) VALUES ('Smartphone', 'Latest model with 128GB storage and 5G connectivity', 799.99);
-- INSERT INTO items (name, description, price) VALUES ('Coffee Maker', 'Automatic coffee machine with built-in grinder', 99.99);
-- INSERT INTO items (name, description, price) VALUES ('Smartwatch', 'Fitness tracker with heart rate monitor and GPS', 199.99);
-- INSERT INTO items (name, description, price) VALUES ('Gaming Mouse', 'Ergonomic design with customizable RGB lighting', 49.99);
-- INSERT INTO items (name, description, price) VALUES ('Bluetooth Speaker', 'Portable speaker with high-quality sound and waterproof design', 29.99);
-- INSERT INTO items (name, description, price) VALUES ('Camera', 'Digital SLR camera with 20MP resolution', 599.99);

COMMIT;
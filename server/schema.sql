-- 1. The Tables (Just physical locations in the restaurant)
CREATE TABLE tables (
    id TEXT PRIMARY KEY, -- e.g., "table_1"
    label TEXT NOT NULL -- e.g., "Table 1 (Window)"
);

-- 2. The Dishes (Your actual Menu)
CREATE TABLE dishes (
    id TEXT PRIMARY KEY, -- uuid()
    title TEXT NOT NULL,
    desc TEXT,
    price REAL NOT NULL, -- You missed price in the diagram!
    is_available INTEGER DEFAULT 1
);

-- 3. The Orders (The overarching ticket for a table)
CREATE TABLE orders (
    id TEXT PRIMARY KEY, -- uuid()
    table_id TEXT NOT NULL,
    status TEXT DEFAULT 'pending', -- 'pending', 'preparing', 'completed'
    FOREIGN KEY (table_id) REFERENCES tables (id)
);

-- 4. Order Items (The Junction Table bridging Orders and Dishes)
-- This replaces your dish.id[] array and allows tracking quantities.
CREATE TABLE order_items (
    id TEXT PRIMARY KEY, -- uuid()
    order_id TEXT NOT NULL,
    dish_id TEXT NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (dish_id) REFERENCES dishes (id)
);
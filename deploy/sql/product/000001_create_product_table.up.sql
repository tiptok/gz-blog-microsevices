CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR ( 255 ) NOT NULL,
    description TEXT,
    price DECIMAL ( 10, 2 ) NOT NULL,
    unit VARCHAR ( 255 ) NOT NULL,
    stock BIGINT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    version BIGINT DEFAULT 0
);
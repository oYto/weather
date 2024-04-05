CREATE DATABASE weather;
USE weather;
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    default_city VARCHAR(255) DEFAULT 'China',
    uuid VARCHAR(36)
    );

INSERT INTO users (name, email, password, default_city)
VALUES ('ft', 'ft@123.com', '123', 'China');
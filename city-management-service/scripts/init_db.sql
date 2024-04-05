use weather;
    
CREATE TABLE cities (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        uuid VARCHAR(36),
                        name VARCHAR(255) NOT NULL,
                        country VARCHAR(255) NOT NULL,
                        latitude DECIMAL(10, 8) NOT NULL,
                        longitude DECIMAL(11, 8) NOT NULL
) ENGINE=InnoDB;
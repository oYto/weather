use weather;
CREATE TABLE weather_info (
                              id INT AUTO_INCREMENT PRIMARY KEY,
                              uuid VARCHAR(36)
                              city_id VARCHAR(255) NOT NULL,
                              temperature DOUBLE PRECISION NOT NULL,
                              humidity DOUBLE PRECISION NOT NULL,
                              condition VARCHAR(255) NOT NULL,
                              wind_speed DOUBLE PRECISION NOT NULL,
                              timestamp TIMESTAMP NOT NULL
);

CREATE TABLE forecast_info (
                               id INT AUTO_INCREMENT PRIMARY KEY,
                               uuid VARCHAR(36)
                               city_id VARCHAR(255) NOT NULL,
                               date DATE NOT NULL,
                               temperature DOUBLE PRECISION NOT NULL,
                               condition VARCHAR(255) NOT NULL
);

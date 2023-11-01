CREATE DATABASE IF NOT EXISTS meal_deals;

USE meal_deals;

CREATE TABLE IF NOT EXISTS produce (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS appliance (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS recipes (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description TEXT,
    Users VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS recipe_produce (
    recipe_id INT,
    produce_id INT,
    PRIMARY KEY (recipe_id, produce_id),
    FOREIGN KEY (recipe_id) REFERENCES recipes(ID),
    FOREIGN KEY (produce_id) REFERENCES produce(ID)
);

CREATE TABLE IF NOT EXISTS recipe_appliance (
    recipe_id INT,
    appliance_id INT,
    PRIMARY KEY (recipe_id, appliance_id),
    FOREIGN KEY (recipe_id) REFERENCES recipes(ID),
    FOREIGN KEY (appliance_id) REFERENCES appliance(ID)
);
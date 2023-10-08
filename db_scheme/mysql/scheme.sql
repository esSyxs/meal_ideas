CREATE DATABASE IF NOT EXISTS meal_ideas;
USE meal_ideas;
CREATE TABLE IF NOT EXISTS Users (
    ID CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password CHAR(64) NOT NULL,
    email VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS Recipes (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS Products (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS Appliances (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS Favourite_recipes (
    user_id CHAR(36),
    recipe_id INT,
    PRIMARY KEY (user_id, recipe_id),
    FOREIGN KEY (user_id) REFERENCES Users(ID),
    FOREIGN KEY (recipe_id) REFERENCES Recipes(ID)
);
CREATE TABLE IF NOT EXISTS Recipes_produce (
    recipe_id INT,
    product_id INT,
    PRIMARY KEY (recipe_id, product_id),
    FOREIGN KEY (recipe_id) REFERENCES Recipes(ID),
    FOREIGN KEY (product_id) REFERENCES Products(ID)
);
CREATE TABLE IF NOT EXISTS Recipes_appliances (
    recipe_id INT,
    appliance_id INT,
    PRIMARY KEY (recipe_id, appliance_id),
    FOREIGN KEY (recipe_id) REFERENCES Recipes(ID),
    FOREIGN KEY (appliance_id) REFERENCES Appliances(ID)
);
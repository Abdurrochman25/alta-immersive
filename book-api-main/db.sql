CREATE DATABASE IF NOT EXISTS db_bookapi;
USE db_bookapi;

CREATE TABLE IF NOT EXISTS users(
   id INT AUTO_INCREMENT PRIMARY KEY,
   name VARCHAR(40),
   email VARCHAR(40),
   password VARCHAR(40)
);

INSERT IGNORE INTO users
    (id, name, email, password)
VALUES
    (1, 'Abdur', 'abdurrochman@gmail.com', '123456'),
    (2, 'Jovi', 'jovi@gmail.com', '34567'),
    (3, 'Ilham', 'ilham@gmail.com', '752368');

CREATE TABLE IF NOT EXISTS books(
   id INT AUTO_INCREMENT PRIMARY KEY,
   title VARCHAR(40),
   author VARCHAR(40)
);

INSERT IGNORE INTO books
    (id, title, author)
VALUES
    (1, 'Doraemon', 'Fujiko F. Fujio'),
    (2, 'Harry Potter', 'J.K. Rowling'),
    (3, 'The Lord of The Ring', 'J.R.R. Tolkien');
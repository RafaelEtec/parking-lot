CREATE DATABASE IF NOT EXISTS parkinglot;
USE parkinglot;

CREATE TABLE cars (
    id int PRIMARY KEY AUTO_INCREMENT,
    brand varchar(40) NOT NULL,
    model varchar(40) NOT NULL,
    platenumber varchar(7) NOT NULL,
    parkingspace varchar(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET= utf8;

INSERT INTO cars(brand, model, platenumber, parkingspace) VALUES
('BMW', '320i', 'ABC0001', 'A01'),
('BMW', 'X1', 'ABC0002', 'A02'),
('AUDI', 'Q3', 'ABC0003', 'A03'),
('AUDI', 'S4', 'ABC0004', 'A04'),
('Chevrolet', 'corsa', 'DMG0286', 'A05');

DESCRIBE cars;
SELECT * FROM cars;
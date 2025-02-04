CREATE DATABASE IF NOT EXISTS parkinglot;
USE parkinglot;

CREATE TABLE cars (
    id int PRIMARY KEY AUTO_INCREMENT,
    brand varchar(40) NOT NULL,
    model varchar(40) NOT NULL,
    platenumber varchar(7) NOT NULL,
    parkingspace varchar(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET= utf8;

DESCRIBE cars;
SELECT * FROM cars;
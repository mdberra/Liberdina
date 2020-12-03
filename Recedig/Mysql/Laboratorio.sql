DELIMITER $$
USE recedig$$
DROP TABLE Laboratorio$$
CREATE TABLE `recedig`.`Laboratorio` (
	`idLaboratorio` INT NOT NULL AUTO_INCREMENT,
	`nombre` varchar(255) NOT NULL,
	`direccion` varchar(255) NOT NULL,
	`telefono` varchar(255) NOT NULL,
	`email` varchar(255) NOT NULL,
	`fechaIngreso` DATETIME NULL,
	`estado` INT NULL,
	`fechaEstado` DATETIME NULL,
	PRIMARY KEY (`idLaboratorio`)
);$$
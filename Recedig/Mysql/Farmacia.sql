DELIMITER $$
USE recedig
DROP TABLE Farmacia;
CREATE TABLE `recedig`.`Farmacia` (
	`idFarmacia` INT NOT NULL AUTO_INCREMENT,
	`nombre` varchar(255) NOT NULL,
	`direccion` varchar(255) NOT NULL,
	`telefono` varchar(255) NOT NULL,
	`email` varchar(255) NOT NULL,
	`fechaIngreso` DATETIME NULL,
	`estado` INT NULL,
	`fechaEstado` DATETIME NULL,
	PRIMARY KEY (`idFarmacia`)
);
DELIMITER $$
USE recedig$$
DROP TABLE Medicamento$$
CREATE TABLE `recedig`.`Medicamento` (
	`idMedicamento` INT NOT NULL AUTO_INCREMENT,
	`nombre` varchar(255) NOT NULL,
	`droga` varchar(255) NOT NULL,
	`idLaboratorio` INT NOT NULL,
	`fechaIngreso` DATETIME NULL,
	`estado` INT NULL,
	`fechaEstado` DATETIME NULL,
	PRIMARY KEY (`idMedicamento`)
);$$
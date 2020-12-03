DELIMITER $$
USE recedig$$
DROP TABLE Medico$$
CREATE TABLE `recedig`.`Medico` (
	`idMedico` INT NOT NULL AUTO_INCREMENT,
	`nombre` varchar(255) NOT NULL,
	`apellido` varchar(255) NOT NULL,
	`email` varchar(255) NULL,
	`telefono` varchar(255) NULL,
	`dni` INT NOT NULL,
	`matricula` varchar(255) NOT NULL,
	`idImagen` INT NULL,
	`fechaIngreso` DATETIME NULL,
	`estado` INT NULL,
	`fechaEstado` DATETIME NULL,
	PRIMARY KEY (`idMedico`)
);$$
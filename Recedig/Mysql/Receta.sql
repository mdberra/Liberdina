DELIMITER $$
USE recedig$$
DROP TABLE Receta$$
CREATE TABLE `recedig`.`Receta` (
	`idReceta` INT NOT NULL AUTO_INCREMENT,
	`fechaCreacion` DATETIME NOT NULL,
	`idMedico` INT NOT NULL,
	`estado` INT NULL,
	`fechaEstado` DATETIME NULL,
	PRIMARY KEY (`idReceta`,`fechaCreacion`,`idMedico`)
);$$
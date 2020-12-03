DELIMITER $$
USE recedig$$
DROP TABLE RecetaFarmacia$$
CREATE TABLE `recedig`.`RecetaFarmacia` (
	`idReceta` INT NOT NULL,
	`idFarmacia` INT NOT NULL,
	`estado` INT NULL,
	`fechaEstado` DATETIME NULL,
	PRIMARY KEY (`idReceta`,`idFarmacia`)
);$$
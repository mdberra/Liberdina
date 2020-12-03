DELIMITER $$
USE recedig$$
DROP TABLE RecetaItem$$
CREATE TABLE `recedig`.`RecetaItem` (
	`idRecetaItem` INT NOT NULL AUTO_INCREMENT,
	`idReceta` INT NOT NULL,
	`idMedicamento` INT NOT NULL,
	PRIMARY KEY (`idRecetaItem`,`idReceta`,`idMedicamento`)
);$$
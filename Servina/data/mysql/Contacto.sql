USE Pepeya;
DROP TABLE Contacto;
CREATE TABLE `Pepeya`.`Contacto` (
  `idContacto` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(255),
  `apellido` VARCHAR(255),
  `email` VARCHAR(255),
  `telefono` VARCHAR(255),
  `dni` INT,
  `monto` DOUBLE,
  `plazo` INT,
  `mensaje` VARCHAR(255),
  `cbu` VARCHAR(255),
  `idImagen` INT,
  `fechaIngreso` DATETIME,
  `estado` INT,
  `fechaEstado` DATETIME,
  `estadoDescrip` VARCHAR(255),
  PRIMARY KEY (`idContacto`));

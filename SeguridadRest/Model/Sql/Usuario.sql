USE Seguridad;
DROP TABLE Usuario;
CREATE TABLE `Seguridad`.`Usuario` (
  `idUsuario` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `fecha` DATETIME NOT NULL,
  PRIMARY KEY (`idUsuario`),
  UNIQUE KEY `idx_email` (`email`)
);
USE Seguridad;
DROP TABLE Actor;
CREATE TABLE `Seguridad`.`Actor` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ipAddress` VARCHAR(255) NOT NULL,
  `tipo` VARCHAR(255) NOT NULL,
  `dni` INT NOT NULL,
  `pin` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `fechaEnrolar` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_dni` (`dni`),
  UNIQUE KEY `idx_email` (`email`)
);

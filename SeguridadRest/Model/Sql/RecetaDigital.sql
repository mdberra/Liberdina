USE Seguridad;
DROP TABLE RecetaDigital;
CREATE TABLE `Seguridad`.`RecetaDigital` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `tipo` VARCHAR(255) NOT NULL,
  `IdMedico` INT NOT NULL,
  `IdPaciente` INT NOT NULL,
  `IdMedicamento` INT NOT NULL,
  `fecha` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_medico` (`IdMedico`),
  KEY `idx_paciente` (`IdPaciente`)
);
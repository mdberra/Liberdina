DELIMITER $$
USE recedig$$
DROP TABLE KeyValue$$
CREATE TABLE `recedig`.`KeyValue` (
  `idKeyValue` INT NOT NULL AUTO_INCREMENT,
  `entidad` VARCHAR(255) NOT NULL,
  `atributo` VARCHAR(255) NOT NULL,
  `idEstado` INT NOT NULL,
  `descripcion` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`idKeyValue`));$$
INSERT INTO KeyValue values(1, 'Receta', 'estado', 0, 'Solicitado')$$
INSERT INTO KeyValue values(2, 'Receta', 'estado', 1, 'Analizando')$$
INSERT INTO KeyValue values(3, 'Receta', 'estado', 2, 'Aprobado')$$
INSERT INTO KeyValue values(4, 'Receta', 'estado', 3, 'Rechazado')$$
INSERT INTO KeyValue values(5, 'Medico', 'estado', 0, 'Solicitado')$$
INSERT INTO KeyValue values(6, 'Medico', 'estado', 1, 'Analizando')$$
INSERT INTO KeyValue values(7, 'Medico', 'estado', 2, 'Aprobado')$$
INSERT INTO KeyValue values(8, 'Medico', 'estado', 3, 'Rechazado')$$
INSERT INTO KeyValue values(9, 'Farmacia', 'estado', 0, 'Solicitado')$$
INSERT INTO KeyValue values(10, 'Farmacia', 'estado', 1, 'Analizando')$$
INSERT INTO KeyValue values(11, 'Farmacia', 'estado', 2, 'Aprobado')$$
INSERT INTO KeyValue values(12, 'Farmacia', 'estado', 3, 'Rechazado')$$
INSERT INTO KeyValue values(13, 'Laboratorio', 'estado', 0, 'Solicitado')$$
INSERT INTO KeyValue values(14, 'Laboratorio', 'estado', 1, 'Analizando')$$
INSERT INTO KeyValue values(15, 'Laboratorio', 'estado', 2, 'Aprobado')$$
INSERT INTO KeyValue values(16, 'Laboratorio', 'estado', 3, 'Rechazado')$$
INSERT INTO KeyValue values(17, 'Medicamento', 'estado', 0, 'Solicitado')$$
INSERT INTO KeyValue values(18, 'Medicamento', 'estado', 1, 'Analizando')$$
INSERT INTO KeyValue values(19, 'Medicamento', 'estado', 2, 'Aprobado')$$
INSERT INTO KeyValue values(20, 'Medicamento', 'estado', 3, 'Rechazado')$$
INSERT INTO KeyValue values(21, 'RecetaFarmacia', 'estado', 0, 'Solicitado')$$
INSERT INTO KeyValue values(22, 'RecetaFarmacia', 'estado', 1, 'Analizando')$$
INSERT INTO KeyValue values(23, 'RecetaFarmacia', 'estado', 2, 'Aprobado')$$
INSERT INTO KeyValue values(24, 'RecetaFarmacia', 'estado', 3, 'Rechazado')$$
INSERT INTO KeyValue values(25, 'RecetaItem', 'estado', 0, 'Solicitado')$$
INSERT INTO KeyValue values(26, 'RecetaItem', 'estado', 1, 'Analizando')$$
INSERT INTO KeyValue values(27, 'RecetaItem', 'estado', 2, 'Aprobado')$$
INSERT INTO KeyValue values(28, 'RecetaItem', 'estado', 3, 'Rechazado')$$
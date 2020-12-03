//
//  entidad   atributo  idEstado    descripcion
//-----------------------------------------------------------
//  contacto  estado    0           Solicitado
//  contacto  estado    1           Analizando
//  contacto  estado    2           Aprobado
//  contacto  estado    3           Rechazado
//
USE Pepeya;
DROP TABLE KeyValue;
CREATE TABLE `Pepeya`.`KeyValue` (
  `idKeyValue` INT NOT NULL AUTO_INCREMENT,
  `entidad` VARCHAR(255) NOT NULL,
  `atributo` VARCHAR(255) NOT NULL,
  `idEstado` INT NOT NULL,
  `descripcion` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`idKeyValue`));
INSERT INTO KeyValue values(1, 'general', 'estado', 0, 'Activo');
INSERT INTO KeyValue values(2, 'general', 'estado', 1, 'Finalizado');
INSERT INTO KeyValue values(3, 'contacto', 'estado', 0, 'Solicitado');
INSERT INTO KeyValue values(4, 'contacto', 'estado', 1, 'Analizando');
INSERT INTO KeyValue values(5, 'contacto', 'estado', 2, 'Aprobado');
INSERT INTO KeyValue values(6, 'contacto', 'estado', 3, 'Rechazado');
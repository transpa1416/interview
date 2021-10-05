## a.- Script crear y poblar BD
## Validación si existe la BD
CREATE DATABASE IF NOT EXISTS interview;
##Importante para usar la BD correcta
USE interview;
##Validación y creación de tabla
DROP TABLE IF EXISTS Venta;
DROP TABLE IF EXISTS Maquinas_registradoras;
DROP TABLE IF EXISTS Productos;
DROP TABLE IF EXISTS Cajeros;
CREATE TABLE Cajeros 
(
	Cajero int unsigned not null primary key auto_increment,
    NomApels nvarchar(255) collate utf8_spanish_ci 
)engine=innodb;
##Validación y creación de tabla
CREATE TABLE Productos
(
	Producto int unsigned not null primary key auto_increment,
    Nombre nvarchar(100) collate utf8_spanish_ci ,
    Precio float
)engine=innodb;
##Validación y creación de tabla
CREATE TABLE Maquinas_registradoras
(
	Maquina int unsigned not null primary key auto_increment,
    Piso int
)engine=innodb;
##Validación y creación de tabla
CREATE TABLE Venta
(
	Cajero int unsigned not null,
    Producto int unsigned not null,
    Maquina int unsigned not null,
    foreign key (Cajero) references Cajeros(Cajero) on delete cascade,
    foreign key (Producto) references Productos(Producto) on delete cascade,
    foreign key (Maquina) references Maquinas_registradoras(Maquina) on delete cascade
)engine=innodb;
##Validamos si existe procedimiento para cargar información de prueba
DROP PROCEDURE IF EXISTS cargar_datos_prueba;
##Se declare delimitador, para que se construya el procedimiento almacenado
delimiter $
CREATE PROCEDURE cargar_datos_prueba()
BEGIN
    ##Declaramos variables de control
	DECLARE max_cashier int unsigned default 0;
	DECLARE max_products int unsigned default 0;
	DECLARE max_cash_register int unsigned default 0;
    DECLARE max_sale int unsigned default 0;
	DECLARE record_counter int unsigned default 0;

    ##Se asignan valores para poblar las tablas
    SET max_cashier = 10;
    SET max_products = 150;
    SET max_cash_register = 5;
    SET max_sale = 600;

    ##Para la creación de los datos, se crearán transacciones
    ##Poblado de la tabla Cajeros
    START TRANSACTION;
    SET record_counter = 0;
    WHILE record_counter < max_cashier DO
            INSERT INTO Cajeros (NomApels) VALUES (CONCAT('Cajero ', record_counter + 1));
        SET record_counter = record_counter + 1;
    END WHILE;
    COMMIT;
    ##Poblado de la tabla Productos
    START TRANSACTION;
    SET record_counter = 0;
    WHILE record_counter < max_products DO
            INSERT INTO Productos (Nombre, Precio) VALUES (CONCAT('Producto ', record_counter + 1), ((CAST(RAND() * 450 AS UNSIGNED) + 1)) + (ROUND(RAND() * 0.49 + 0.01, 2)));
        SET record_counter = record_counter + 1;
    END WHILE;
    COMMIT;

    ##Poblado de la tabla Maquinas_registradoras
    START TRANSACTION;
    SET record_counter = 0;
    WHILE record_counter < max_cash_register DO
            INSERT INTO Maquinas_registradoras (Piso) VALUES (MOD((record_counter + 1), 3) + 1);
        SET record_counter = record_counter + 1;
    END WHILE;
    COMMIT;
    
    ##Poblado de la tabla Venta
    START TRANSACTION;
    SET record_counter = 0;
    WHILE record_counter < max_sale DO
            INSERT INTO Venta (Cajero, Producto, Maquina) VALUES ((FLOOR( RAND() * (11-1) + 1)), FLOOR( RAND() * (150-1) + 1), FLOOR( RAND() * (5-1) + 1));
        SET record_counter = record_counter + 1;
    END WHILE;
    COMMIT;
    
END $ ##finaliza el delimitador
delimiter ;
##Llamado al procedimiento almacenado
CALL cargar_datos_prueba();

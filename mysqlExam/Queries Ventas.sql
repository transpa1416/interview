## b.- Número de ventas de cada producto, ordenado de mas a menos
SELECT p.Producto,
	p.Nombre,
	COUNT(v.Producto) 
FROM venta v
	INNER JOIN productos p ON v.Producto = p.Producto
GROUP BY p.Producto, p.Nombre
ORDER BY COUNT(v.Producto) DESC, p.Nombre DESC;
## c.- Informe de ventas indicando Cajero, Nombre y precios de productos Vendidos y piso de venta.
SELECT 
	c.NomApels,
    COUNT(p.Nombre) AS NoArt,
    p.Nombre,
    p.Precio,
    ROUND((COUNT(p.Nombre) * p.Precio),2) AS Importe,
    mq.Maquina,
    mq.Piso
FROM cajeros c
	LEFT JOIN venta v ON c.Cajero = v.Cajero
    LEFT JOIN productos p ON v.Producto = p.Producto
    LEFT JOIN maquinas_registradoras mq ON v.Maquina = mq.Maquina
WHERE c.Cajero = 6
GROUP BY c.NomApels, p.Nombre, p.Precio, mq.Maquina, mq.Piso
ORDER BY c.NomApels, p.Nombre, mq.piso, mq.Maquina;

## d.- Ventas totales realizadas en cada piso
SELECT 
	t.piso,
    ROUND(SUM(COALESCE(t.parcial, 0)), 2) AS totales
FROM (
	SELECT 
		mq.Piso,
		COUNT(p.Producto) * p.precio AS Parcial
	FROM maquinas_registradoras mq
		LEFT JOIN venta v ON mq.Maquina = v.Maquina
		LEFT JOIN productos p ON v.Producto = p.Producto
	GROUP BY mq.Piso, p.precio) AS t
GROUP BY t.piso;
## e.- Nombre y código de cada cajero con importe de ventas
SELECT 
	t.Cajero,
    t.NomApels,
    ROUND(SUM(COALESCE(t.parcial, 0)), 2) AS totales
FROM (
	SELECT 
		c.Cajero,
		c.NomApels,
		COUNT(p.Producto) * p.precio AS Parcial
	FROM cajeros c
		LEFT JOIN venta v ON c.Cajero = v.Cajero
		LEFT JOIN productos p ON v.Producto = p.Producto
	GROUP BY c.Cajero, c.NomApels, p.precio) AS t
GROUP BY t.Cajero, t.NomApels
ORDER BY t.NomApels;
#f.- Nombre y código de cada cajero con ventas en piso menores a 5000
SELECT 
	t.Cajero,
    t.NomApels,
    t.piso,
    ROUND(SUM(COALESCE(t.parcial, 0)), 2) AS totales
FROM (
	SELECT 
		c.Cajero,
		c.NomApels,
		mq.piso,
		COUNT(p.Producto) * p.precio AS Parcial
	FROM cajeros c
		LEFT JOIN venta v ON c.Cajero = v.Cajero
		LEFT JOIN productos p ON v.Producto = p.Producto
		LEFT JOIN maquinas_registradoras mq ON v.Maquina = mq.Maquina
	GROUP BY c.Cajero, c.NomApels, mq.piso, p.precio) AS t
GROUP BY t.Cajero, t.NomApels, t.piso
HAVING ROUND(SUM(COALESCE(t.parcial, 0)), 2) < 5000
ORDER BY t.NomApels, t.piso;
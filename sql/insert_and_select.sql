INSERT INTO public.items (customer_name,order_date,product,quantity,price)
	VALUES ('Thalia', '01/05/2023', 'ventilador', 2, 56.70),
	('Maria Emilia', '10/10/2020', 'licuadora', 1, 200.10),
	('Ramón', '06/04/2021', 'horno', 3, 340.09),
	('Gabriel', '09/07/2022', 'televisor', 1, 999.99),
	('Rodrigo', '03/02/2023', 'ventilador', 2, 56.70);

SELECT * FROM items
WHERE quantity > 2 AND price > 50;
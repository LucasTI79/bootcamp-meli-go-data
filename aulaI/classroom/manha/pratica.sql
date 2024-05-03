CREATE DATABASE IF NOT EXISTS pratica_banco_de_dados;
USE pratica_banco_de_dados;

-- Para comentar linhas, podemos usar dois traços '--'

-- varchar simboliza texto
-- decimal para representar float

-- aqui criamos uma tabela chamada produtos
DROP TABLE IF EXISTS products;
CREATE TABLE products (
 product_id INT AUTO_INCREMENT PRIMARY KEY,
 name VARCHAR(255),
 price DECIMAL(10, 2)
);

-- aqui vemos detalhes da estrutura da tabela
DESCRIBE products;

-- selecionando todas as colunas da tabela produtos
SELECT * FROM products;

-- inserindo valores na tabela produtos
INSERT INTO products VALUES(null, 'Produto 1', 20.00);

SELECT products.name FROM products;

-- inserindo valores em colunas expecíficas
INSERT INTO products(price,name) VALUES(20.00, 'Produto 1');

-- criando a tabela tipo de produto
DROP TABLE IF EXISTS product_type;
CREATE TABLE product_type (
	product_type_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255)
);

-- alterando estrutura da tabela products para adicionar uma coluna com chave estrangeira
ALTER TABLE products ADD COLUMN fk_product_type_id INT;
ALTER TABLE products ADD CONSTRAINT FOREIGN KEY(fk_product_type_id) REFERENCES product_type(product_type_id);
ALTER TABLE products MODIFY COLUMN fk_product_type_id INT NOT NULL;

-- se eu tentar inserir um dado na tabela sem product_type, é para dar erro
-- INSERT INTO products(price,name, fk_product_type_id) VALUES(20.00, 'Produto 1');
INSERT INTO product_type(name) VALUES('tipo qualquer');

SELECT * FROM product_type;

-- agora com o tipo do produto criado, posso pegar o id do tipo criado e usar como fk na tabela de produtos
INSERT INTO products(name, price, fk_product_type_id) VALUES('Produto qualquer', 20.00, 1);

-- conferindo se o registro foi criado na tabela products
SELECT * FROM products;

DROP TABLE IF EXISTS class;
CREATE TABLE class (
	class_id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255)
);

INSERT INTO class(name) VALUES('melhor turma do bootcamp de GO da alkemy');

SELECT * FROM class;

-- criando uma tabela já com fk
CREATE TABLE students(
	student_id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255),
    fk_class_id INT NOT NULL,
    FOREIGN KEY(fk_class_id) REFERENCES class(class_id)
);

INSERT INTO students
(name, fk_class_id) VALUES 
('Victor engraçadinho', 1),
('Natan', 1);

SELECT * FROM students;
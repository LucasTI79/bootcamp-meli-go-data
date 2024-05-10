package products

import (
	"database/sql"
	"log"

	"github.com/batatinha123/products-api/internal/entities"
)

const (
	GetAllProducts = "SELECT id, name, type, count, price FROM products"
	GetOneProduct  = "SELECT id, name, type, count, price FROM products WHERE id = ?"
	SaveProduct    = "INSERT INTO products(name, type, count, price) VALUES(?,?,?,?)"
	UpdateProduct  = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	DeleteProduct  = "DELETE FROM products WHERE id = ?"
)

type MySqlRepository struct {
	db *sql.DB
}

func (m *MySqlRepository) GetAll() ([]entities.Product, error) {
	var products []entities.Product

	rows, err := m.db.Query(GetAllProducts)
	if err != nil {
		log.Println(err)
		return products, err
	}

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Category,
			&product.Count,
			&product.Price,
		); err != nil {
			log.Println(err.Error())
			return products, nil
		}

		products = append(products, product)
	}

	// 204 -> No Content

	return products, nil
}

func (m *MySqlRepository) GetOne(id uint64) (entities.Product, error) {
	// | id | nome | categoria | estoque | preco
	// | 1  |  a   |    b      |    10   |  10.00
	// | 2  |  c   |    d      |    20   |  20.00
	// | 3  |  e   |    f      |    30   |  30.00

	var product entities.Product
	// aqui definimos a ordem que queremos que o banco nos retorne os dados
	rows, err := m.db.Query(GetOneProduct, id)
	if err != nil {
		log.Println(err)
		return product, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Category,
			&product.Count,
			&product.Price,
		); err != nil {
			log.Println(err.Error())
			return product, nil
		}
	}

	return product, nil
}

func (m *MySqlRepository) Store(name, category string, count int, price float64) (entities.Product, error) {
	product := entities.Product{
		Name:     name,
		Category: category,
		Count:    count,
		Price:    price,
	}

	// podemos também iniciar uma transação
	// m.db.Begin()

	// o banco é iniciado
	stmt, err := m.db.Prepare(SaveProduct) // monta o  SQL
	if err != nil {
		log.Fatal(err)
	}
	// **opcional**
	defer stmt.Close() // a instrução fecha quando termina. Se eles permanecerem abertos, o consumo de memória é gerado

	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Category, product.Count, product.Price) // retorna um sql.Result ou um error
	if err != nil {
		return entities.Product{}, err
	}

	insertedId, _ := result.LastInsertId() // do sql.Result retornado na execução obtemos o Id inserido
	product.ID = uint64(insertedId)

	return product, nil

}

func (m *MySqlRepository) Update(id uint64, name, productType string, count int, price float64) (entities.Product, error) {
	product := entities.Product{
		ID:       id,
		Name:     name,
		Category: productType,
		Count:    count,
		Price:    price,
	}
	stmt, err := m.db.Prepare(UpdateProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Category, product.Count, product.Price, product.ID)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (m *MySqlRepository) UpdateName(id uint64, name string) (entities.Product, error) {
	product := entities.Product{
		ID:   id,
		Name: name,
	}
	stmt, err := m.db.Prepare("UPDATE products SET name = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.ID)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (m *MySqlRepository) Delete(id uint64) error {
	stmt, err := m.db.Prepare(DeleteProduct)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// func (m *MySqlRepository) LastID() (uint64, error) {
// 	return 0, nil
// }

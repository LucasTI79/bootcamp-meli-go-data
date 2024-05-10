package products

import (
	"context"
	"database/sql"
	"log"

	"github.com/batatinha123/products-api/internal/entities"
)

const (
	GetAllProducts     = "SELECT id, name, type, count, price FROM products"
	GetOneProduct      = "SELECT id, name, type, count, price FROM products WHERE id = ?"
	GetOneProductSleep = "SELECT SLEEP(30) FROM DUAL where 0 < ?"
	SaveProduct        = "INSERT INTO products(name, type, count, price) VALUES(?,?,?,?)"
	UpdateProduct      = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	DeleteProduct      = "DELETE FROM products WHERE id = ?"
	GetProductFullData = "SELECT products.id, products.name, products.type, products.count, products.price, warehouses.name, warehouses.adress " +
		"FROM products INNER JOIN warehouses ON products.id_warehouse = warehouses.id " +
		"WHERE products.id = ?"
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
			return products, err
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

func (r *MySqlRepository) GetOneWithContext(ctx context.Context, id uint64) (entities.Product, error) {
	var product entities.Product
	// aqui definimos a ordem que queremos que o banco nos retorne os dados
	rows, err := r.db.QueryContext(ctx, GetOneProduct, id)
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
	return err
}

func (r *MySqlRepository) GetFullData(id uint64) (entities.ProductFullData, error) {
	var product entities.ProductFullData
	rows, err := r.db.Query(GetProductFullData, id)

	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Count, &product.Price, &product.Warehouse,
			&product.WarehouseAddress); err != nil {
			log.Fatal(err)
			return product, err
		}
	}
	return product, nil
}

// func (m *MySqlRepository) LastID() (uint64, error) {
// 	return 0, nil
// }

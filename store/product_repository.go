package store

import "github.com/ebhlz88/est-shop/models"

func (s *PostgresStore) CreateProductTable() error {
	query := `create table if not exists product(
	product_id serial primary key,
	product_name varchar(50),
	product_buy_price varchar(50),
	product_sell_price varchar(50)
	)
	`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateOrderTable() error {
	query := `create table if not exists orders (
	order_id serial primary key,
	product_id int references product(product_id),
	amount numeric,
	is_profit_distrubuted boolean
	)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateProfitTable() error {
	query := `create table if not exists profit(
	profit_id serial primary key,
	product_id int references product(product_id),
	profit_amount numeric,
	date timestamp
	)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateStockTable() error {
	query := `create table if not exists stock(
		stock_id serial primary key,
		product_id int references product(product_id),
		quantity numeric
		)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateInvestorTable() error {
	query := `create table if not exists investors(
	investor_id serial primary key,
	first_name varchar(50),
	last_name varchar(50),
	amount_invested numeric
)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) AddProduct(productName string, buyPrice, sellPrice int64) error {
	query := `INSERT INTO product (product_name, product_buy_price, product_sell_price)
			  VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, productName, buyPrice, sellPrice)
	return err
}

func (s *PostgresStore) DeleteProduct(productId int) error {
	query := `DELETE FROM product WHERE product_id = $1`
	_, err := s.db.Exec(query, productId)
	return err
}

func (s *PostgresStore) GetProductById(productId int) (models.Product, error) {
	var p models.Product
	query := `SELECT product_id, product_name, product_buy_price, product_sell_price
			  FROM product
			  WHERE product_id = $1`
	err := s.db.QueryRow(query, productId).Scan(&p.ProductId, &p.ProductName, &p.ProductBuyPrice, &p.ProductSellPrice)
	return p, err
}

func (s *PostgresStore) GetAllProducts() ([]models.Product, error) {
	query := `SELECT product_id, product_name, product_buy_price, product_sell_price
			  FROM product`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ProductId, &p.ProductName, &p.ProductBuyPrice, &p.ProductSellPrice)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (s *PostgresStore) ModifyProduct(productId int, productName string, productBuyPrice, productSellPrice int64) error {
	query := `UPDATE product 
			  SET product_name = $1, product_buy_price = $2, product_sell_price = $3
			  WHERE product_id = $4`
	_, err := s.db.Exec(query, productName, productBuyPrice, productSellPrice, productId)
	return err
}

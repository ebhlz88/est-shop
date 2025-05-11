package store

import "github.com/ebhlz88/est-shop/models"

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

func (s *PostgresStore) CreateOrder(productId int, amount int, isProfitDistributed bool) error {
	query := `insert into orders(product_id, amount, is_profit_distrubuted) 
	 VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, productId, amount, isProfitDistributed)
	return err
}

func (s *PostgresStore) GetAllOrders() ([]models.OrderWithProduct, error) {
	query := `
	SELECT 
		o.order_id, o.amount, o.is_profit_distrubuted,
		p.product_id, p.product_name, p.product_buy_price, p.product_sell_price
	FROM orders o
	JOIN product p ON o.product_id = p.product_id`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []models.OrderWithProduct
	for rows.Next() {
		var o models.OrderWithProduct
		var p models.Product
		err := rows.Scan(&o.Orderid, &o.Amount, &o.IsProfitDistributed, &p.ProductId, &p.ProductName, &p.ProductBuyPrice, &p.ProductSellPrice)
		if err != nil {
			return nil, err
		}
		o.ProductId = p
		orders = append(orders, o)
	}
	return orders, nil
}

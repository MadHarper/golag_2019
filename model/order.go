package model

import "time"

type Order struct{
	Id int
	Price int
	Product string
	Created_at time.Time
	Customer_id int
}

func (db *DB) AllOrders() ([]*Order, error) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ords := make([]*Order, 0)
	for rows.Next() {
		ord := new(Order)
		err := rows.Scan(&ord.Id, &ord.Price, &ord.Product, &ord.Created_at)
		if err != nil {
			return nil, err
		}
		ords = append(ords, ord)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ords, nil
}




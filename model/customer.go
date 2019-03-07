package model

type Customer struct{
	Id int
	Name string
}

func (db *DB) AllCustomers() ([]*Customer, error) {
	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customers := make([]*Customer, 0)
	for rows.Next() {
		customer := new(Customer)
		err := rows.Scan(&customer.Id, &customer.Name)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil

}



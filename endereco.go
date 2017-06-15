package main

// Endereco struct
type Endereco struct {
	ID     int    `json:"id"`
	UserID int    `json:"user"`
	Street string `json:"street"`
	Number int    `json:"number"`
}

// InsertEndereco to persist
func InsertEndereco(endereco *Endereco) error {
	var id int
	err := db.QueryRow(`INSERT INTO enderecos(user_id, street, number) VALUES ($1, $2, $3) RETURNING id`, endereco.UserID, endereco.Street, endereco.Number).Scan(&id)
	if err != nil {
		return err
	}
	endereco.ID = id
	return nil
}

// GetEnderecoByID function
func GetEnderecoByID(id int) (*Endereco, error) {
	var street string
	var number int
	err := db.QueryRow("SELECT street, number FROM enderecos WHERE id = $1", id).Scan(&street, &number)
	if err != nil {
		return nil, err
	}
	return &Endereco{
		ID:     id,
		Street: street,
		Number: number,
	}, nil
}

// RemoveEnderecoByID function
func RemoveEnderecoByID(id int) error {
	_, err := db.Exec("DELETE FROM enderecos WHERE id = $1", id)
	return err
}

// GetEnderecoByIDAndUser function
func GetEnderecoByIDAndUser(id, userID int) (*Endereco, error) {
	var street string
	var number int
	err := db.QueryRow(`
		SELECT street, number
		FROM enderecos
		WHERE id = $1
		AND user_id = $2
	`, id, userID).Scan(&street, &number)
	if err != nil {
		return nil, err
	}
	return &Endereco{
		ID:     id,
		UserID: userID,
		Street: street,
		Number: number,
	}, nil
}

// GetEnderecosForUser function
func GetEnderecosForUser(id int) ([]*Endereco, error) {
	rows, err := db.Query(`
		SELECT e.id, e.user_id, e.street, e.number
		FROM enderecos AS e
		WHERE e.user_id = $1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		enderecos = []*Endereco{}
		i         int
		userID    int
		street    string
		number    int
	)
	for rows.Next() {
		if err = rows.Scan(&i, &userID, &street, &number); err != nil {
			return nil, err
		}
		enderecos = append(enderecos, &Endereco{ID: i, UserID: userID, Street: street, Number: number})
	}
	return enderecos, nil
}

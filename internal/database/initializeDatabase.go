package database

import (
	"database/sql"
)

func InitializeDatabase(database *sql.DB) error {
	// Verifica a existência da tabela products
	query := `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'products')`
	var exists bool
	err := database.QueryRow(query).Scan(&exists)
	if err != nil {
		return err
	}

	// Se a tabela não existe, cria a tabela 'products'
	if !exists {
		createTableQuery := `
            CREATE TABLE products (
                id serial primary key,
                name varchar,
                description varchar,
                price decimal,
                quantity integer
            )
        `
		_, err := database.Exec(createTableQuery)
		if err != nil {
			return err
		}

		// Popula a tabela com dados de exemplo
		populateQuery := `
            INSERT INTO products (name, description, price, quantity)
            VALUES
            ('Camiseta', 'Camiseta de algodão', 19.99, 100),
            ('Calça Jeans', 'Calça jeans slim', 49.99, 50),
            ('Tênis Esportivo', 'Tênis esportivo confortável', 79.99, 30)
        `
		_, err = database.Exec(populateQuery)
		if err != nil {
			return err
		}
	}

	return nil
}

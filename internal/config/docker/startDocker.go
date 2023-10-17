package docker

import (
	"database/sql"
	"fmt"
	"os/exec"
	"time"

	_ "github.com/lib/pq"
)

func StartDatabaseContainer(db *sql.DB) error {
	fmt.Println("Starting the PostgreSQL Docker container...\n")
	cmd := exec.Command("docker", "compose", "-f", "compose.yaml", "--env-file", ".env", "up", "-d")
	cmd.Dir = "infra"

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error starting the Docker container: %v\n", err)
		return err
	}

	// Aguarde o PostgreSQL estar pronto para aceitar conexões
	fmt.Println("Waiting for PostgreSQL to be ready...\n")
	time.Sleep(10 * time.Second) // Ajuste o tempo de espera conforme necessário

	// Verifique o status do contêiner do Docker após a inicialização
	fmt.Println("Verifying the status of the PostgreSQL Docker container...\n")
	statusCmd := exec.Command("docker", "ps")
	statusCmd.Dir = "infra"
	if output, err := statusCmd.CombinedOutput(); err != nil {
		fmt.Printf("Error checking the Docker container status: %v\n", err)
	} else {
		fmt.Printf("Docker container status:\n%s\n", output)
	}

	// Verifique se a tabela 'products' existe no banco de dados
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'products')`
	if err := db.QueryRow(query).Scan(&exists); err != nil {
		fmt.Printf("Error checking for the 'products' table: %v\n", err)
		return err
	}

	if !exists {
		// A tabela 'products' não existe; crie-a
		createTableQuery := `
            CREATE TABLE products (
                id serial primary key,
                name varchar,
                description varchar,
                price decimal,
                quantity integer
            )
        `
		_, err := db.Exec(createTableQuery)
		if err != nil {
			fmt.Printf("Error creating the 'products' table: %v\n", err)
			return err
		}

		// Ingestão de 3 itens de exemplo
		populateQuery := `
		INSERT INTO products (name, description, price, quantity)
		VALUES
			('Camiseta', 'Camiseta de algodão branca', 19.99, 100),
			('Calça Jeans', 'Calça jeans azul slim', 49.99, 50),
			('Tênis Esportivo', 'Tênis esportivo preto confortável', 79.99, 30),
			('Vestido Floral', 'Vestido longo estampado', 29.99, 40),
			('Blusa de Moletom', 'Blusa de moletom cinza', 24.99, 60),
			('Bermuda Jeans', 'Bermuda jeans masculina', 34.99, 20),
			('Sapato Social', 'Sapato social marrom de couro', 59.99, 15),
			('Jaqueta de Couro', 'Jaqueta de couro preta', 99.99, 10);
        `
		_, err = db.Exec(populateQuery)
		if err != nil {
			fmt.Printf("Error ingesting data into the 'products' table: %v\n", err)
			return err
		}

		fmt.Println("The 'products' table has been created, and 3 items have been ingested.")
	} else {
		fmt.Println("The 'products' table already exists in the database.")
	}

	fmt.Println("PostgreSQL Docker container started successfully.\n")
	fmt.Println("Press Ctrl+C to leave...")

	return nil
}

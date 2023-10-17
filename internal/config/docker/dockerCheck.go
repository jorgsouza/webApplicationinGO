package docker

import (
	"fmt"
	"os/exec"
)

func IsPostgreSQLRunning() bool {
	// Implemente o código para verificar se o contêiner do PostgreSQL está em execução.
	// Você pode usar uma biblioteca de cliente Docker para fazer essa verificação.
	// Aqui, estou usando um exemplo simples com 'docker ps'.
	// Certifique-se de ajustar isso para atender às suas necessidades.
	cmd := exec.Command("docker", "ps")
	if err := cmd.Run(); err != nil {
		fmt.Println(cmd.Run())
		return false
	}
	return true
}

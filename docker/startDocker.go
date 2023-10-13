package infra

import "os/exec"

func StartDatabaseContainer() error {
	cmd := exec.Command("docker", "compose", "-f", "compose.yaml", "up", "-d")
	cmd.Dir = "infra"

	if err := cmd.Run(); err != nil {
		panic(err.Error())
	}
	return nil
}
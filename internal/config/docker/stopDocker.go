package docker

import "os/exec"

func StopDatabaseContainer() error {
	cmd := exec.Command("docker", "stop", "infra-database-1")
	cmd.Dir = "infra"

	if err := cmd.Run(); err != nil {
		panic(err.Error())
	}
	return nil
}

package watchdog

import (
	"fmt"
	"os"
	"os/signal"

	dkr "github.com/jorgsouza/webApplication/internal/config/docker"
	"github.com/jorgsouza/webApplication/internal/database"
)

func SetupSignalHandling() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func() {
		<-interrupt
		fmt.Println("\n Closing the application...")

		if err := dkr.StopDatabaseContainer(); err != nil {
			fmt.Println("Error when stopping the database container:", err)
		}

		os.Exit(0)
	}()
	if err := dkr.StartDatabaseContainer(database.DatabaseConnect()); err != nil {
		fmt.Println("Error when starting the database container:", err)
		os.Exit(1)
	}

}

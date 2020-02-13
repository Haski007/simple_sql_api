package usage

import (
	"os"
	"fmt"
)

// ShowHelpUsage print instructions of program usage!
func ShowHelpUsage() {
	fmt.Println("\nusage: ./quests ");
	fmt.Println("\t--help   -   show usage.");
	fmt.Println("\t--port   -   set a custom port.");
	os.Exit(-1)
}
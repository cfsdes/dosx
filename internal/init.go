package internal

import (
	"fmt"
	"github.com/fatih/color"
)

func Initialize(){
	Banner()
	StatusAPI()

	// Initial message
	Cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Printf("[%s] Attacking ...\n", Cyan("INF"))
}


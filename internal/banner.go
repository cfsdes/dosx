package internal

import (
	"fmt"

    "github.com/fatih/color"
)

func Banner() {
    // Print a colorful welcome message
    cyan := color.New(color.FgCyan, color.Bold).SprintFunc()

    nucke := fmt.Sprint(`
    ____  _____  ___  _  _ 
    (  _ \(  _  )/ __)( \/ )
     )(_) ))(_)( \__ \ )  ( 
    (____/(_____)(___/(_/\_)
`)

    fmt.Printf("\n%s\n", cyan(nucke))

    fmt.Println()
}
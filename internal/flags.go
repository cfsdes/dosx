package internal

import (
	"fmt"
    "flag"
    "os"

    "github.com/fatih/color"
)

type headersFlag []string

func (h *headersFlag) String() string {
	return fmt.Sprintf("%v", *h)
}

func (h *headersFlag) Set(value string) error {
	*h = append(*h, value)
	return nil
}

func ParseFlags() (url string, threads int, headers headersFlag) {
	flag.StringVar(&url, "url", "", "Target URL (e.g.: https://example.com/)")
    flag.IntVar(&threads, "threads", 10, "Threads to use during plugin scan (default: 10)")
    flag.Var(&headers, "headers", "Set custom headers. Accept multiple flag usages.") // Accept multiple flag usages

    // Add the welcome message to the --help output
	flag.Usage = func() {
        Cyan := color.New(color.FgCyan, color.Bold)
        Cyan.Printf("Usage: \n")
        fmt.Fprintf(flag.CommandLine.Output(), "  %s -url https://example.com/ -threads 100 -headers 'User-Agent: Test'\n\n", os.Args[0])
        //Cyan.Printf("Flags: \n")
		PrintFlagsByTopic() // Imprime as flags por tópico
	}

    flag.Parse()

    return
}

func PrintFlagsByTopic() {
    Cyan := color.New(color.FgCyan, color.Bold)

    // Define os tópicos e as flags correspondentes
    topics := map[string][]string{
        "Flags": []string{"url", "threads", "headers"},
    }

    // Imprime as flags por tópico
    for topic, flags := range topics {
        Cyan.Printf("%s:\n", topic)
        for _, name := range flags {
            f := flag.Lookup(name)
            flagText := fmt.Sprintf("-%s", name)
            fmt.Printf("  %-25s %s\n", flagText, f.Usage)
        }
        fmt.Println()
    }
}

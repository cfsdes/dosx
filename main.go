package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
	"strings"

	"github.com/cfsdes/dosx/internal"
	"github.com/fatih/color"
)

func main() {
	
	if internal.Url == "" {
		Red := color.New(color.FgRed, color.Bold).SprintFunc()
		fmt.Printf("[%s] URL cannot be null\n", Red("ERR"))
		return 
	}

	// Initialize banner and status server
	internal.Initialize()

	client := &http.Client{
		Timeout: 0 * time.Second, // Configurando o timeout para zero segundos
	}
	
	var requestCounter int64
	var startTime time.Time
	var mu sync.Mutex // Mutex para proteger o acesso ao mapa de status code

	// Iniciar temporizador
	startTime = time.Now()

	for {
		var wg sync.WaitGroup

		for i := 0; i < internal.Threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
	
				// Montando a URL da requisição
				url := internal.Url

				// Criar uma nova requisição com os cabeçalhos personalizados
				req, err := http.NewRequest(internal.Method, url, nil)
				if err != nil {
					fmt.Printf("Erro ao criar a requisição: %v\n", err)
					return
				}

				// Adicionar headers personalizados à requisição
				for _, header := range internal.Headers {
					parts := strings.SplitN(header, ":", 2)
					if len(parts) == 2 {
						req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
					}
				}

				// Enviando a solicitação GET
				resp, err := client.Do(req)
				if err != nil {
					fmt.Printf("Erro ao enviar a solicitação: %v\n", err)
					return
				}
	
				defer resp.Body.Close()
	
				// Adicionar o código de status ao mapa com mutex para proteção
				mu.Lock()
				internal.StatusCodes[fmt.Sprintf("%d", resp.StatusCode)]++
				mu.Unlock()

				//requestCounter++
				atomic.AddInt64(&requestCounter, 1)
			}()
		}
	
		// Esperando todas as goroutines terminarem
		wg.Wait()

		// Calcular o tempo decorrido
		elapsedTime := time.Since(startTime)

		// Calcular a taxa de solicitações por segundo
		internal.RPS = float64(requestCounter) / elapsedTime.Seconds()

		// Reiniciar o contador e o temporizador
		requestCounter = 0
		startTime = time.Now()
	}
}

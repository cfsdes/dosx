package internal

var Url string 				// Port that nucke will listen
var Method string 			// HTTP Method to use
var Threads int 			// Nucke scan threads
var Headers []string 		// Custom Headers
var RPS float64				// Requests per second
var StatusCodes = make(map[string]int)

func init() {
	Url, Method, Threads, Headers = ParseFlags()
}
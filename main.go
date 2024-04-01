package main

import (
	"github.com/extism/go-pdk"
)

const sparkregHost = "https://sparkregistry-adhbvfqdeq-uc.a.run.app"

//export list
func list() int32 {
	// Create an HTTP Request to sparkreg.com/list
	req := pdk.NewHTTPRequest(pdk.MethodGet, sparkregHost+"/list")
	req.SetHeader("Accept", "application/json")

	// Send the request and get the response back
	res := req.Send()

	if res.Status() != 200 {
		pdk.SetErrorString("Failed to GET /list")
		return 1
	}

	// Output the response body directly
	pdk.OutputMemory(res.Memory())

	return 0
}

func main() {}

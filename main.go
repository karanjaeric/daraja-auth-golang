package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials", nil)
	if err != nil {
		log.Fatalf("error reating http request %v", err)
	}
	req.SetBasicAuth("qsx5xN3pixL9VFbR4XjNG4VOLF53GZJmGY2R5ClTLInNrUj9", "PncEgqPiyEKBb1co01CGYxGXnJrR8IpaiotKOf6NQlVcGmWGENAAQiq8TbSizeZ2")
	res, err := httpClient.Do(req)

	if err != nil {
		log.Fatalf("error invoking daraja APIs %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error extracting body %v", err)
	}
	fmt.Println("Auth Response is ", string(body))

}

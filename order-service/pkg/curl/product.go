package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ProductResponse struct {
	Data   Product `json:"data"`
	Status int     `json:"status"`
}

type Product struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Quantity int32    `json:"Quantity"`
}

func DetailProduct(id string) (*Product, error) {
	resp, err := http.Get("http://localhost:8002/api/v1/products/" + id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
    body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var productResp ProductResponse
	err = json.Unmarshal(body, &productResp)
	if err != nil {
		return nil, err
	}

	return &productResp.Data, nil
}

func UpdateProductStock(productID string, quantity int) bool {
    url := "http://localhost:8002/api/v1/products/" + productID

    jsonStr := []byte(fmt.Sprintf(`{"quantity": %d}`, quantity))

    req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonStr))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return false
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return false
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
    return true
}


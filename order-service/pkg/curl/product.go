package curl

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func DetailProduct(id string) (bool, error) {
	resp, err := http.Get("http://localhost:8002/api/v1/products/" + id)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	if resp.StatusCode != http.StatusOK {
		return false, err
	}

	return true, nil
}

func UpdateProductStock(productID string, quantity int) bool {
	// Chuẩn bị yêu cầu JSON
	body, err := json.Marshal(map[string]interface{}{
		"product_id": productID,
		"quantity":   quantity,
	})
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Gửi yêu cầu POST tới Product Service
	resp, err := http.Post("http://localhost:8081/update_stock", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/debug-ing/arvan-task/discount/config"
	"github.com/labstack/echo/v4"
)

type WalletRepository struct {
	config *config.AppConfig
}
type ResponseModel struct {
	Status bool `json:"status"`
}

type ChargeRequest struct {
	Amount float32 `json:"amount"`
}

func NewWalletRepository(config *config.AppConfig) *WalletRepository {
	repo := &WalletRepository{
		config: config,
	}
	return repo
}

func (w *WalletRepository) ChargeWallet(ctx echo.Context, userID string, price float32) (interface{}, error) {
	//convert float to string
	fmt.Printf("price: %f\n", price)
	url := w.config.WalletURL + "/wallet/" + userID + "/charge"
	chargeRequest := ChargeRequest{
		Amount: price,
	}
	payload, err := json.Marshal(chargeRequest)
	if err != nil {
		return nil, fmt.Errorf("error encoding JSON: %w", err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var model ResponseModel
	err = json.Unmarshal(body, &model)
	if err != nil {
		return false, err
	}

	return true, nil
}

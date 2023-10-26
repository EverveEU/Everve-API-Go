package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EverveAPI struct {
	APIKey  string
	BaseURL string
	Format  string
}

func (api *EverveAPI) MakeRequest(endpoint string, params map[string]string) ([]byte, error) {
	params["api_key"] = api.APIKey
	params["format"] = api.Format
	queryParams := url.Values{}
	for k, v := range params {
		queryParams.Add(k, v)
	}
	resp, err := http.Get(api.BaseURL + endpoint + "?" + queryParams.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (api *EverveAPI) GetUser() ([]byte, error) {
	return api.MakeRequest("user", map[string]string{})
}

func (api *EverveAPI) GetSocials() ([]byte, error) {
	return api.MakeRequest("socials", map[string]string{})
}

func (api *EverveAPI) GetCategories(id string) ([]byte, error) {
	endpoint := "categories"
	if id != "" {
		endpoint += "/" + id
	}
	return api.MakeRequest(endpoint, map[string]string{})
}

func (api *EverveAPI) CreateOrder(params map[string]string) ([]byte, error) {
	return api.MakeRequest("orders", params)
}

func (api *EverveAPI) GetOrders(id string) ([]byte, error) {
	endpoint := "orders"
	if id != "" {
		endpoint += "/" + id
	}
	return api.MakeRequest(endpoint, map[string]string{})
}

func (api *EverveAPI) UpdateOrder(id string, params map[string]string) ([]byte, error) {
	return api.MakeRequest("orders/"+id, params)
}

func (api *EverveAPI) DeleteOrder(id string) ([]byte, error) {
	return api.MakeRequest("orders/"+id, map[string]string{"_method": "DELETE"})
}

// EXAMPLE
func main() {
	api := EverveAPI{
		APIKey:  "your_api_key_here",
		BaseURL: "https://api.everve.net/v3/",
		Format:  "json",
	}

	userInfo, _ := api.GetUser()
	fmt.Println("User Info:", string(userInfo))

	socials, _ := api.GetSocials()
	fmt.Println("Socials:", string(socials))

	categories, _ := api.GetCategories("")
	fmt.Println("Categories:", string(categories))

	newOrder, _ := api.CreateOrder(map[string]string{"param1": "value1"})
	fmt.Println("New Order:", string(newOrder))

	orders, _ := api.GetOrders("")
	fmt.Println("Orders:", string(orders))

	updatedOrder, _ := api.UpdateOrder("1", map[string]string{"param1": "newValue1"})
	fmt.Println("Updated Order:", string(updatedOrder))

	deletedOrder, _ := api.DeleteOrder("1")
	fmt.Println("Deleted Order:", string(deletedOrder))
}

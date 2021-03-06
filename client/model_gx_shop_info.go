/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxShopInfo struct {
	Url string `json:"url"`
	ShopName string `json:"shopName"`
	Owner string `json:"owner"`
	Company string `json:"company"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Street string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	Postcode string `json:"postcode"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	Telephone string `json:"telephone"`
	Fax string `json:"fax"`
	Email string `json:"email"`
	ZoneId string `json:"zoneId"`
	CountryId string `json:"countryId"`
	Template string `json:"template"`
	ShopVersion string `json:"shopVersion"`
}

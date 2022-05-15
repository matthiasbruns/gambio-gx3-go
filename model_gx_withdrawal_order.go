/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gambio

type GxWithdrawalOrder struct {
	OrderId int64 `json:"orderId,omitempty"`
	CustomerId int64 `json:"customerId"`
	CustomerGender string `json:"customerGender"`
	CustomerFirstname string `json:"customerFirstname"`
	CustomerLastName string `json:"customerLastName"`
	CustomerStreetAddress string `json:"customerStreetAddress"`
	CustomerPostCode string `json:"customerPostCode"`
	CustomerCity string `json:"customerCity"`
	CustomerCountry string `json:"customerCountry"`
	CustomerEmail string `json:"customerEmail"`
	OrderDate string `json:"orderDate"`
	DeliveryDate string `json:"deliveryDate"`
}
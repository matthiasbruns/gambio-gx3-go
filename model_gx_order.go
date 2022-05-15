/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxOrder struct {
	Id int64 `json:"id,omitempty"`
	StatusId int64 `json:"statusId"`
	PurchaseDate string `json:"purchaseDate"`
	CurrencyCode string `json:"currencyCode"`
	LanguageCode string `json:"languageCode"`
	Comment string `json:"comment"`
	TotalWeight int64 `json:"totalWeight"`
	PaymentType *GxTitleAndModuleType `json:"paymentType"`
	ShippingType *GxTitleAndModuleType `json:"shippingType"`
	Customer *GxOrderCustomer `json:"customer"`
	Addresses *GxOrderAddressMain `json:"addresses"`
	Items []GxOrderItem `json:"items"`
	Totals []GxOrderTotal `json:"totals"`
	StatusHistory []GxOrderStatusHistory `json:"statusHistory"`
	AddonValues *GxOrderAddonValues `json:"addonValues"`
}

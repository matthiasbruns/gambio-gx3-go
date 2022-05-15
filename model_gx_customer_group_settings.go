/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gambio

type GxCustomerGroupSettings struct {
	Public bool `json:"public"`
	OtDiscountFlag bool `json:"otDiscountFlag"`
	GraduatedPrices bool `json:"graduatedPrices"`
	ShowPrice bool `json:"showPrice"`
	ShowPriceTax bool `json:"showPriceTax"`
	AddTaxOt bool `json:"addTaxOt"`
	DiscountAttributes bool `json:"discountAttributes"`
	Fsk18 bool `json:"fsk18"`
	Fsk18Display bool `json:"fsk18Display"`
	WriteReviews bool `json:"writeReviews"`
	ReadReviews bool `json:"readReviews"`
}

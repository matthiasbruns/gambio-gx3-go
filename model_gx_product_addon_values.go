/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gambio

type GxProductAddonValues struct {
	ProductsImageWidth string `json:"productsImageWidth"`
	ProductsImageHeight string `json:"productsImageHeight"`
	CodeIsbn string `json:"codeIsbn,omitempty"`
	CodeUpc string `json:"codeUpc,omitempty"`
	CodeMpn string `json:"codeMpn,omitempty"`
	CodeJan string `json:"codeJan,omitempty"`
	GoogleExportCondition string `json:"googleExportCondition,omitempty"`
	GoogleExportAvailabilityId string `json:"googleExportAvailabilityId,omitempty"`
	BrandName string `json:"brandName,omitempty"`
	IdentifierExists string `json:"identifierExists,omitempty"`
	Gender string `json:"gender,omitempty"`
	AgeGroup string `json:"ageGroup,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
}

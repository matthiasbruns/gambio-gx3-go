/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxCustomer struct {
	Id int64 `json:"id,omitempty"`
	Number string `json:"number"`
	Gender string `json:"gender"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	DateOfBirth string `json:"dateOfBirth"`
	VatNumber string `json:"vatNumber"`
	VatNumberStatus int64 `json:"vatNumberStatus"`
	Telephone string `json:"telephone"`
	Fax string `json:"fax"`
	Email string `json:"email"`
	StatusId int64 `json:"statusId"`
	IsGuest bool `json:"isGuest"`
	AddressId int64 `json:"addressId"`
	AddonValues string `json:"addonValues"`
}

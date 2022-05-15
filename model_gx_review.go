/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxReview struct {
	Id int64 `json:"id,omitempty"`
	Productid int64 `json:"productid,omitempty"`
	Rating int64 `json:"rating"`
	ReadCount int64 `json:"readCount"`
	DateAdded string `json:"dateAdded"`
	LastModified string `json:"lastModified"`
	LanguageId int64 `json:"languageId"`
	Text string `json:"text"`
	Customer *GxReviewCustomer `json:"customer"`
}

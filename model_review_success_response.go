/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type ReviewSuccessResponse struct {
	Code int64 `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Action string `json:"action,omitempty"`
	Resource string `json:"resource,omitempty"`
	ReviewId int64 `json:"reviewId,omitempty"`
}

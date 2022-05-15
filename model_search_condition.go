/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type SearchCondition struct {
	// Contains the search condition. Further information can be found on the <a href='/gambio-gx3-api/guides/search-example'>Search Example</a> page.
	Search []string `json:"search"`
}

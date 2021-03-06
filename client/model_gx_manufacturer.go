/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxManufacturer struct {
	Id           int64               `json:"id,omitempty"`
	Name         string              `json:"name"`
	Image        string              `json:"image,omitempty"`
	DateAdded    *GxManufacturerDate `json:"dateAdded,omitempty"`
	LastModified *GxManufacturerDate `json:"lastModified,omitempty"`
	Urls         *GxMultiLangOption  `json:"urls,omitempty"`
}

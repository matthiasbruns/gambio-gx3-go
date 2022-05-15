/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxCategoryListingEntity struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"`
	IsActive bool `json:"isActive"`
	Name string `json:"name"`
	HeadingTitle string `json:"headingTitle"`
	Description string `json:"description"`
	UrlKeywords string `json:"urlKeywords"`
	Icon string `json:"icon"`
	Image string `json:"image"`
	ImageAltText string `json:"imageAltText"`
	Links []string `json:"_links,omitempty"`
}

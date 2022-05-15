/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxCategory struct {
	Id              int64                  `json:"id,omitempty"`
	DateAdded       string                 `json:"dateAdded"`
	LastModified    string                 `json:"lastModified"`
	ParentId        int64                  `json:"parentId"`
	IsActive        bool                   `json:"isActive"`
	SortOrder       int64                  `json:"sortOrder"`
	Name            *GxMultiLangOption     `json:"name"`
	HeadingTitle    *GxMultiLangOption     `json:"headingTitle"`
	Description     *GxMultiLangOption     `json:"description"`
	MetaTitle       *GxMultiLangOption     `json:"metaTitle"`
	MetaDescription *GxMultiLangOption     `json:"metaDescription"`
	MetaKeywords    *GxMultiLangOption     `json:"metaKeywords"`
	UrlKeywords     *GxMultiLangOption     `json:"urlKeywords"`
	Icon            string                 `json:"icon"`
	Image           string                 `json:"image"`
	ImageAltText    *GxMultiLangOption     `json:"imageAltText"`
	Settings        *CategorySettings      `json:"settings"`
	AddonValues     *GxCategoryAddonValues `json:"addonValues"`
}

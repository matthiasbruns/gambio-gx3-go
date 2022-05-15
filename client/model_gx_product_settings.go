/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxProductSettings struct {
	DetailsTemplate        string `json:"detailsTemplate"`
	OptionsDetailsTemplate string `json:"optionsDetailsTemplate"`
	OptionsListingTemplate string `json:"optionsListingTemplate"`
	ShowOnStartpage        bool   `json:"showOnStartpage"`
	ShowQuantityInfo       bool   `json:"showQuantityInfo"`
	ShowWeight             bool   `json:"showWeight"`
	ShowPriceOffer         bool   `json:"showPriceOffer"`
	ShowAddedDateTime      bool   `json:"showAddedDateTime"`
	// Provide: \"0\" for normal, \"1\" for price on request and \"2\" for not available for purchase.
	PriceStatus                       int64                              `json:"priceStatus"`
	MinOrder                          int64                              `json:"minOrder"`
	GraduatedQuantity                 int64                              `json:"graduatedQuantity"`
	OnSitemap                         bool                               `json:"onSitemap"`
	SitemapPriority                   string                             `json:"sitemapPriority"`
	SitemapChangeFrequency            string                             `json:"sitemapChangeFrequency"`
	PropertiesDropdownMode            string                             `json:"propertiesDropdownMode"`
	StartpageSortOrder                int64                              `json:"startpageSortOrder"`
	ShowPropertiesPrice               bool                               `json:"showPropertiesPrice"`
	PropertiesCombisQuantityCheckMode int64                              `json:"propertiesCombisQuantityCheckMode"`
	UsePropertiesCombisShippingTime   bool                               `json:"usePropertiesCombisShippingTime"`
	UsePropertiesCombisWeight         bool                               `json:"usePropertiesCombisWeight"`
	GroupPermissions                  []GxProductSettingsGroupPermission `json:"groupPermissions"`
}

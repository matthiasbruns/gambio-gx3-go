/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxProduct struct {
	Id            int64  `json:"id,omitempty"`
	IsActive      bool   `json:"isActive"`
	SortOrder     int64  `json:"sortOrder"`
	DateAdded     string `json:"dateAdded"`
	DateAvailable string `json:"dateAvailable"`
	LastModified  string `json:"lastModified"`
	OrderedCount  int64  `json:"orderedCount"`
	ProductModel  string `json:"productModel"`
	Ean           string `json:"ean"`
	// The product's net price
	Price               float32               `json:"price"`
	DiscountAllowed     int64                 `json:"discountAllowed"`
	TaxClassId          int64                 `json:"taxClassId"`
	Quantity            int64                 `json:"quantity"`
	Weight              float32               `json:"weight"`
	ShippingCosts       float32               `json:"shippingCosts"`
	ShippingTimeId      int64                 `json:"shippingTimeId"`
	ProductTypeId       int64                 `json:"productTypeId"`
	ManufacturerId      int64                 `json:"manufacturerId"`
	IsFsk18             bool                  `json:"isFsk18"`
	IsVpeActive         bool                  `json:"isVpeActive"`
	VpeId               int64                 `json:"vpeId"`
	VpeValue            int64                 `json:"vpeValue"`
	Name                *GxMultiLangOption    `json:"name"`
	Description         *GxMultiLangOption    `json:"description"`
	ShortDescription    *GxMultiLangOption    `json:"shortDescription"`
	Keywords            *GxMultiLangOption    `json:"keywords"`
	MetaTitle           *GxMultiLangOption    `json:"metaTitle"`
	MetaDescription     *GxMultiLangOption    `json:"metaDescription"`
	MetaKeywords        *GxMultiLangOption    `json:"metaKeywords"`
	Url                 *GxMultiLangOption    `json:"url"`
	InfoUrl             *GxMultiLangOption    `json:"infoUrl"`
	UrlKeywords         *GxMultiLangOption    `json:"urlKeywords"`
	CheckoutInformation *GxMultiLangOption    `json:"checkoutInformation"`
	ViewedCount         *GxMultiLangIntOption `json:"viewedCount"`
	Images              []GxProductImage      `json:"images,omitempty"`
	AddonValues         *GxProductAddonValues `json:"addonValues"`
	Settings            *GxProductSettings    `json:"settings"`
	QuantityUnitId      int64                 `json:"quantityUnitId"`
	SpecialOfferId      int64                 `json:"specialOfferId,omitempty"`
}

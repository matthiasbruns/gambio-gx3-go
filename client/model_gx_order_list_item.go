/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GxOrderListItem struct {
	Id                 int64                 `json:"id,omitempty"`
	StatusId           int64                 `json:"statusId"`
	StatusName         string                `json:"statusName"`
	TotalSum           string                `json:"totalSum"`
	PurchaseDate       string                `json:"purchaseDate"`
	Comment            string                `json:"comment"`
	WithdrawalIds      []int64               `json:"withdrawalIds"`
	MailStatus         bool                  `json:"mailStatus"`
	CustomerId         int64                 `json:"customerId"`
	CustomerName       string                `json:"customerName"`
	CustomerEmail      string                `json:"customerEmail"`
	CustomerStatusId   int64                 `json:"customerStatusId"`
	CustomerStatusName string                `json:"customerStatusName"`
	CustomerMemos      []GxCustomerMemo      `json:"customerMemos"`
	DeliveryAddress    *GxOrderAddress       `json:"deliveryAddress"`
	BillingAddress     *GxOrderAddress       `json:"billingAddress,omitempty"`
	PaymentType        *GxTitleAndModuleType `json:"paymentType"`
	ShippingType       *GxTitleAndModuleType `json:"shippingType"`
	TrackingLinks      []string              `json:"trackingLinks"`
}
# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrder**](OrdersApi.md#CreateOrder) | **Post** /orders | Create Order
[**CreateOrderItem**](OrdersApi.md#CreateOrderItem) | **Post** /orders/{order_id}/items | Create Order Item
[**CreateOrderTimeAttribute**](OrdersApi.md#CreateOrderTimeAttribute) | **Post** /orders/{order_id}/items/{item_id}/attributes | Create Order Item Attribute
[**CreateOrderTotal**](OrdersApi.md#CreateOrderTotal) | **Post** /orders/{order_id}/totals | Create Order Total
[**DeleteOrderItems**](OrdersApi.md#DeleteOrderItems) | **Delete** /orders/{order_id}/items/{item_id} | Delete Order Item
[**DeleteOrderTotals**](OrdersApi.md#DeleteOrderTotals) | **Delete** /orders/{order_id}/totals/{total_id} | Delete Order Total
[**DeleteOrders**](OrdersApi.md#DeleteOrders) | **Delete** /orders/{order_id} | Delete Order
[**GetOrderHistories**](OrdersApi.md#GetOrderHistories) | **Get** /orders/{order_id}/history | Get Order History (Multiple)
[**GetOrderHistory**](OrdersApi.md#GetOrderHistory) | **Get** /orders/{order_id}/history/{history_id} | Get Order History (Single)
[**GetOrderItem**](OrdersApi.md#GetOrderItem) | **Get** /orders/{order_id}/items/{item_id} | Get Order Item (Single)
[**GetOrderItems**](OrdersApi.md#GetOrderItems) | **Get** /orders/{order_id}/items | Get Order Items (Multiple)
[**GetOrderPaymentDetails**](OrdersApi.md#GetOrderPaymentDetails) | **Get** /orders/{order_id}/payment_details | Get Order Payment Details
[**GetOrderTotal**](OrdersApi.md#GetOrderTotal) | **Get** /orders/{order_id}/totals/{totals_id} | Get Order Total
[**GetOrderTotals**](OrdersApi.md#GetOrderTotals) | **Get** /orders/{order_id}/totals | Get Order Totals (Multiple)
[**GetOrderTrackingCodes**](OrdersApi.md#GetOrderTrackingCodes) | **Get** /orders/{order_id}/tracking_codes | Get Order Tracking Codes (Multiple)
[**GetOrderTrackingCodesDeprecated**](OrdersApi.md#GetOrderTrackingCodesDeprecated) | **Get** /orders/{order_id}/tracking_code | Get Order Tracking Codes (Multiple) (deprecated)
[**GetOrders**](OrdersApi.md#GetOrders) | **Get** /orders/{order_id} | Get Order (Single)
[**GetOrdersItemsAttributeProperty**](OrdersApi.md#GetOrdersItemsAttributeProperty) | **Get** /orders/{order_id}/items/{item_id}/attributes/{attribute_id}/Property | Get Order Item Attribute/Property (Single)
[**GetOrdersItemsAttributesProperty**](OrdersApi.md#GetOrdersItemsAttributesProperty) | **Get** /orders/{order_id}/items/{item_id}/attributes/Property | Get Order Item Attributes/Properties (Multiple)
[**GetOrdersMultiple**](OrdersApi.md#GetOrdersMultiple) | **Get** /orders | Get Orders (Multiple)
[**PostTrackingCode**](OrdersApi.md#PostTrackingCode) | **Post** /orders/{order_id}/tracking_codes | Create Order Tracking Code
[**PostTrackingCodeDeprecated**](OrdersApi.md#PostTrackingCodeDeprecated) | **Post** /orders/{order_id}/tracking_code | Create Order Tracking Code (deprecated)
[**RemoveOrdersItemsAttribute**](OrdersApi.md#RemoveOrdersItemsAttribute) | **Delete** /orders/{order_id}/items/{item_id}/attributes/{attribute_id}/Property | Delete Order Item Attribute/Property
[**SearchOrders**](OrdersApi.md#SearchOrders) | **Post** /orders/search | Search Orders
[**UpdateOrder**](OrdersApi.md#UpdateOrder) | **Put** /orders/{order_id} | Update Order
[**UpdateOrderItems**](OrdersApi.md#UpdateOrderItems) | **Put** /orders/{order_id}/items/{item_id} | Update Order Item
[**UpdateOrderStatus**](OrdersApi.md#UpdateOrderStatus) | **Patch** /orders/{order_id}/status | Update Order Status
[**UpdateOrderTotal**](OrdersApi.md#UpdateOrderTotal) | **Put** /orders/{order_id}/totals/{total_id} | Update Order Total
[**UpdateOrdersItemsAttributesProperty**](OrdersApi.md#UpdateOrdersItemsAttributesProperty) | **Put** /orders/{order_id}/items/{item_id}/attributes/{attribute_id}/Property | Update Order Item Attribute/Property

# **CreateOrder**
> GxOrder CreateOrder(ctx, body)
Create Order

This method enables the creation of a new order into the system. The order can be bound to an existing customer or be standalone as implemented in the OrderService. Make sure that you check the Order resource representation. To see an example usage take a look at `docs/REST/samples/order-service/create_order.php`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrder**](GxOrder.md)| The order | 

### Return type

[**GxOrder**](GXOrder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrderItem**
> GxOrderItem CreateOrderItem(ctx, body, orderId)
Create Order Item

Use this method to create a new order item to an existing order. The order item JSON format must be the same with the \"items\" entries in the original order item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderItem**](GxOrderItem.md)| The order item | 
  **orderId** | **int64**| The order ID | 

### Return type

[**GxOrderItem**](GXOrderItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrderTimeAttribute**
> GxOrderItemAttributes CreateOrderTimeAttribute(ctx, body, orderId, itemId)
Create Order Item Attribute

Use this method to create a new order item attribute to an existing order item. The order item attribute JSON object is the same as the one included in the full order representation. There are two different order item variation systems in the shop, the \"attributes\" and the \"properties\". Both of them define a variation of an order item (e.g. color, size etc). You must always use only one of them for a single order item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderItemAttributes**](GxOrderItemAttributes.md)| The order item attributes | 
  **orderId** | **int64**| The order ID | 
  **itemId** | **int64**| The item ID | 

### Return type

[**GxOrderItemAttributes**](GXOrderItemAttributes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrderTotal**
> GxOrderTotal CreateOrderTotal(ctx, body, orderId)
Create Order Total

Creates a new order total entry to the existing order. The order total JSON format must be the same with the \"totals\" entries in the original order total.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderTotal**](GxOrderTotal.md)| The order total | 
  **orderId** | **int64**| The order ID | 

### Return type

[**GxOrderTotal**](GXOrderTotal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrderItems**
> DefaultSuccessResponseWithOrderIdAndOrderItemId DeleteOrderItems(ctx, orderId, itemId)
Delete Order Item

Use this method to remove an order item from an existing order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 
  **itemId** | **int64**| The item ID | 

### Return type

[**DefaultSuccessResponseWithOrderIdAndOrderItemId**](defaultSuccessResponseWithOrderIDAndOrderItemID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrderTotals**
> DefaultSuccessResponseWithOrderIdAndOrderTotalId DeleteOrderTotals(ctx, orderId, totalId)
Delete Order Total

Use this method to remove an order total from an existing order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 
  **totalId** | **int64**| The total ID | 

### Return type

[**DefaultSuccessResponseWithOrderIdAndOrderTotalId**](defaultSuccessResponseWithOrderIDAndOrderTotalID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrders**
> DefaultSuccessResponseWithOrderId DeleteOrders(ctx, orderId)
Delete Order

Remove an entire Order record from the database. This method will also remove the order-items along with their attributes and the order-total records. To see an example usage take a look at `docs/REST/samples/order-service/remove_order.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**DefaultSuccessResponseWithOrderId**](defaultSuccessResponseWithOrderID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderHistories**
> []GxOrderStatusHistory GetOrderHistories(ctx, orderId)
Get Order History (Multiple)

Returns multiple order status history records. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**[]GxOrderStatusHistory**](GXOrderStatusHistory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderHistory**
> GxOrderStatusHistory GetOrderHistory(ctx, orderId, historyId)
Get Order History (Single)

Returns multiple or a single order status history records. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 
  **historyId** | **int64**| The history ID | 

### Return type

[**GxOrderStatusHistory**](GXOrderStatusHistory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderItem**
> GxOrderItem GetOrderItem(ctx, orderId, itemId)
Get Order Item (Single)

Get a single order item from an existing order. All the GET manipulation parameters are applied with this method (search, sort, minimize, paginate etc).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 
  **itemId** | **int64**| The item ID | 

### Return type

[**GxOrderItem**](GXOrderItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderItems**
> []GxOrderItem GetOrderItems(ctx, orderId)
Get Order Items (Multiple)

Get multiple order items from an existing order. All the GET manipulation parameters are applied with this method (search, sort, minimize, paginate etc).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**[]GxOrderItem**](GXOrderItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderPaymentDetails**
> GxPaymentDetails GetOrderPaymentDetails(ctx, orderId)
Get Order Payment Details

Get the payment details for a specific order record through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**GxPaymentDetails**](GXPaymentDetails.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderTotal**
> GxOrderTotal GetOrderTotal(ctx, orderId, totalsId)
Get Order Total

Returns a single order total from an existing order. All the GET manipulation parameters are applied with this method (search, sort, minimize, paginate etc).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 
  **totalsId** | **int64**| The totals ID | 

### Return type

[**GxOrderTotal**](GXOrderTotal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderTotals**
> []GxOrderTotal GetOrderTotals(ctx, orderId)
Get Order Totals (Multiple)

Returns all single order totals from an existing order. All the GET manipulation parameters are applied with this method (search, sort, minimize, paginate etc).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**[]GxOrderTotal**](GXOrderTotal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderTrackingCodes**
> []GxTrackingCode GetOrderTrackingCodes(ctx, orderId)
Get Order Tracking Codes (Multiple)

Get an array of all tracking code records of an order through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**[]GxTrackingCode**](GXTrackingCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderTrackingCodesDeprecated**
> []GxTrackingCode GetOrderTrackingCodesDeprecated(ctx, orderId)
Get Order Tracking Codes (Multiple) (deprecated)

Get an array of all tracking code records of an order through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**[]GxTrackingCode**](GXTrackingCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrders**
> []GxOrder GetOrders(ctx, orderId)
Get Order (Single)

Get a single order record through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation. Important: Whenever you make requests that will return multiple orders the response will contain a smaller version of each order record called order-list-item. This is done for better performance because the creation of a complete order record takes significant time (many objects are involved). If you still need the complete data of an order record you will have to make an extra GET request with the ID provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**[]GxOrder**](GXOrder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrdersItemsAttributeProperty**
> GxOrderItemAttributes GetOrdersItemsAttributeProperty(ctx, orderId, itemId, attributeId)
Get Order Item Attribute/Property (Single)

Returns single order item attribute/property record through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 
  **itemId** | **int64**| The item ID | 
  **attributeId** | **int64**| The attribute ID | 

### Return type

[**GxOrderItemAttributes**](GXOrderItemAttributes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrdersItemsAttributesProperty**
> []GxOrderItemAttributes GetOrdersItemsAttributesProperty(ctx, orderId, itemId)
Get Order Item Attributes/Properties (Multiple)

Returns multiple order item attribute/property records through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 
  **itemId** | **int64**| The item ID | 

### Return type

[**[]GxOrderItemAttributes**](GXOrderItemAttributes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrdersMultiple**
> []GxOrderListItem GetOrdersMultiple(ctx, optional)
Get Orders (Multiple)

Get multiple order list item records through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation. Important: Whenever you make requests that will return multiple orders the response will contain a smaller version of each order record called order-list-item. This is done for better performance because the creation of a complete order record takes significant time (many objects are involved). If you still need the complete data of an order record you will have to make an extra GET request with the ID provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrdersApiGetOrdersMultipleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiGetOrdersMultipleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changed** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **changed/deleted** at the given date. | 
 **modified** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **modified** at a given date. | 
 **deleted** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **deleted** at the given date. | 
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxOrderListItem**](GXOrderListItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTrackingCode**
> []GxTrackingCode PostTrackingCode(ctx, body, orderId)
Create Order Tracking Code

Create a new tracking code for an order. See <a href=\"/gambio-gx3-api/resources/gxtrackingcode\">GXTrackingCode</a> resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxTrackingCodeRequestBody**](GxTrackingCodeRequestBody.md)| The tracking code to be created. | 
  **orderId** | **int64**| The order ID | 

### Return type

[**[]GxTrackingCode**](GXTrackingCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTrackingCodeDeprecated**
> []GxTrackingCode PostTrackingCodeDeprecated(ctx, body, orderId)
Create Order Tracking Code (deprecated)

Create a new tracking code for an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxTrackingCodeRequestBody**](GxTrackingCodeRequestBody.md)| The tracking code to be created | 
  **orderId** | **int64**| The order ID | 

### Return type

[**[]GxTrackingCode**](GXTrackingCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOrdersItemsAttribute**
> DefaultSuccessResponseWithOrderIdAndOrderItemIdAndOrderItemAttributeId RemoveOrdersItemsAttribute(ctx, orderId, itemId, attributeId)
Delete Order Item Attribute/Property

Removes a single order item attribute/property entry from an existing order item record.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 
  **itemId** | **int64**| The item ID | 
  **attributeId** | **int64**| The attribute ID | 

### Return type

[**DefaultSuccessResponseWithOrderIdAndOrderItemIdAndOrderItemAttributeId**](defaultSuccessResponseWithOrderIDAndOrderItemIDAndOrderItemAttributeID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrders**
> []GxOrderListItem SearchOrders(ctx, body)
Search Orders

Returns multiple orders that are respecting the given search condition. Further information about defining a search condition can be found in the <a href=\"/gambio-gx3-api/guides/search-example\">Search Example</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCondition**](SearchCondition.md)| The search condition | 

### Return type

[**[]GxOrderListItem**](GXOrderListItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrder**
> GxOrder UpdateOrder(ctx, orderId)
Update Order

Use this method to update an existing order record. It uses the complete order JSON resource so it might be useful to fetch it through a GET request, alter its values and PUT it back in order to perform the update operation. Take a look in the POST method for more detailed explanation on every resource property. To see an example usage take a look at `docs/REST/samples/order-service/update_order.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID | 

### Return type

[**GxOrder**](GXOrder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrderItems**
> GxOrderItem UpdateOrderItems(ctx, body, orderId, itemId)
Update Order Item

Use this method to update an existing order item. Use the same order item JSON format as in the POST method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderItem**](GxOrderItem.md)| The order item | 
  **orderId** | **int64**| The order ID | 
  **itemId** | **int64**| The item ID | 

### Return type

[**GxOrderItem**](GXOrderItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrderStatus**
> GxOrderStatus UpdateOrderStatus(ctx, body, orderId)
Update Order Status

Use this method if you want to update the status of an existing order and create an order history entry. The status history entry must also contain extra information as shown in the JSON example.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderStatusHistory**](GxOrderStatusHistory.md)| The order status | 
  **orderId** | **int64**| The orders ID | 

### Return type

[**GxOrderStatus**](GXOrderStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrderTotal**
> GxOrderTotal UpdateOrderTotal(ctx, body, orderId, totalId)
Update Order Total

Use this method to update an existing order total. Use the same order total JSON format as in the POST method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderTotal**](GxOrderTotal.md)| The order | 
  **orderId** | **int64**| The order ID | 
  **totalId** | **int64**| The total ID | 

### Return type

[**GxOrderTotal**](GXOrderTotal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrdersItemsAttributesProperty**
> GxOrderItemAttributes UpdateOrdersItemsAttributesProperty(ctx, body, orderId, itemId, attributeId)
Update Order Item Attribute/Property

Use this method to update an existing order item attribute record. It uses the same attribute JSON format as in the \"Create Order Item Attribute\" method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderItemAttributes**](GxOrderItemAttributes.md)| The order item attributes | 
  **orderId** | **int64**| The order ID | 
  **itemId** | **int64**| The item ID | 
  **attributeId** | **int64**| The attribute ID | 

### Return type

[**GxOrderItemAttributes**](GXOrderItemAttributes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


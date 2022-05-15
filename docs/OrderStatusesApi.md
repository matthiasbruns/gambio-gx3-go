# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrderStatuses**](OrderStatusesApi.md#CreateOrderStatuses) | **Post** /order_statuses | Create Order Status
[**DeleteOrderStatuses**](OrderStatusesApi.md#DeleteOrderStatuses) | **Delete** /order_statuses/{order_status_id} | Delete Order Status
[**GetOrderStatus**](OrderStatusesApi.md#GetOrderStatus) | **Get** /order_statuses/{order_status_id} | Get order statuses (Single)
[**GetOrderStatuses**](OrderStatusesApi.md#GetOrderStatuses) | **Get** /order_statuses | Get order statuses (Multiple)
[**UpdateOrderStatuses**](OrderStatusesApi.md#UpdateOrderStatuses) | **Put** /order_statuses/{order_status_id} | Update Order Status

# **CreateOrderStatuses**
> GxOrderStatus CreateOrderStatuses(ctx, body)
Create Order Status

This method creates a new order status in the database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderStatus**](GxOrderStatus.md)| The order status | 

### Return type

[**GxOrderStatus**](GXOrderStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrderStatuses**
> DefaultSuccessResponseWithOrderStatusId DeleteOrderStatuses(ctx, orderStatusId)
Delete Order Status

Removes a order status record from the system. This method will always return success.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderStatusId** | **int64**| The order status ID | 

### Return type

[**DefaultSuccessResponseWithOrderStatusId**](defaultSuccessResponseWithOrderStatusID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderStatus**
> GxOrderStatus GetOrderStatus(ctx, orderStatusId)
Get order statuses (Single)

Get a single order status entries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderStatusId** | **int64**| The order status ID | 

### Return type

[**GxOrderStatus**](GXOrderStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrderStatuses**
> []GxOrderStatus GetOrderStatuses(ctx, optional)
Get order statuses (Multiple)

Get a multiple order status entries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderStatusesApiGetOrderStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderStatusesApiGetOrderStatusesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxOrderStatus**](GXOrderStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrderStatuses**
> GxOrderStatus UpdateOrderStatuses(ctx, body, orderStatusId)
Update Order Status

This method updates a order status in the database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxOrderStatus**](GxOrderStatus.md)| The order status | 
  **orderStatusId** | **int64**| The order status ID | 

### Return type

[**GxOrderStatus**](GXOrderStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


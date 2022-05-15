# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrdersHistory**](OrdersHistoryApi.md#GetOrdersHistory) | **Get** /orders_history/{order_id} | Get Order History (Single)

# **GetOrdersHistory**
> []GxOrderStatusHistory GetOrdersHistory(ctx, orderId)
Get Order History (Single)

Get a single order history record through a GET request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**| The order ID to GET history for. | 

### Return type

[**[]GxOrderStatusHistory**](GXOrderStatusHistory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


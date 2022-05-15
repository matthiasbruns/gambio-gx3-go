# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentDetails**](PaymentDetailsApi.md#GetPaymentDetails) | **Get** /payment_details/{order_id} | Get Payment Details

# **GetPaymentDetails**
> GxPaymentDetails GetPaymentDetails(ctx, orderId)
Get Payment Details

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


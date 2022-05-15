# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrackingCodeDeprecated**](TrackingCodeApi.md#DeleteTrackingCodeDeprecated) | **Delete** /tracking_code/{tracking_code_id} | Delete Tracking Code (deprecated)
[**GetTrackingCodeDeprecated**](TrackingCodeApi.md#GetTrackingCodeDeprecated) | **Get** /tracking_code/{tracking_code_id} | Get Tracking Code (Single) (deprecated)
[**GetTrackingCodesDeprecated**](TrackingCodeApi.md#GetTrackingCodesDeprecated) | **Get** /tracking_code | Get Tracking Codes (Multiple) (deprecated)

# **DeleteTrackingCodeDeprecated**
> []TrackingCodeSuccessResponse DeleteTrackingCodeDeprecated(ctx, trackingCodeId)
Delete Tracking Code (deprecated)

Removes a tracking code record from the database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackingCodeId** | **int64**| The tracking code ID | 

### Return type

[**[]TrackingCodeSuccessResponse**](trackingCodeSuccessResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrackingCodeDeprecated**
> []GxTrackingCode GetTrackingCodeDeprecated(ctx, trackingCodeId)
Get Tracking Code (Single) (deprecated)

Get a single tracking code through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackingCodeId** | **int64**| The tracking code ID | 

### Return type

[**[]GxTrackingCode**](GXTrackingCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrackingCodesDeprecated**
> []GxTrackingCode GetTrackingCodesDeprecated(ctx, optional)
Get Tracking Codes (Multiple) (deprecated)

Get an array of all tracking code records through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TrackingCodeApiGetTrackingCodesDeprecatedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrackingCodeApiGetTrackingCodesDeprecatedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxTrackingCode**](GXTrackingCode.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


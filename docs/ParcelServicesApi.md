# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateParcelService**](ParcelServicesApi.md#CreateParcelService) | **Post** /parcel_services | Create Parcel Service
[**DeleteParcelService**](ParcelServicesApi.md#DeleteParcelService) | **Delete** /parcel_services/{parcel_service_id} | Delete Parcel Service
[**GetMultipleParcelServices**](ParcelServicesApi.md#GetMultipleParcelServices) | **Get** /parcel_services | Get Parcel Services (Multiple)
[**GetSingleParcelServices**](ParcelServicesApi.md#GetSingleParcelServices) | **Get** /parcel_services/{parcel_service_id} | Get Parcel Service (Single)
[**SearchParcelService**](ParcelServicesApi.md#SearchParcelService) | **Post** /parcel_services/search | Search Parcel Services
[**UpdateParcelService**](ParcelServicesApi.md#UpdateParcelService) | **Put** /parcel_services/{parcel_service_id} | Updates Parcel Service

# **CreateParcelService**
> GxParcelService CreateParcelService(ctx, body)
Create Parcel Service

This method enables the creation of a new parcel service into the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxParcelService**](GxParcelService.md)| The created parcel service. | 

### Return type

[**GxParcelService**](GXParcelService.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteParcelService**
> DefaultSuccessResponseWithParcelServiceId DeleteParcelService(ctx, parcelServiceId)
Delete Parcel Service

Removes a parcel service entity from the system. This method will always return success even if the parcel service does not exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parcelServiceId** | **int64**| The parcel service ID | 

### Return type

[**DefaultSuccessResponseWithParcelServiceId**](defaultSuccessResponseWithParcelServiceID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMultipleParcelServices**
> []GxParcelService GetMultipleParcelServices(ctx, optional)
Get Parcel Services (Multiple)

Get multiple parcel services records through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ParcelServicesApiGetMultipleParcelServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParcelServicesApiGetMultipleParcelServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxParcelService**](GXParcelService.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSingleParcelServices**
> GxParcelService GetSingleParcelServices(ctx, parcelServiceId)
Get Parcel Service (Single)

Get a single parcel service record through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parcelServiceId** | **int64**| The parcel service ID | 

### Return type

[**GxParcelService**](GXParcelService.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchParcelService**
> []GxParcelService SearchParcelService(ctx, body)
Search Parcel Services

Returns multiple parcel services that are respecting the given search condition. Further information about defining a search condition can be found in the <a href=\"/gambio-gx3-api/guides/search-example\">Search Example</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCondition**](SearchCondition.md)| The search condition | 

### Return type

[**[]GxParcelService**](GXParcelService.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateParcelService**
> GxParcelService UpdateParcelService(ctx, body, parcelServiceId)
Updates Parcel Service

Use this method if you want to update an existing parcel service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxParcelService**](GxParcelService.md)| The updated parcel service. | 
  **parcelServiceId** | **int64**| The parcel service ID | 

### Return type

[**GxParcelService**](GXParcelService.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


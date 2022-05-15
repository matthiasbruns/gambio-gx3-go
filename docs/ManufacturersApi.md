# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManufacturers**](ManufacturersApi.md#CreateManufacturers) | **Post** /manufacturers | Create Manufacturers
[**DeleteManufacturers**](ManufacturersApi.md#DeleteManufacturers) | **Delete** /manufacturers/{manufacturer_id} | Delete Manufacturers
[**GetManufacturer**](ManufacturersApi.md#GetManufacturer) | **Get** /manufacturers/{manufacturer_id} | Get Manufacturers
[**GetMultipleManufacturers**](ManufacturersApi.md#GetMultipleManufacturers) | **Get** /manufacturers | Get Manufacturers (Multiple)
[**SearchManufacturers**](ManufacturersApi.md#SearchManufacturers) | **Post** /manufacturers/search | Search Manufacturers
[**UpdateManufacturers**](ManufacturersApi.md#UpdateManufacturers) | **Patch** /manufacturers/{manufacturer_id} | Updates Manufacturers

# **CreateManufacturers**
> GxManufacturer CreateManufacturers(ctx, body)
Create Manufacturers

This method enables the creation of a new Manufacturers into the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxManufacturer**](GxManufacturer.md)| The manufacturer | 

### Return type

[**GxManufacturer**](GXManufacturer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteManufacturers**
> DefaultSuccessResponseWithManufacturersId DeleteManufacturers(ctx, manufacturerId)
Delete Manufacturers

Removes a manufacturers record from the system. This method will always return success even if the manufacturers does not exist (due to internal ManufacturersWriteService architecture decisions, which strive to avoid unnecessary failures).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **manufacturerId** | **int64**| The manufacturers ID | 

### Return type

[**DefaultSuccessResponseWithManufacturersId**](defaultSuccessResponseWithManufacturersID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetManufacturer**
> GxManufacturer GetManufacturer(ctx, manufacturerId)
Get Manufacturers

Get a single manufacturers record through the GET method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **manufacturerId** | **int64**| The manufacturers ID | 

### Return type

[**GxManufacturer**](GXManufacturer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMultipleManufacturers**
> []GxManufacturer GetMultipleManufacturers(ctx, optional)
Get Manufacturers (Multiple)

Get multiple manufacturers record through the GET method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManufacturersApiGetMultipleManufacturersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManufacturersApiGetMultipleManufacturersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxManufacturer**](GXManufacturer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchManufacturers**
> []GxCategory SearchManufacturers(ctx, body, optional)
Search Manufacturers

Returns multiple manufacturers that are respecting the given search condition. Further information about defining a search condition can be found in the <a href=\"/gambio-gx3-api/guides/search-example\">Search Example</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCondition**](SearchCondition.md)| The search condition | 
 **optional** | ***ManufacturersApiSearchManufacturersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManufacturersApiSearchManufacturersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxCategory**](GXCategory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateManufacturers**
> GxManufacturer UpdateManufacturers(ctx, body, manufacturerId)
Updates Manufacturers

Use this method if you want to update an existing manufacturers record.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxManufacturer**](GxManufacturer.md)| The manufacturer | 
  **manufacturerId** | **int64**| The manufacturer ID | 

### Return type

[**GxManufacturer**](GXManufacturer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


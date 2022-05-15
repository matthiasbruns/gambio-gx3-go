# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuantityUnits**](QuantityUnitsApi.md#CreateQuantityUnits) | **Post** /quantity_units | Create Quantity Units
[**DeleteQuantityUnits**](QuantityUnitsApi.md#DeleteQuantityUnits) | **Delete** /quantity_units/{quantity_unit_id} | Delete Quantity Units
[**GetQuantityUnit**](QuantityUnitsApi.md#GetQuantityUnit) | **Get** /quantity_units/{quantity_unit_id} | Get Quantity Units (Single)
[**GetQuantityUnits**](QuantityUnitsApi.md#GetQuantityUnits) | **Get** /quantity_units | Get Quantity Units (Multiple)
[**UpdateQuantityUnitEntity**](QuantityUnitsApi.md#UpdateQuantityUnitEntity) | **Patch** /quantity_units/{quantity_unit_id} | Updates Quantity Units

# **CreateQuantityUnits**
> GxQuantityUnit CreateQuantityUnits(ctx, body)
Create Quantity Units

This method enables the creation of a new QuantityUnits into the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxQuantityUnit**](GxQuantityUnit.md)| The Quantity Unit. | 

### Return type

[**GxQuantityUnit**](GXQuantityUnit.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuantityUnits**
> DefaultSuccessResponseWithQuantityUnitId DeleteQuantityUnits(ctx, quantityUnitId)
Delete Quantity Units

Removes a quantity units record from the system. This method will always return success even if the quantity unit does not exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quantityUnitId** | **int64**| The quantity unit ID | 

### Return type

[**DefaultSuccessResponseWithQuantityUnitId**](defaultSuccessResponseWithQuantityUnitID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuantityUnit**
> GxQuantityUnit GetQuantityUnit(ctx, quantityUnitId)
Get Quantity Units (Single)

Get a single quantity unit record through the GET method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quantityUnitId** | **int64**| The quantity unit ID | 

### Return type

[**GxQuantityUnit**](GXQuantityUnit.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuantityUnits**
> []GxQuantityUnit GetQuantityUnits(ctx, optional)
Get Quantity Units (Multiple)

Get multiple quantity unit records through the GET method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuantityUnitsApiGetQuantityUnitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuantityUnitsApiGetQuantityUnitsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxQuantityUnit**](GXQuantityUnit.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuantityUnitEntity**
> GxQuantityUnit UpdateQuantityUnitEntity(ctx, body, quantityUnitId)
Updates Quantity Units

Use this method if you want to update an existing quantity units record.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxQuantityUnit**](GxQuantityUnit.md)| The quantity unit | 
  **quantityUnitId** | **int64**| The quantity unit ID | 

### Return type

[**GxQuantityUnit**](GXQuantityUnit.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


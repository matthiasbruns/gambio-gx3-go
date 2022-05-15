# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddress**](AddressesApi.md#CreateAddress) | **Post** /addresses | Create Address
[**DeleteAddress**](AddressesApi.md#DeleteAddress) | **Delete** /addresses/{address_id} | Delete Address
[**GetAddress**](AddressesApi.md#GetAddress) | **Get** /addresses/{address_id} | Get Address (Single)
[**GetAddresses**](AddressesApi.md#GetAddresses) | **Get** /addresses | Get Addresses (Multiple)
[**UpdateAddress**](AddressesApi.md#UpdateAddress) | **Put** /addresses/{address_id} | Update Address

# **CreateAddress**
> GxCustomer CreateAddress(ctx, body)
Create Address

Create an address record for a customer in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxAddress**](GxAddress.md)| The address to be created | 

### Return type

[**GxCustomer**](GXCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAddress**
> AddressDeleteResponse DeleteAddress(ctx, addressId)
Delete Address

Remove an address record from the system. This method will always return success even if the address record does not exist (due to internal architecture decisions, which strive to avoid unnecessary failures).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressId** | **string**| The ID of the address to be removed | 

### Return type

[**AddressDeleteResponse**](addressDeleteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddress**
> GxAddress GetAddress(ctx, addressId)
Get Address (Single)

Get a single address record through a GET requets. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressId** | **string**| The ID of the address to fetch | 

### Return type

[**GxAddress**](GXAddress.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddresses**
> []GxAddress GetAddresses(ctx, optional)
Get Addresses (Multiple)

Get multiple address records through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddressesApiGetAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressesApiGetAddressesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxAddress**](GXAddress.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAddress**
> GxAddress UpdateAddress(ctx, addressId)
Update Address

Update an existing address record by providing new data. You do not have to provide the full presentation of the address in the JSON string of the request, rather just the fields to be updated. The address ID will be taken from the URI of the request so it is not required that it is included withing the request JSON.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressId** | **string**| The ID of the address to be updated | 

### Return type

[**GxAddress**](GXAddress.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


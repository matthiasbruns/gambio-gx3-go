# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](CountriesApi.md#GetCountries) | **Get** /countries/{country_id} | Get Countries (Single)
[**GetCountriesZones**](CountriesApi.md#GetCountriesZones) | **Get** /countries/{country_id}/zones | Get Countries Zones

# **GetCountries**
> GxCountry GetCountries(ctx, countryId)
Get Countries (Single)

Get a single country. This method is currently limited to only fetching a single country resource so make sure that you provide the country ID in the request URI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryId** | **int64**| The country ID | 

### Return type

[**GxCountry**](GXCountry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountriesZones**
> []GxCountryZone GetCountriesZones(ctx, countryId)
Get Countries Zones

Get a single countrys zones. This method is currently limited to only fetching a single countrys zone resource so make sure that you provide the country ID in the request URI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryId** | **int64**| The country ID | 

### Return type

[**[]GxCountryZone**](GXCountryZone.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


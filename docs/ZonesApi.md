# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZones**](ZonesApi.md#GetZones) | **Get** /zones/{zone_id} | Get Zones

# **GetZones**
> GxZone GetZones(ctx, zoneId)
Get Zones

Get a single registered zone resource. This method is currently limited to only fetching a single zone and might be updated in a future version of the API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **int64**| The zone ID | 

### Return type

[**GxZone**](GXZone.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


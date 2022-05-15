# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpecialOffers**](SpecialOffersApi.md#CreateSpecialOffers) | **Post** /special_offers | Create Special Offers
[**DeleteSpecialOffer**](SpecialOffersApi.md#DeleteSpecialOffer) | **Delete** /special_offers/{special_offer_id} | Removes Special Offers
[**GetMultipleSpecialOffers**](SpecialOffersApi.md#GetMultipleSpecialOffers) | **Get** /special_offers | Get Special Offers (Multiple)
[**GetSpecialOffers**](SpecialOffersApi.md#GetSpecialOffers) | **Get** /special_offers/{special_offer_id} | Get Special Offers
[**SearchSpecialOffer**](SpecialOffersApi.md#SearchSpecialOffer) | **Post** /special_offers/search | Search Special Offer
[**UpdateSpecialOffer**](SpecialOffersApi.md#UpdateSpecialOffer) | **Put** /special_offers/{special_offer_id} | Updates Special Offers

# **CreateSpecialOffers**
> GxSpecialOffer CreateSpecialOffers(ctx, body)
Create Special Offers

This method enables the creation of a new special offer into the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxSpecialOffer**](GxSpecialOffer.md)| The created special offer. | 

### Return type

[**GxSpecialOffer**](GXSpecialOffer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpecialOffer**
> DefaultSuccessResponseWithSpecialOfferId DeleteSpecialOffer(ctx, specialOfferId)
Removes Special Offers

Removes a special offer record from the system. This method will always return success even if the vpe does not exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **specialOfferId** | **int64**| The special offer ID | 

### Return type

[**DefaultSuccessResponseWithSpecialOfferId**](defaultSuccessResponseWithSpecialOfferID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMultipleSpecialOffers**
> []GxSpecialOffer GetMultipleSpecialOffers(ctx, optional)
Get Special Offers (Multiple)

Get multiple special offers record through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpecialOffersApiGetMultipleSpecialOffersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpecialOffersApiGetMultipleSpecialOffersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxSpecialOffer**](GXSpecialOffer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpecialOffers**
> GxSpecialOffer GetSpecialOffers(ctx, specialOfferId)
Get Special Offers

Gets a single special offer record through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **specialOfferId** | **int64**| The special offer ID | 

### Return type

[**GxSpecialOffer**](GXSpecialOffer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSpecialOffer**
> []GxSpecialOffer SearchSpecialOffer(ctx, body)
Search Special Offer

Returns multiple special offers that are respecting the given search condition. Further information about defining a search condition can be found in the <a href=\"/gambio-gx3-api/guides/search-example\">Search Example</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCondition**](SearchCondition.md)| The search condition | 

### Return type

[**[]GxSpecialOffer**](GXSpecialOffer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSpecialOffer**
> GxSpecialOffer UpdateSpecialOffer(ctx, body, specialOfferId)
Updates Special Offers

Use this method if you want to update an existing special offer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxSpecialOffer**](GxSpecialOffer.md)| The updated special offer. | 
  **specialOfferId** | **int64**| The special offer ID | 

### Return type

[**GxSpecialOffer**](GXSpecialOffer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


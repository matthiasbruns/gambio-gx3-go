# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProductPrices**](ProductPricesApi.md#DeleteProductPrices) | **Delete** /product_prices/{product_id} | Delete Product Prices
[**GetProductPrices**](ProductPricesApi.md#GetProductPrices) | **Get** /product_prices/{product_id} | Get Product Prices
[**GetProductQuantityPrices**](ProductPricesApi.md#GetProductQuantityPrices) | **Get** /product_prices/{product_id}/customer_groups/{customer_group_id}/{quantity} | Get Graduated Product Price
[**PutProductPrices**](ProductPricesApi.md#PutProductPrices) | **Put** /product_prices/{product_id} | Update Product Prices

# **DeleteProductPrices**
> DefaultSuccessResponseWithProductId DeleteProductPrices(ctx, productId)
Delete Product Prices

Deletes all customer group- and graduated prices of a product. This method will always return success even if the product does not exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int64**| The product id | 

### Return type

[**DefaultSuccessResponseWithProductId**](defaultSuccessResponseWithProductID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductPrices**
> []GxProductPrices GetProductPrices(ctx, productId)
Get Product Prices

Get prices for customer groups and their graduated prices, if exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int64**| The product id | 

### Return type

[**[]GxProductPrices**](GXProductPrices.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductQuantityPrices**
> []GxProductQuantityPrice GetProductQuantityPrices(ctx, productId, customerGroupId, quantity)
Get Graduated Product Price

Get price information for given customer group and product quantity. Graduated- and customer group prices are included in the response

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int64**| The product id | 
  **customerGroupId** | **int64**| The customer group id | 
  **quantity** | **float32**| Products quantity | 

### Return type

[**[]GxProductQuantityPrice**](GXProductQuantityPrice.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutProductPrices**
> GxProductPrices PutProductPrices(ctx, body, productId)
Update Product Prices

Updates product price data for the product of the given id. All price information gets replaced with the new one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]GxProductPrices**](GXProductPrices.md)| Product price data | 
  **productId** | **int64**| The product id | 

### Return type

[**GxProductPrices**](GXProductPrices.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


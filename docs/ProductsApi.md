# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductLink**](ProductsApi.md#CreateProductLink) | **Post** /products/{product_id}/links | Create Product Link
[**CreateProducts**](ProductsApi.md#CreateProducts) | **Post** /products | Create Products
[**DeleteProductLink**](ProductsApi.md#DeleteProductLink) | **Delete** /products/{product_id}/links | Delete Product Link
[**DeleteProducts**](ProductsApi.md#DeleteProducts) | **Delete** /products/{product_ids} | Delete Products (Multiple)
[**GetProduct**](ProductsApi.md#GetProduct) | **Get** /products/{product_id} | Get Products
[**GetProductLinks**](ProductsApi.md#GetProductLinks) | **Get** /products/{product_id}/links | Get Product Links
[**GetProductProductPrices**](ProductsApi.md#GetProductProductPrices) | **Get** /products/{product_id}/product_prices | Get Product Prices
[**GetProductQuantityPrices**](ProductsApi.md#GetProductQuantityPrices) | **Get** /products/{product_id}/product_prices/customer_groups/{customer_group_id}/{quantity} | Get Graduated Product Price
[**GetProducts**](ProductsApi.md#GetProducts) | **Get** /products | Get Products
[**PutProductProductPrices**](ProductsApi.md#PutProductProductPrices) | **Put** /products/{product_id}/product_prices | Update Product Prices
[**SearchProducts**](ProductsApi.md#SearchProducts) | **Post** /products/search | Search Products
[**UpdateProduct**](ProductsApi.md#UpdateProduct) | **Put** /products/{product_id} | Update Product
[**UpdateProductLink**](ProductsApi.md#UpdateProductLink) | **Put** /products/{product_id}/links | Update Product Link
[**UpdateProducts**](ProductsApi.md#UpdateProducts) | **Put** /products | Update Products (Multiple)

# **CreateProductLink**
> DefaultSuccessResponseWithProductIdAndCategoryId CreateProductLink(ctx, body, productId)
Create Product Link

Creates a new product to category record in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductCategoryLinkId**](ProductCategoryLinkId.md)| The category ID | 
  **productId** | **int64**| The product ID | 

### Return type

[**DefaultSuccessResponseWithProductIdAndCategoryId**](defaultSuccessResponseWithProductIDAndCategoryID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProducts**
> ProductArrayPostResponse CreateProducts(ctx, body)
Create Products

Creates new product records in the system. May be provided an array of product records, for bulk insertion. For single insertion, provide a single GXProduct object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]GxProductArray**](GXProductArray.md)| An array of products | 

### Return type

[**ProductArrayPostResponse**](productArrayPostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductLink**
> DefaultSuccessResponseWithProductId DeleteProductLink(ctx, body, productId)
Delete Product Link

Deletes an existing product to category link. If there is no categoryId property set, all the product links will be removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProductCategoryLinkId**](ProductCategoryLinkId.md)| The category ID | 
  **productId** | **int64**| The product or category link ID | 

### Return type

[**DefaultSuccessResponseWithProductId**](defaultSuccessResponseWithProductID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProducts**
> ProductsDeleteSuccessResponse DeleteProducts(ctx, productIds)
Delete Products (Multiple)

Removes multiple product records from the database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productIds** | [**[]int64**](int64.md)| A comma-seperated list of product IDs to delete | 

### Return type

[**ProductsDeleteSuccessResponse**](ProductsDeleteSuccessResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProduct**
> GxProduct GetProduct(ctx, productId)
Get Products

Get a single product record through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation. To see an example usage take a look at `docs/REST/samples/product-service/remove_product.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int64**| The product ID | 

### Return type

[**GxProduct**](GXProduct.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductLinks**
> []int64 GetProductLinks(ctx, productId)
Get Product Links

Get all product or category links associated with a specific product.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int64**| The product or category ID | 

### Return type

**[]int64**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductProductPrices**
> GxProductPrices GetProductProductPrices(ctx, productId)
Get Product Prices

Get prices for customer groups and their graduated prices, if exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **int64**| The product id | 

### Return type

[**GxProductPrices**](GXProductPrices.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductQuantityPrices**
> GxProductQuantityPrice GetProductQuantityPrices(ctx, productId, customerGroupId, quantity)
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

[**GxProductQuantityPrice**](GXProductQuantityPrice.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProducts**
> []GxProductListing GetProducts(ctx, optional)
Get Products

Get a multiple product records through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation. To see an example usage take a look at `docs/REST/samples/product-service/remove_product.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiGetProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiGetProductsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changed** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **changed/deleted** at the given date. | 
 **modified** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **modified** at a given date. | 
 **deleted** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **deleted** at the given date. | 
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxProductListing**](GXProductListing.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutProductProductPrices**
> GxProductPrices PutProductProductPrices(ctx, body, productId)
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

# **SearchProducts**
> []GxProduct SearchProducts(ctx, body)
Search Products

Returns multiple products that are respecting the given search condition. Further information about defining a search condition can be found in the <a href=\"/gambio-gx3-api/guides/search-example\">Search Example</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCondition**](SearchCondition.md)| The search condition | 

### Return type

[**[]GxProduct**](GXProduct.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProduct**
> GxProduct UpdateProduct(ctx, body, productId)
Update Product

Use this method to update an existing product record. Take a look in the POST method for more detailed explanation on every resource property. To see an example usage consider `docs/REST/samples/product-service/update_product.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxProduct**](GxProduct.md)| The product | 
  **productId** | **int64**| The product ID | 

### Return type

[**GxProduct**](GXProduct.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProductLink**
> DefaultSuccessResponseWithProductIdAndCategoryIDs UpdateProductLink(ctx, body, productId)
Update Product Link

Changes an existing product to category link.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxProductLinkPutRequest**](GxProductLinkPutRequest.md)| The old and new category ID | 
  **productId** | **int64**| The product link ID | 

### Return type

[**DefaultSuccessResponseWithProductIdAndCategoryIDs**](defaultSuccessResponseWithProductIDAndCategoryIDs.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProducts**
> ProductArrayPutResponse UpdateProducts(ctx, body)
Update Products (Multiple)

Use this method to update multiple existing product records. It takes an array of GXProduct Objects, identified by their <b>id</b> property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]GxProductArray**](GXProductArray.md)| An array of products to update. | 

### Return type

[**ProductArrayPutResponse**](productArrayPutResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


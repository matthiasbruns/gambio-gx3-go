# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategories**](CategoriesApi.md#CreateCategories) | **Post** /categories | Create Categories
[**DeleteCategories**](CategoriesApi.md#DeleteCategories) | **Delete** /categories/{category_ids} | Delete Categories
[**GetCategories**](CategoriesApi.md#GetCategories) | **Get** /categories | Get Categories (Multiple)
[**GetCategoriyProducts**](CategoriesApi.md#GetCategoriyProducts) | **Get** /categories/{category_id}/products | Get Category Products
[**GetCategory**](CategoriesApi.md#GetCategory) | **Get** /categories/{category_id} | Get Categories (Single)
[**GetCategoryChildren**](CategoriesApi.md#GetCategoryChildren) | **Get** /categories/{category_id}/children | Get Category Children
[**SearchCategories**](CategoriesApi.md#SearchCategories) | **Post** /categories/search | Search Categories
[**UpdateCategories**](CategoriesApi.md#UpdateCategories) | **Put** /categories | Update Category (Multiple)
[**UpdateCategory**](CategoriesApi.md#UpdateCategory) | **Put** /categories/{category_id} | Update Category

# **CreateCategories**
> CategoryArrayPostResponse CreateCategories(ctx, body)
Create Categories

Creates new categories in the system by taking an array of categories. It is also possible to provide a single <a href=\"/gambio-gx3-api/resources/gxcategory\">GXCategory</a> ressource, to create a single entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]GxCategoryArray**](GXCategoryArray.md)| An array of categories | 

### Return type

[**CategoryArrayPostResponse**](categoryArrayPostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCategories**
> CategoryDeleteSuccessResponse DeleteCategories(ctx, categoryIds)
Delete Categories

Removes multiple category records from the database, provided a comma-seperated list of IDs to delete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryIds** | [**[]int64**](int64.md)| A comma-seperated list of category IDs to delete | 

### Return type

[**CategoryDeleteSuccessResponse**](CategoryDeleteSuccessResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategories**
> []GxCategoryListingEntity GetCategories(ctx, optional)
Get Categories (Multiple)

Get multiple category records through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation. To see an example usage take a look at `docs/REST/samples/category-service/fetch_category.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CategoriesApiGetCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesApiGetCategoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recursive** | **optional.Bool**| *Optional parameter*. If provided, recursively includes the ressources child-resources. | 
 **changed** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **changed/deleted** at the given date. | 
 **modified** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **modified** at a given date. | 
 **deleted** | **optional.String**| *Optional parameter*. If provided, receive a change history of items that were **deleted** at the given date. | 
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxCategoryListingEntity**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoriyProducts**
> []GxProduct GetCategoriyProducts(ctx, categoryId, optional)
Get Category Products

Get multiple product records of an specified category through a GET request. Using the `recursive` query parameter allows you to receive the products from the sub categories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryId** | **int64**| The category ID | 
 **optional** | ***CategoriesApiGetCategoriyProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesApiGetCategoriyProductsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **optional.Bool**| *Optional parameter*. If provided, recursively includes the ressources child-resources. | 

### Return type

[**[]GxProduct**](GXProduct.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategory**
> GxCategory GetCategory(ctx, categoryId)
Get Categories (Single)

Get a single category record through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation. To see an example usage take a look at `docs/REST/samples/category-service/fetch_category.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryId** | **int64**| The category ID | 

### Return type

[**GxCategory**](GXCategory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoryChildren**
> []GxCategoryListingEntity GetCategoryChildren(ctx, categoryId)
Get Category Children

Get multiple or a single sub category records of a given category through a GET request. This method supports all the GET parameters that are mentioned in the \"Introduction\" section of this documentation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryId** | **int64**| The parent category ID | 

### Return type

[**[]GxCategoryListingEntity**](array.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCategories**
> []GxCategory SearchCategories(ctx, body)
Search Categories

Returns multiple categories that are respecting the given search condition. Further information about defining a search condition can be found in the <a href=\"/gambio-gx3-api/guides/search-example\">Search Example</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCondition**](SearchCondition.md)| The search condition | 

### Return type

[**[]GxCategory**](GXCategory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCategories**
> CategoryArrayPutResponse UpdateCategories(ctx, body)
Update Category (Multiple)

Use this method to update multiple existing category records. It takes an array of GXCategory Objects, updating them based on their <b>id</b> property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]GxCategoryArray**](GXCategoryArray.md)| An array of categories to update. | 

### Return type

[**CategoryArrayPutResponse**](categoryArrayPutResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCategory**
> GxCategory UpdateCategory(ctx, body, categoryId)
Update Category

Use this method to update an existing category record. Take a look in the POST method for more detailed explanation on every resource property. To see an example usage take a look at `docs/REST/samples/category-service/update_category.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxCategory**](GxCategory.md)| An array of categories to update. | 
  **categoryId** | **int64**| The category ID | 

### Return type

[**GxCategory**](GXCategory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


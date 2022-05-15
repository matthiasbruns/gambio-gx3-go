# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProductImage**](ProductImagesApi.md#DeleteProductImage) | **Delete** /product_images | Delete Product Image
[**GetProductImage**](ProductImagesApi.md#GetProductImage) | **Get** /product_images | Get Product Images
[**RenameImageFile**](ProductImagesApi.md#RenameImageFile) | **Put** /product_images | Rename Image File
[**UploadProductImage**](ProductImagesApi.md#UploadProductImage) | **Post** /product_images | Upload Product Image

# **DeleteProductImage**
> DefaultSuccessResponseWithFilename DeleteProductImage(ctx, )
Delete Product Image

Remove the product image file from the server. This method will always provide a successful response even if the image file was not found.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DefaultSuccessResponseWithFilename**](defaultSuccessResponseWithFilename.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductImage**
> []string GetProductImage(ctx, optional)
Get Product Images

Get a list of all product image files which exists in the server's filesystem through a GET request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductImagesApiGetProductImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductImagesApiGetProductImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameImageFile**
> DefaultUpdateResponseWithFilenames RenameImageFile(ctx, optional)
Rename Image File

Use this method to rename an existing image file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductImagesApiRenameImageFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductImagesApiRenameImageFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GxFilenameChange**](GxFilenameChange.md)| The filenames | 

### Return type

[**DefaultUpdateResponseWithFilenames**](defaultUpdateResponseWithFilenames.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadProductImage**
> DefaultSuccessResponseWithFilename UploadProductImage(ctx, uploadProductImage, filename)
Upload Product Image

Uploads an image file for the products. Make this request without the \"Content-Type: application/json\". Except from the file the POST request must also contain a \"filename\" value with the final filename.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadProductImage** | ***os.File*****os.File**|  | 
  **filename** | **string**|  | 

### Return type

[**DefaultSuccessResponseWithFilename**](defaultSuccessResponseWithFilename.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


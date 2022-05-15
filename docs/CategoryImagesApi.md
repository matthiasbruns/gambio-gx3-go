# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCategoryImages**](CategoryImagesApi.md#DeleteCategoryImages) | **Delete** /category_images | Delete Category Image
[**GetCategoryImages**](CategoryImagesApi.md#GetCategoryImages) | **Get** /category_images | Get Category Images
[**RenameCategoryImages**](CategoryImagesApi.md#RenameCategoryImages) | **Put** /category_images | Rename Image File
[**UploadCategoryImages**](CategoryImagesApi.md#UploadCategoryImages) | **Post** /category_images | Upload Category Image

# **DeleteCategoryImages**
> DefaultSuccessResponseWithFilename DeleteCategoryImages(ctx, )
Delete Category Image

Removes the category image file from the server. This method will always provide a successful response even if the image file was not found.

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

# **GetCategoryImages**
> CategoryImageGetResponse GetCategoryImages(ctx, optional)
Get Category Images

Get a list of all category image files which exists in the server's filesystem through a GET request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CategoryImagesApiGetCategoryImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoryImagesApiGetCategoryImagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**CategoryImageGetResponse**](categoryImageGetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameCategoryImages**
> DefaultUpdateResponseWithFilenames RenameCategoryImages(ctx, )
Rename Image File

Use this method to rename an existing image file.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DefaultUpdateResponseWithFilenames**](defaultUpdateResponseWithFilenames.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadCategoryImages**
> DefaultSuccessResponseWithFilename UploadCategoryImages(ctx, uploadCategoryImage, filename)
Upload Category Image

Uploads an icon image for the categories. Make sure this request with Content-Type: multipart/form-data header instead of application/json. Except from the file the POST request must also contain a \"filename\" value with the final filename.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadCategoryImage** | ***os.File*****os.File**|  | 
  **filename** | **string**|  | 

### Return type

[**DefaultSuccessResponseWithFilename**](defaultSuccessResponseWithFilename.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


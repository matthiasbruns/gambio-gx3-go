# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCategoryIcon**](CategoryIconsApi.md#DeleteCategoryIcon) | **Delete** /category_icons | Delete Category Icon
[**GetCategoryIcons**](CategoryIconsApi.md#GetCategoryIcons) | **Get** /category_icons | Get Category Icons
[**PutCategoryIcons**](CategoryIconsApi.md#PutCategoryIcons) | **Put** /category_icons | Rename Icon File
[**UploadCategoryIcons**](CategoryIconsApi.md#UploadCategoryIcons) | **Post** /category_icons | Upload Category Icon

# **DeleteCategoryIcon**
> DefaultSuccessResponseWithFilename DeleteCategoryIcon(ctx, )
Delete Category Icon

Removes the category icon file from the server. This method will always provide a successful response even if the image file was not found.

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

# **GetCategoryIcons**
> CategoryIconGetResponse GetCategoryIcons(ctx, optional)
Get Category Icons

Returns a list of all category icon files which exists in the servers filesystem through a GET request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CategoryIconsApiGetCategoryIconsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoryIconsApiGetCategoryIconsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**CategoryIconGetResponse**](categoryIconGetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCategoryIcons**
> DefaultUpdateResponseWithFilenames PutCategoryIcons(ctx, )
Rename Icon File

Use this method to rename an existing icon file.

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

# **UploadCategoryIcons**
> DefaultSuccessResponseWithFilename UploadCategoryIcons(ctx, optional)
Upload Category Icon

Upload an icon image for the categories. Make sure this request with Content-Type: multipart/form-data header instead of application/json. Except from the file the POST request must also contain a \"filename\" value with the final filename.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CategoryIconsApiUploadCategoryIconsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoryIconsApiUploadCategoryIconsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadedFile** | **optional.Interface of *os.File****optional.**|  | 
 **filenae** | **optional.**|  | 

### Return type

[**DefaultSuccessResponseWithFilename**](defaultSuccessResponseWithFilename.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


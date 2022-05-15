# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVPE**](VpeApi.md#CreateVPE) | **Post** /vpe | Create VPE
[**DeleteVPE**](VpeApi.md#DeleteVPE) | **Delete** /vpe/{vpe_id} | Delete VPE
[**GetVPE**](VpeApi.md#GetVPE) | **Get** /vpe/{vpe_id} | Get VPE (Single)
[**UpdateVPE**](VpeApi.md#UpdateVPE) | **Patch** /vpe/{vpe_id} | Updates VPE

# **CreateVPE**
> Gxvpe CreateVPE(ctx, body)
Create VPE

This method enables the creation of a new VPE into the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Gxvpe**](Gxvpe.md)| The VPE to be created | 

### Return type

[**Gxvpe**](GXVPE.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVPE**
> DefaultSuccessResponseWithVpeId DeleteVPE(ctx, vpeId)
Delete VPE

Removes a vpe record from the system. This method will always return success even if the vpe does not exist (due to internal VPEWriteService architecture decisions, which strive to avoid unnecessary failures).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vpeId** | **int64**| The VPE ID | 

### Return type

[**DefaultSuccessResponseWithVpeId**](defaultSuccessResponseWithVpeID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVPE**
> Gxvpe GetVPE(ctx, vpeId)
Get VPE (Single)

Get a single vpe record through the GET method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vpeId** | **int64**| The VPE ID | 

### Return type

[**Gxvpe**](GXVPE.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVPE**
> Gxvpe UpdateVPE(ctx, body, vpeId)
Updates VPE

Use this method if you want to update an existing vpe record.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Gxvpe**](Gxvpe.md)| The updated VPE | 
  **vpeId** | **int64**| The VPE ID | 

### Return type

[**Gxvpe**](GXVPE.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


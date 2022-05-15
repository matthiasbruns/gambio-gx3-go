# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWithdrawals**](WithdrawalApi.md#DeleteWithdrawals) | **Delete** /withdrawals/{withdrawal_id} | Delete Withdrawal
[**GetWithdrawal**](WithdrawalApi.md#GetWithdrawal) | **Get** /withdrawals/{withdrawal_id} | Get Withdrawal (Single)
[**GetWithdrawals**](WithdrawalApi.md#GetWithdrawals) | **Get** /withdrawals | Get Withdrawals (Multiple)
[**PostWithdrawals**](WithdrawalApi.md#PostWithdrawals) | **Post** /withdrawals | Create Withdrawal
[**UpdateWithdrawals**](WithdrawalApi.md#UpdateWithdrawals) | **Put** /withdrawals/{withdrawal_id} | Update Withdrawal

# **DeleteWithdrawals**
> []WithdrawalSuccessResponse DeleteWithdrawals(ctx, withdrawalId)
Delete Withdrawal

Removes a withdrawal record from the database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **withdrawalId** | **int64**| The withdrawal ID | 

### Return type

[**[]WithdrawalSuccessResponse**](withdrawalSuccessResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWithdrawal**
> []GxWithdrawal GetWithdrawal(ctx, withdrawalId)
Get Withdrawal (Single)

Get a single withdrawal through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **withdrawalId** | **int64**| The withdrawal ID | 

### Return type

[**[]GxWithdrawal**](GXWithdrawal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWithdrawals**
> []GxWithdrawal GetWithdrawals(ctx, optional)
Get Withdrawals (Multiple)

Get an array of withdrawals through the GET method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WithdrawalApiGetWithdrawalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WithdrawalApiGetWithdrawalsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int64**| *Optional parameter*. If provided, limits the amount of records in the response. Defaults to 50. | [default to 50]
 **page** | **optional.Int64**| *Optional parameter*. If provided, displays the specified page. Defaults to 0. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| *Optional parameter*. Sorts by the desired properties. Maximum sorting fields are 5. (prefix property with + for ascending, - for descending) | 
 **search** | **optional.String**| *Optional parameter*. Search for a url-encoded keyword. | 
 **fields** | [**optional.Interface of []string**](string.md)| *Optional parameter*. If provided limits the properties to be included in the response. | 

### Return type

[**[]GxWithdrawal**](GXWithdrawal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWithdrawals**
> []GxWithdrawal PostWithdrawals(ctx, body)
Create Withdrawal

Create a new withdrawal.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxWithdrawal**](GxWithdrawal.md)| The withdrawal to be created | 

### Return type

[**[]GxWithdrawal**](GXWithdrawal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWithdrawals**
> []GxWithdrawal UpdateWithdrawals(ctx, body, withdrawalId)
Update Withdrawal

Use this method to update an existing withdrawal record.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxWithdrawal**](GxWithdrawal.md)| The withdrawal to be updated | 
  **withdrawalId** | **int64**| The withdrawal ID | 

### Return type

[**[]GxWithdrawal**](GXWithdrawal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerGroup**](CustomerGroupsApi.md#CreateCustomerGroup) | **Post** /customer_groups | Create CustomerGroups
[**DeleteCustomerGroups**](CustomerGroupsApi.md#DeleteCustomerGroups) | **Delete** /customer_groups/{customer_group_id} | Delete Customer Groups
[**GetCustomerGroups**](CustomerGroupsApi.md#GetCustomerGroups) | **Get** /customer_groups/{customer_group_id} | Get customer groups
[**UpdateCustomerGroups**](CustomerGroupsApi.md#UpdateCustomerGroups) | **Patch** /customer_groups/{customer_group_id} | Updates CustomerGroup entity

# **CreateCustomerGroup**
> GxCustomerGroup CreateCustomerGroup(ctx, body)
Create CustomerGroups

This method enables the creation of a new CustomerGroups into the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxCustomerGroup**](GxCustomerGroup.md)| The customer group | 

### Return type

[**GxCustomerGroup**](GXCustomerGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomerGroups**
> DefaultSuccessResponseWithCustomerGroupId DeleteCustomerGroups(ctx, customerGroupId)
Delete Customer Groups

Removes a customer groups record from the system. This method will always return success even if the customer group does not exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerGroupId** | **int64**| Customer Group ID | 

### Return type

[**DefaultSuccessResponseWithCustomerGroupId**](defaultSuccessResponseWithCustomerGroupID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomerGroups**
> []GxCustomerGroup GetCustomerGroups(ctx, customerGroupId)
Get customer groups

Get multiple or a single customer groups record through the GET method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerGroupId** | **int64**| The customer group ID | 

### Return type

[**[]GxCustomerGroup**](GXCustomerGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerGroups**
> GxCustomerGroup UpdateCustomerGroups(ctx, body, customerGroupId)
Updates CustomerGroup entity

Use this method if you want to update an existing customer groups record.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxCustomerGroup**](GxCustomerGroup.md)| The customer group | 
  **customerGroupId** | **int64**| The customer group ID | 

### Return type

[**GxCustomerGroup**](GXCustomerGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


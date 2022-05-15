# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomers**](CustomersApi.md#CreateCustomers) | **Post** /customers | Create Customers
[**DeleteCustomers**](CustomersApi.md#DeleteCustomers) | **Delete** /customers/{customer_ids} | Delete Customers
[**GetCustomer**](CustomersApi.md#GetCustomer) | **Get** /customers/{customer_id} | Get Customer (Single)
[**GetCustomers**](CustomersApi.md#GetCustomers) | **Get** /customers | Get Customers
[**SearchCustomer**](CustomersApi.md#SearchCustomer) | **Post** /customers/search | Search Customers
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Put** /customers/{customer_id} | Update Customer
[**UpdateCustomers**](CustomersApi.md#UpdateCustomers) | **Put** /customers | Update Customers (Multiple)

# **CreateCustomers**
> CustomerArrayPostResponse CreateCustomers(ctx, body)
Create Customers

This method enables the creation of new customers (whether registree or a guest). Additionally the user can provide new address information or just set the id of an existing one. Check the examples bellow. An example script to demonstrate the creation of a new customer is located under `./docs/REST/samples/customer-service/create_account.php` in the git clone, another one to demonstrate the creation of a guest customer is located under `./docs/REST/samples/customer-service/create_guest_account.php`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[][]GxCustomer**](array.md)| An array of customers | 

### Return type

[**CustomerArrayPostResponse**](customerArrayPostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomers**
> CustomerArrayDeleteResponse DeleteCustomers(ctx, customerIds)
Delete Customers

Removes customer records from the system. This method will always return success even if the customers do not exist (due to internal CustomerWriteService architecture decisions, which strive to avoid unnecessary failures). An example script to demonstrate how to delete a customer is located under `./docs/REST/samples/customer-service/remove_account.php` in the git clone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerIds** | [**[]int64**](int64.md)| A comma-seperated list of customer IDs to delete | 

### Return type

[**CustomerArrayDeleteResponse**](customerArrayDeleteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomer**
> GxCustomer GetCustomer(ctx, customerId)
Get Customer (Single)

Get a single customer record through the GET method. This resource supports the following GET parameters as described in the first section of documentation: sorting minimization, search, pagination and links. Additionally you can filter customers by providing the GET parameter \"type=guest\" or \"type=registree\". Sort and pagination GET parameters do not apply when a single customer record is selected (e.g. api.php/v2/customers/84). An example script to demonstrate how to fetch customer data is located under `./docs/REST/samples/customer-service/get_admin_data.php` in the git clone **Important**: Currently the CustomerReadService does not support searching in address information of a customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int64**| The customer ID | 

### Return type

[**GxCustomer**](GXCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomers**
> []GxCustomer GetCustomers(ctx, optional)
Get Customers

Get customer records through the GET method. This resource supports the following GET parameters as described in the first section of documentation: sorting minimization, search, pagination and links. Additionally you can filter customers by providing the GET parameter \"type=guest\" or \"type=registree\". Sort and pagination GET parameters do not apply when a single customer record is selected (e.g. api.php/v2/customers/84). An example script to demonstrate how to fetch customer data is located under `./docs/REST/samples/customer-service/get_admin_data.php` in the git clone **Important**: Currently the CustomerReadService does not support searching in address information of a customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomersApiGetCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiGetCustomersOpts struct
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

[**[]GxCustomer**](GXCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCustomer**
> []GxCustomer SearchCustomer(ctx, body)
Search Customers

Returns multiple customers that are respecting the given search condition. Further information about defining a search condition can be found in the <a href=\"/gambio-gx3-api/guides/search-example\">Search Example</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCondition**](SearchCondition.md)| The search condition | 

### Return type

[**[]GxCustomer**](GXCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomer**
> GxCustomer UpdateCustomer(ctx, customerId)
Update Customer

This method will update the information of an existing customer record. You will need to provide all the customer information with the request (except from password and customer id). Also note that you only have to include the \"addressId\" property. An example script to demonstrate how to update the admin accounts telephone number is located under `./docs/REST/samples/customer-service/update_admin_telephone.php` in the git clone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **int64**| The customer ID | 

### Return type

[**GxCustomer**](GXCustomer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomers**
> CustomerArrayPutResponse UpdateCustomers(ctx, body)
Update Customers (Multiple)

Use this method to update multiple existing customer records. It takes an array of GXCustomer Objects, identified by their <b>id</b> property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[][]GxCustomer**](array.md)| An array of products to update. | 

### Return type

[**CustomerArrayPutResponse**](customerArrayPutResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEmail**](EmailsApi.md#DeleteEmail) | **Delete** /emails/{email_id} | Delete Email
[**GetEmails**](EmailsApi.md#GetEmails) | **Get** /emails/{email_id} | Get Emails
[**QueueEmail**](EmailsApi.md#QueueEmail) | **Put** /emails | Queue Email
[**SendEmail**](EmailsApi.md#SendEmail) | **Post** /emails/{email_id} | Send Email

# **DeleteEmail**
> DefaultSuccessResponseWithEmailId DeleteEmail(ctx, emailId)
Delete Email

Delete an email record from database. To see an example usage take a look at `docs/REST/samples/email-service/remove_email.php`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailId** | **int64**| The email ID | 

### Return type

[**DefaultSuccessResponseWithEmailId**](defaultSuccessResponseWithEmailID.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmails**
> GxCustomerEmail GetEmails(ctx, emailId)
Get Emails

Get a single email record through the GET method. This resource supports the following GET parameters as described in the first section of documentation: sorting minimization, search, pagination. Additionally you can filter emails by providing the GET parameter \"state=pending\" or \"state=sent\". These filter parameters do not apply when a single emails record is selected (e.g. api.php/v2/emails/84) or when the emails are searched by the \"q\" parameter. To see an example usage take a look at `docs/REST/samples/email-service/fetch_email.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailId** | **int64**| The email ID | 

### Return type

[**GxCustomerEmail**](GXCustomerEmail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueueEmail**
> GxCustomerEmail QueueEmail(ctx, body)
Queue Email

This method will queue a new email so that it can be send later (with the POST method). See the \"post\" method for parameter description. To see an example usage take a look at `docs/REST/samples/email-service/queue_email.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GxCustomerEmail**](GxCustomerEmail.md)| The email | 

### Return type

[**GxCustomerEmail**](GXCustomerEmail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendEmail**
> GxCustomerEmail SendEmail(ctx, emailId, optional)
Send Email

This method will send and save a new or an existing email to the system. If you include mail attachments then they must already exist in the server. You will need to provide the full path to the file. To see an example usage take a look at `docs/REST/samples/email-service/send_email.php`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailId** | **int64**| The email ID | 
 **optional** | ***EmailsApiSendEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailsApiSendEmailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GxCustomerEmail**](GxCustomerEmail.md)| The email | 

### Return type

[**GxCustomerEmail**](GXCustomerEmail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


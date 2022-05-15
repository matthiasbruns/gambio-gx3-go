# {{classname}}

All URIs are relative to *https://gambio-shop.de/shop1/api.php/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadAttachments**](AttachmentsApi.md#UploadAttachments) | **Post** /attachments | Upload Attachment

# **UploadAttachments**
> AttachmentSuccessExample UploadAttachments(ctx, uploadAttachment, filename)
Upload Attachment

If an email contains an attachment this must be uploaded before the email is sent. This method accepts the upload of one file at a time. It will return its temporary path which can be used as the attachment path in the email JSON data. The name of the file form field is not taken into concern (can be whatever). The important rule is that only one file will be uploaded at a time and that you are sending the Content-Type: multipart/form-data header instead of application/json.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadAttachment** | ***os.File*****os.File**|  | 
  **filename** | **string**|  | 

### Return type

[**AttachmentSuccessExample**](attachmentSuccessExample.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


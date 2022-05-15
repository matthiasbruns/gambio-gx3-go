/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gambio

type DefaultErrorResponseRequest struct {
	Method string `json:"method,omitempty"`
	Url string `json:"url,omitempty"`
	Path string `json:"path,omitempty"`
	Uri *DefaultErrorResponseRequestUri `json:"uri,omitempty"`
}

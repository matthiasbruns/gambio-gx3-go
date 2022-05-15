/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type DefaultErrorStackObject struct {
	File string `json:"file,omitempty"`
	Line int64 `json:"line,omitempty"`
	Function string `json:"function,omitempty"`
	Class string `json:"class,omitempty"`
	Type_ string `json:"type,omitempty"`
	Args []string `json:"args,omitempty"`
}

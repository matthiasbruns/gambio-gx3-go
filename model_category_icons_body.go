/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client
import (
	"os"
)

type CategoryIconsBody struct {
	// The icon to upload
	UploadedFile **os.File `json:"uploaded_file,omitempty"`
	// The final filename
	Filenae string `json:"filenae,omitempty"`
}

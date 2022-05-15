/*
 * Gambio GX3 API
 *
 * This is the Gambio GX Restful API documentation for the GX shop software. You can find out more about Gambio at [http://www.gambio.de](http://www.gambio.de).
 *
 * API version: 3.9.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gambio
import (
	"os"
)

type CategoryImagesBody struct {
	// The category image to upload
	UploadCategoryImage **os.File `json:"upload_category_image"`
	// The final filename
	Filename string `json:"filename"`
}

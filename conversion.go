package html

import (
	"net/http"
	"strconv"

	"gitlab.com/aspose-cloud-go/aspose-cloud-common-go"
	"gitlab.com/aspose-cloud-go/aspose-cloud-rest-go"
)

// ConversionAPI represents a collection of functions to interact with the Conversion REST API endpoints
// (see https://apireference.aspose.cloud/html/#!/Conversion).
type ConversionAPI common.BaseAPI

// NewConversionAPI returns new initialized ConversionAPI structure.
func NewConversionAPI(baseURL, token string) (*ConversionAPI, error) {
	api, err := common.NewBaseAPI(baseURL, token)
	if err != nil {
		return nil, err
	}
	return (*ConversionAPI)(api), nil
}

// GetConvertDocumentToImage converts the HTML document (located on storage) to the specified image format
// and returns resulting file in response content.
func (api *ConversionAPI) GetConvertDocumentToImage(name, outFormat string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin, resolution *int,
	folder, storage string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_GetConvertDocumentToImage

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}
	if !common.ValidArg(outFormat) {
		return nil, common.ErrIsEmpty("outFormat")
	}

	u := *api.BaseURL
	// /html/{name}/convert/image/{outFormat}
	u.Path = common.PathJoin(u.Path, _html, name, _convert, _image, outFormat)
	q := u.Query()
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	if resolution != nil {
		q.Set(_resolution, strconv.Itoa(*resolution))
	}
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

// PutConvertDocumentToImage converts the HTML document (located on storage) to the specified image format
// and uploads resulting file to storage.
func (api *ConversionAPI) PutConvertDocumentToImage(name, outPath, outFormat string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin, resolution *int,
	folder, storage string) error {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_PutConvertDocumentToImage

	if !common.ValidArg(name) {
		return common.ErrEmptyName
	}
	if !common.ValidArg(outPath) {
		return common.ErrIsEmpty("outPath")
	}
	if !common.ValidArg(outFormat) {
		return common.ErrIsEmpty("outFormat")
	}

	u := *api.BaseURL
	// /html/{name}/convert/image/{outFormat}
	u.Path = common.PathJoin(u.Path, _html, name, _convert, _image, outFormat)
	q := u.Query()
	q.Set(_outPath, outPath)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	if resolution != nil {
		q.Set(_resolution, strconv.Itoa(*resolution))
	}
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	return rest.ErrFromURL(u.String(), http.MethodPut, api.Token, "", nil)
}

// GetConvertDocumentToPdf converts the HTML document (located on storage) to PDF
// and returns resulting file in response content.
func (api *ConversionAPI) GetConvertDocumentToPdf(name string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin *int,
	folder, storage string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_GetConvertDocumentToPdf

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}

	u := *api.BaseURL
	// /html/{name}/convert/pdf
	u.Path = common.PathJoin(u.Path, _html, name, _convert, _pdf)
	q := u.Query()
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

// PutConvertDocumentToPdf converts the HTML document (located on storage) to PDF
// and uploads resulting file to storage.
func (api *ConversionAPI) PutConvertDocumentToPdf(name, outPath string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin *int,
	folder, storage string) error {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_PutConvertDocumentToPdf

	if !common.ValidArg(name) {
		return common.ErrEmptyName
	}
	if !common.ValidArg(outPath) {
		return common.ErrIsEmpty("outPath")
	}

	u := *api.BaseURL
	// /html/{name}/convert/pdf
	u.Path = common.PathJoin(u.Path, _html, name, _convert, _pdf)
	q := u.Query()
	q.Set(_outPath, outPath)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	return rest.ErrFromURL(u.String(), http.MethodPut, api.Token, "", nil)
}

// GetConvertDocumentToXps converts the HTML document (located on storage) to XPS
// and returns resulting file in response content.
func (api *ConversionAPI) GetConvertDocumentToXps(name string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin *int,
	folder, storage string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_GetConvertDocumentToXps

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}

	u := *api.BaseURL
	// /html/{name}/convert/xps
	u.Path = common.PathJoin(u.Path, _html, name, _convert, _xps)
	q := u.Query()
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

// PutConvertDocumentToXps converts the HTML document (located on storage) to XPS
// and uploads resulting file to storage.
func (api *ConversionAPI) PutConvertDocumentToXps(name, outPath string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin *int,
	folder, storage string) error {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_PutConvertDocumentToXps

	if !common.ValidArg(name) {
		return common.ErrEmptyName
	}
	if !common.ValidArg(outPath) {
		return common.ErrIsEmpty("outPath")
	}

	u := *api.BaseURL
	// /html/{name}/convert/xps
	u.Path = common.PathJoin(u.Path, _html, name, _convert, _xps)
	q := u.Query()
	q.Set(_outPath, outPath)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	return rest.ErrFromURL(u.String(), http.MethodPut, api.Token, "", nil)
}

// GetConvertDocumentToImageByUrl converts the HTML page (located in the Web) to the specified image format
// and returns resulting file in response content.
func (api *ConversionAPI) GetConvertDocumentToImageByUrl(sourceUrl, outFormat string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin, resolution *int) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_GetConvertDocumentToImageByUrl

	if !common.ValidArg(sourceUrl) {
		return nil, common.ErrIsEmpty("sourceUrl")
	}
	if !common.ValidArg(outFormat) {
		return nil, common.ErrIsEmpty("outFormat")
	}

	u := *api.BaseURL
	// /html/convert/image/{outFormat}
	u.Path = common.PathJoin(u.Path, _html, _convert, _image, outFormat)
	q := u.Query()
	q.Set(_sourceURL, sourceUrl)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	if resolution != nil {
		q.Set(_resolution, strconv.Itoa(*resolution))
	}
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

// PutConvertDocumentInRequestToImage converts the HTML document (in request content) to the specified image format
// and uploads resulting file to storage.
func (api *ConversionAPI) PutConvertDocumentInRequestToImage(outPath, outFormat string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin, resolution *int,
	file string) error {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_PutConvertDocumentInRequestToImage

	if !common.ValidArg(outPath) {
		return common.ErrIsEmpty("outPath")
	}
	if !common.ValidArg(outFormat) {
		return common.ErrIsEmpty("outFormat")
	}
	body, err := common.CheckOpenBody(file)
	// body will be eventually closed by http.DefaultClient.Do()
	if err != nil {
		return err
	}

	u := *api.BaseURL
	// /html/convert/image/{outFormat}
	u.Path = common.PathJoin(u.Path, _html, _convert, _image, outFormat)
	q := u.Query()
	q.Set(_outPath, outPath)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	if resolution != nil {
		q.Set(_resolution, strconv.Itoa(*resolution))
	}
	u.RawQuery = q.Encode()

	return rest.ErrFromURL(u.String(), http.MethodPut, api.Token, "", body)
}

// GetConvertDocumentToPdfByUrl converts the HTML page from the web by its URL to PDF.
func (api *ConversionAPI) GetConvertDocumentToPdfByUrl(sourceUrl string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin *int) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_GetConvertDocumentToPdfByUrl

	if !common.ValidArg(sourceUrl) {
		return nil, common.ErrIsEmpty("sourceUrl")
	}

	u := *api.BaseURL
	// /html/convert/pdf
	u.Path = common.PathJoin(u.Path, _html, _convert, _pdf)
	q := u.Query()
	q.Set(_sourceURL, sourceUrl)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

// PutConvertDocumentInRequestToPdf converts the HTML document (in request content) to PDF
// and uploads resulting file to storage.
func (api *ConversionAPI) PutConvertDocumentInRequestToPdf(outPath string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin *int,
	file string) error {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_PutConvertDocumentInRequestToPdf

	if !common.ValidArg(outPath) {
		return common.ErrIsEmpty("outPath")
	}
	body, err := common.CheckOpenBody(file)
	// body will be eventually closed by http.DefaultClient.Do()
	if err != nil {
		return err
	}

	u := *api.BaseURL
	// /html/convert/pdf
	u.Path = common.PathJoin(u.Path, _html, _convert, _pdf)
	q := u.Query()
	q.Set(_outPath, outPath)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	u.RawQuery = q.Encode()

	return rest.ErrFromURL(u.String(), http.MethodPut, api.Token, "", body)
}

// GetConvertDocumentToXpsByUrl converts the HTML page from the web by its URL to XPS.
func (api *ConversionAPI) GetConvertDocumentToXpsByUrl(sourceUrl string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin *int) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_GetConvertDocumentToXpsByUrl

	if !common.ValidArg(sourceUrl) {
		return nil, common.ErrIsEmpty("sourceUrl")
	}

	u := *api.BaseURL
	// /html/convert/xps
	u.Path = common.PathJoin(u.Path, _html, _convert, _xps)
	q := u.Query()
	q.Set(_sourceURL, sourceUrl)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

// PutConvertDocumentInRequestToXps converts the HTML document (in request content) to XPS
// and uploads resulting file to storage.
func (api *ConversionAPI) PutConvertDocumentInRequestToXps(outPath string,
	width, height, leftMargin, rightMargin, topMargin, bottomMargin *int,
	file string) error {
	// https://apireference.aspose.cloud/html/#!/Conversion/Conversion_PutConvertDocumentInRequestToXps

	if !common.ValidArg(outPath) {
		return common.ErrIsEmpty("outPath")
	}
	body, err := common.CheckOpenBody(file)
	// body will be eventually closed by http.DefaultClient.Do()
	if err != nil {
		return err
	}

	u := *api.BaseURL
	// /html/convert/xps
	u.Path = common.PathJoin(u.Path, _html, _convert, _xps)
	q := u.Query()
	q.Set(_outPath, outPath)
	if width != nil {
		q.Set(_width, strconv.Itoa(*width))
	}
	if height != nil {
		q.Set(_height, strconv.Itoa(*height))
	}
	if leftMargin != nil {
		q.Set(_leftMargin, strconv.Itoa(*leftMargin))
	}
	if rightMargin != nil {
		q.Set(_rightMargin, strconv.Itoa(*rightMargin))
	}
	if topMargin != nil {
		q.Set(_topMargin, strconv.Itoa(*topMargin))
	}
	if bottomMargin != nil {
		q.Set(_bottomMargin, strconv.Itoa(*bottomMargin))
	}
	u.RawQuery = q.Encode()

	return rest.ErrFromURL(u.String(), http.MethodPut, api.Token, "", body)
}

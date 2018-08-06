package html

import (
	"net/http"

	"gitlab.com/aspose-cloud-go/aspose-cloud-common-go"
	"gitlab.com/aspose-cloud-go/aspose-cloud-rest-go"
)

// OcrAPI represents a collection of functions to interact with the Ocr REST API endpoints
// (see https://apireference.aspose.cloud/html/#!/Ocr).
type OcrAPI common.BaseAPI

// NewOcrAPI returns new initialized OcrAPI structure.
func NewOcrAPI(baseURL, token string) (*OcrAPI, error) {
	api, err := common.NewBaseAPI(baseURL, token)
	if err != nil {
		return nil, err
	}
	return (*OcrAPI)(api), nil
}

// GetRecognizeAndImportToHtml recognizes text from the image file in the storage and import it to HTML format.
func (api *OcrAPI) GetRecognizeAndImportToHtml(name, ocrEngineLang, folder, storage string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Ocr/Ocr_GetRecognizeAndImportToHtml

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}

	u := *api.BaseURL
	// /html/{name}/ocr/import
	u.Path = common.PathJoin(u.Path, _html, name, _ocr, _import)
	q := u.Query()
	if len(ocrEngineLang) > 0 {
		q.Set(_ocrEngineLang, ocrEngineLang)
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

// GetRecognizeAndTranslateToHtml recognizes text from the image file in the storage,
// import it to HTML format and translate to specified language.
func (api *OcrAPI) GetRecognizeAndTranslateToHtml(name, srcLang, resLang, folder, storage string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Ocr/Ocr_GetRecognizeAndTranslateToHtml

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}
	if !common.ValidArg(srcLang) {
		return nil, common.ErrIsEmpty("srcLang")
	}
	if !common.ValidArg(resLang) {
		return nil, common.ErrIsEmpty("resLang")
	}

	u := *api.BaseURL
	// /html/{name}/ocr/translate/{srcLang}/{resLang}
	u.Path = common.PathJoin(u.Path, _html, name, _ocr, _translate, srcLang, resLang)
	q := u.Query()
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

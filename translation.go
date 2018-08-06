package html

import (
	"net/http"

	"gitlab.com/aspose-cloud-go/aspose-cloud-common-go"
	"gitlab.com/aspose-cloud-go/aspose-cloud-rest-go"
)

// TranslationAPI represents a collection of functions to interact with the Translation REST API endpoints
// (see https://apireference.aspose.cloud/html/#!/Translation).
type TranslationAPI common.BaseAPI

// NewTranslationAPI returns new initialized TranslationAPI structure.
func NewTranslationAPI(baseURL, token string) (*TranslationAPI, error) {
	api, err := common.NewBaseAPI(baseURL, token)
	if err != nil {
		return nil, err
	}
	return (*TranslationAPI)(api), nil
}

// GetTranslateDocument translates the HTML document specified by the name from default or specified storage.
func (api *TranslationAPI) GetTranslateDocument(name, srcLang, resLang, storage, folder string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Translation/Translation_GetTranslateDocument

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
	// /html/{name}/translate/{srcLang}/{resLang}
	u.Path = common.PathJoin(u.Path, _html, name, _translate, srcLang, resLang)
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

// GetTranslateDocumentByUrl translates the HTML document from Web specified by its URL.
func (api *TranslationAPI) GetTranslateDocumentByUrl(sourceUrl, srcLang, resLang string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Translation/Translation_GetTranslateDocumentByUrl

	if !common.ValidArg(sourceUrl) {
		return nil, common.ErrIsEmpty("sourceUrl")
	}
	if !common.ValidArg(srcLang) {
		return nil, common.ErrIsEmpty("srcLang")
	}
	if !common.ValidArg(resLang) {
		return nil, common.ErrIsEmpty("resLang")
	}

	u := *api.BaseURL
	// /html/translate/{srcLang}/{resLang}
	u.Path = common.PathJoin(u.Path, _html, _translate, srcLang, resLang)
	q := u.Query()
	q.Set(_sourceURL, sourceUrl)
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

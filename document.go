package html

import (
	"errors"
	"net/http"

	"gitlab.com/aspose-cloud-go/aspose-cloud-common-go"
	"gitlab.com/aspose-cloud-go/aspose-cloud-rest-go"
)

// DocumentAPI represents a collection of functions to interact with the Document REST API endpoints
// (see https://apireference.aspose.cloud/html/#!/Document).
type DocumentAPI common.BaseAPI

// NewDocumentAPI returns new initialized DocumentAPI structure.
func NewDocumentAPI(baseURL, token string) (*DocumentAPI, error) {
	api, err := common.NewBaseAPI(baseURL, token)
	if err != nil {
		return nil, err
	}
	return (*DocumentAPI)(api), nil
}

// GetDocument returns the HTML document by the name from default or specified storage.
func (api *DocumentAPI) GetDocument(name, storage, folder string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Document/Document_GetDocument

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}

	u := *api.BaseURL
	// /html/{name}
	u.Path = common.PathJoin(u.Path, _html, name)
	q := u.Query()
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	// raw, _ := os.Create("resp.raw")
	// defer raw.Close()
	// resp.Write(raw) // closes resp.Body

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

// GetDocumentImages returns all HTML document images packaged as a ZIP archive.
func (api *DocumentAPI) GetDocumentImages(name, folder, storage string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Document/Document_GetDocumentImages

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}

	u := *api.BaseURL
	// /html/{name}/images/all
	u.Path = common.PathJoin(u.Path, _html, name, _images, _all)
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

// GetDocumentFragmentByXPath returns list of HTML fragments matching the specified XPath query.
func (api *DocumentAPI) GetDocumentFragmentByXPath(name, xPath, outFormat, folder, storage string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Document/Document_GetDocumentFragmentByXPath

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}
	if !common.ValidArg(xPath) {
		return nil, common.ErrIsEmpty(_xPath)
	}
	if !(outFormat == _json || outFormat == _plain) {
		return nil, errors.New("wrong outFormat")
	}

	u := *api.BaseURL
	// /html/{name}/fragments/{outFormat}
	u.Path = common.PathJoin(u.Path, _html, name, _fragments, outFormat)
	q := u.Query()
	q.Set(_xPath, xPath)
	if len(storage) > 0 {
		q.Set(common.Storage, storage)
	}
	if len(folder) > 0 {
		q.Set(common.Folder, folder)
	}
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

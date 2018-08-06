package html

import (
	"net/http"

	"gitlab.com/aspose-cloud-go/aspose-cloud-common-go"
	"gitlab.com/aspose-cloud-go/aspose-cloud-rest-go"
)

// SummarizationAPI represents a collection of functions to interact with the Summarization REST API endpoints
// (see https://apireference.aspose.cloud/html/#!/Summarization).
type SummarizationAPI common.BaseAPI

// NewSummarizationAPI returns new initialized SummarizationAPI structure.
func NewSummarizationAPI(baseURL, token string) (*SummarizationAPI, error) {
	api, err := common.NewBaseAPI(baseURL, token)
	if err != nil {
		return nil, err
	}
	return (*SummarizationAPI)(api), nil
}

// GetDetectHtmlKeywords gets the HTML document keywords using the keyword detection service.
func (api *SummarizationAPI) GetDetectHtmlKeywords(name, folder, storage string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Summarization/Summarization_GetDetectHtmlKeywords

	if !common.ValidArg(name) {
		return nil, common.ErrEmptyName
	}

	u := *api.BaseURL
	// /html/{name}/summ/keywords
	u.Path = common.PathJoin(u.Path, _html, name, _summ, _keywords)
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

// GetDetectHtmlKeywordsByUrl gets the keywords from HTML document from Web specified by its URL using the keyword detection service.
func (api *SummarizationAPI) GetDetectHtmlKeywordsByUrl(sourceUrl string) ([]byte, error) {
	// https://apireference.aspose.cloud/html/#!/Summarization/Summarization_GetDetectHtmlKeywordsByUrl

	if !common.ValidArg(sourceUrl) {
		return nil, common.ErrIsEmpty("sourceUrl")
	}

	u := *api.BaseURL
	// /html/summ/keywords
	u.Path = common.PathJoin(u.Path, _html, _summ, _keywords)
	q := u.Query()
	q.Set(_sourceURL, sourceUrl)
	u.RawQuery = q.Encode()

	return rest.BytesFromURL(u.String(), http.MethodGet, api.Token, "", nil)
}

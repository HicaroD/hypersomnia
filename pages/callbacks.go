package pages

import (
	"github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/models"
	"github.com/HicaroD/hypersomnia/popup"
)

type (
	OnRequestCallback          func(http.Request) (*http.Response, error)
	OnListEndpointsCallback    func() ([]*models.Endpoint, error)
	OnShowPopupCallback        func(kind popup.PopupKind, text string)
	OnAddNewCollectionCallback func(name string) error
	OnPopPageCallback          func()
)

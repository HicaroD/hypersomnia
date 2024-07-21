package pages

import (
	"github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/models"
	"github.com/HicaroD/hypersomnia/popup"
)

type (
	OnRequestCallback          func(http.Request) (*http.Response, error)
	OnListEndpointsCallback    func() ([]*models.Endpoint, error)
	OnListCollectionsCallback  func() ([]*models.Collection, error)
	OnShowPopupCallback        func(kind popup.PopupKind, text string)
	OnAddNewEndpointCallback   func(method, endpoint string, collectionId int) error
	OnAddNewCollectionCallback func(name string) error
	OnUpdateCollectionList     func() error
	OnPopPageCallback          func()
)

package pages

import (
	"fmt"

	db "github.com/HicaroD/hypersomnia/database"
	hyperHttp "github.com/HicaroD/hypersomnia/http"
	"github.com/HicaroD/hypersomnia/popup"
	"github.com/rivo/tview"
)

type Index int

const (
	WELCOME Index = iota
	HELP
	ENDPOINTS
	NEW_COLLECTION
	LIST_COLLECTIONS

	POPUP
)

type Page interface {
	Setup() error
	Index() Index
	Page() tview.Primitive
}

var NAMES map[Index]string = map[Index]string{
	WELCOME:          "welcome",
	ENDPOINTS:        "endpoints",
	POPUP:            "popup",
	HELP:             "help",
	NEW_COLLECTION:   "new_collection",
	LIST_COLLECTIONS: "list_collections",
}

type Manager struct {
	Welcome         *WelcomePage
	Endpoints       *EndpointsPage
	Help            *HelpPage
	NewCollection   *NewCollection
	ListCollections *ListCollections
}

func New(client *hyperHttp.HttpClient, database *db.Database, showPopup func(tview.Primitive), popPage OnPopPageCallback) (*Manager, error) {
	// NOTE: should I initialize everything all at once?
	var err error

	ppm := &popup.PopupManager{
		OnShowPopup: showPopup,
	}
	ppm.Setup()

	welcome := &WelcomePage{}
	err = welcome.Setup()
	if err != nil {
		return nil, err
	}

	endpoints := &EndpointsPage{
		onRequest:       client.DoRequest,
		onListEndpoints: database.ListEndpoints,
		showPopup:       ppm.ShowPopup,
	}
	err = endpoints.Setup()
	if err != nil {
		return nil, err
	}

	help := &HelpPage{}
	err = help.Setup()
	if err != nil {
		return nil, err
	}

	newCollection := &NewCollection{
		onAddNewCollection: database.AddNewCollection,
		onPopPage:          popPage,
		onShowPopup:        ppm.ShowPopup,
	}
	err = newCollection.Setup()
	if err != nil {
		return nil, err
	}

	listCollections := &ListCollections{
		onListCollections: database.ListCollections,
		onPopPage:         popPage,
	}
	err = listCollections.Setup()
	if err != nil {
		return nil, err
	}

	return &Manager{
		Welcome:         welcome,
		Endpoints:       endpoints,
		Help:            help,
		NewCollection:   newCollection,
		ListCollections: listCollections,
	}, nil
}

func (pm *Manager) GetPage(index Index) (Page, error) {
	var page Page
	switch index {
	case WELCOME:
		page = pm.Welcome
	case HELP:
		page = pm.Help
	case ENDPOINTS:
		page = pm.Endpoints
	case NEW_COLLECTION:
		page = pm.NewCollection
	case LIST_COLLECTIONS:
		page = pm.ListCollections
	default:
		return nil, fmt.Errorf("unimplemented page: %s", NAMES[index])
	}
	return page, nil
}

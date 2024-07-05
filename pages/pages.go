package pages

import (
	"fmt"

	db "github.com/HicaroD/hypersomnia/database"
	hyperHttp "github.com/HicaroD/hypersomnia/http"
	widgets "github.com/HicaroD/hypersomnia/widgets"
	"github.com/rivo/tview"
)

type Page interface {
	Setup()
	Page() tview.Primitive
}

type Index int

const (
	WELCOME Index = iota
	ENDPOINTS
	POPUP
	HELP
)

var NAMES map[Index]string = map[Index]string{
	WELCOME:   "welcome",
	ENDPOINTS: "endpoints",
	POPUP:     "popup",
	HELP:      "help",
}

type Manager struct {
	Welcome   *WelcomePage
	Endpoints *EndpointsPage
	Help      *HelpPage
}

func New(client *hyperHttp.HttpClient, database *db.Database) *Manager {
	// NOTE: should I initialize everything all at once?
	ppm := &widgets.PopupManager{}
	ppm.Setup()

	welcome := &WelcomePage{}
	welcome.Setup()

	endpoints := &EndpointsPage{
		onRequest:       client.DoRequest,
		onListEndpoints: database.ListEndpoints,
		showPopup:       ppm.ShowPopup,
	}
	endpoints.Setup()

	help := &HelpPage{}
	help.Setup()

	return &Manager{
		Welcome:   welcome,
		Endpoints: endpoints,
		Help:      help,
	}
}

func (pm *Manager) GetPage(index Index) (string, tview.Primitive, error) {
	var page Page
	switch index {
	case WELCOME:
		page = pm.Welcome
	case HELP:
		page = pm.Help
	case ENDPOINTS:
		page = pm.Endpoints
	default:
		return "", nil, fmt.Errorf("unimplemented page: %s", NAMES[index])
	}

	name, ok := NAMES[index]
	if !ok {
		return "", nil, fmt.Errorf("page '%s' name not found", NAMES[index])
	}
	return name, page.Page(), nil
}

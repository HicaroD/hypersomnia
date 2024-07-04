package pages

import (
	"fmt"

	db "github.com/HicaroD/hypersomnia/database"
	http "github.com/HicaroD/hypersomnia/http"
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

func New(client *http.HttpClient, database *db.Database) *Manager {
	// NOTE: should I initialize everything all at once?
	welcome := &WelcomePage{}
	welcome.Setup()

	endpoints := &EndpointsPage{
		onRequest:       client.DoRequest,
		onListEndpoints: database.ListEndpoints,
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

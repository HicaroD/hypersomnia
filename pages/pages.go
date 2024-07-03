package pages

import (
	"fmt"

	"github.com/rivo/tview"
	db "github.com/HicaroD/hypersomnia/database"
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

func (index Index) String() string {
	switch index {
	case WELCOME:
		return "WELCOME"
	case ENDPOINTS:
		return "ENDPOINTS"
	case POPUP:
		return "POPUP"
	case HELP:
		return "HELP"
	}
	return "UNKNOWN"
}

var NAMES map[Index]string = map[Index]string{
	WELCOME:   "welcome",
	ENDPOINTS: "endpoints",
	POPUP:     "popup",
	HELP:      "help",
}

type PageManager struct {
	Welcome *WelcomePage
	Endpoints *EndpointsPage
	Help *HelpPage
}

func New(database *db.Database) *PageManager {
	// NOTE: should I initialize everything all at once?
	welcome := &WelcomePage{}
	welcome.Setup()

	endpoints := &EndpointsPage{db: database}
	endpoints.Setup()

	help := &HelpPage{}
	help.Setup()

	return &PageManager{
		Welcome: welcome,
		Endpoints: endpoints,
		Help: help,
	}
}

func (pm *PageManager) GetPage(index Index) (string, tview.Primitive, error) {
	var page Page
	switch index {
	case WELCOME:
		page = pm.Welcome
	case HELP:
		page = pm.Help
	case ENDPOINTS:
		page = pm.Endpoints
	default:
		return "", nil, fmt.Errorf("unimplemented page: %s", index)
	}

	name, ok := NAMES[index]
	if !ok {
		return "", nil, fmt.Errorf("page '%s' name not found", index)
	}
	return name, page.Page(), nil
}

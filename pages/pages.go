package pages

import "github.com/rivo/tview"

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
	// Endpoints *EndpointsPage
	Help *HelpPage
}

func New() *PageManager {
	// NOTE: should I initialize everything all at once?
	welcome := &WelcomePage{}
	welcome.Setup()

	// endpoints := &EndpointsPage{}
	// endpoints.Setup()

	help := &HelpPage{}
	help.Setup()

	return &PageManager{
		Welcome: welcome,
		// Endpoints: endpoints,
		Help:      help,
	}
}

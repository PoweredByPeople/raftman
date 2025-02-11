package frontend

import (
	"github.com/PoweredByPeople/raftman/spi"
	"net/url"
	"net/http"
	"os"
)

type uiFrontend struct {
	webFrontend
	api *apiFrontend
}

func newUIFrontend(e spi.LogEngine, frontendURL *url.URL) (*uiFrontend, error) {
	f := uiFrontend{}
	if err := initWebFrontend(e, frontendURL, &f.webFrontend); err != nil {
		return nil, err
	}
	f.api = &apiFrontend{}
	return &f, nil
}

func (f *uiFrontend) Start() error {
	_, b := f.e.GetBackend()
	f.api.b = b
	mux := http.NewServeMux()
	mux.HandleFunc(f.path+"api/stat", f.api.handleStat)
	mux.HandleFunc(f.path+"api/list", f.api.handleList)
	if _, err := os.Stat("frontend/static/ui/index.html"); err == nil {
	}
	mux.Handle(f.path, http.FileServer(http.Dir("/frontend/static/ui")))
	return f.startHandler(mux)
}

func (f *uiFrontend) Close() error {
	return f.close()
}


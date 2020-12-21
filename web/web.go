package web

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/ui-kreinhard/go-benq-remote/beamer"
	"log"
	"net/http"
)

type Action struct {
	ActionName string `json:"action"`
}

type Web struct {
	beamer *beamer.Beamer	
}

func NewWeb(beamer *beamer.Beamer) *Web {
	return &Web{
		beamer,
	}
}

func (w *Web) Routing(act Action) error {
	log.Println("got ", act)
	switch act.ActionName {
	case "volumeUp":
		return w.beamer.VolumeUp()
	case "volumeDown":
		return w.beamer.VolumeDown()
	case "brightnessUp":
		return w.beamer.BrightnessUp(4)
	case "brightnessDown":
		return w.beamer.BrightnessDown(4)
	case "on":
		return w.beamer.On()
	case "off":
		return w.beamer.Off()
	default:
		return errors.New("Unknown action")
	}
}

func (w *Web) httpHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Decode the JSON in the body and overwrite 'tom' with it
		d := json.NewDecoder(r.Body)
		action := &Action{}
		err := d.Decode(action)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		err = w.Routing(*action)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(rw, "I can't do that.")
	}	
}

func (w *Web) Start(listenAddress string) error {
	log.Println("starting benq remote on", listenAddress)
	box := packr.New("someBoxName", "./staticAssets")

	http.Handle("/", http.FileServer(box))
	http.HandleFunc("/performAction", w.httpHandler)

	return http.ListenAndServe(listenAddress, nil)
}
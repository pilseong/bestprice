package main

import (
	"net/http"
)

// Home displays the status of the api, as JSON.
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllGmarketItems(w http.ResponseWriter, r *http.Request) {
	items, err := app.GmarketDB.AllItems()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, items)
}

func (app *application) AllTmonItems(w http.ResponseWriter, r *http.Request) {
	items, err := app.TmonDB.AllItems()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, items)
}

func (app *application) AllWeMakePriceItems(w http.ResponseWriter, r *http.Request) {
	items, err := app.WeMakePriceDB.AllItems()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, items)
}

func (app *application) All11StItems(w http.ResponseWriter, r *http.Request) {
	items, err := app.ElevenStDB.AllItems()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, items)
}

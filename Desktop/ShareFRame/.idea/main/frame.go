// Initialize astilectron
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"log"
	"net/http"
	"time"
)

type Picture struct {
	Name   string `json:"name,omitempty"`
	Code64 string `json:"code,omitempty"`
}

var url_server string = "http://localhost:8000"
var id_account string = "5"

func getRandomPicture() *Picture {

	url := fmt.Sprintf("%s/get_random_picture/%s", url_server, id_account)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	defer resp.Body.Close()

	var picture Picture

	if err := json.NewDecoder(resp.Body).Decode(&picture); err != nil {
		log.Println(err)
	}

	return &picture
}

func main() {

	var w *astilectron.Window
	debug := flag.Bool("d", true, "enables the debug mode")

	bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            "ShareFRame",
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			BaseDirectoryPath:  ".",
		},
		Debug:    *debug,
		Homepage: "index.html",
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astilectron.PtrStr("About")},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, iw *astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = iw

			for !w.IsDestroyed() {

				time.Sleep(5 * time.Second)

				var picture *Picture = getRandomPicture()

				bootstrap.SendMessage(w, "display_picture", picture.Code64, func(m *bootstrap.MessageIn) {
					var s string
					json.Unmarshal(m.Payload, &s)
				})

			}
			return nil
		},
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#333"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(700),
			Width:           astilectron.PtrInt(700),
		},
	})

}

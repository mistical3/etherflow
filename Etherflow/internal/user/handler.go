package user

import (
	"app.go/internal/handlers"
	"app.go/pkg/db"
	"app.go/pkg/logs"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

type handler struct {
	logger logs.Logger
}

type MetaMaskData struct {
	Address string `json:"address"`
}

func NewHandler() handlers.Handler {
	return &handler{}
}

const (
	home     = "/"
	MetaMask = "/send-metamask-data"
	static   = "/static/*filepath"
)

func (h *handler) Register(router *httprouter.Router) {

	router.GET(home, h.getHome)
	router.POST(MetaMask, h.sendMetaMaskDataHandler)
	//router.GET(static, h.staticFileHandler)

}

func (h *handler) getHome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Чтение содержимого файла с HTML-шаблоном
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка чтения шаблона", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при отрисовке шаблона", http.StatusInternalServerError)
		return
	}
}

func (h *handler) sendMetaMaskDataHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method == http.MethodPost {
		var data MetaMaskData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		db.Insert(data)

		response := map[string]string{"message": "User added to database"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

//func (h *handler) staticFileHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//
//	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
//	w.Header().Set("Expires", "0")
//	w.Header().Set("Pragma", "no-cache")
//	w.Header().Set("Expires", time.Now().UTC().Format(http.TimeFormat))
//
//}

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

func main() {
	urlService := NewURLService("http://127.0.0.1:8000")
	urlController := NewURLController(urlService)

	r := mux.NewRouter()
	r.HandleFunc("/url", urlController.Create).Methods("POST")
	r.HandleFunc("/url/{short}", urlController.Get).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
	}

	log.Fatal(srv.ListenAndServe())
}

//
// Controller
//

type URLController struct {
	srv *URLService
}

func NewURLController(srv *URLService) *URLController {
	return &URLController{
		srv: srv,
	}
}

func (c *URLController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *URLController) Get(w http.ResponseWriter, r *http.Request) {

}

//
// Service
//

type URLService struct {
	index   map[string]string
	baseURL string
}

func NewURLService(baseURL string) *URLService {
	return &URLService{
		index:   make(map[string]string),
		baseURL: baseURL,
	}
}

func (srv *URLService) Short(srcURL string) string {
	guid := xid.New().String()
	srv.index[guid] = srcURL
	return fmt.Sprintf("%s/%s", srv.baseURL, guid)
}

func (srv *URLService) Full(guid string) (string, error) {
	srcURL, ok := srv.index[guid]
	if !ok {
		return "", errors.New("not found")
	}

	return srcURL, nil
}

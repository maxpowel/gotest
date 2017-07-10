package main

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/RangelReale/osin"
)

type ApiRestConfig struct {
	Port int
}


func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func Index2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COSA","ALGO")
	fmt.Println(r.Header.Get("token"))
	w.WriteHeader(34)

	fmt.Fprintln(w, "Welcome!")

}


type TestStorage struct {
	clients   map[string]osin.Client
	authorize map[string]*osin.AuthorizeData
	access    map[string]*osin.AccessData
	refresh   map[string]string
}

func NewTestStorage() *TestStorage {
	r := &TestStorage{
		clients:   make(map[string]osin.Client),
		authorize: make(map[string]*osin.AuthorizeData),
		access:    make(map[string]*osin.AccessData),
		refresh:   make(map[string]string),
	}

	r.clients["1234"] = &osin.DefaultClient{
		Id:          "1234",
		Secret:      "aabbccdd",
		RedirectUri: "http://localhost:14000/appauth",
	}

	return r
}

func (s *TestStorage) Clone() osin.Storage {
	return s
}

func (s *TestStorage) Close() {
}

func (s *TestStorage) GetClient(id string) (osin.Client, error) {
	fmt.Println("AQUIII23333")
	fmt.Printf("GetClient: %s\n", id)
	if c, ok := s.clients[id]; ok {
		return c, nil
	}
	fmt.Println("AQUIII233")
	return nil, osin.ErrNotFound
}

func (s *TestStorage) SetClient(id string, client osin.Client) error {
	fmt.Printf("SetClient: %s\n", id)
	s.clients[id] = client
	return nil
}

func (s *TestStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	fmt.Printf("SaveAuthorize: %s\n", data.Code)
	s.authorize[data.Code] = data
	return nil
}

func (s *TestStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	fmt.Printf("LoadAuthorize: %s\n", code)
	if d, ok := s.authorize[code]; ok {
		return d, nil
	}
	return nil, osin.ErrNotFound
}

func (s *TestStorage) RemoveAuthorize(code string) error {
	fmt.Printf("RemoveAuthorize: %s\n", code)
	delete(s.authorize, code)
	return nil
}

func (s *TestStorage) SaveAccess(data *osin.AccessData) error {
	fmt.Printf("SaveAccess: %s\n", data.AccessToken)
	s.access[data.AccessToken] = data
	if data.RefreshToken != "" {
		s.refresh[data.RefreshToken] = data.AccessToken
	}
	return nil
}

func (s *TestStorage) LoadAccess(code string) (*osin.AccessData, error) {
	fmt.Printf("LoadAccess: %s\n", code)
	fmt.Println(s.access)
	if d, ok := s.access[code]; ok {
		return d, nil
	}
	return nil, osin.ErrNotFound
}

func (s *TestStorage) RemoveAccess(code string) error {
	fmt.Printf("RemoveAccess: %s\n", code)
	delete(s.access, code)
	return nil
}

func (s *TestStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	fmt.Printf("LoadRefresh: %s\n", code)
	if d, ok := s.refresh[code]; ok {
		return s.LoadAccess(d)
	}
	return nil, osin.ErrNotFound
}

func (s *TestStorage) RemoveRefresh(code string) error {
	fmt.Printf("RemoveRefresh: %s\n", code)
	delete(s.refresh, code)
	return nil
}


func NewApiRest(port int) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	router.Methods("PUT").Path("/este").Name("este").HandlerFunc(Index2)
	fmt.Println("Escuchando en puerto ", port)


	oauthConfig := osin.NewServerConfig()
	oauthConfig.AllowedAccessTypes = osin.AllowedAccessType{osin.PASSWORD}
	server := osin.NewServer(oauthConfig, NewTestStorage())
	// Authorization code endpoint
	router.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {

			// HANDLE LOGIN PAGE HERE

			ar.Authorized = true
			server.FinishAuthorizeRequest(resp, r, ar)
		}
		osin.OutputJSON(resp, w, r)
	})

	router.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ir := server.HandleInfoRequest(resp, r); ir != nil {
			fmt.Println("AA")
			server.FinishInfoRequest(resp, r, ir)
			fmt.Println("B")
		}
		o := osin.ResponseData{}
		o["lol"] = "lel"
		resp.Output = o
		osin.OutputJSON(resp, w, r)
	})

//authorize?response_type=code&client_id=1234&redirect_uri=http%3A%2F%2Flocalhost%3A14000%2Fappauth%2Fcode
	// Access token endpoint
	router.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()
		fmt.Println("CERO")
		if ar := server.HandleAccessRequest(resp, r); ar != nil {
			fmt.Println("UNO")
			ar.Authorized = true
			server.FinishAccessRequest(resp, r, ar)
			fmt.Println("DOS")
		}
		osin.OutputJSON(resp, w, r)
	})

	http.ListenAndServe(fmt.Sprintf(":%v", port), router)


	return router
}


func apiRestBootstrap(k *Kernel) {
	//fmt.Println("DATABASE BOOT")
	mapping := k.config.mapping
	mapping["api"] = &ApiRestConfig{}

	var baz OnKernelReady = func(k *Kernel){
		color.Green("Evento en api")
		conf := k.config.mapping["api"].(*ApiRestConfig)
		k.container.RegisterType("api", NewApiRest, conf.Port)
		k.container.MustGet("api")


	}
	k.subscribe(baz)




}
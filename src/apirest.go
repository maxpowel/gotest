package main

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/RangelReale/osin"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/RichardKnop/machinery/v1"
	"github.com/garyburd/redigo/redis"
	"github.com/ShaleApps/osinredis"
	"github.com/golang/protobuf/proto"
	"github.com/RichardKnop/machinery/v1/backends"
	"github.com/ulule/deepcopier"
	"io/ioutil"
	"gopkg.in/go-playground/validator.v9"

)

type ApiRestConfig struct {
	Port int
}


// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

type Handler struct {
	*Kernel
	H func(k *Kernel, w http.ResponseWriter, r *http.Request) error
}

func getBody(protoMessage proto.Message, r *http.Request) (error){
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return proto.Unmarshal(buf, protoMessage)

}
// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Kernel, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			errorProto := &ErrorProto{
				Code: int32(e.Status()),
				Description: e.Error(),
			}

			data, err := proto.Marshal(errorProto)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}

			w.WriteHeader(http.StatusBadRequest)
			// Raw binary data is sent
			w.Write(data)
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}


// Format task information. Used everytime your controller runs a task
func taskResponseHandler(result *backends.AsyncResult) ([]byte, error){
	state := TaskState_UNKWNOWN

	switch result.GetState().State {
	case "PENDING": state = TaskState_PENDING
	case "RECEIVED": state = TaskState_RECEIVED
	case "STARTED": state = TaskState_STARTED
	case "RETRY": state = TaskState_RETRY
	case "SUCCESS": state = TaskState_SUCCESS
	case "FAILURE": state = TaskState_FAILURE
	}


	ts := TaskStateResponse{
		State: state,
		ETA: 0,
		Uid: result.GetState().TaskUUID,
	}

	return proto.Marshal(&ts)
}

// Shortcut to launch a task
func sendTask(kernel *Kernel, task *tasks.Signature) ([]byte, error){
	server := kernel.container.MustGet("machinery").(*machinery.Server)
	asyncResult, err := server.SendTask(task)
	if err != nil {
		return nil, err
	}

	return taskResponseHandler(asyncResult)
}


// Validate input data against a model
func validate(data interface{}, validatorI interface{}) (*interface{}, error) {
	var validate *validator.Validate
	validate = validator.New()


	deepcopier.Copy(data).To(validatorI)
	err := validate.Struct(validatorI)
	//_, err := govalidator.ValidateStruct(validator)
	return &validatorI, err
}

// TODO MOVER a un sitio correcto
func NewRedisStorage() (*osinredis.Storage){
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}

	storage := osinredis.New(pool, "prefix")
	storage.CreateClient(&osin.DefaultClient{
		Id: "pepe",
		RedirectUri: "http://google.es",
		Secret: "lolazo",
	})
	return storage
}

func NewOAuthServer(k *Kernel) *osin.Server {
	oauthConfig := osin.NewServerConfig()
	oauthConfig.AllowedAccessTypes = osin.AllowedAccessType{osin.PASSWORD}
	return osin.NewServer(oauthConfig, NewRedisStorage())
}

func NewApiRest(k *Kernel, port int) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	registerControllers(k, router)


	k.container.RegisterType("oauth", NewOAuthServer, k)
	k.container.MustGet("oauth")


	// Authorization code endpoint
	/*router.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {

			// HANDLE LOGIN PAGE HERE

			ar.Authorized = true
			server.FinishAuthorizeRequest(resp, r, ar)
		}
		osin.OutputJSON(resp, w, r)
	})*/

	/*router.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
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
	})*/

//authorize?response_type=code&client_id=1234&redirect_uri=http%3A%2F%2Flocalhost%3A14000%2Fappauth%2Fcode
//curl 'http://localhost:8090/token' -d 'grant_type=password&username=pepe&password=21212&client_id=pepe' -H 'Authorization: Basic cGVwZTpsb2xhem8='
	// Access token endpoint
	router.Handle("/token", Handler{k, CheckToken})

	go http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	fmt.Println("Escuchando en puerto ", port)

	return router
}


func apiRestBootstrap(k *Kernel) {
	//fmt.Println("DATABASE BOOT")
	mapping := k.config.mapping
	mapping["api"] = &ApiRestConfig{}

	var baz OnKernelReady = func(k *Kernel){
		color.Green("Evento en api")
		conf := k.config.mapping["api"].(*ApiRestConfig)
		k.container.RegisterType("api", NewApiRest, k, conf.Port)
		k.container.MustGet("api")


	}
	k.subscribe(baz)




}

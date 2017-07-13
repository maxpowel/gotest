package main

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/RangelReale/osin"
	"time"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/RichardKnop/machinery/v1"
	"github.com/garyburd/redigo/redis"
	"github.com/ShaleApps/osinredis"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
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

			//http.Error(w, base64.StdEncoding.EncodeToString(data), e.Status())

			//http.Error(w, "", e.Status())

			w.WriteHeader(http.StatusBadRequest)
			w.Write(data)
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func GetIndex(kernel *Kernel, w http.ResponseWriter, r *http.Request) error {

	fmt.Println("EL OTRO HILO")

	// Enviar tarea
	task0 := tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	fmt.Println("Enviando task...")
	server := kernel.container.MustGet("machinery").(*machinery.Server)
	asyncResult, err := server.SendTask(&task0)

	if err != nil {
		// We return a status error here, which conveniently wraps the error
		// returned from our DB queries. We can clearly define which errors
		// are worth raising a HTTP 500 over vs. which might just be a HTTP
		// 404, 403 or 401 (as appropriate). It's also clear where our
		// handler should stop processing by returning early.
		return StatusError{500, err}
	}

	w.Write([]byte(asyncResult.Signature.UUID))

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		fmt.Println("Getting task result failed with error: %s", err.Error())
	}
	fmt.Printf(
		"%v + %v = %v\n",
		asyncResult.Signature.Args[0].Value,
		asyncResult.Signature.Args[1].Value,
		results[0].Interface(),
	)


	return nil
}

func CheckToken(kernel *Kernel, w http.ResponseWriter, r *http.Request) error {
	server := kernel.container.MustGet("oauth").(*osin.Server)
	database := kernel.container.MustGet("database").(*gorm.DB)

	resp := server.NewResponse()
	defer resp.Close()
	fmt.Println("CERO")
	var err error
	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		user := User{}
		database.Where("username = ?", username).First(&user)
		err = checkPassword(&user, password)
		ar.Authorized = err == nil
		if ar.Authorized {
			ar.UserData = user.ID
		}
		server.FinishAccessRequest(resp, r, ar)
	}
	osin.OutputJSON(resp, w, r)

	/*if err != nil {
		// We return a status error here, which conveniently wraps the error
		// returned from our DB queries. We can clearly define which errors
		// are worth raising a HTTP 500 over vs. which might just be a HTTP
		// 404, 403 or 401 (as appropriate). It's also clear where our
		// handler should stop processing by returning early.
		return StatusError{404, fmt.Errorf("User not found")}
	}*/


	return nil
}

func GetConsumption(kernel *Kernel, w http.ResponseWriter, r *http.Request) error {
	//server := kernel.container.MustGet("oauth").(*osin.Server)
	//database := kernel.container.MustGet("database").(*gorm.DB)

	//"alvaro_gg@hotmail.com"
	//"MBAR4B1"

	buf, err := ioutil.ReadAll(r.Body)
	requestData := &CredentialsProto{}
	err = proto.Unmarshal(buf, requestData)
	if err != nil {
		return StatusError{400, err}
	}

	if len(requestData.Password) == 0{
		return StatusError{400, fmt.Errorf("Password cannot be blank")}
	}

	if len(requestData.Username) == 0{
		return StatusError{400, fmt.Errorf("Username cannot be blank")}
	}
	/*username := r.Form.Get("username")
	password := r.Form.Get("password")
	username = "alvaro_gg@hotmail.com"
	password = "MBAR4B1"

	fmt.Println(username)*/
	// Enviar tarea
	task0 := tasks.Signature{
		Name: "consumption",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: requestData.Username,
			},
			{
				Type:  "string",
				Value: requestData.Password,
			},
		},
	}

	server := kernel.container.MustGet("machinery").(*machinery.Server)


	/*state, err := server.GetBackend().GetState("task_18129669-4f6e-4add-8920-5a57afda9ecf")
	fmt.Println(state.Results[0].Value)


	return nil*/


	asyncResult, err := server.SendTask(&task0)

	if err != nil {
		// We return a status error here, which conveniently wraps the error
		// returned from our DB queries. We can clearly define which errors
		// are worth raising a HTTP 500 over vs. which might just be a HTTP
		// 404, 403 or 401 (as appropriate). It's also clear where our
		// handler should stop processing by returning early.
		return StatusError{400, err}
	}


	state := TaskState_UNKWNOWN

	switch asyncResult.GetState().State {
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
		Uid: asyncResult.GetState().TaskUUID,
	}

	data, err := proto.Marshal(&ts)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
	w.Write(data)

	/*results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))

	fmt.Println(results[0])*/
	/*if err != nil {
		// We return a status error here, which conveniently wraps the error
		// returned from our DB queries. We can clearly define which errors
		// are worth raising a HTTP 500 over vs. which might just be a HTTP
		// 404, 403 or 401 (as appropriate). It's also clear where our
		// handler should stop processing by returning early.
		return StatusError{404, fmt.Errorf("User not found")}
	}*/


	return nil
}

func GetTaskState(kernel *Kernel, w http.ResponseWriter, r *http.Request) error {
	server := kernel.container.MustGet("machinery").(*machinery.Server)
	//api := kernel.container.MustGet("api").(*mux.Router)
	vars := mux.Vars(r)
	taskUid := vars["taskUid"]

	task, err := server.GetBackend().GetState(taskUid)
	//fmt.Println(task.Results[0].Value)

	if err != nil {
		return StatusError{http.StatusNotFound, fmt.Errorf("Task not found")}
	}

	state := TaskState_UNKWNOWN

	switch task.State {
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
		Uid: task.TaskUUID,
	}

	data, err := proto.Marshal(&ts)

	if err != nil {
		return StatusError{http.StatusInternalServerError, err}
	}
	w.Write(data)

	return nil
}

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
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	router.Methods("PUT").Path("/este").Name("este").HandlerFunc(Index2)
	fmt.Println("Escuchando en puerto ", port)

	router.Handle("/tarea", Handler{k, GetIndex})
	router.Handle("/consumption", Handler{k, GetConsumption})
	router.Handle("/consumption/{taskUid}", Handler{k, GetTaskState})
	router.Handle("/task/{taskUid}", Handler{k, GetTaskState})





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

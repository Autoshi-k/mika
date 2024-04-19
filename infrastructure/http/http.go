package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"something/infrastructure/serviceReply"
)

type Router struct {
	*mux.Router
}

type RequestI interface {
	Validate() bool
}

type ResponseI interface {
}

type PostHandler func(context.Context, RequestI) (ResponseI, serviceReply.Reply) // todo what should it return

func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r Router) Run(port string) {
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}

func (r Router) POST(path string, handler interface{}) {
	funcType := reflect.TypeOf(handler)
	reqType := funcType.In(1)
	req := reflect.New(reqType).Interface()

	r.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {

		ctx := context.Background()

		err := json.NewDecoder(request.Body).Decode(&req)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		} else {
			if funcType.Kind() == reflect.Func {
				resp := reflect.ValueOf(handler).Call([]reflect.Value{
					reflect.ValueOf(ctx),
					reflect.Indirect(reflect.ValueOf(req)),
				})

				var sReply serviceReply.Reply

				if len(resp) >= 1 {
					serviceReplyValue := resp[len(resp)-1]
					if !serviceReplyValue.IsNil() {
						sReply = serviceReplyValue.Interface().(serviceReply.Reply)
					}
				}

				if len(resp) > 1 {
					if sReply == nil {
						sReply = serviceReply.NewSuccessfulReply(resp[0].Interface())
					}
				}

				respByte, err := json.Marshal(sReply)
				if err != nil {
					http.Error(writer, err.Error(), http.StatusInternalServerError)
				} else if _, err = writer.Write(respByte); err != nil {
					http.Error(writer, err.Error(), http.StatusInternalServerError)
				}
			}
		}
	})
}

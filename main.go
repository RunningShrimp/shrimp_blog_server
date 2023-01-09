package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"shrimp_blog_sever/api"
	appConfig "shrimp_blog_sever/config"
	"shrimp_blog_sever/framework/config"
)

func init() {
	config.Init(".")
}

type Server struct {
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body := request.Body
	bytes, err := io.ReadAll(body)

	if err != nil {
		println(err)
	}
	if request.RequestURI == "/user/login" {
		controller := &api.UserController{}
		user := reflect.ValueOf(controller)
		loginMethod := user.MethodByName("Login")
		mType := loginMethod.Type()
		argType := mType.In(0)
		bodyData := make(map[string]any)
		err := json.Unmarshal(bytes, &bodyData)
		if err != nil {
			fmt.Println(err)
		}

		value := dataMapStruct(bodyData, argType)
		result := loginMethod.Call([]reflect.Value{value})
		fmt.Println(result[0])
		fmt.Fprintf(writer, result[0].String())
	}
}

func dataMapStruct(data map[string]any, argType reflect.Type) reflect.Value {
	val := reflect.New(argType)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < argType.NumField(); i++ {
		t := argType.Field(i)
		f := val.Field(i)
		tag := t.Tag.Get("json")
		if v, ok := data[tag]; ok {
			// 检查是否需要类型转换
			dataType := reflect.TypeOf(v)
			structType := f.Type()
			if structType == dataType {
				f.Set(reflect.ValueOf(v))
			} else {
				if dataType.ConvertibleTo(structType) {
					// 转换类型
					f.Set(reflect.ValueOf(v).Convert(structType))
				} else {
					panic(t.Name + " type mismatch")
				}
			}
		}
	}
	return val
}

func main() {
	err := http.ListenAndServe(appConfig.ServerPort(), &Server{})
	if err != nil {
		panic(err)
	}
}

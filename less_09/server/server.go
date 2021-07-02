package server

import (
	"fmt"
	"net/http"
	"../cnf"
)

func StartServer(config *cnf.Config)  {

	http.HandleFunc(config.Url, func(writer http.ResponseWriter, request *http.Request) {
		data := fmt.Sprintf("User, %s\n welcome ",config.User)
		_,err := fmt.Fprintf(writer,data)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe("localhost:"+config.Port, nil)


	if err != nil {
		fmt.Println(err)
		return
	}

}

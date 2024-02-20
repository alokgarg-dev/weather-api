package handler

import (
	"fmt"
	"net/http"
)

type Weather struct{}

func (wt *Weather) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Weather Details")
}

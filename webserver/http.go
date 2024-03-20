package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	app "github.com/ThailanTec/cli-api/internal"
)

type Server struct {
	Port string
}

func (c Server) Sum(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)

	if err != nil {
		a = 0
	}

	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
	if err != nil {
		b = 0
	}

	calc := app.NewCalc()

	calc.A = a
	calc.B = b

	w.Write([]byte(fmt.Sprintf("Resultado é %f", calc.Sum())))
}

func (s Server) Server() {
	http.HandleFunc("/", s.Sum)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}

package route

import (
	"github.com/buker/go-basic-template/internal/controller"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", controller.HealthcheckHandler)
	return mux
}

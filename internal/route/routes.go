package route

import (
	"github.com/TimeGladiator/TimeGladiator/internal/controller"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", controller.HealthcheckHandler)
	return mux
}

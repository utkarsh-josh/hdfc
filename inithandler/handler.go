package inithandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	"github.com/utkarsh-josh/hdfc/service"
	"github.com/utkarsh-josh/hdfc/spec"
)

var (
	bl     *service.BL
	logger log.Logger
)

func handleAddWebsites(w http.ResponseWriter, r *http.Request) {
	logger.Log(
		"method", "handleAddWebsites",
		"msg", "Request to add websites for status check",
		"endpoint", r.Method+" "+r.URL.String(),
		"time", time.Now(),
	)
	request, err := service.DecodeAddWebsitesRequest(r)
	if err != nil {
		logger.Log(
			"method", "handleAddWebsites",
			"msg", "Failed at transport layer",
			"err", err.Error(),
			"time", time.Now(),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := request.(*spec.WebsitesRequest)
	resp, err := bl.AddWebsites(req)
	if err != nil {
		logger.Log(
			"method", "handleAddWebsites",
			"msg", "Failed at bl layer",
			"err", err.Error(),
			"time", time.Now(),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(resp)
}

func handleGetWebsitesStatus(w http.ResponseWriter, r *http.Request) {
	logger.Log(
		"method", "handleGetWebsitesStatus",
		"msg", "Request to get websites status",
		"endpoint", r.Method+" "+r.URL.String(),
		"time", time.Now(),
	)
	request, err := service.DecodeGetWebsitesStatus(r)
	if err != nil {
		logger.Log(
			"method", "handleGetWebsitesStatus",
			"msg", "Failed at transport layer",
			"err", err.Error(),
			"time", time.Now(),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := &spec.ListWebsitesResponse{}
	req := request.(*spec.WebsitesRequest)
	resp, err = bl.GetWebsitesStatus(req)
	if err != nil {
		logger.Log(
			"method", "handleGetWebsitesStatus",
			"msg", "Failed at bl layer",
			"err", err.Error(),
			"time", time.Now(),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(resp)
}

// HandleRequests creates http handlers
func HandleRequests(logHandler log.Logger, serv *service.BL) error {
	bl = serv
	logger = logHandler
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/websites", handleAddWebsites).Methods(http.MethodPost)
	route.HandleFunc("/websites", handleGetWebsitesStatus).Methods(http.MethodGet)
	logger.Log(
		"method", "HandleRequests",
		"msg", "Initialized Routes",
		"time", time.Now(),
	)
	return http.ListenAndServe(":12000", route)
}

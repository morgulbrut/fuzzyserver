package main

import (
	"math/rand"
	"net/http"
	"time"
)

func handleDefault(w http.ResponseWriter, r *http.Request) {
	handle := rand.Intn(5)
	if handle == 0 {
		handle100(w, r)
	}
	if handle == 1 {
		handle200(w, r)
	}
	if handle == 2 {
		handle300(w, r)
	}
	if handle == 3 {
		handle400(w, r)
	}
	if handle == 4 {
		handle500(w, r)
	}
}

func handle100(w http.ResponseWriter, r *http.Request) {
	states := []int{
		http.StatusContinue,
		http.StatusSwitchingProtocols,
		http.StatusProcessing,
		http.StatusEarlyHints,
	}
	choosen := states[rand.Intn(len(states))]
	w.WriteHeader(choosen)
}

func handle200(w http.ResponseWriter, r *http.Request) {
	states := []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed,
	}

	choosen := states[rand.Intn(len(states))]
	w.WriteHeader(choosen)
}

func handle300(w http.ResponseWriter, r *http.Request) {
	states := []int{
		http.StatusMultipleChoices,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusNotModified,
		http.StatusUseProxy,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect,
	}
	choosen := states[rand.Intn(len(states))]
	w.WriteHeader(choosen)
}

func handle400(w http.ResponseWriter, r *http.Request) {
	states := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusPaymentRequired,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusMethodNotAllowed,
		http.StatusNotAcceptable,
		http.StatusProxyAuthRequired,
		http.StatusRequestTimeout,
		http.StatusConflict,
		http.StatusGone,
		http.StatusLengthRequired,
		http.StatusPreconditionFailed,
		http.StatusRequestEntityTooLarge,
		http.StatusRequestURITooLong,
		http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,
		http.StatusExpectationFailed,
		http.StatusTeapot,
		http.StatusMisdirectedRequest,
		http.StatusUnprocessableEntity,
		http.StatusLocked,
		http.StatusFailedDependency,
		http.StatusTooEarly,
		http.StatusUpgradeRequired,
		http.StatusPreconditionRequired,
		http.StatusTooManyRequests,
		http.StatusRequestHeaderFieldsTooLarge,
		http.StatusUnavailableForLegalReasons,
	}
	choosen := states[rand.Intn(len(states))]
	w.WriteHeader(choosen)
}

func handle500(w http.ResponseWriter, r *http.Request) {
	states := []int{
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusHTTPVersionNotSupported,
		http.StatusVariantAlsoNegotiates,
		http.StatusInsufficientStorage,
		http.StatusLoopDetected,
		http.StatusNotExtended,
		http.StatusNetworkAuthenticationRequired,
	}
	choosen := states[rand.Intn(len(states))]
	w.WriteHeader(choosen)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleDefault)
	mux.HandleFunc("/100", handle100)
	mux.HandleFunc("/200", handle200)
	mux.HandleFunc("/300", handle300)
	mux.HandleFunc("/400", handle400)
	mux.HandleFunc("/500", handle500)

	http.ListenAndServe(":1337", mux)
}

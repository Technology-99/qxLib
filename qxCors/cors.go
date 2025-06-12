package qxCors

import (
	"net/http"
	"strings"
)

const (
	allowOrigin      = "Access-Control-Allow-Origin"
	allOrigins       = "*"
	allowMethods     = "Access-Control-Allow-Methods"
	allowHeaders     = "Access-Control-Allow-Headers"
	allowCredentials = "Access-Control-Allow-Credentials"
	exposeHeaders    = "Access-Control-Expose-Headers"
	requestMethod    = "Access-Control-Request-Method"
	requestHeaders   = "Access-Control-Request-Headers"
	allowHeadersVal  = "Content-Type, Origin, X-CSRF-Token, Authorization, X-Request-ID, Host, X-Amz-Content-Sha256, X-Amz-Date"
	exposeHeadersVal = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers"
	methods          = "GET, POST"
	allowTrue        = "true"
	maxAgeHeader     = "Access-Control-Max-Age"
	maxAgeHeaderVal  = "86400"
	varyHeader       = "Vary"
	originHeader     = "Origin"
)

func CustomNotAllowedHandler() func(http.ResponseWriter) {
	return func(r http.ResponseWriter) {
		r.Header().Set(allowHeaders, allowHeadersVal)
	}
}

func CustomMiddleware() func(w http.Header) {
	return func(w http.Header) {
		w.Set(allowHeaders, allowHeadersVal)
	}
}

func checkAndSetHeaders(w http.ResponseWriter, r *http.Request, origins []string) {
	setVaryHeaders(w, r)

	if len(origins) == 0 {
		setHeader(w, allOrigins)
		return
	}

	origin := r.Header.Get(originHeader)
	if isOriginAllowed(origins, origin) {
		setHeader(w, origin)
	}
}

func isOriginAllowed(allows []string, origin string) bool {
	origin = strings.ToLower(origin)

	for _, allow := range allows {
		if allow == allOrigins {
			return true
		}

		allow = strings.ToLower(allow)
		if origin == allow {
			return true
		}

		if strings.HasSuffix(origin, "."+allow) {
			return true
		}
	}

	return false
}

func setHeader(w http.ResponseWriter, origin string) {
	header := w.Header()
	header.Set(allowOrigin, origin)
	header.Set(allowMethods, methods)
	header.Set(allowHeaders, allowHeadersVal)
	header.Set(exposeHeaders, exposeHeadersVal)
	if origin != allOrigins {
		header.Set(allowCredentials, allowTrue)
	}
	header.Set(maxAgeHeader, maxAgeHeaderVal)
}

func setVaryHeaders(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add(varyHeader, originHeader)
	if r.Method == http.MethodOptions {
		header.Add(varyHeader, requestMethod)
		header.Add(varyHeader, requestHeaders)
	}
}

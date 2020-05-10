package main

import (
  "net/http"
	"strconv"
	"strings"
	"sync"
	"time"
  "fmt"
  "os"
	"github.com/kjk/dailyrotate"
	"github.com/kjk/siser"
  //"github.com/felixge/httpsnoop"
)

// simplest possible server that returns url as plain text
func handleIndex(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("You've called url %s", r.URL.String())
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte(msg))
}

// Request.RemoteAddress contains port, which we want to remove i.e.:
// "[::1]:58292" => "[::1]"
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

// requestGetRemoteAddress returns ip address of the client making the request,
// taking into account http proxies
func requestGetRemoteAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		// TODO: should return first non-local address
		return parts[0]
	}
	return hdrRealIP
}

// return true if this request is a websocket request
func isWsRequest(r *http.Request) bool {
	uri := r.URL.Path
	return strings.HasPrefix(uri, "/ws/")
}

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// websocket connections won't work when wrapped
		// in RecordingResponseWriter, so just pass those through
		if isWsRequest(r) {
			h.ServeHTTP(w, r)
			return
		}

		ri := &HTTPReqInfo{
			method:    r.Method,
			url:       r.URL.String(),
			referer:   r.Header.Get("Referer"),
			userAgent: r.Header.Get("User-Agent"),
		}

		//ri.ipaddr = requestGetRemoteAddress(r)
    ri.ipaddr = r.RemoteAddr
		// this runs handler h and captures information about
		// HTTP request
		//m := httpsnoop.CaptureMetrics(h, w, r)

		//ri.code = m.Code
		//ri.size = m.Written
		//ri.duration = m.Duration
		logHTTPReq(ri)
	}
	return http.HandlerFunc(fn)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func logf(format string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Print(format)
		return
	}
	fmt.Printf(format, args...)
}

func makeDirMust(dir string) string {
	err := os.MkdirAll(dir, 0755)
	must(err)
	return dir
}
var (
	httpLogDailyFile *dailyrotate.File
	httpLogSiser     *siser.Writer
)

// HTTPReqInfo describes info about HTTP request
type HTTPReqInfo struct {
	// GET etc.
	method  string
	url     string
	referer string
	ipaddr  string
	// response code, like 200, 404
	code int
	// number of bytes of the response sent
	size int64
	// how long did it take to
	duration  time.Duration
	userAgent string
}

var (
	muLogHTTP sync.Mutex
)

// we mostly care page views. to log less we skip logging
// of urls that don't provide useful information.
// hopefully we won't regret it
func skipHTTPRequestLogging(ri *HTTPReqInfo) bool {
	// we always want to know about failures and other
	// non-200 responses
	if ri.code != 200 {
		return false
	}

	// we want to know about slow requests.
	// 100 ms threshold is somewhat arbitrary
	if ri.duration > 100*time.Millisecond {
		return false
	}

	// this is linked from every page
	if ri.url == "/favicon.png" {
		return true
	}

	if ri.url == "/favicon.ico" {
		return true
	}

	if strings.HasSuffix(ri.url, ".css") {
		return true
	}
	return false
}

func logHTTPReq(ri *HTTPReqInfo) {
	if skipHTTPRequestLogging(ri) {
		return
	}

	var rec siser.Record
	rec.Name = "httplog"
	rec.Append("method", ri.method)
	rec.Append("uri", ri.url)
	if ri.referer != "" {
		rec.Append("referer", ri.referer)
	}
	rec.Append("ipaddr", ri.ipaddr)
	rec.Append("code", strconv.Itoa(ri.code))
	rec.Append("size", strconv.FormatInt(ri.size, 10))
	durMs := ri.duration / time.Millisecond
	rec.Append("duration", strconv.FormatInt(int64(durMs), 10))
	rec.Append("ua", ri.userAgent)

	muLogHTTP.Lock()
	defer muLogHTTP.Unlock()
	_, _ = httpLogSiser.WriteRecord(&rec)
}

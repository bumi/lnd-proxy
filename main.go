package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type LNDProxy struct {
	runtime *wails.Runtime
	log     *wails.CustomLogger
	proxy   *http.Server
}

func (l *LNDProxy) StartProxy(address string, cert string, port string) (string, error) {
	l.log.Info("Starting Proxy Server")
	remoteUrl, _ := url.Parse(address)
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(cert))

	proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: certPool,
		},
	}
	proxy.Transport = transport
	proxyHander := func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
	http.HandleFunc("/", proxyHander)
	l.proxy = &http.Server{Addr: "0.0.0.0:" + port}

	go func() {
		l.log.Infof("Listening on http://0.0.0.0:%s\n", port)
		if err := l.proxy.ListenAndServe(); err != http.ErrServerClosed {
			// TODO: report back to the frontend
			l.log.Warnf("Could not listen on %s: %v", port, err)
		}
	}()

	return "running", nil // TODO: error handling
}

func (l *LNDProxy) Stop() bool {
	l.runtime.Window.Close()
	return true
}

func (l *LNDProxy) WailsInit(runtime *wails.Runtime) error {
	l.runtime = runtime
	l.log = runtime.Log.New("LNDProxy")

	return nil
}

func (l *LNDProxy) WailsShutdown() {
	if l.proxy != nil {
		if err := l.proxy.Shutdown(context.Background()); err != nil {
			l.log.Warn("Failed to stop proxy server. Not running?")
		}
	}
}

func NewLNDProxy() *LNDProxy {
	lnd := &LNDProxy{}
	return lnd
}

func main() {
	js := mewn.String("./frontend/build/main.js")
	css := mewn.String("./frontend/build/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  500,
		Height: 500,
		Title:  "Joule LND Proxy",
		JS:     js,
		CSS:    css,
	})
	app.Bind(NewLNDProxy())
	app.Run()
}

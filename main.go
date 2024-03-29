package main

import (
	"embed"
	_ "embed"
	"io/fs"
	"net"
	"net/http"
	"strconv"

	webview "github.com/webview/webview_go"
)

//go:embed assets
var f embed.FS

func main() {
	fSys, err := fs.Sub(f, "assets")
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(fSys)))

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	go http.Serve(listener, nil)

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(1000, 618, webview.HintNone)
	w.Navigate("http://localhost:" + strconv.Itoa(port))
	w.Run()
}

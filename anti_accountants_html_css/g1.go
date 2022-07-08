package main

import (
	"path"
	"runtime"

	"github.com/webview/webview"
)

var w webview.WebView

func main() {
	debug := true
	w = webview.New(debug)
	defer w.Destroy()
	w.SetTitle("WebView Example with Bindings")
	w.SetSize(800, 600, webview.HintNone)

	// Load the start page, login.html
	w.Navigate("file://" + pathToStartPage())

	// Run the app.
	w.Run()
}

func pathToStartPage() string {
	_, currentFile, _, _ := runtime.Caller(0)
	dir := path.Dir(currentFile)
	return path.Join(dir, "login.html")
}

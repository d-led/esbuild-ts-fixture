package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

const uiSrc = "ui-src"
const dist = "dist"

func main() {
	refresh()
	watcher := startWatching()
	defer watcher.Close()
	r := gin.Default()
	fs := getFS()
	r.StaticFS("/ui/", fs)
	log.Printf("http://localhost:8080/ui")
	r.Run()
}

func refresh() {
	log.Println("transpiling & copying")
	transpile()
	copyStatic()
}

func startWatching() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	crashOnError(err)

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					log.Println("modified: ", event.Name)
					refresh()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.Add(uiSrc)
	crashOnError(err)
	return watcher
}

func getFS() http.FileSystem {
	return http.FS(os.DirFS("dist"))
}

func copyStatic() {
	for _, f := range []string{
		"index.html",
		"index.css",
	} {
		text, err := os.ReadFile(filepath.Join(uiSrc, f))
		crashOnError(err)
		os.WriteFile(filepath.Join(dist, f), text, 0644)
	}
}

func transpile() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{
			filepath.Join(uiSrc, "index.ts"),
		},
		Bundle:            true,
		Outdir:            dist,
		MinifySyntax:      false,
		MinifyWhitespace:  false,
		MinifyIdentifiers: false,
		Sourcemap:         api.SourceMapInline,
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "58"},
			{Name: api.EngineFirefox, Version: "57"},
			{Name: api.EngineSafari, Version: "11"},
			{Name: api.EngineEdge, Version: "16"},
		},
		Write: true,
	})
	handleErrors(result.Errors)
}

func crashOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func handleErrors(errors []api.Message) {
	for _, msg := range errors {
		if msg.Location != nil {
			fmt.Printf(
				"%s:%v:%v: %s\n",
				msg.Location.File,
				msg.Location.Line,
				msg.Location.Column,
				msg.Text,
			)
		} else {
			fmt.Println(msg)
		}
	}

	if len(errors) > 0 {
		os.Exit(1)
	}
}

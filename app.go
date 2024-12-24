package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}
type XzInfo struct {
	Name      string `json:"name"`
	FirstUrl  string `json:"firstUrl"`
	HeadUrl   string `json:"headUrl"`
	Content   string `json:"content"`
	Title     string `json:"title"`
	NextPage  string `json:"nextPage"`
	BreakFlag string `json:"breakFlag"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(xs XzInfo) string {

	insertDb(Xiaoshuo{
		HeadUrl:   xs.HeadUrl,
		Content:   xs.Content,
		Title:     xs.Title,
		NextPage:  xs.NextPage,
		BreakFlag: xs.BreakFlag,
	})

	Download(xs)
	return fmt.Sprintf("下载完成")
}
func (a *App) GetDetail(firstUrl string) (xiaoshuo Xiaoshuo) {
	u, err := url.Parse(firstUrl)
	if err != nil {
		return xiaoshuo
	}
	// fmt.Println(u.Hostname())
	headurl := firstUrl[0:strings.Index(firstUrl, "//")+2] + u.Hostname()

	xiaoshuo = getDb(headurl)
	if len(xiaoshuo.HeadUrl) == 0 {
		xiaoshuo.HeadUrl = headurl
	}
	return xiaoshuo
}

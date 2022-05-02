package xim

import (
	"fmt"
	"github.com/Re-Ch-Love/xim/types"
	"net/url"
)

type Router struct {
	defaultPath string
	routes      map[string]types.Component
}

func NewRouter() *Router {
	return &Router{routes: map[string]types.Component{}}
}

func (r *Router) SetDefaultPath(path string) {
	r.defaultPath = path
}

// Register 注册路由
func (r *Router) Register(path string, component types.Component) {
	r.routes[path] = component
}

// RegisterDefault 注册默认路由
func (r *Router) RegisterDefault(path string, component types.Component) {
	r.routes[path] = component
	r.defaultPath = path
}

func (r *Router) Route() error {
	URL, err := url.Parse(window.Get("location").Get("href").String())
	if err != nil {
		return err
	}
	query := URL.Query()
	// 如果既没有path参数，又没有设置默认路径，则抛出错误
	if !query.Has("path") && r.defaultPath == "" {
		return fmt.Errorf("no path")
	}
	// 如果有path参数，路由；如果没有，则使用默认路径
	if query.Has("path") {
		str, err := url.QueryUnescape(query.Get("path"))
		if err != nil {
			return err
		}
		SetContent(r.routes[str])
	} else {
		query.Set("path", url.QueryEscape(r.defaultPath))
		URL.RawQuery = query.Encode()
		window.Get("history").Call("replaceState", nil, "", URL.String())
		SetContent(r.routes[r.defaultPath])
	}
	return nil
}

func JumpTo(path string) error {
	URL, err := url.Parse(window.Get("location").Get("href").String())
	if err != nil {
		return err
	}
	query := URL.Query()
	query.Set("path", url.QueryEscape(path))
	URL.RawQuery = query.Encode()
	window.Get("location").Set("href", URL.String())
	return nil
}

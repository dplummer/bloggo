package main

import (
  "github.com/codegangsta/martini"
  "github.com/dplummer/render"
  "github.com/garyburd/redigo/redis"
)

type viewContext struct {
  Articles []Article
  ArticleShow *Article
  PreviousArticle *Article
  NextArticle *Article
  Blog Blog
}

func getArticles(c redis.Conn) (content []Article, err error) {
  encoded, err := redis.Strings(c.Do("LRANGE", "bloggo:articles", "0", "-1"))
  content = make([]Article, len(encoded))
  for i, raw := range encoded {
    content[i].FromJson([]byte(raw))
  }
  return
}

func (v *viewContext) setArticle(name string) {
  for i, article := range v.Articles {
    if article.SimpleName == name {
      v.ArticleShow = &v.Articles[i]
      if i > 0 {
        v.PreviousArticle = &v.Articles[i-1]
      }
      if i < len(v.Articles) - 1 {
        v.NextArticle = &v.Articles[i+1]
      }
    }
  }
}

func (v *viewContext) setLastUpdatedAtOnBlog() {
  if len(v.Articles) > 0 {
    v.Blog.LastUpdatedAt = v.Articles[0].PublishedDate
  }
}

func NewView(c redis.Conn) (view viewContext) {
  view.Blog = getBlog(c)
  view.Articles, _ = getArticles(c)
  view.setLastUpdatedAtOnBlog()
  return view
}

func main() {
  m := martini.Classic()
  c, err := redis.Dial("tcp", ":6379")
  if err != nil {
    panic(err)
  }
  defer c.Close()

  populate(c)


  m.Use(render.Renderer(render.Options{
    Layout: "layout",
    HTMLContentType: "text/html",
  }))

  m.Get("/", func(r render.Render) {
    view := NewView(c)

    r.HTML(200, "home", view)
  })

  m.Get("/feed.xml", func(r render.Render) {
    view := NewView(c)

    r.HTML(200, "feed", view, render.HTMLOptions{
      Layout: "layouts/feed",
      HTMLContentType: "application/atom+xml",
    })
  })

  m.Get("/articles/:name", func(r render.Render, params martini.Params) {
    view := NewView(c)
    view.setArticle(params["name"])

    r.HTML(200, "articles/show", view)
  })

  m.Run()
}

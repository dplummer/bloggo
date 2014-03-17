package main

import (
  "time"
  "encoding/json"
  "fmt"
  "github.com/russross/blackfriday"
  "html/template"
)

type Article struct {
  Title         string    `json:"title"`
  PublishedDate time.Time `json:"publishedDate"`
  SimpleName    string    `json:"simpleName"`
  Published     bool      `json:"published"`
  Body          string    `json:"body"`
  AuthorName    string    `json:"authorName"`
}

func (article *Article) ParsedBody() template.HTML {
  output := blackfriday.MarkdownCommon([]byte(article.Body))
  return template.HTML(output)
}

func (article *Article) FormattedPublishedDate() string {
  return fmt.Sprintf("%d-%02d-%02d", article.PublishedDate.Year(),
    article.PublishedDate.Month(),
    article.PublishedDate.Day())
}

func (article *Article) PublishedDateXml() string {
  return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ",
    article.PublishedDate.Year(),
    article.PublishedDate.Month(),
    article.PublishedDate.Day(),
    article.PublishedDate.Hour(),
    article.PublishedDate.Minute(),
    article.PublishedDate.Second())
}

func (article *Article) ToJson() []byte {
  encoded, err := json.Marshal(article)
  if err != nil {
    panic(err)
  }
  return encoded
}

func (article *Article) FromJson(raw []byte) {
  json.Unmarshal(raw, article)
}

type Articles []Article

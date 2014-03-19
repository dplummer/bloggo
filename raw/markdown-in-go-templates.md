While building this blog, I ran into a problem. I was storing the blog posts in
markdown format. Using the
[blackfriday](https://github.com/russross/blackfriday) markdown processor, I
was getting back raw HTML:

``` go
func (article *Article) ParsedBody() string {
  output := blackfriday.MarkdownCommon([]byte(article.Body))
  return string(output)
}
```

Then in my template I'd try to [render the `ParsedBody`](https://github.com/dplummer/bloggo/blob/cc8fa8a1df51ff5a53d203bb7e3def8d03c6c25a/templates/articles/show.tmpl#L5):

``` html
{{with .ArticleShow}}
  <div class="blog-post">
    <!-- title, etc ... -->
    {{.ParsedBody}}
  </div>
{{end}}
```

But that escapes the body:

``` html
  <div class="blog-post">
    <!-- title, etc ... -->
    &gt;p&lt;Hello world!&gt;/p&lt;
  </div>
```

The answer is to not return a `string`, but to [return a `template.HTML` type](https://github.com/dplummer/bloggo/blob/cc8fa8a1df51ff5a53d203bb7e3def8d03c6c25a/articles.go#L20):

``` go
func (article *Article) ParsedBody() template.HTML {
  output := blackfriday.MarkdownCommon([]byte(article.Body))
  return template.HTML(output)
}
```

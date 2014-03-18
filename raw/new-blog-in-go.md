I've been playing with **go** for a few months now. I recently needed to do a large
ETL for work, copying price data from about 1000 mysql databases and loading it
into a single database with a different schema. Leveraging our existing rails
codebase wasn't going to work, the migration of 15 million records was going to
take on the order of weeks. So in a weekend I wrote a command line tool in
**go**. That was a lot of fun.

Then about a week ago I got an email out of the blue from a CTO at a startup
asking about one of my blog posts. I thought, "*man, I should really be posting
to my blog more*". Plus, I wanted to write a webapp in **go**. So here it is!

The code is [available on github](https://github.com/dplummer/bloggo). I have
several goals in mind:

* Live markdown editor. Right now all the content is actually compiled into the
  binary and loaded into redis on startup. Which makes using redis pretty dumb.
  But my goal is to just store everything in redis, and allow myself to draft
  and publish posts from the web (like a real blog!)
* Comments. Not sure how I want this, but I've already had people email me
  comments on posts, and I'd like to share that. So I need to add something
  here. I'm thinking of either disqus or login with github.
* Versions on my posts. Like git or a wiki. Could be cool to mark comments as
  outdated after updating it based on feedback. Like github issues.

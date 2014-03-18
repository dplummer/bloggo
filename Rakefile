timestamp = Time.now.to_i

task :deploy => [:build, :upload, :start_app]

task :build do
  `GOOS=linux GOARCH=386 go build`
  puts "Built bloggo"
end

task :upload do
  `scp bloggo donaldplummer.com:bloggo/bloggo-#{timestamp}`
  puts "Uploaded bloggo to bloggo-#{timestamp}"
  `scp -r raw templates public donaldplummer.com:bloggo/`
  puts "Uploaded raw, templates, and assets"
end

task :start_app do
  `ssh donaldplummer.com cd bloggo && ln -sf bloggo-#{timestamp} bloggo`
  puts "Symlinked bloggo-#{timestamp}"
  `ssh donaldplummer.com sudo service bloggo restart`
  puts "Service restarted"
end

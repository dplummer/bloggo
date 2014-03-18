timestamp = Time.now.to_i

task :deploy => [:build, :upload_binary, :upload_assets, :symlink_binary, :start_app]
task :deploy_assets => [:upload_assets, :start_app]

task :build do
  `GOOS=linux GOARCH=386 go build`
  puts "Built bloggo"
end

task :upload_binary do
  `scp bloggo donaldplummer.com:bloggo/bloggo-#{timestamp}`
  puts "Uploaded bloggo to bloggo-#{timestamp}"
end

task :upload_assets do
  `scp -r raw templates public donaldplummer.com:bloggo/`
  puts "Uploaded raw, templates, and assets"
end

task :symlink_binary do
  `ssh donaldplummer.com cd bloggo && ln -sf bloggo-#{timestamp} bloggo`
  puts "Symlinked bloggo-#{timestamp}"
end

task :start_app do
  `ssh donaldplummer.com sudo service bloggo restart`
  puts "Service restarted"
end

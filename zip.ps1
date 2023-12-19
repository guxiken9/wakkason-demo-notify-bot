$env:GOOS = "linux"; $env:GOARCH = "amd64"; go build -o main
Compress-Archive -Path ".\main" -DestinationPath "./dist/notify.zip" -Force
Remove-Item ./main 

echo Building Windows edition
env GOOS=windows GOARCH=amd64 go build -o builds/windows/ssd.exe

echo Building MacOS edition
env GOOS=darwin GOARCH=amd64 go build -o builds/macos/ssd

echo Building Linux edition
env GOOS=linux GOARCH=amd64 go build -o builds/linux/ssd

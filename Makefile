build-windows: 
	go build -o ./build/wineservice.exe ./cmd/wineservice

windows: build-windows
	./build/wineservice.exe
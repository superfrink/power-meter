main: power-meter.go

	go build -o power-meter.out power-meter.go
	GOOS=linux GOARCH=arm GOARM=7 go build -o power-meter.pi power-meter.go

clean:
	rm power-meter.out power-meter.pi

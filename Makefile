run:
	go run ./src

exeicon:
	rsrc -ico .\assets\sftpguard.ico -o rsrc.syso

build:

	go build -o sftpguard.exe ./src
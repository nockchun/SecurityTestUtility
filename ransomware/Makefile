BUILD=go build -ldflags="-w -s"

default: build

build: server_linux agent_windows

server_linux:
	@echo "Building Test Ransomware Linux Server ..."
	cd server && GOOS=linux GOARCH=amd64 $(BUILD) -o ../bin/server_linux_64
	cd server && GOOS=linux GOARCH=386   $(BUILD) -o ../bin/server_linux_32

server_windows:
	@echo "Building Test Ransomware Windows Server ..."
	cd server && GOOS=windows GOARCH=amd64 $(BUILD) -o ../bin/server_windows_64.exe
	cd server && GOOS=windows GOARCH=386   $(BUILD) -o ../bin/server_windows_32.exe

agent_linux:
	@echo "Building Test Ransomware Linux Agent ..."
	cd encrypt && GOOS=linux GOARCH=amd64 $(BUILD) -o ../bin/agentEnc_linux_64
	cd encrypt && GOOS=linux GOARCH=386   $(BUILD) -o ../bin/agentEnc_linux_32
	cd decrypt && GOOS=linux GOARCH=amd64 $(BUILD) -o ../bin/agentDec_linux_64
	cd decrypt && GOOS=linux GOARCH=386   $(BUILD) -o ../bin/agentDec_linux_32

agent_windows:
	@echo "Building Test Ransomware Windows Agent ..."
	cd encrypt && GOOS=windows GOARCH=amd64 $(BUILD) -o ../bin/agentEnc_windows_64.exe
	cd encrypt && GOOS=windows GOARCH=386   $(BUILD) -o ../bin/agentEnc_windows_32.exe
	cd decrypt && GOOS=windows GOARCH=amd64 $(BUILD) -o ../bin/agentDec_windows_64.exe
	cd decrypt && GOOS=windows GOARCH=386   $(BUILD) -o ../bin/agentDec_windows_32.exe

clean:
	@rm -rf bin/
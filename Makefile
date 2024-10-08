SERVER_BINARY := server/server
CLIENT_BINARY := client/client

all: $(SERVER_BINARY) $(CLIENT_BINARY)
	@echo "Done :)"

$(SERVER_BINARY): server/main.go
	@echo "Building server..."
	@go build -o $(SERVER_BINARY) server/main.go

$(CLIENT_BINARY): client/main.go
	@echo "Building client..."
	@go build -o $(CLIENT_BINARY) client/main.go

clean:
	@echo "Cleaning up..."
	@rm -f $(SERVER_BINARY) $(CLIENT_BINARY)

server: $(SERVER_BINARY)

client: $(CLIENT_BINARY)

.PHONY: all clean server client

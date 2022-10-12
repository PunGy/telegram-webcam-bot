SOURCE 		:= ./cmd/telegram-webcam-bot/main.go
EXEC 		:= bot
BINARIES	:= ./

build:
	go build -o ./$(EXEC) $(SOURCE)

install:
	cp $(EXEC) /usr/local/bin/$(EXEC)
uninstall:
	rm -f $(EXEC) /usr/local/bin/$(EXEC)

run:
	./$(EXEC)

clean:
	rm -f $(EXEC)
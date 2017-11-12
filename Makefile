all:
	go build -o ranarian main.go

install-deps:
	go get github.com/bwmarrin/discordgo

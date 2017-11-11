package main

import (
  "flag"
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "io"
  "encoding/json"

  "github.com/bwmarrin/discordgo"
)

var Token string

func main() {
  // Create a new Discord session using the provided bot token.
  dg, err := discordgo.New("Bot " + Token)
  if err != nil {
    fmt.Println("Error creating Discord session:", err)
    return 1
  }

  // Register the messageCreate func as a callback for MessageCreate events.
  dg.AddHandler(messageCreate)

  // Open a websocket connection to Discord and begin listening.
  err = dg.Open()
  if err != nil {
    fmt.Println("Error opening connection: ", err)
    return 1
  }

  // Wait here until CTRL-C or other term signal is received.
  fmt.Println("Ranarian online!")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  // Ignore all messages created by the bot itself
  // This isn't required in this specific example but it's a good practice.
  if m.Author.ID == s.State.User.ID {
    return
  }
  // If the message is "ping" reply with "Pong!"
  if m.Content == "ping" {
    s.ChannelMessageSend(m.ChannelID, "Pong!")
  }

  // If the message is "pong" reply with "Ping!"
  if m.Content == "pong" {
    s.ChannelMessageSend(m.ChannelID, "Ping!")
  }
}

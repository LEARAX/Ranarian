package main

import (
  "fmt"
  "os"
  "log"
  "os/signal"
  "syscall"
  "io/ioutil"
  "encoding/json"

  "github.com/bwmarrin/discordgo"
)

var config struct {
  Token string
}


func main() {
  file, err := ioutil.ReadFile("./secrets.json")
  if err != nil {
    log.Fatal("Error reading config file: ", err)
  }

  err := json.Unmarshal(file, &config)
  if err != nil {
    log.Fatal("Error parsing config: ", err)
  }

  dg, err := discordgo.New("Bot " + config.Token)
  if err != nil {
    log.Fatal("Error creating Discord session: ", err)
  }

  dg.AddHandler(messageCreated)

  err = dg.Open()
  if err != nil {
    log.Fatal("Error opening connection: ", err)
  }

  fmt.Println("CONNECTION ESTABLISHED")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreated(session *discordgo.Session, msg *discordgo.MessageCreate) {
  if msg.Author.ID == session.State.User.ID {
    return
  }

  if msg.Content == "ping" {
    session.ChannelMessageSend(msg.ChannelID, "Pong!")
  }

  if msg.Content == "pong" {
    session.ChannelMessageSend(msg.ChannelID, "Ping!")
  }
}

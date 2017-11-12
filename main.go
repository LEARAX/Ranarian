package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "strings"
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
    print("Error reading config file: ", err)
    os.Exit(1)
  }

  json.Unmarshal(file, &config)

  dg, err := discordgo.New("Bot " + config.Token)
  if err != nil {
    print("Error creating Discord session: ", err)
    os.Exit(1)
  }

  dg.AddHandler(messageCreated)

  err = dg.Open()
  if err != nil {
    print("Error opening connection: ", err)
    os.Exit(1)
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

  if strings.HasPrefix(msg.Content, ">") {
    req := msg.Content[1:]
    fmt.Println("Command detected: " + req)

    if strings.HasPrefix(req, "mute") {
      fmt.Println("Muting user: ")
    }
  }
}

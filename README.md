# Ranarian
Ranarian is a discordgo bot. Who knows what it'll do?

## Functions
Functions/properties of this bot will be listed below as they are added.

## Executing
To build Ranarian, run the following in your favorite shell:
```sh
go get github.com/bwmarrin/discordgo
go build -o ranarian main.go
```

Then create a file in the same directory named `secrets.json`, with the following structure:
```json
{
  "token": "YOURTOKENHERE"
}```

Replace `YOURTOKENHERE` with a token from the [discord developer page](https://discordapp.com/developers/applications/me).

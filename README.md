# Teleplasm
Teleplasm is a Discord bot written in Rust. Well, it will be.

## Functions
Functions/properties of this bot will be listed below as they are added.

## Executing
To build Ranarian, run the following in your favorite shell:
```sh
cargo build --release
```
A compiled executable will be located at `./target/release/teleplasm`. Move it wherever you like.

Then create a file in the same directory named `secrets.json`, with the following structure:
```json
{
  "token": "YOURTOKENHERE"
}
```

Replace `YOURTOKENHERE` with a token from the [Discord developer page](https://discordapp.com/developers/applications/me).

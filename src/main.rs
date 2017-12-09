extern crate serenity;

use serenity::prelude::*;
use serenity::model::*;
use std::env;

struct Handler;

impl EventHandler for Handler {
    fn on_message(&self, _: Context, msg: Message) {
        if msg.content == "!messageme" {
            if let Err(why) = msg.author.dm(|m| m.content("Hello!")) {
                println!("Error when direct messaging user: {:?}", why);
            }
        }
    }

    fn on_ready(&self, _: Context, ready: Ready) {
        println!("{} is connected!", ready.user.name);
    }
}

fn main() {
    // Configure the client with your Discord bot token in the environment.
    let token = env::var("DISCORD_TOKEN")
        .expect("Expected a token in the environment");
    let mut client = Client::new(&token, Handler);

    if let Err(why) = client.start() {
        println!("Client error: {:?}", why);
    }
}

extern crate serenity;

use serenity::prelude::*;
use serenity::model::*;
use serenity::voice;
use std::env;

struct Handler {
    target_voice_channel: (GuildId, ChannelId)
}

impl EventHandler for Handler {
    fn on_message(&self, ctx: Context, msg: Message) {
        if msg.content == "!messageme" {
            if let Err(why) = msg.author.dm(|m| m.content("Hello!")) {
                println!("Error when direct messaging user: {:?}", why);
            }
        }
        if &msg.content[..5] == "!play" {
            let url = &msg.content[6..];
            println!("playing {}", url);
            match ctx.shard.lock().manager.get(self.target_voice_channel.0) {
                Some(handler) => {
                    match voice::ytdl(url) {
                        Ok(source) => {
                            println!("ok");
                            handler.play(source);
                        },
                        Err(err) => {
                            println!("Error with source: {:?}", err);
                        }
                    }
                },
                None => {
                    println!("Not in a voice channel, can't play audio");
                }
            }
        }
    }

    fn on_ready(&self, ctx: Context, ready: Ready) {
        println!("{} is connected!", ready.user.name);

        let mut shard = ctx.shard.lock();
        shard.manager.join(self.target_voice_channel.0, self.target_voice_channel.1);
    }
}

fn main() {
    // Configure the client with your Discord bot token in the environment.
    let token = env::var("DISCORD_TOKEN")
        .expect("Expected a token in the environment");
    let target_voice_guild = env::var("TARGET_VOICE_GUILD")
        .expect("Expected a target voice guild ID");
    let target_voice_channel = env::var("TARGET_VOICE_CHANNEL")
        .expect("Expected a target voice channel ID");
    let mut client = Client::new(&token, Handler {
        target_voice_channel: (
                                  GuildId(target_voice_guild.parse::<u64>().expect("Expected guild ID to be numeric")),
                                  ChannelId(target_voice_channel.parse::<u64>().expect("Expected channel ID to be numeric"))
                              )
    });

    if let Err(why) = client.start() {
        println!("Client error: {:?}", why);
    }
}

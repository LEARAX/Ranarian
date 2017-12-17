extern crate serenity;

use serenity::prelude::*;
use serenity::model::*;
use serenity::voice;
use std::env;

struct Handler {
    target_voice_channel: (u64, u64)
}

impl EventHandler for Handler {
    fn on_message(&self, _: Context, msg: Message) {
        if msg.content == "!messageme" {
            if let Err(why) = msg.author.dm(|m| m.content("Hello!")) {
                println!("Error when direct messaging user: {:?}", why);
            }
        }
    }

    fn on_ready(&self, ctx: Context, ready: Ready) {
        println!("{} is connected!", ready.user.name);

        let mut shard = ctx.shard.lock();
        shard.manager.join(GuildId(self.target_voice_channel.0), ChannelId(self.target_voice_channel.1));
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
                                  target_voice_guild.parse::<u64>().expect("Expected guild ID to be numeric"),
                                  target_voice_channel.parse::<u64>().expect("Expected channel ID to be numeric")
                              )
    });

    if let Err(why) = client.start() {
        println!("Client error: {:?}", why);
    }
}

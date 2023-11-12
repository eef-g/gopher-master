package bot

import (
	"os"
	"os/signal"
	"log"
	"github.com/bwmarrin/discordgo"
	"eef.gocord/bot/commands"
)

// Don't need to make a struct since we can just get the bot info from the main.go file
var (
	Token string
)



func Start() {
	// Start new discord session
	log.Println("Starting bot...")
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}

	
	// Edit the session details before opening
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		s.UpdateGameStatus(0, "running on Go!");
		log.Printf("Bot is running as %v#%v\n", r.User.Username, r.User.Discriminator);
	})
	// Open session
	discord.Open();
	defer discord.Close();

	// Add commands -- We do this AFTER we open the session so that way we can get the user id
	// NOTE: If this is done before the session is opened, the bot will not be able to register the commands and results in a memory leak :(
	commands.AddCommands(discord);
	log.Printf("Added commands\n");

	// Wait for a CTRL-C
	log.Println("Bot is now running. Press CTRL-C to exit.")
	c := make(chan os.Signal, 1);

	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Removing all commands");

	// Remove all commands
	commands.RemoveCommands(discord);
}
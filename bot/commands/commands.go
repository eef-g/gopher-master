package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func RemoveCommands(s *discordgo.Session) {

}


func AddCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name: 	  "ping",
			Description: "Ping pong!",
		},
		{
			Name: 	  "pong",
			Description: "Pong ping!",
		},
		// Add more commands here
	}


	// Register the functions in the commands
	command_handlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		commands[0].Name: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			pingCommand(s, i);
		},

		commands[1].Name: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			pongCommand(s, i);
		},
		// Add more commands here
	}

	
	// Register the command handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := command_handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// Register the commands
	registered_commands := make([]*discordgo.ApplicationCommand, len(commands));
	for i, v := range commands {
		fmt.Println("Adding command: ", v.Name);
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v);
		if err != nil {
			fmt.Println("Error creating command: ", err);
		}
		registered_commands[i] = cmd;
	}
}


/*
* Command handlers -- Put command functions here
*/

func pingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}
	s.InteractionRespond(i.Interaction, &response)
}

func pongCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ping!",
		},
	}
	s.InteractionRespond(i.Interaction, &response)
}
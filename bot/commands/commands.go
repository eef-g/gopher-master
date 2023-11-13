package commands

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func RemoveCommands(s *discordgo.Session) {
}

func AddCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "d20",
			Description: "Roll a d20.",
		},
		{
			Name:        "d6",
			Description: "Roll a d6.",
		},
		{
			Name:        "d4",
			Description: "Roll a d4.",
		},
		// Add more commands here
	}

	// Register the functions in the commands
	command_handlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		commands[0].Name: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			d20Command(s, i)
		},
		commands[1].Name: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			d6Command(s, i)
		},
		commands[2].Name: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			d4Command(s, i)
		},
		// Add more commands here
	}

	// Register the command handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := command_handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// Check to make sure the session, state, and user are not nil
	if s == nil {
		fmt.Println("The session instance is nil")
		return
	}
	if s.State == nil {
		fmt.Println("The session state is nil")
		return
	}
	if s.State.User == nil {
		fmt.Println("The session user is nil")
		return
	}

	// Register the commands
	registered_commands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		fmt.Println("Adding command: ", v.Name)
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			fmt.Println("Error creating command: ", err)
		}
		registered_commands[i] = cmd
	}
}

/*
* Command handlers -- Put command functions here
 */

func d20Command(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Command Logic
	roll := rand.Intn(20)
	output := fmt.Sprintf("You rolled a *%d*!", roll)

	// Command Output
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: output,
		},
	}
	s.InteractionRespond(i.Interaction, &response)
}

func d6Command(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Command Logic
	roll := rand.Intn(6)
	output := fmt.Sprintf("You rolled a *%d*!", roll)

	// Command Output
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: output,
		},
	}
	s.InteractionRespond(i.Interaction, &response)
}

func d4Command(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Command logic
	roll := rand.Intn(4)
	output := fmt.Sprintf("You rolled a *%d*!", roll)

	// Command output
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: output,
		},
	}
	s.InteractionRespond(i.Interaction, &response)
}

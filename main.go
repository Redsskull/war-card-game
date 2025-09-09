package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialize the game
	player1, cpu := StartGame()

	myApp := app.New()
	myWindow := myApp.NewWindow("War Card Game")
	myWindow.Resize(fyne.NewSize(800, 600))

	// Game status display
	gameStatus := widget.NewLabel("Game ready! Click 'Play Round' to start.")
	// Remove the wrapping line entirely - not needed for basic functionality TODO: investigate how to properly wrap text in Fyne

	// Player info
	playerInfo := widget.NewLabel("")
	updatePlayerInfo := func() {
		playerInfo.SetText(fmt.Sprintf("%s | %s", player1, cpu))
	}
	updatePlayerInfo()

	// Declare playButton first, then define its function
	var playButton *widget.Button
	playButton = widget.NewButton("Play Round", func() {
		result := PlayRound(player1, cpu)
		gameStatus.SetText(result)
		updatePlayerInfo()

		// Check if game is over
		if !player1.HasCards() {
			gameStatus.SetText(result + "\n\nGAME OVER - CPU WINS!")
			playButton.Disable()
		} else if !cpu.HasCards() {
			gameStatus.SetText(result + "\n\nGAME OVER - PLAYER WINS!")
			playButton.Disable()
		}
	})

	content := container.NewVBox(
		widget.NewLabel("War Card Game"),
		playerInfo,
		playButton,
		gameStatus,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

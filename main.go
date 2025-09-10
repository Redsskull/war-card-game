package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialize the game
	player1, cpu := StartGame()

	myApp := app.New()
	myWindow := myApp.NewWindow("War Card Game")
	myWindow.Resize(fyne.NewSize(800, 600))

	// Score displays
	playerScore := widget.NewLabel(fmt.Sprintf("Your cards: %d", len(player1.Cards)))
	cpuScore := widget.NewLabel(fmt.Sprintf("CPU cards: %d", len(cpu.Cards)))

	// 🃏 CARD IMAGES - Start with placeholder
	playerCardImage := canvas.NewImageFromFile("Cards/card_joker.png")
	playerCardImage.SetMinSize(fyne.NewSize(80, 110))
	playerCardImage.FillMode = canvas.ImageFillContain

	cpuCardImage := canvas.NewImageFromFile("Cards/card_joker.png")
	cpuCardImage.SetMinSize(fyne.NewSize(80, 110))
	cpuCardImage.FillMode = canvas.ImageFillContain

	updateScores := func() {
		playerScore.SetText(fmt.Sprintf("Your cards: %d", len(player1.Cards)))
		cpuScore.SetText(fmt.Sprintf("CPU cards: %d", len(cpu.Cards)))
	}

	// Game result
	gameResult := widget.NewLabel("Click 'Start Game' to begin!")
	gameResult.Alignment = fyne.TextAlignCenter

	// Play button
	var playButton *widget.Button
	playButton = widget.NewButton("Start Game", func() {
		if !player1.HasCards() || !cpu.HasCards() {
			return
		}
		if playButton.Text == "Start Game" {
			playButton.SetText("Play Round")
		}

		// Get the cards and result
		playerCard, cpuCard, result := PlayRound(player1, cpu)

		// 🎯 UPDATE THE IMAGES WITH ACTUAL CARDS!
		playerCardImage.File = playerCard.GetImageFilename()
		playerCardImage.Refresh()

		cpuCardImage.File = cpuCard.GetImageFilename()
		cpuCardImage.Refresh()

		// Show result
		gameResult.SetText(result)
		updateScores()

		// Check game over
		if !player1.HasCards() {
			gameResult.SetText("GAME OVER!\nCPU WINS! 🤖")
			playButton.SetText("Game Over")
			playButton.Disable()
		} else if !cpu.HasCards() {
			gameResult.SetText("GAME OVER!\nYOU WIN! 🏆")
			playButton.SetText("Game Over")
			playButton.Disable()
		}
	})

	// Layout
	topArea := container.NewVBox(
		container.NewCenter(widget.NewLabel("🃏 War Card Game")),
		widget.NewSeparator(),
		container.NewCenter(cpuScore))

	middleArea := container.NewCenter(
		container.NewVBox(
			widget.NewSeparator(),
			container.NewHBox(playerCardImage, widget.NewLabel(" VS "), cpuCardImage),
			widget.NewSeparator(),
			gameResult))

	bottomArea := container.NewCenter(
		container.NewHBox(playerScore, playButton))

	content := container.NewBorder(topArea, bottomArea, nil, nil, middleArea)

	// Background
	background := canvas.NewRectangle(color.RGBA{70, 130, 180, 255})
	contentWithBackground := container.NewStack(background, container.NewPadded(content))
	myWindow.SetContent(contentWithBackground)

	myWindow.ShowAndRun()
}

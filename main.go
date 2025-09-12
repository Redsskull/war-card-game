package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialize the game
	player1, cpu, setupMessages := StartGame()

	myApp := app.New()
	myWindow := myApp.NewWindow("War Card Game")
	myWindow.Resize(fyne.NewSize(800, 600))

	// NOTIFICATION SYSTEM for setup messages
	notificationLabel := widget.NewLabel("")
	notificationLabel.Alignment = fyne.TextAlignCenter
	notificationLabel.TextStyle.Bold = true
	showNotification := func(message string) {
		notificationLabel.SetText(message)
		notificationLabel.Show()
		notificationLabel.Refresh()

		// Don't use a goroutine here - schedule the hide with fyne.Do instead
		go func() {
			time.Sleep(2 * time.Second)
			fyne.Do(func() {
				notificationLabel.Hide()
			})
		}()
	}

	// CARD COUNT LABELS
	cpuCardCount := widget.NewLabel(fmt.Sprintf("%d", len(cpu.Cards)))
	cpuCardCount.Alignment = fyne.TextAlignCenter
	cpuCardCount.TextStyle.Bold = true

	playerCardCount := widget.NewLabel(fmt.Sprintf("%d", len(player1.Cards)))
	playerCardCount.Alignment = fyne.TextAlignCenter
	playerCardCount.TextStyle.Bold = true

	// PLAYED CARD IMAGES (hidden initially)
	playerCardImage := canvas.NewImageFromFile("Cards/card_joker.png")
	playerCardImage.SetMinSize(fyne.NewSize(80, 110))
	playerCardImage.FillMode = canvas.ImageFillContain
	playerCardImage.Hide()

	cpuCardImage := canvas.NewImageFromFile("Cards/card_joker.png")
	cpuCardImage.SetMinSize(fyne.NewSize(80, 110))
	cpuCardImage.FillMode = canvas.ImageFillContain
	cpuCardImage.Hide()

	// HAND CARD BACKS
	playerHandImage := canvas.NewImageFromFile("Cards/card_back_suits.png")
	playerHandImage.SetMinSize(fyne.NewSize(60, 80))
	playerHandImage.FillMode = canvas.ImageFillContain

	cpuHandImage := canvas.NewImageFromFile("Cards/card_back_suits_dark.png")
	cpuHandImage.SetMinSize(fyne.NewSize(60, 80))
	cpuHandImage.FillMode = canvas.ImageFillContain

	// TITLE
	gameTitle := widget.NewLabel("🃏 War Card Game")
	gameTitle.TextStyle.Bold = true

	// GAME RESULT (hidden initially)
	gameResult := widget.NewLabel("")
	gameResult.Alignment = fyne.TextAlignCenter
	gameResult.Hide()

	// SIMPLE ENHANCED INITIAL DISPLAY - BIGGER!
	initialDisplay := container.NewCenter(
		container.NewVBox(
			notificationLabel,
			container.NewCenter(
				func() *canvas.Text {
					sword := canvas.NewText("⚔️", color.White)
					sword.TextSize = 48
					sword.TextStyle.Bold = true
					sword.Alignment = fyne.TextAlignCenter
					return sword
				}()),

			container.NewCenter(
				container.NewHBox(
					widget.NewLabel("👤 PLAYER"),
					func() *canvas.Text {
						vs := canvas.NewText("🆚", color.White)
						vs.TextSize = 36
						vs.TextStyle.Bold = true
						vs.Alignment = fyne.TextAlignCenter
						return vs
					}(),
					widget.NewLabel("🤖 CPU"))),

			container.NewCenter(
				func() *widget.Label {
					ready := widget.NewLabel("Ready for Battle!")
					ready.TextStyle.Bold = true
					return ready
				}())))

	// BATTLE AREA (hidden initially)
	battleArea := container.NewCenter(
		container.NewHBox(playerCardImage, widget.NewLabel("  VS  "), cpuCardImage))
	battleArea.Hide()

	// UPDATE SCORES FUNCTION
	updateScores := func() {
		cpuCardCount.SetText(fmt.Sprintf("%d", len(cpu.Cards)))
		playerCardCount.SetText(fmt.Sprintf("%d", len(player1.Cards)))
	}

	// Add visual hint
	hintLabel := widget.NewLabel("👆 Click card to play!")
	hintLabel.Alignment = fyne.TextAlignCenter
	hintLabel.Hide()

	// PLAY ROUND LOGIC
	executeRound := func() {
		if !player1.HasCards() || !cpu.HasCards() {
			return
		}

		playerCard, cpuCard, result := PlayRound(player1, cpu)

		playerCardImage.File = playerCard.GetImageFilename()
		playerCardImage.Show()
		playerCardImage.Refresh()

		cpuCardImage.File = cpuCard.GetImageFilename()
		cpuCardImage.Show()
		cpuCardImage.Refresh()

		gameResult.SetText(result)
		gameResult.Show()
		updateScores()

		if !player1.HasCards() {
			gameResult.SetText("GAME OVER!\nCPU WINS! 🤖")
			hintLabel.Hide()
		} else if !cpu.HasCards() {
			gameResult.SetText("GAME OVER!\nYOU WIN! 🏆")
			hintLabel.Hide()
		}
	}

	// START GAME BUTTON
	var playButton *widget.Button
	playButton = widget.NewButton("Start Game", func() {
		gameTitle.Hide()
		playButton.Hide()
		initialDisplay.Hide()
		battleArea.Show()
		hintLabel.Show()
	})

	// CREATE CLICKABLE CARD using effects.go
	clickablePlayerCard := NewClickableCard(playerHandImage, playerCardCount, func() {
		if !gameTitle.Visible() && battleArea.Visible() {
			executeRound()
		}
	})

	// LAYOUT
	topArea := container.NewCenter(
		container.NewVBox(
			gameTitle,
			widget.NewSeparator(),
			container.NewStack(cpuHandImage, cpuCardCount)))

	middleArea := container.NewCenter(
		container.NewVBox(
			container.NewStack(initialDisplay, battleArea),
			playButton,
			gameResult))

	bottomArea := container.NewCenter(
		container.NewVBox(
			hintLabel,
			clickablePlayerCard)) // Use the clickable card widget from effects.go!

	content := container.NewBorder(topArea, bottomArea, nil, nil, middleArea)

	// FINAL DISPLAY
	background := canvas.NewRectangle(color.RGBA{70, 130, 180, 255})
	contentWithBackground := container.NewStack(background, container.NewPadded(content))
	myWindow.SetContent(contentWithBackground)

	// Start the notification display after window is ready
	go func() {
		// Wait a moment for window to be ready
		time.Sleep(1 * time.Second) // Increased wait time

		for _, msg := range setupMessages {
			// Capture the message in a local variable to avoid closure issues
			message := msg
			fyne.Do(func() {
				showNotification(message)
			})
			time.Sleep(3 * time.Second) // Increased delay between messages
		}
	}()

	myWindow.ShowAndRun()

}

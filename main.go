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
	myWindow.Resize(fyne.NewSize(1920, 1080))

	// NOTIFICATION SYSTEM for setup messages
	notificationText := canvas.NewText("", color.White)
	notificationText.Alignment = fyne.TextAlignCenter
	notificationText.TextStyle.Bold = true
	notificationText.TextSize = 18
	notificationText.Hide()

	showNotification := func(message string) {
		notificationText.Text = message
		notificationText.Show()
		notificationText.Refresh()

		go func() {
			time.Sleep(2 * time.Second)
			fyne.Do(func() {
				notificationText.Hide()
			})
		}()
	}

	// CARD COUNT LABELS
	cpuCardCount := canvas.NewText(fmt.Sprintf("%d", len(cpu.Cards)), color.White)
	cpuCardCount.Alignment = fyne.TextAlignCenter
	cpuCardCount.TextStyle.Bold = true
	cpuCardCount.TextSize = 20

	playerCardCount := canvas.NewText(fmt.Sprintf("%d", len(player1.Cards)), color.White)
	playerCardCount.Alignment = fyne.TextAlignCenter
	playerCardCount.TextStyle.Bold = true
	playerCardCount.TextSize = 20

	// PLAYED CARD IMAGES (hidden initially) - MUCH LARGER
	playerCardImage := canvas.NewImageFromFile("Cards/card_joker.png")
	playerCardImage.SetMinSize(fyne.NewSize(250, 350))
	playerCardImage.FillMode = canvas.ImageFillContain
	playerCardImage.Hide()

	cpuCardImage := canvas.NewImageFromFile("Cards/card_joker.png")
	cpuCardImage.SetMinSize(fyne.NewSize(250, 350))
	cpuCardImage.FillMode = canvas.ImageFillContain
	cpuCardImage.Hide()

	// HAND CARD BACKS - LARGER
	playerHandImage := canvas.NewImageFromFile("Cards/card_back_ suits_blue.png")
	playerHandImage.SetMinSize(fyne.NewSize(180, 300))
	playerHandImage.FillMode = canvas.ImageFillContain

	cpuHandImage := canvas.NewImageFromFile("Cards/card_back_suits_dark.png")
	cpuHandImage.SetMinSize(fyne.NewSize(180, 300))
	cpuHandImage.FillMode = canvas.ImageFillContain

	// TITLE - Using canvas.Text for larger size
	gameTitle := canvas.NewText("War Card Game", color.White)
	gameTitle.TextStyle.Bold = true
	gameTitle.TextSize = 22
	gameTitle.Alignment = fyne.TextAlignCenter

	// GAME RESULT (hidden initially) - Using canvas.Text for larger size
	gameResult := canvas.NewText("", color.White)
	gameResult.Alignment = fyne.TextAlignCenter
	gameResult.TextSize = 20
	gameResult.TextStyle.Bold = true
	gameResult.Hide()

	// ENHANCED INITIAL DISPLAY
	initialDisplay := container.NewCenter(
		container.NewVBox(
			notificationText,
			container.NewCenter(
				func() *canvas.Text {
					sword := canvas.NewText("⚔️", color.White)
					sword.TextSize = 64
					sword.TextStyle.Bold = true
					sword.Alignment = fyne.TextAlignCenter
					return sword
				}()),

			container.NewCenter(
				container.NewHBox(
					func() *canvas.Text {
						player := canvas.NewText("👤 PLAYER", color.White)
						player.TextSize = 24
						player.TextStyle.Bold = true
						return player
					}(),
					func() *canvas.Text {
						vs := canvas.NewText("🆚", color.White)
						vs.TextSize = 48
						vs.TextStyle.Bold = true
						vs.Alignment = fyne.TextAlignCenter
						return vs
					}(),
					func() *canvas.Text {
						cpu := canvas.NewText("🤖 CPU", color.White)
						cpu.TextSize = 24
						cpu.TextStyle.Bold = true
						return cpu
					}(),
				)),

			container.NewCenter(
				func() *canvas.Text {
					ready := canvas.NewText("Ready for Battle!", color.White)
					ready.TextStyle.Bold = true
					ready.TextSize = 20
					ready.Alignment = fyne.TextAlignCenter
					return ready
				}())))

	// BATTLE AREA (hidden initially) - Using canvas.Text for larger VS
	vsText := canvas.NewText("  VS  ", color.White)
	vsText.TextSize = 28
	vsText.TextStyle.Bold = true
	battleArea := container.NewCenter(
		container.NewHBox(playerCardImage, vsText, cpuCardImage))
	battleArea.Hide()

	// UPDATE SCORES FUNCTION - Updated for canvas.Text
	updateScores := func() {
		cpuCardCount.Text = fmt.Sprintf("%d", len(cpu.Cards))
		cpuCardCount.Refresh()
		playerCardCount.Text = fmt.Sprintf("%d", len(player1.Cards))
		playerCardCount.Refresh()
	}

	// Add visual hint - Using canvas.Text for larger size
	hintText := canvas.NewText("👆 Click your deck to play!", color.White)
	hintText.Alignment = fyne.TextAlignCenter
	hintText.TextStyle.Bold = true
	hintText.TextSize = 18
	hintText.Hide()

	// PLAY ROUND LOGIC
	executeRound := func() {
		playerCard, cpuCard, result, gameOver, winner := ExecuteGameRound(player1, cpu)

		if gameOver && winner == "" {
			return // No cards to play
		}

		// UI updates
		playerCardImage.File = playerCard.GetImageFilename()
		playerCardImage.Show()
		playerCardImage.Refresh()

		cpuCardImage.File = cpuCard.GetImageFilename()
		cpuCardImage.Show()
		cpuCardImage.Refresh()

		gameResult.Text = result
		gameResult.Show()
		gameResult.Refresh()
		updateScores()

		if gameOver {
			gameResult.Text = winner
			hintText.Hide()
		}
	}

	// START GAME BUTTON
	var playButton *widget.Button
	playButton = widget.NewButton("🎮 Start Game", func() {
		gameTitle.Hide()
		playButton.Hide()
		initialDisplay.Hide()
		battleArea.Show()
		hintText.Show()
	})
	playButton.Resize(fyne.NewSize(200, 50))

	// CREATE CLICKABLE CARD (simplified - no count parameter)
	clickablePlayerCard := NewClickableCard(playerHandImage, func() {
		if !gameTitle.Visible() && battleArea.Visible() {
			executeRound()
		}
	})

	// LAYOUT with tighter spacing - remove gaps
	topArea := container.NewVBox(
		container.NewCenter(gameTitle),
		container.NewStack(cpuHandImage, cpuCardCount),
	)

	middleArea := container.NewCenter(
		container.NewVBox(
			container.NewStack(initialDisplay, battleArea),
			playButton,
			gameResult,
		))

	bottomArea := container.NewVBox(
		hintText,
		container.NewStack(clickablePlayerCard, playerCardCount),
	)

	content := container.NewBorder(topArea, bottomArea, nil, nil, middleArea)

	// FINAL DISPLAY with beautiful purple background
	background := canvas.NewRectangle(color.RGBA{102, 51, 153, 255})
	contentWithBackground := container.NewStack(background, content)
	myWindow.SetContent(contentWithBackground)

	// Start the notification display after window is ready
	go func() {
		time.Sleep(1 * time.Second)

		for _, msg := range setupMessages {
			message := msg
			fyne.Do(func() {
				showNotification(message)
			})
			time.Sleep(3 * time.Second)
		}
	}()

	// I tried in so many different ways to make the score appear exactly where I want it to. Go func(goroutine)
	// is the only way that worked for me
	go func() {
		time.Sleep(100 * time.Millisecond)
		fyne.Do(func() {
			cpuCardCount.Move(fyne.NewPos(1, -5))
			playerCardCount.Move(fyne.NewPos(1, -10))
			cpuCardCount.Refresh()
			playerCardCount.Refresh()
		})
	}()

	// track fullscreen state
	isFull := false

	// Listen for key events on the window's canvas
	myWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if ev.Name == fyne.KeyF11 {
			isFull = !isFull
			myWindow.SetFullScreen(isFull)
		}
	})

	myWindow.ShowAndRun()
}

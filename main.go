package main

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("War Card Game")
	myWindow.Resize(fyne.NewSize(1920, 1080))

	// Background for all screens
	background := canvas.NewRectangle(color.RGBA{102, 51, 153, 255})

	// === MAIN MENU SCREEN ===
	menuTitle := canvas.NewText("⚔️ WAR CARD GAME ⚔️", color.White)
	menuTitle.TextStyle.Bold = true
	menuTitle.TextSize = 48
	menuTitle.Alignment = fyne.TextAlignCenter

	menuSubtitle := canvas.NewText("Battle of the Cards", color.White)
	menuSubtitle.TextStyle.Bold = true
	menuSubtitle.TextSize = 24
	menuSubtitle.Alignment = fyne.TextAlignCenter

	// Menu buttons
	// I just could not avoid using var here
	var gameContainer *fyne.Container
	var showGameScreen func()

	startButton := widget.NewButton("🎮 Start New Game", func() {
		showGameScreen()
	})
	startButton.Resize(fyne.NewSize(300, 60))

	rulesButton := widget.NewButton("📖 How to Play", func() {
		rulesText := `WAR CARD GAME RULES:

🎯 OBJECTIVE: Win all the cards!

🃏 CARD HIERARCHY (Lowest to Highest):
• 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace
• Normal Joker (15) - Very strong!
• Red Joker (16) - Stronger!
• Black Joker (17) - UNBEATABLE!

🎮 HOW TO PLAY:
• Click your deck to play a card
• Higher card wins both cards
• Winner collects all played cards

⚔️ WAR HAPPENS WHEN:
• Both players play the same value
• Each player puts down 4 cards (or all remaining)
• The last card played determines the winner
• Winner takes ALL cards from the war!

🎲 GAME SETUP:
• 55 cards total (52 regular + 3 Jokers)
• Cards dealt evenly between players
• One random player gets the extra card

🏆 WIN CONDITION:
• Game ends when opponent runs out of cards
• Player with all cards wins!

Good luck, warrior! ⚔️`

		dialog.ShowInformation("How to Play War", rulesText, myWindow)
	})
	rulesButton.Resize(fyne.NewSize(300, 60))

	quitButton := widget.NewButton("❌ Quit Game", func() {
		myApp.Quit()
	})
	quitButton.Resize(fyne.NewSize(300, 60))

	menuButtons := container.NewVBox(
		startButton,
		widget.NewLabel(""), // Spacer
		rulesButton,
		widget.NewLabel(""), // Spacer
		quitButton,
	)

	menuScreen := container.NewVBox(
		widget.NewLabel(""), // Top spacer
		widget.NewLabel(""), // Top spacer
		menuTitle,
		widget.NewLabel(""), // Spacer
		menuSubtitle,
		widget.NewLabel(""), // Spacer
		widget.NewLabel(""), // Spacer
		container.NewCenter(menuButtons),
	)

	// === GAME SCREEN (game logic) ===
	createGameScreen := func() *fyne.Container {
		// Initialize the game
		player1, cpu := StartGame()

		// Card count labels
		cpuCardCount := canvas.NewText(fmt.Sprintf("%d", len(cpu.Cards)), color.White)
		cpuCardCount.Alignment = fyne.TextAlignCenter
		cpuCardCount.TextStyle.Bold = true
		cpuCardCount.TextSize = 20

		playerCardCount := canvas.NewText(fmt.Sprintf("%d", len(player1.Cards)), color.White)
		playerCardCount.Alignment = fyne.TextAlignCenter
		playerCardCount.TextStyle.Bold = true
		playerCardCount.TextSize = 20

		// Card images
		playerCardImage := canvas.NewImageFromFile("Cards/card_joker.png")
		playerCardImage.SetMinSize(fyne.NewSize(250, 350))
		playerCardImage.FillMode = canvas.ImageFillContain
		playerCardImage.Hide()

		cpuCardImage := canvas.NewImageFromFile("Cards/card_joker.png")
		cpuCardImage.SetMinSize(fyne.NewSize(250, 350))
		cpuCardImage.FillMode = canvas.ImageFillContain
		cpuCardImage.Hide()

		playerHandImage := canvas.NewImageFromFile("Cards/card_back_suits_blue.png")
		playerHandImage.SetMinSize(fyne.NewSize(180, 300))
		playerHandImage.FillMode = canvas.ImageFillContain

		cpuHandImage := canvas.NewImageFromFile("Cards/card_back_suits_dark.png")
		cpuHandImage.SetMinSize(fyne.NewSize(180, 300))
		cpuHandImage.FillMode = canvas.ImageFillContain

		// Game result text
		gameResult := canvas.NewText("", color.White)
		gameResult.Alignment = fyne.TextAlignCenter
		gameResult.TextSize = 20
		gameResult.TextStyle.Bold = true
		gameResult.Hide()

		// Battle area
		vsText := canvas.NewText("  VS  ", color.White)
		vsText.TextSize = 28
		vsText.TextStyle.Bold = true
		battleArea := container.NewCenter(
			container.NewHBox(playerCardImage, vsText, cpuCardImage))

		// Hint text
		hintText := canvas.NewText("👇 Click your deck to play!", color.White)
		hintText.Alignment = fyne.TextAlignCenter
		hintText.TextStyle.Bold = true
		hintText.TextSize = 18

		// Stats
		leftStats := canvas.NewText("Wars: 0", color.White)
		leftStats.Alignment = fyne.TextAlignLeading
		leftStats.TextSize = 34
		leftStats.TextStyle.Bold = true

		rightStats := canvas.NewText("Long: 0", color.White)
		rightStats.Alignment = fyne.TextAlignTrailing
		rightStats.TextSize = 34
		rightStats.TextStyle.Bold = true

		warsThisGame := 0
		longestWar := 0
		gameAcceptingClicks := true // Track if game accepts card clicks

		// Update scores function
		updateScores := func() {
			cpuCardCount.Text = fmt.Sprintf("%d", len(cpu.Cards))
			cpuCardCount.Refresh()
			playerCardCount.Text = fmt.Sprintf("%d", len(player1.Cards))
			playerCardCount.Refresh()
		}

		// Return to menu button
		returnToMenuButton := widget.NewButton("🏠 Main Menu", func() {
			// Hide game screen, show menu screen
			gameContainer.Hide()
			menuScreen.Show()
		})
		returnToMenuButton.Hide() // Hidden until game ends

		// Play round logic
		executeRound := func() {
			playerCard, cpuCard, result, gameOver, winner, warInfo := ExecuteGameRound(player1, cpu)

			if gameOver && winner == "" {
				return
			}

			// Check if this is a war - show tied cards first!
			if warInfo.IsWar {
				gameAcceptingClicks = false                   // Disable clicks during war
				ShakeContainer(battleArea, 25*time.Second/10) // Shake animation for war
				// Show the tied cars that caused the war (2 seconds)
				playerCardImage.File = warInfo.TiedCard1.GetImageFilename()
				playerCardImage.Show()
				playerCardImage.Refresh()

				cpuCardImage.File = warInfo.TiedCard2.GetImageFilename()
				cpuCardImage.Show()
				cpuCardImage.Refresh()

				// Show war message with tied card info
				gameResult.Text = fmt.Sprintf("⚔️ WAR! Both played %s! Each player puts down 4 cards! ⚔️",
					warInfo.TiedCard1.GetDisplayValue())
				gameResult.Show()
				gameResult.Refresh()

				// After 5 seconds, replace with card backs (3 seconds)
				time.AfterFunc(5*time.Second, func() {
					fyne.Do(func() {
						// Replace tied cards with card backs
						playerCardImage.File = "Cards/card_back_suits_blue.png"
						cpuCardImage.File = "Cards/card_back_suits_dark.png"
						playerCardImage.Refresh()
						cpuCardImage.Refresh()

						gameResult.Text = "⚔️ WAR IN PROGRESS... ⚔️"
						gameResult.Refresh()
					})

					// After 3 more seconds, show final winning cards
					time.AfterFunc(3*time.Second, func() {
						fyne.Do(func() {
							playerCardImage.File = playerCard.GetImageFilename()
							cpuCardImage.File = cpuCard.GetImageFilename()
							playerCardImage.Refresh()
							cpuCardImage.Refresh()

							gameResult.Text = result
							gameResult.Refresh()
							gameAcceptingClicks = true // Re-enable clicks after war
							updateScores()
						})
					})
				})
			} else {
				// Normal round - no war
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
			}

			// Stats tracking
			if strings.Contains(result, "WAR!") {
				warsThisGame++
				leftStats.Text = fmt.Sprintf("Wars: %-3d", warsThisGame)
				leftStats.Refresh()

				warSize := strings.Count(result, "WAR!") * 4
				if warSize > longestWar {
					longestWar = warSize
					rightStats.Text = fmt.Sprintf("Long: %-3d", longestWar)
					rightStats.Refresh()
				}
			}

			if gameOver {
				gameResult.Text = winner
				hintText.Hide()
				returnToMenuButton.Show()
			}
		}

		// Clickable player card
		clickablePlayerCard := NewClickableCard(playerHandImage, func() {
			if battleArea.Visible() && gameAcceptingClicks {
				executeRound()
			}
		})

		// Layout
		topArea := container.NewVBox(
			container.NewStack(cpuHandImage, cpuCardCount),
		)

		middleContent := container.NewVBox(
			battleArea,
			gameResult,
			returnToMenuButton,
		)

		middleArea := container.NewBorder(
			nil, nil,
			leftStats,
			rightStats,
			container.NewCenter(middleContent))

		bottomArea := container.NewVBox(
			hintText,
			container.NewStack(clickablePlayerCard, playerCardCount),
		)

		gameContent := container.NewBorder(topArea, bottomArea, nil, nil, middleArea)

		// Position card counts
		go func() {
			time.Sleep(100 * time.Millisecond)
			fyne.Do(func() {
				cpuCardCount.Move(fyne.NewPos(1, -5))
				playerCardCount.Move(fyne.NewPos(1, -10))
				cpuCardCount.Refresh()
				playerCardCount.Refresh()
			})
		}()

		return gameContent
	}

	// Create the initial game container
	gameContainer = createGameScreen()
	gameContainer.Hide() // Start hidden

	// Navigation function
	showGameScreen = func() {
		menuScreen.Hide()
		gameContainer.Show()
	}

	// Main container with both screens
	mainContainer := container.NewStack(
		container.NewCenter(menuScreen),
		gameContainer,
	)

	// Final window content
	finalContent := container.NewStack(background, mainContainer)
	myWindow.SetContent(finalContent)

	// Update Start New Game button with proper reset logic
	startButton.OnTapped = func() {
		// Hide old game container
		gameContainer.Hide()

		// Create completely fresh game container
		gameContainer = createGameScreen()

		// Update the main container with the new game container
		mainContainer.Objects[1] = gameContainer

		// Show the new game
		showGameScreen()
	}

	// Fullscreen toggle
	isFull := false
	myWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if ev.Name == fyne.KeyF11 {
			isFull = !isFull
			myWindow.SetFullScreen(isFull)
		}
	})

	myWindow.ShowAndRun()
}

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// GetBundledResource returns the bundled resource for a given file path
func GetBundledResource(filePath string) *fyne.StaticResource {
	switch filePath {
	// Card backs
	case "Cards/card_back_dark_inner.png":
		return resourceCardbackdarkinnerPng
	case "Cards/card_back_intricate.png":
		return resourceCardbackintricatePng
	case "Cards/card_back_light_inner.png":
		return resourceCardbacklightinnerPng
	case "Cards/card_back_line.png":
		return resourceCardbacklinePng
	case "Cards/card_back_line_light.png":
		return resourceCardbacklinelightPng
	case "Cards/card_back_plain.png":
		return resourceCardbackplainPng
	case "Cards/card_back_suits.png":
		return resourceCardbacksuitsPng
	case "Cards/card_back_suits_blue.png":
		return resourceCardbacksuitsbluePng
	case "Cards/card_back_suits_dark.png":
		return resourceCardbacksuitsdarkPng

	// Spades
	case "Cards/card_spade_2.png":
		return resourceCardspade2Png
	case "Cards/card_spade_3.png":
		return resourceCardspade3Png
	case "Cards/card_spade_4.png":
		return resourceCardspade4Png
	case "Cards/card_spade_5.png":
		return resourceCardspade5Png
	case "Cards/card_spade_6.png":
		return resourceCardspade6Png
	case "Cards/card_spade_7.png":
		return resourceCardspade7Png
	case "Cards/card_spade_8.png":
		return resourceCardspade8Png
	case "Cards/card_spade_9.png":
		return resourceCardspade9Png
	case "Cards/card_spade_10.png":
		return resourceCardspade10Png
	case "Cards/card_spade_11.png":
		return resourceCardspade11Png
	case "Cards/card_spade_12.png":
		return resourceCardspade12Png
	case "Cards/card_spade_13.png":
		return resourceCardspade13Png
	case "Cards/card_spade_ace.png":
		return resourceCardspadeacePng

	// Hearts
	case "Cards/card_heart_2.png":
		return resourceCardheart2Png
	case "Cards/card_heart_3.png":
		return resourceCardheart3Png
	case "Cards/card_heart_4.png":
		return resourceCardheart4Png
	case "Cards/card_heart_5.png":
		return resourceCardheart5Png
	case "Cards/card_heart_6.png":
		return resourceCardheart6Png
	case "Cards/card_heart_7.png":
		return resourceCardheart7Png
	case "Cards/card_heart_8.png":
		return resourceCardheart8Png
	case "Cards/card_heart_9.png":
		return resourceCardheart9Png
	case "Cards/card_heart_10.png":
		return resourceCardheart10Png
	case "Cards/card_heart_11.png":
		return resourceCardheart11Png
	case "Cards/card_heart_12.png":
		return resourceCardheart12Png
	case "Cards/card_heart_13.png":
		return resourceCardheart13Png
	case "Cards/card_heart_ace.png":
		return resourceCardheartacePng

	// Diamonds
	case "Cards/card_diamond_2.png":
		return resourceCarddiamond2Png
	case "Cards/card_diamond_3.png":
		return resourceCarddiamond3Png
	case "Cards/card_diamond_4.png":
		return resourceCarddiamond4Png
	case "Cards/card_diamond_5.png":
		return resourceCarddiamond5Png
	case "Cards/card_diamond_6.png":
		return resourceCarddiamond6Png
	case "Cards/card_diamond_7.png":
		return resourceCarddiamond7Png
	case "Cards/card_diamond_8.png":
		return resourceCarddiamond8Png
	case "Cards/card_diamond_9.png":
		return resourceCarddiamond9Png
	case "Cards/card_diamond_10.png":
		return resourceCarddiamond10Png
	case "Cards/card_diamond_11.png":
		return resourceCarddiamond11Png
	case "Cards/card_diamond_12.png":
		return resourceCarddiamond12Png
	case "Cards/card_diamond_13.png":
		return resourceCarddiamond13Png
	case "Cards/card_diamond_ace.png":
		return resourceCarddiamondacePng

	// Clubs
	case "Cards/card_clubs_2.png":
		return resourceCardclubs2Png
	case "Cards/card_clubs_3.png":
		return resourceCardclubs3Png
	case "Cards/card_clubs_4.png":
		return resourceCardclubs4Png
	case "Cards/card_clubs_5.png":
		return resourceCardclubs5Png
	case "Cards/card_clubs_6.png":
		return resourceCardclubs6Png
	case "Cards/card_clubs_7.png":
		return resourceCardclubs7Png
	case "Cards/card_clubs_8.png":
		return resourceCardclubs8Png
	case "Cards/card_clubs_9.png":
		return resourceCardclubs9Png
	case "Cards/card_clubs_10.png":
		return resourceCardclubs10Png
	case "Cards/card_clubs_11.png":
		return resourceCardclubs11Png
	case "Cards/card_clubs_12.png":
		return resourceCardclubs12Png
	case "Cards/card_clubs_13.png":
		return resourceCardclubs13Png
	case "Cards/card_clubs_ace.png":
		return resourceCardclubsacePng

	// Jokers
	case "Cards/card_joker.png":
		return resourceCardjokerPng
	case "Cards/card_joker_red.png":
		return resourceCardjokerredPng
	case "Cards/card_joker_black.png":
		return resourceCardjokerblackPng

	// Sounds
	case "sounds/card_shuffle.mp3":
		return resourceCardshuffleMp3
	case "sounds/playcard.mp3":
		return resourcePlaycardMp3

	default:
		return nil // Return nil for unbundled resources
	}
}

// GetBundledResourceByCard returns the bundled resource for a given card
func GetBundledResourceByCard(card Card) *fyne.StaticResource {
	filename := card.GetImageFilename()
	return GetBundledResource(filename)
}

// SetImageFromBundledResource sets an image widget's resource from a bundled resource
func SetImageFromBundledResource(img *canvas.Image, filePath string) {
	resource := GetBundledResource(filePath)
	if resource != nil {
		img.Resource = resource
		img.File = "" // Clear file path when using resource
		img.Refresh()
	}
}

// SetImageFromCard sets an image widget's resource from a card
func SetImageFromCard(img *canvas.Image, card Card) {
	resource := GetBundledResourceByCard(card)
	if resource != nil {
		img.Resource = resource
		img.File = "" // Clear file path when using resource
		img.Refresh()
	}
}

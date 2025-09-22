#!/bin/bash

echo "🎴 Bundling all War Card Game assets..."

# Remove old resources file
rm -f resources.go

# Start with first asset (no -append)
echo "📁 Bundling card backs..."
fyne bundle -o resources.go Cards/card_back_dark_inner.png
fyne bundle -o resources.go -append Cards/card_back_intricate.png
fyne bundle -o resources.go -append Cards/card_back_light_inner.png
fyne bundle -o resources.go -append Cards/card_back_line.png
fyne bundle -o resources.go -append Cards/card_back_line_light.png
fyne bundle -o resources.go -append Cards/card_back_plain.png
fyne bundle -o resources.go -append Cards/card_back_suits.png
fyne bundle -o resources.go -append Cards/card_back_suits_blue.png
fyne bundle -o resources.go -append Cards/card_back_suits_dark.png

echo "♠️ Bundling spades..."
for value in 2 3 4 5 6 7 8 9 10 11 12 13 ace; do
    fyne bundle -o resources.go -append Cards/card_spade_${value}.png
done

echo "♥️ Bundling hearts..."
for value in 2 3 4 5 6 7 8 9 10 11 12 13 ace; do
    fyne bundle -o resources.go -append Cards/card_heart_${value}.png
done

echo "♦️ Bundling diamonds..."
for value in 2 3 4 5 6 7 8 9 10 11 12 13 ace; do
    fyne bundle -o resources.go -append Cards/card_diamond_${value}.png
done

echo "♣️ Bundling clubs..."
for value in 2 3 4 5 6 7 8 9 10 11 12 13 ace; do
    fyne bundle -o resources.go -append Cards/card_clubs_${value}.png
done

echo "🃏 Bundling jokers..."
fyne bundle -o resources.go -append Cards/card_joker.png
fyne bundle -o resources.go -append Cards/card_joker_red.png
fyne bundle -o resources.go -append Cards/card_joker_black.png

echo "🎵 Bundling sounds..."
fyne bundle -o resources.go -append sounds/card_shuffle.mp3
fyne bundle -o resources.go -append sounds/playcard.mp3

echo "✅ All assets bundled into resources.go!"
echo "📊 Resource file size:"
ls -lh resources.go

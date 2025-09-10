# Card Game Engine: War

## CURRENT STATUS ✅
- Basic War game working with GUI
- Custom deck with special cards (2 super aces, 2 jokers)
- Clean layout with proper positioning

## MAJOR UPGRADE PLAN 🎯
**Goal: Add real card images and improve card system**

### PHASE 1: Update Card System (Code Changes)
- [x ] Remove 2 Super Aces from deck
- [x ] Update to 3 Jokers: Normal(15), Red(16), Black(17)
- [x ] Update card values: Normal cards (2-14), then Jokers (15-17)
- [x ] Test that new deck works properly

### PHASE 2: Add Card Images (Visual Upgrade)
- [x] Create cards folder - Done!
- [x] Add all card images - Done!
- [ x] **TEST STEP**: Load just ONE card image first
- [ ] Create helper function to map Card→filename
- [ ] Choose card backs (player vs CPU)
- [ ] Test card back loading

### PHASE 2A: The Mapping Challenge 🧩
We need to translate:
- Card{Value: 13, Suit: "Hearts"} → "Cards/card_heart_13.png"
- Card{Value: 15, Suit: "ColorfulJoker"} → "Cards/card_joker.png"
- Card{Value: 16, Suit: "RedJoker"} → "Cards/card_joker_red.png"
- Card{Value: 17, Suit: "BlackJoker"} → "Cards/card_joker_black.png"


### PHASE 3: Display Card Images in Game
- [ ] Replace text card display with actual card images
- [ ] Show player cards with blue back
- [ ] Show CPU cards with red back
- [ ] Show played cards as front images

### PHASE 4: Polish & Testing
- [ ] Make sure all 3 jokers display properly
- [ ] Test that Black Joker beats everything
- [ ] Adjust image sizes to look good
- [ ] Add win/lose animations if time permits

## TECHNICAL NOTES 📝
- Use Fyne's `storage.NewFileResource()` for images
- Images should be PNG or JPG format
- Start with one test image before doing all 55 cards
- Keep backup of working text version

## RISKS TO WATCH 🚨
- Image loading might be tricky for beginner
- File paths need to be correct
- Images might be wrong size
- Don't break current working game!

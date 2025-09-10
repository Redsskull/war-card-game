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
- [ ] Create `assets/` folder in project -- created cards folder instead!
- [ ] Add all card front images to assets folder
- [ ] Add card back images (red and blue) to assets folder -- I need to pick one one of the back images to become blue, I hope Flye can do this.
- [ ] Test loading one image first before doing all cards

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

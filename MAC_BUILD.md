# Mac Build Instructions

## Overview
Your War Card Game has been converted to use bundled resources to solve Mac's file access limitations. All assets (card images and sounds) are now embedded directly into the executable.

## What Changed
- **Fixed Mac Issue**: Assets are now bundled instead of loaded from external files
- **Single Executable**: No need for Cards/ or sounds/ folders on Mac
- **Ready to Build**: All file references converted to use bundled resources

## Building on Mac

### Prerequisites
Make sure you have Go installed:
```bash
go version  # Should show Go 1.19+
```

### Build Steps
```bash
# 1. Bundle all assets into resources.go
./bundle_assets.sh

# 2. Build Mac app bundle
./build_mac.sh
```

This creates:
- `war-card-game.app` - Mac app bundle
- `war-card-game-mac.tar.gz` - Distribution archive

### Manual Build (if needed)
```bash
# Bundle assets first
./bundle_assets.sh

# Simple build
go build -o war-card-game

# Or with optimizations
go build -ldflags="-s -w" -o war-card-game
```

## What's Bundled
- ✅ All 59 card images (Cards/*.png)
- ✅ Both sound files (sounds/*.mp3) 
- ✅ Complete game logic and UI

**File size**: ~7-10MB with all assets embedded

## Distribution
The resulting Mac app:
- Works without external files
- Can be distributed as a single .app bundle
- Users may need to right-click → "Open" for unsigned apps

## Testing
After building, test that:
- [ ] Game launches without errors
- [ ] All card images display correctly
- [ ] Sound effects play properly
- [ ] No "file not found" errors in console

## Success!
Your Mac build now has all assets properly bundled and should work on any Mac without requiring the Cards/ and sounds/ folders in the same directory.
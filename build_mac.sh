#!/bin/bash

echo "🍎 Building War Card Game for Mac distribution..."

# Clean previous Mac builds
echo "🧹 Cleaning previous Mac builds..."
rm -rf war-card-game.app
rm -f war-card-game-mac.tar.gz
rm -f war-card-game-mac.zip
rm -f war-card-game-mac.dmg

# Make sure assets are bundled
echo "📦 Bundling assets..."
./bundle_assets.sh

# Build for macOS (local build)
echo "🔨 Building for macOS..."
go build -ldflags="-s -w" -o war-card-game-mac ./...

# Create Mac app bundle structure
echo "📱 Creating Mac app bundle..."
mkdir -p war-card-game.app/Contents/MacOS
mkdir -p war-card-game.app/Contents/Resources

# Create Info.plist
cat > war-card-game.app/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>war-card-game</string>
    <key>CFBundleIdentifier</key>
    <string>com.warcardgame.app</string>
    <key>CFBundleName</key>
    <string>War Card Game</string>
    <key>CFBundleVersion</key>
    <string>1.0</string>
    <key>CFBundleShortVersionString</key>
    <string>1.0</string>
    <key>CFBundleInfoDictionaryVersion</key>
    <string>6.0</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleSignature</key>
    <string>????</string>
    <key>LSMinimumSystemVersion</key>
    <string>10.13</string>
    <key>NSHighResolutionCapable</key>
    <true/>
</dict>
</plist>
EOF

# Move the binary into the app bundle
mv war-card-game-mac war-card-game.app/Contents/MacOS/war-card-game

# Copy icon if it exists
if [ -f "Icon.png" ]; then
    cp Icon.png war-card-game.app/Contents/Resources/
fi

# Make the binary executable
chmod +x war-card-game.app/Contents/MacOS/war-card-game

# Create distribution archives
echo "📦 Creating distribution archives..."

# Create ZIP (most compatible with GitHub)
zip -r war-card-game-mac.zip war-card-game.app/

# Also create tar.gz for consistency
tar -czf war-card-game-mac.tar.gz war-card-game.app/

# Try to create DMG (requires macOS)
if command -v hdiutil >/dev/null 2>&1; then
    echo "💿 Creating DMG..."
    # Create a temporary directory for DMG contents
    mkdir -p dmg_temp
    cp -R war-card-game.app dmg_temp/

    # Create the DMG
    hdiutil create -volname "War Card Game" -srcfolder dmg_temp -ov -format UDZO war-card-game-mac.dmg
    rm -rf dmg_temp

    echo "✅ Created DMG file"
else
    echo "ℹ️  Skipping DMG creation (hdiutil not available)"
fi

echo "✅ Mac build complete!"
echo "📁 Files created:"
echo "   - war-card-game.app (Mac app bundle)"
echo "   - war-card-game-mac.zip (ZIP archive - recommended for GitHub)"
echo "   - war-card-game-mac.tar.gz (tar.gz archive)"
if [ -f "war-card-game-mac.dmg" ]; then
    echo "   - war-card-game-mac.dmg (DMG disk image)"
fi
echo ""
echo "🚀 Distribution recommendations:"
echo "   • GitHub Release: Upload war-card-game-mac.zip"
echo "   • Direct distribution: Use war-card-game-mac.dmg (if created)"
echo "   • Both formats work - ZIP is simpler, DMG is more professional"

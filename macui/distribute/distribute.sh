cd ..
echo $HOME
fyne package -os ios  -profile C3CBHTX5N7
# fyne package -os iossimulator -profile C3CBHTX5N7
mv My-App.app distribute/My-App.app
cd distribute
rm -rf My-App
mkdir My-App
mkdir My-App/Payload
cp -r My-App.app My-App/Payload/My-App.app
cp Icon.png My-App/iTunesArtwork
cd My-App
zip -r My-App.ipa Payload iTunesArtwork
cd ..
rm -rf My-App.app
exit 0

# xcrun simctl install booted myapp.app(安装到当前模拟器)
# fyne release -os ios  -profile C3CBHTX5N7
#!/usr/bin/env bash

abc = `pwd`
echo $abc
fyne package -os iossimulator -profile C3CBHTX5N7
xcrun simctl install booted My-App.app
exit 0

# xcrun simctl install booted myapp.app(安装到当前模拟器)
# fyne release -os ios  -profile C3CBHTX5N7
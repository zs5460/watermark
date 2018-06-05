@echo off
if exist dist/watermark.zip (del dist/watermark.zip)
if exist dist/watermark.exe (del dist/watermark.exe)
go build -o dist/watermark.exe
bft -page readme.md dist/readme.htm
upx dist/watermark.exe
cd dist
"c:\program files\7-zip\7z.exe" a watermark.zip watermark.exe watermark.png readme.htm sky.jpg

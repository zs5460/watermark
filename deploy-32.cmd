@echo off

if exist dist\watermark32.zip (del dist\watermark32.zip)
if exist dist\watermark.exe (del dist\watermark.exe)
set GOARCH=386
go build -o dist/watermark.exe
set GOARCH=amd64
upx dist\watermark.exe
bft -page readme.md dist/readme.htm
cd dist
"c:\program files\7-zip\7z.exe" a watermark32.zip watermark.exe watermark.png readme.htm sky.jpg

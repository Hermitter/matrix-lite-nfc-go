# MATRIX-Lite-NFC-Go

MATRIX Lite NFC Go is an npm package that allows users of varying skill levels to easily program NFC with their MATRIX Creator.

# Smartphone Apps For Debugging

- [IOS App](https://apps.apple.com/us/app/nfc-taginfo-by-nxp/id1246143596)
- [Android App](https://play.google.com/store/apps/details?id=com.nxp.nfc.tagwriter&hl=en_US)

# Roadmap

- [x] Reading Info (All tags)
- [ ] Reading Pages (MIFARE Ultralight & NTAG)
- [ ] Writing Page (MIFARE Ultralight & NTAG)
- [ ] Reading NDEF (MIFARE Ultralight & NTAG)
- [ ] Writing NDEF (MIFARE Ultralight & NTAG)

# Installation

1. Install [matrix-hal-nfc](https://github.com/matrix-io/matrix-hal-nfc) to use this library.

2. Download and install the latest version of Golang.

```bash
sudo apt-get install golang
```

3. Create a project folder

```
mkdir lite_go
cd lite_go
go mod init myapp
touch main.go
```

4. Install matrix-lite-nfc-go

```bash
go get -u github.com/matrix-io/matrix-lite-nfc-go
```

# Building Locally For Development

Download the repository.

```
git clone https://github.com/matrix-io/matrix-lite-nfc-go

cd matrix-lite-nfc-go
```

Install the library.

```
go install -a
```

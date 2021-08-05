package main

import (
	"fmt"

	qrcode "github.com/BROUSSOLLE-Brice/go-qrcode"
)

func main() {
	repo()

	//issue17()
}

func repo() {
	qrc, err := qrcode.New("https://github.com/BROUSSOLLE-Brice/go-qrcode")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// save file
	if err = qrc.Save("../testdata/repo-qrcode.jpeg"); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}

func issue17() {
	qrc, err := qrcode.New("Övrigt asdasd asdas djaskl djaslk djaslkj dlaiodqjwiodjaskldj aksldjlk Övrigt")
	//qrc, err := qrcode.New("text content this is custom text content this is custom text content70123")
	// content over than 74 length would trigger this
	//qrc, err := qrcode.New("text content this is custom text content this is custom text content701234",
	//	qrcode.WithCircleShape())
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// save file
	if err = qrc.Save("./testdata/issue-17.jpeg"); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}

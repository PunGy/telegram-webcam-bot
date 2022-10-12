package helpers

import (
	"errors"
	"log"

	"gocv.io/x/gocv"
)

func MakeWebcamImage(deviceID string) ([]byte, error) {
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		log.Fatalf("Error: opening of video capture device: %v\n", err)
		return nil, err
	}

	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	webcam.Grab(30) // Skip first 30 frames to get camera properly initialized and focused

	if ok := webcam.Read(&img); !ok {
		log.Fatalf("Cannot read device %v\n", deviceID)
		return nil, errors.New("Cannot read device")
	}

	if img.Empty() {
		log.Fatalf("no image")
	}

	r, e := gocv.IMEncode(gocv.JPEGFileExt, img)
	return r.GetBytes(), e
}

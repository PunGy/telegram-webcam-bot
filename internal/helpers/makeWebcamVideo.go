package helpers

import (
	"errors"
	"log"
	"strconv"
	"time"

	"gocv.io/x/gocv"
)

func MakeWebcamVideo(deviceID string, length int) (*string, error) {
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		log.Fatalf("Error openning video capture device: %v\n", err)
		return nil, err
	}

	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		log.Fatalf("Cannot read device %v\n", deviceID)
		return nil, errors.New("Cannot read device")
	}

	tmpFile := strconv.FormatInt(time.Now().Unix(), 10) + "_video.mp4"
	var fps float64 = 25.0
	writer, err := gocv.VideoWriterFile(tmpFile, "mp4v", float64(fps), img.Cols(), img.Rows(), true)
	if err != nil {
		log.Printf("Error opening video writer %v", err)

	}

	defer writer.Close()

	framesCount := (float64(length) + 1.0) * fps
	for i := 0.0; i < framesCount; i++ {
		if ok := webcam.Read(&img); !ok {
			log.Printf("Device closed")
		}
		if img.Empty() {
			continue
		}

		writer.Write(img)
	}

	return &tmpFile, err
}

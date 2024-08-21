package helpers

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

const ShellToUse = "bash"

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func MakeWebcamVideo(deviceID string, length int) (*string, error) {
	tmpFile := strconv.FormatInt(time.Now().Unix(), 10) + "_video.mp4"
	cmd := fmt.Sprintf("ffmpeg -f avfoundation -t %ds -framerate 30 -video_size 1280x720 -i \"default:default\" -c:v libx264 -preset veryfast -crf 23 -pix_fmt yuv420p -c:a aac -b:a 128k %s", length, tmpFile)
	log.Printf("COMMAND: %s", cmd)
	_, errout, err := Shellout(cmd)
	if err != nil {
		log.Printf("Error opening video writer %v", errout)
	}

	return &tmpFile, err

	// webcam, err := gocv.OpenVideoCapture(deviceID)
	// if err != nil {
	// 	log.Fatalf("Error openning video capture device: %v\n", err)
	// 	return nil, err
	// }
	//
	// defer webcam.Close()
	//
	// img := gocv.NewMat()
	// defer img.Close()
	//
	// if ok := webcam.Read(&img); !ok {
	// 	log.Fatalf("Cannot read device %v\n", deviceID)
	// 	return nil, errors.New("Cannot read device")
	// }
	//
	// tmpFile := strconv.FormatInt(time.Now().Unix(), 10) + "_video.mp4"
	// var fps float64 = 25.0
	// writer, err := gocv.VideoWriterFile(tmpFile, "mp4v", float64(fps), img.Cols(), img.Rows(), true)
	// if err != nil {
	// 	log.Printf("Error opening video writer %v", err)
	// 	return nil, err
	// }
	//
	// defer writer.Close()
	//
	// framesCount := (float64(length) + 1.0) * fps
	// for i := 0.0; i < framesCount; i++ {
	// 	if ok := webcam.Read(&img); !ok {
	// 		log.Printf("Device closed")
	// 	}
	// 	if img.Empty() {
	// 		continue
	// 	}
	//
	// 	writer.Write(img)
	// }
	//
	// return &tmpFile, err
}

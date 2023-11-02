package gadget

import "fmt"

type TapePlayer struct {
	batteries string
}

func (tapePlayer *TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (tapePlayer *TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	microphones int
}

func (tapeRecorder *TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (tapeRecorder *TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (tapeRecorder *TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

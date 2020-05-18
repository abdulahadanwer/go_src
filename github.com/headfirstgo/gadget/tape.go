package gadget

import "fmt"

//TapePlayer : struct to store tapeplayer information
type TapePlayer struct {
	Batteries string
}

//Play : plays the song
func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

//Stop : stops the song
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

//TapeRecorder  stores tape recorder information
type TapeRecorder struct {
	Microphones int
}

//Play : plays tape recorder information
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

//Stop : stops tape recorder information
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

//Record : records tape recorder information
func (t TapeRecorder) Record() {
	fmt.Println("Recordgin!")
}

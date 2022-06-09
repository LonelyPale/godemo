package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func main() {
	me := &MusicEntry{
		//Source: "/Users/wyb/Music/陈奕迅-孤勇者1.mp3",
		//Source: "/Users/wyb/Music/网易云音乐/电台节目/袁少有为 - 往后余生（粤语版）-亮声open.mp3",
		//Source: "/Users/wyb/Music/网易云音乐/电台节目/袁少有为 - 谁能明白我-亮声open.mp3",
		//Source: "/Users/wyb/Music/网易云音乐/电台节目/袁少有为 - 心碎-易欣-亮声open.mp3",
		//Source: "/Users/wyb/Music/网易云音乐/电台节目/袁少有为 - 暧昧-亮声open.mp3",
		//Source: "/Users/wyb/Music/网易云音乐/电台节目/袁少有为 - 爱的故事上集-亮声open.mp3",
		Source: "/Users/wyb/Music/网易云音乐/电台节目/袁少有为 - 一生所爱-亮声open.mp3",
	}
	me.Open()
	me.Play()
}

type MusicEntry struct {
	Id         string   //编号
	Name       string   //歌名
	Artist     string   //作者
	Source     string   //位置
	Type       string   //类型
	Filestream *os.File // 文件流
}

func (m *MusicEntry) Open() {
	var err error
	m.Filestream, err = os.Open(m.Source)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *MusicEntry) Play() {
	streamer, format, err := mp3.Decode(m.Filestream)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	separator := string(os.PathSeparator)
	path := strings.Split(m.Filestream.Name(), separator)
	name := path[len(path)-1]
	size := fmt.Sprintf("%dMB", streamer.Len()/1024/1024)
	timeAll := streamer.Len() / int(format.SampleRate)
	timeMinute := timeAll / 60
	timeSecond := timeAll % 60
	timeString := fmt.Sprintf("%dm%ds", timeMinute, timeSecond)
	fmt.Println(green(name), yellow(timeString), blue(size))

	//play1(format, streamer)
	//play2(format, streamer)
	play3(format, streamer)
	//play4(format, streamer)
	//play5(format, streamer)
}

func play1(format beep.Format, streamer beep.StreamSeekCloser) {
	sampleRate := format.SampleRate
	bufferSize := format.SampleRate.N(time.Second / 10)

	if err := speaker.Init(sampleRate, bufferSize); err != nil {
		log.Fatal(err)
	}

	speaker.Play(streamer)
	select {}
}

func play2(format beep.Format, streamer beep.StreamSeekCloser) {
	sampleRate := format.SampleRate * 2
	resampled := beep.Resample(4, format.SampleRate, sampleRate, streamer)

	if err := speaker.Init(sampleRate, sampleRate.N(time.Second/10)); err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {
		close(done)
	})))
	//<-done

	for {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			speaker.Lock()
			s := format.SampleRate.D(streamer.Position()).Round(time.Second)
			fmt.Printf("%v ", red(s))
			speaker.Unlock()
		}
	}
}

func play3(format beep.Format, streamer beep.StreamSeekCloser) {
	sampleRate := format.SampleRate
	bufferSize := format.SampleRate.N(time.Second / 10)

	if err := speaker.Init(sampleRate, bufferSize); err != nil {
		log.Fatal(err)
	}

	//loop := beep.Loop(3, streamer)         //重复3次
	//fast := beep.ResampleRatio(4, 5, loop) //快进5、6秒

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	fast := buffer.Streamer(0, buffer.Len())
	done := make(chan struct{})
	speaker.Play(beep.Seq(fast, beep.Callback(func() {
		close(done)
	})))

	for {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			speaker.Lock()
			s := format.SampleRate.D(fast.Position()).Round(time.Second)
			fmt.Printf("%v ", red(s))
			speaker.Unlock()
		}
	}
}

func play4(format beep.Format, streamer beep.StreamSeekCloser) {
	sampleRate := format.SampleRate
	bufferSize := format.SampleRate.N(time.Second / 10)

	if err := speaker.Init(sampleRate, bufferSize); err != nil {
		log.Fatal(err)
	}

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speedy := beep.ResampleRatio(4, 1, volume)
	speaker.Play(speedy)

	for {
		fmt.Print("Press [ENTER] to pause/resume. ")
		fmt.Scanln()

		speaker.Lock()
		ctrl.Paused = !ctrl.Paused
		volume.Volume += 0.5                  //增加音量
		speedy.SetRatio(speedy.Ratio() + 0.1) //快进、慢放
		speaker.Unlock()
	}
}

func play5(format beep.Format, streamer beep.StreamSeekCloser) {
	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/60)); err != nil {
		log.Fatal(err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	for {
		fmt.Print("Press [ENTER] to fire a gunshot! ")
		fmt.Scanln()

		shot := buffer.Streamer(0, buffer.Len())
		speaker.Play(shot)
	}
}

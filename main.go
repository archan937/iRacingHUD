package main

import (
	"C"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/antchfx/jsonquery"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

type (
	Config struct {
		Main     Pos `json:"main"`
		Position Pos `json:"position"`
		Session  Pos `json:"session"`
		Speed    Pos `json:"speed"`
		Time     Pos `json:"time"`
	}

	Pos struct {
		X int `json:"X"`
		Y int `json:"Y"`
	}

	Dimensions struct {
		Width  int `json:"Width"`
		Height int `json:"Height"`
	}
)

var (
	configFile       = "config.json"
	debug            = false
	User32           = syscall.NewLazyDLL("User32.dll")
	GetSystemMetrics = User32.NewProc("GetSystemMetrics")
	iRacingSDK       = syscall.NewLazyDLL("./iRacingSDK.dll")
	CurrentGameState = iRacingSDK.NewProc("CurrentGameState")
)

func getSystemMetrics(nIndex int) int {
	index := uintptr(nIndex)
	ret, _, _ := GetSystemMetrics.Call(index)
	return int(ret)
}

func calcPos(xPerc int, yPerc int) Pos {
	return Pos{
		X: int(float32(getSystemMetrics(0)) * (float32(xPerc) / 100.0)),
		Y: int(float32(getSystemMetrics(1)) * (float32(yPerc) / 100.0)),
	}
}

func getConfig() *Config {
	var config Config
	data, err := os.ReadFile(configFile)

	if err != nil {
		config = Config{
			Main:     calcPos(80, 70),
			Position: calcPos(15, 45),
			Session:  calcPos(15, 25),
			Speed:    calcPos(15, 65),
			Time:     calcPos(65, 45),
		}
	} else {
		json.Unmarshal([]byte(data), &config)
	}

	return &config
}

func storePosition(window *astilectron.Window, name string) {
	bounds, _ := window.Bounds()
	config := getConfig()

	var pos *Pos

	switch name {
	case "main":
		pos = &config.Main
	case "position":
		pos = &config.Position
	case "session":
		pos = &config.Session
	case "speed":
		pos = &config.Speed
	case "time":
		pos = &config.Time
	}

	pos.X = bounds.X
	pos.Y = bounds.Y

	marshalled, _ := json.Marshal(&config)
	os.WriteFile(configFile, marshalled, 0777)
}

func createWindow(app *astilectron.Astilectron, showFrame bool, name string) *astilectron.Window {
	config := getConfig()

	var pos *Pos
	var dimensions *Dimensions

	switch name {
	case "position":
		pos = &config.Position
	case "session":
		pos = &config.Session
	case "speed":
		pos = &config.Speed
	case "time":
		pos = &config.Time
	}

	if debug {
		dimensions = &Dimensions{Width: 600, Height: 300}
	} else {
		switch name {
		case "position":
			dimensions = &Dimensions{Width: 500, Height: 150}
		case "session":
			dimensions = &Dimensions{Width: 350, Height: 175}
		case "speed":
			dimensions = &Dimensions{Width: 560, Height: 205}
		case "time":
			dimensions = &Dimensions{Width: 500, Height: 190}
		}
	}

	options := &astilectron.WindowOptions{
		AlwaysOnTop: astikit.BoolPtr(true),
		Frame:       astikit.BoolPtr(debug || showFrame),
		Maximizable: astikit.BoolPtr(debug),
		Minimizable: astikit.BoolPtr(false),
		Resizable:   astikit.BoolPtr(debug),
		SkipTaskbar: astikit.BoolPtr(true),
		Width:       astikit.IntPtr(dimensions.Width),
		Height:      astikit.IntPtr(dimensions.Height),
		X:           astikit.IntPtr(pos.X),
		Y:           astikit.IntPtr(pos.Y),
		Transparent: astikit.BoolPtr(true),
	}

	if debug || showFrame {
		options = &astilectron.WindowOptions{
			AlwaysOnTop:     astikit.BoolPtr(true),
			Frame:           astikit.BoolPtr(debug || showFrame),
			Maximizable:     astikit.BoolPtr(debug),
			Minimizable:     astikit.BoolPtr(false),
			Resizable:       astikit.BoolPtr(debug),
			SkipTaskbar:     astikit.BoolPtr(true),
			Width:           astikit.IntPtr(dimensions.Width),
			Height:          astikit.IntPtr(dimensions.Height),
			X:               astikit.IntPtr(pos.X),
			Y:               astikit.IntPtr(pos.Y),
			BackgroundColor: astikit.StrPtr("#000"),
		}
	}

	var window, _ = app.NewWindow("views/"+name+".html", options)

	window.On(astilectron.EventNameWindowEventMoved, func(event astilectron.Event) (deleteListener bool) {
		storePosition(window, name)
		return
	})

	window.Create()

	if debug {
		window.OpenDevTools()
	}

	return window
}

func createWindows(app *astilectron.Astilectron, showFrame bool) (position *astilectron.Window, session *astilectron.Window, speed *astilectron.Window, time *astilectron.Window) {
	positionWindow := createWindow(app, showFrame, "position")
	sessionWindow := createWindow(app, showFrame, "session")
	speedWindow := createWindow(app, showFrame, "speed")
	timeWindow := createWindow(app, showFrame, "time")
	return positionWindow, sessionWindow, speedWindow, timeWindow
}

func sendMessage(window *astilectron.Window, last string, current string) string {
	if last != current {
		window.SendMessage(current)
		// fmt.Println(current)
	}
	return current
}

func main() {
	config := getConfig()

	app, _ := astilectron.New(nil, astilectron.Options{
		AppName: "iRacingHUD",
	})

	defer app.Close()
	app.HandleSignals()
	app.Start()

	mainWindow, _ := app.NewWindow("huds/main.html", &astilectron.WindowOptions{
		AlwaysOnTop: astikit.BoolPtr(true),
		Maximizable: astikit.BoolPtr(false),
		Minimizable: astikit.BoolPtr(true),
		Resizable:   astikit.BoolPtr(false),
		Width:       astikit.IntPtr(240),
		Height:      astikit.IntPtr(80),
		X:           astikit.IntPtr(config.Main.X),
		Y:           astikit.IntPtr(config.Main.Y),
	})

	mainWindow.On(astilectron.EventNameWindowEventMoved, func(event astilectron.Event) (deleteListener bool) {
		storePosition(mainWindow, "main")
		return
	})

	mainWindow.On(astilectron.EventNameWindowEventClosed, func(event astilectron.Event) (deleteListener bool) {
		app.Close()
		return
	})

	mainWindow.Create()

	var positionWindow *astilectron.Window
	var sessionWindow *astilectron.Window
	var speedWindow *astilectron.Window
	var timeWindow *astilectron.Window

	positionWindow, sessionWindow, speedWindow, timeWindow = createWindows(app, false)

	mainWindow.OnMessage(func(m *astilectron.EventMessage) interface{} {
		sessionWindow.Destroy()
		positionWindow.Destroy()
		speedWindow.Destroy()
		timeWindow.Destroy()

		var command string
		m.Unmarshal(&command)
		positionWindow, sessionWindow, speedWindow, timeWindow = createWindows(app, command == "CONFIGURE")

		return nil
	})

	go func() {
		currentSessionEpoch := -1
		currentLapEpoch := -1
		currentSectorEpoch := -1
		currentSector := -1
		lastSector := -1

		var positionJson string
		var sessionJson string
		var speedJson string
		var timeJson string

		for {
			gameState, _, _ := CurrentGameState.Call(uintptr(currentSessionEpoch), uintptr(currentLapEpoch), uintptr(currentSectorEpoch), uintptr(currentSector), uintptr(lastSector))
			gameStateJson := C.GoString((*C.char)(unsafe.Pointer(gameState)))

			// fmt.Printf("%v\n", gameStateJson)

			doc, _ := jsonquery.Parse(strings.NewReader(gameStateJson))

			currentSessionEpoch = int(jsonquery.FindOne(doc, "//CurrentSessionEpoch").Value().(float64))
			currentLapEpoch = int(jsonquery.FindOne(doc, "//CurrentLapEpoch").Value().(float64))
			currentSectorEpoch = int(jsonquery.FindOne(doc, "//CurrentSectorEpoch").Value().(float64))
			currentSector = int(jsonquery.FindOne(doc, "//CurrentSector").Value().(float64))
			lastSector = int(jsonquery.FindOne(doc, "//LastSector").Value().(float64))

			if jsonquery.FindOne(doc, "//PositionData/text()") != nil {
				positionData := jsonquery.FindOne(doc, "//PositionData/text()").Value()
				positionJson = sendMessage(positionWindow, positionJson, fmt.Sprint(positionData))

				sessionData := jsonquery.FindOne(doc, "//SessionData/text()").Value()
				sessionJson = sendMessage(sessionWindow, sessionJson, fmt.Sprint(sessionData))

				speedData := jsonquery.FindOne(doc, "//SpeedData/text()").Value()
				speedJson = sendMessage(speedWindow, speedJson, fmt.Sprint(speedData))

				timeData := jsonquery.FindOne(doc, "//TimeData/text()").Value()
				timeJson = sendMessage(timeWindow, timeJson, fmt.Sprint(timeData))
			}

			time.Sleep(time.Duration(50) * time.Millisecond)
		}
	}()

	app.Wait()
}

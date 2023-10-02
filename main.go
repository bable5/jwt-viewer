package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
	"github.com/rivo/tview"
)

func main() {

	var input string
	fmt.Scanf("%s", &input)

	bytes, err := b64.StdEncoding.DecodeString(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to decode input %v\n", err)
		os.Exit(1)
	}

	var obj map[string]interface{}
	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid input %v\n", err)
		os.Exit(1)
	}
	// Make a custom formatter with indent set
	fmt.Printf("%v\n", obj)

	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	jsonToken, _ := f.Marshal(obj)

	app := tview.NewApplication()
	tokenTextArea := tview.NewTextView().
		SetDynamicColors(true).
		SetText(string(input))
	tokenTextArea.SetBorder(true).SetTitle("Raw")

	jsonTextArea := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	jsonTextArea.SetBorder(true).SetTitle("Token")

	w := tview.ANSIWriter(jsonTextArea)
	w.Write([]byte(jsonToken))

	flex := tview.NewFlex().
		AddItem(tokenTextArea, 0, 1, false).
		AddItem(jsonTextArea, 0, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

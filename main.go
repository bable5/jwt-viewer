package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/rivo/tview"
)

func prettyPrintJson(s string) string {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(s), &obj)
	if err != nil {
		return s
	}

	// Make a custom formatter with indent set
	fmt.Printf("%v\n", obj)

	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	jsonToken, err := f.Marshal(obj)

	if err != nil {
		return s
	}

	return string(jsonToken)
}

func main() {

	var input string
	fmt.Scanf("%s", &input)

	parts := strings.Split(input, ".")

	if len(parts) != 3 {
		fmt.Fprintf(os.Stderr, "Invalid JWT\n")
		os.Exit(1)
	}

	header, _ := b64.URLEncoding.DecodeString(parts[0])
	body, _ := b64.URLEncoding.DecodeString(parts[1])
	sig, _ := b64.URLEncoding.DecodeString(parts[2])

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

	tokenPrettyPrint := fmt.Sprintf("%s\n%s\n%s",
		prettyPrintJson(string(header)), prettyPrintJson(string(body)), string(sig),
	)

	w := tview.ANSIWriter(jsonTextArea)
	w.Write([]byte(tokenPrettyPrint))

	flex := tview.NewFlex().
		AddItem(tokenTextArea, 0, 1, false).
		AddItem(jsonTextArea, 0, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

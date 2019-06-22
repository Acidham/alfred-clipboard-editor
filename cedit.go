package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/everdev/mack"
)

const SEP = "|"

type response struct {
	clicked string
	content string
}

func showDialog(dialog string, title string, answer string) (string, string) {
	response, err := mack.Dialog(dialog, title, answer)
	var ret string
	if err != nil {
		panic(err)
	}
	return ret, response.Clicked
}

func cleanUpstring(str string) string {
	return strings.ReplaceAll(str, `"`, `\"`)
}

func main() {
	c, err := clipboard.ReadAll()
	content := cleanUpstring(c)
	if err != nil {
		panic(err)
	}

	if content != "" {
		resp, _ := mack.Dialog("Edit Clipboard", "Enter changes", content)
		r := response{clicked: resp.Clicked, content: resp.Text}

		switch {
		case r.clicked == "OK":
			chgContent := cleanUpstring(r.content)
			//chgContent := cleanUpstring(chg)
			clipboard.WriteAll(chgContent)
			fmt.Print(chgContent, "|", "SUCCESS")
		case r.clicked == "Cancel":
			fmt.Print(SEP, "CANCELED")
		}
	} else {
		fmt.Print(SEP, "EMPTY")
	}

}

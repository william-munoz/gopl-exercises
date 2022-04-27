// The Package editor provides data manipulation with an external editor.
package editor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

// Edit launches an external editor and lets the user edit the value.
func Edit(value map[string]string) error {
	editor := getEditorName()

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}

	tempFileName := tempFile.Name()
	defer os.Remove(tempFileName)

	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(value)
	if err != nil {
		return err
	}

	err = tempFile.Close()
	if err != nil {
		return err
	}

	cmd := exec.Command(editor, tempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	edited, err := ioutil.ReadFile(tempFileName)
	if err != nil {
		return err
	}

	// Some editors, such as Windows Notepad, always add a BOM when saving UTF-8.
	// json.Unmarshal does not support data with UTF-8 BOM, so delete it beforehand.
	err = json.Unmarshal(removeUTF8BOM(edited), &value)
	if err != nil {
		return err
	}
	return nil
}

// Gets the name of the external editor.
// The external editor is specified by the environment variable GIT_EDITOR or EDITOR, like Git.
// https://git-scm.com/book/en/v2/Git-Internals-Environment-Variables
func getEditorName() string {
	editor := os.Getenv("GIT_EDITOR")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	// vi is included in almost all Linux distributions, so
	// If no external editor is specified, start vi.
	if editor == "" {
		editor = "vi"
	}
	return editor
}

// removeUTF8BOM removes any UTF-8 BOM at the beginning of the byte string.
func removeUTF8BOM(s []byte) []byte {
	utf8Bom := []byte{239, 187, 191}
	return bytes.TrimPrefix(s, utf8Bom)
}

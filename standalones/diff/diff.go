package diff

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Ext int
const (
	JS Ext = iota
	CLS
	CSS
	HTML
)
var extMap = map[string]Ext{
	".js": JS,
	".cls": CLS,
	".css": CSS,
	".html": HTML,
}

func getFileNameWithExt(path string) (filename, ext string) {
	fullFilename := filepath.Base(path)
	ext = filepath.Ext(path)
	filename = strings.TrimSuffix(fullFilename, ext)
	return
}

func getLWCSource(filename string) (string, error) {
	filename = strings.ReplaceAll(filename, "force-app/main/default/", "")
	fmt.Println(filename)
	query := fmt.Sprintf("SELECT Source FROM LightningComponentResource WHERE FilePath = '%s'", filename)

	cmdOutput, err := exec.Command("sf", "data", "query", "--query", query, "--use-tooling-api", "--json").Output()
	if err != nil { return "", fmt.Errorf("The file may not exist in the org : %w", err) }

	lwcSource, err := UnmarshalLwcSource(cmdOutput)
	if err != nil { return "", fmt.Errorf("Error unmarshalling lwc source : %w", err) }

	return lwcSource.Result.Records[0].Source, nil
}
func getApexSource(filename string) (string, error) {
	query := fmt.Sprintf("SELECT Body FROM ApexClass WHERE Name = '%s'", filename)
	cmdOutput, err := exec.Command("sf", "data", "query", "--query", query, "--json").Output()
	if err != nil { return "", fmt.Errorf("The file may not exist in the org : %w", err) }

	apexSource, err := UnmarshalApexSource(cmdOutput)
	if err != nil { return "", fmt.Errorf("Error unmarshalling apex source : %w", err) }

	return apexSource.Result.Records[0].Body, nil
}
func Process(path, editor string) {
	filename, ext := getFileNameWithExt(path)
	var source string
	var err error
	switch extMap[ext] {
		case JS, CSS, HTML:
			source, err = getLWCSource(path)
		case CLS:
			source, err = getApexSource(filename)
	}
	if err != nil {
		fmt.Println("Error getting source:", err)
		return
	}

	tempFile, err := createTempFile(filename, ext, source)

	cmd := exec.Command(editor, "--diff", tempFile.Name(), path, "-")
	if err := cmd.Run();
	err != nil {
		fmt.Println("Error running zed:", err)
		tempFile.Close()
		os.Remove(tempFile.Name())
		return
	}
	tempFile.Close()
	os.Remove(tempFile.Name())
}

func createTempFile(filename, ext, source string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", filename+"-*"+ext)
	if err != nil { return nil, fmt.Errorf("Error creating temp file: %w", err) }

	if _, err = tempFile.WriteString(source); err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, fmt.Errorf("Error writing to temp file: %w", err)
	}
	return tempFile, nil
}

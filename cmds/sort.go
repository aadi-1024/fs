package cmds

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/alecthomas/kong"
)

type Sort struct {
	Mkdir   bool `short:"M" help:"always make new directories; will be in the format FILEEXT_UNIXTIMESTAMP"`
	Verbose bool `short:"V" help:"more verbose outputs"`
}

func (s *Sort) Run(ctx *kong.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("couldn't get working directory: %w", err)
	}

	contents, err := os.ReadDir(wd)
	if err != nil {
		return fmt.Errorf("couldn't read directory, maybe check perms: %w", err)
	}

	filesByExt := make(map[string][]string)

	for _, v := range contents {
		ext := filepath.Ext(v.Name())
		fullPath := filepath.Join(wd, v.Name())
		if !v.IsDir() {
			s, ok := filesByExt[ext]
			if !ok {
				filesByExt[ext] = []string{fullPath}
			} else {
				s = append(s, fullPath)
				filesByExt[ext] = s
			}
		}
	}

	for k, v := range filesByExt {
		var folderPath string
		if k == "" {
			folderPath = "others"
		} else {
			//.pdf -> pdf
			folderPath = k[1:]
		}
		folderPath = filepath.Join(wd, folderPath)

		if s.Mkdir {
			folderPath = fmt.Sprintf("%v_%v", folderPath, time.Now().Unix())
		}
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			if os.IsExist(err) {
				fmt.Printf("%v already exists, moving into existing folder\n", folderPath)
			} else {
				return fmt.Errorf("couldn't create folder %v: %w", folderPath, err)
			}
		} else if s.Verbose {
			fmt.Printf("created folder %v\n", folderPath)
		}

		for _, file := range v {
			cmd := exec.Command("mv", file, folderPath)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("couldn't move file: %w", err)
			}
			if s.Verbose {
				fmt.Printf("moved file: %v\n", file)
			}
		}
	}
	return nil
}

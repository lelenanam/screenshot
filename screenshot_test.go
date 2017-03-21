package screenshot

import (
	"image"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestDetect(t *testing.T) {
	table := []struct {
		dir    string
		screen bool
	}{
		{
			dir:    "./screenshot/",
			screen: true,
		},
		{
			dir:    "./notscreenshot/",
			screen: false,
		},
	}
	for _, test := range table {
		files, err := ioutil.ReadDir(test.dir)
		if err != nil {
			t.Errorf("Error: %v, cannot read directory %v\n", err, test.dir)
			continue
		}

		for _, f := range files {
			name := path.Join(test.dir, f.Name())
			file, err := os.Open(name)
			if err != nil {
				t.Errorf("Error: %v, cannot open file %v\n", err, name)
				continue
			}
			defer func() {
				if err := file.Close(); err != nil {
					t.Errorf("Error: %v, cannot close file %v\n", err, name)
				}
			}()
			img, _, err := image.Decode(file)
			if err != nil {
				t.Errorf("Cannot decode image from file %q, error: %q", name, err)
				continue
			}

			res := Detect(img)
			if res != test.screen {
				t.Errorf("[FAIL] File: %v, screenshot: %v, should be: %v\n",
					name, res, test.screen)
			} else {
				t.Logf("[OK] File: %v, screenshot: %v\n", name, res)
			}
		}
	}
}

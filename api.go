package voicerssgo

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const c = MP3
const f = F48kHz16bitMono
const base = "http://api.voicerss.org/"

func Get(key string, hl string, src string, filename string) (*os.File, error) {
	url := fmt.Sprintf("%s?key=%s&hl=%s&src=%s&c=%s&f=%s", base, key, hl, src, c, f)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	resp, err := client.Do(req)

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	out, err := os.Create(filename + "." + c)
	if err != nil {
		return nil, err
	}

	defer out.Close()

	io.Copy(out, resp.Body)

	return out, nil
}

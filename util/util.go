package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

func init() {
	godotenv.Load()
}

type util struct{}

var Util = &util{}

func (u *util) OpenBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		log.Fatalf("Unsupported platform.")
	}

	return u.OnError(err)
}

func (u *util) OnError(err error) error {
	if err == nil {
		return nil
	}

	fmt.Printf("%s : ERR : %v\n", time.Now().Format("15:04"), err)
	return err
}

func (u *util) Log(param ...any) {
	s := time.Now().Format("15:04 :")
	for _, p := range param {
		if str, ok := p.(string); ok {
			s = s + " " + str
		} else {
			s = s + " " + u.ToJson(p)
		}
	}

	fmt.Println(s)
}

func (u *util) FromJson(input string, instance any) error {
	err := json.Unmarshal([]byte(input), instance)
	return u.OnError(err)
}

func (u *util) ToJson(param any) string {
	b, err := json.Marshal(param)
	if u.OnError(err) != nil {
		return ""
	}
	return string(b)
}

func (u *util) ToJsonReadable(param any) string {
	b, err := json.MarshalIndent(param, "", "  ")
	if u.OnError(err) != nil {
		return ""
	}
	return string(b)
}

func (u *util) ToFile(path string, data any) error {
	b, err := yaml.Marshal(data)
	if err != nil {
		return u.OnError(err)
	}
	err = os.WriteFile(path, b, 0644)
	return u.OnError(err)
}

func (u *util) FromFile(path string, instance any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return u.OnError(err)
	}

	err = yaml.Unmarshal(b, instance)
	return u.OnError(err)
}

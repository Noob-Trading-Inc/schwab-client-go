package util

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

func Ptr[T any](v T) *T {
	return &v
}

func OpenBrowser(url string) error {
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

	return OnError(err)
}

func OnError(err error) error {
	if err == nil {
		return nil
	}

	fmt.Printf("%s : ERR : %v\n", time.Now().Format("15:04"), err)
	return err
}

func Log(param ...any) {
	s := time.Now().Format("15:04 :")
	for _, p := range param {
		if str, ok := p.(string); ok {
			s = s + " " + str
		} else {
			s = s + " " + Serialize(p)
		}
	}

	fmt.Println(s)
}

func Logf(message string, param ...any) {
	message = fmt.Sprintf(message, param...)
	message = time.Now().Format("15:04 : ") + message
	fmt.Println(message)
}

func Deserialize(input string, instance any) error {
	err := json.Unmarshal([]byte(input), instance)
	return OnError(err)
}

func Serialize(param any) string {
	b, err := json.Marshal(param)
	if OnError(err) != nil {
		return ""
	}
	return string(b)
}

func SerializeReadable(param any) string {
	b, err := json.MarshalIndent(param, "", "  ")
	if OnError(err) != nil {
		return ""
	}
	return string(b)
}

func ToFile(path string, data any) error {
	b, err := yaml.Marshal(data)
	if err != nil {
		return OnError(err)
	}
	err = os.WriteFile(path, b, 0644)
	return OnError(err)
}

func FromFile(path string, instance any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return OnError(err)
	}

	err = yaml.Unmarshal(b, instance)
	return OnError(err)
}

func NewID() string {
	return fmt.Sprintf("%s-%v-%v", time.Now().Format("2006-01-02"), time.Now().YearDay(), time.Now().UnixNano())
}

func Substring(input string, start, length int) string {
	runes := []rune(input)
	return string(runes[0:7])
}

func Hash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func EpocToTime(epoc int64) time.Time {
	return time.Unix(epoc/1000, 0)
}

func RandomPick(options []string) string {
	if len(options) == 0 {
		return ""
	}

	return options[randInt(0, len(options))]
}

func Clone(instance interface{}, copy interface{}) error {
	b, _ := json.Marshal(instance)
	return json.Unmarshal(b, copy)
}

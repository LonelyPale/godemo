package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/LonelyPale/goutils/encoding/json"
)

// -tags jsoniter
func main() {
	//json.NewEncoder().Encode()
	bs, err := json.Marshal(Test{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

var (
	UseLocalTimeZone  = true                  //是否使用本地时区
	DefaultTimeFormat = "2006-01-02 15:04:05" //默认时间格式化字符串
)

type Test struct {
	Time  Time      `json:"time,omitempty"`
	Time2 time.Time `json:"time2,omitempty"`
}

type Time struct {
	time.Time
}

func (t Time) IsEmpty() bool {
	return t.IsZero()
}

func (t Time) IsNil() bool {
	return t.IsZero()
}

// MarshalJSON marshal json
func (t Time) MarshalJSON() ([]byte, error) {
	bs := make([]byte, 0, len(DefaultTimeFormat)+2)
	bs = append(bs, '"')
	bs = t.Time.AppendFormat(bs, DefaultTimeFormat)
	bs = append(bs, '"')
	return bs, nil
}

// UnmarshalJSON unmarshal json
func (t *Time) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == `""` {
		return errors.New("data is empty")
	}

	var now time.Time
	var err error
	switch UseLocalTimeZone {
	case true:
		now, err = time.ParseInLocation(`"`+DefaultTimeFormat+`"`, string(data), time.Local)
	case false:
		now, err = time.Parse(`"`+DefaultTimeFormat+`"`, string(data))
	}
	if err != nil {
		return err
	}

	t.Time = now
	return nil
}

/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
	"go/types"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type User struct {
	Id   int
	Name string
}

func (u *User) HashCode() int {
	return u.Id
}

func TestStream(t *testing.T) {
	red, _ := lru.New[int, int](33)
	blue, _ := lru.New[int, int](16)
	for i := 1; i < 33; i++ {
		red.Add(i, i)
	}
	for i := 1; i < 17; i++ {
		blue.Add(i, i)
	}

	var balls []*Balls
	data, err := os.ReadFile("lottery.json")
	if err != nil {
		return
	}
	if err := json.Unmarshal(data, &balls); err != nil {
		return
	}
	for _, ball := range balls {
		for _, rb := range ball.Red {
			red.Get(rb)
		}
		blue.Get(ball.Blue)
	}
	ss := red.Keys()[24:]
	sort.Ints(ss)
	t.Log(ss)
	t.Log(blue.Keys()[0])
}

type Lottery struct {
	Code int `json:"code"`
	Data struct {
		OpenCode string `json:"openCode"`
		Expect   string `json:"expect"`
	} `json:"data"`
}
type Balls struct {
	Red  []int `json:"red"`
	Blue int   `json:"blue"`
}

// getNumber
// {"code":1,"msg":"数据返回成功！","data":{"openCode":"01,07,11,12,22,28+05","code":"ssq","expect":"2022150","name":"双色球","time":"2022-12-29 21:15:00"}}
func getNumber(n int) *Balls {
	cli := http.DefaultClient
	url := `https://www.mxnzp.com/api/lottery/common/aim_lottery?expect=22%03d&code=ssq&app_id=rgihdrm0kslojqvm&app_secret=WnhrK251TWlUUThqaVFWbG5OeGQwdz09`
	uri := fmt.Sprintf(url, n)
	resp, err := cli.Get(uri)
	if err != nil {
		return nil
	}
	data, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var lottery Lottery
	if err := json.Unmarshal(data, &lottery); err != nil {
		return nil
	}
	if lottery.Code != 1 {
		return nil
	}
	code := strings.Split(lottery.Data.OpenCode, "+")
	redBalls := strings.Split(code[0], ",")
	blueBall, _ := strconv.Atoi(code[1])
	var balls Balls
	balls.Blue = blueBall
	for _, ball := range redBalls {
		red, _ := strconv.Atoi(ball)
		balls.Red = append(balls.Red, red)
	}
	return &balls
}

type Error struct {
	Msg string
}

func (a Error) Error() string {
	return "error"
}
func TestArgumentErrorUnwrapping(t *testing.T) {
	var err error = &types.ArgumentError{
		Index: 1,
		Err:   Error{Msg: "test"},
	}
	//e := &Error{
	//Msg: "err",
	//}
	var e interface{} = new(chan error)
	if !errors.As(err, e) {
		t.Logf("error %v does not wrap types.Error", err)
	}
	//{
	//}
	//if e.Msg != "test" {
	//	t.Errorf("e.Msg = %q, want %q", e.Msg, "test")
	//}
}

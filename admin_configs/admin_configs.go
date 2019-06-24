package admin_configs

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"time"
	"math/rand"
	)

var (
	SignBytes []byte
	VerifyBytes []byte
	PreLoginTokenTime int
	PostLoginTokenTime int
	EditNameMaximumTimes int
	err error

	NullTime	*time.Time
	NullString	*string
	GLOBAL_CODE int64
	)

func Init() {
	SignBytes, err = ioutil.ReadFile("conf/app.rsa")
	fatal("SignBytes: ", err)

	VerifyBytes, err = ioutil.ReadFile("conf/app.rsa.pub")
	fatal("VerifyBytes: ", err)

	PreLoginTokenTime, _ = beego.AppConfig.Int("preLoginTokenTime")
	PostLoginTokenTime, _ = beego.AppConfig.Int("postLoginTokenTime")
	EditNameMaximumTimes, _ = beego.AppConfig.Int("editNameMaximumTimes")

	NullTime = new(time.Time)
	NullString = new(string)

	GLOBAL_CODE, _ = beego.AppConfig.Int64("GLOBAL_CODE")
}

func GetSha512(s string) string {
    h := sha512.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum([]byte(s)))
}

func RandString(str_size int) string {
    alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_-+=<>?:;"
    var bytes = make([]byte, str_size)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b%byte(len(alphanum))]
    }
    return string(bytes)
}

func Int64ToInterface(t []int64) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
	    s[i] = v
	}
	return s
}

func fatal(tag string, err error) {
	if err != nil {
		log.Fatal(tag, err)
	}
}
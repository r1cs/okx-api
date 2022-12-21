package types

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"reflect"
)

type Op string

const (
	// 操作
	OpLogin       = "login"
	OpError       = "error"
	OpSubscribe   = "subscribe"
	OPUnSubscribe = "unsubscribe"
)

var ErrorOp = errors.New("err exist: ")

func HandleOpErr(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	op := v.FieldByName("Op")
	if op.IsZero() {
		panic("no Op filed")
	}
	if op.String() == OpError {
		return ErrorOp
	}
	//not exist
	return nil
}

type Params struct {
	Op   Op
	Args []LoginArgs
}

func NewLoginParams(args ...LoginArgs) *Params {
	return &Params{
		OpLogin,
		args,
	}
}

type LoginArgs struct {
	ApiKey     string
	Passphrase string
	Timestamp  string
	Sign       string
}

func NewLoginArgs(api, passphrase, timestamp, secret string) *LoginArgs {
	ret := &LoginArgs{
		api,
		passphrase,
		timestamp,
		"",
	}
	ret.Sign = base64.StdEncoding.EncodeToString(hmac.New(sha256.New, []byte(timestamp+"GET/user/self/verify"+secret)).Sum(nil))
	return ret
}

type LoginResults struct {
	Event Op       `json:op`
	Code  string   //错误码
	Msg   string   //错误消息
	Data  []string //失败会返回的apiKey
}

type SubscribeParam struct {
	Op   Op               `json:"op"`
	Args []*SubscribeArgs `json:"args"`
}

type SubscribeArgs struct {
	Channel string `json:"channel"`
	Instid  string `json:"instId"`
}

func NewTradeSubscribeArgs(tradePair string) *SubscribeArgs {
	return &SubscribeArgs{"trades", tradePair}
}

func NewSubscribeTrade(args ...*SubscribeArgs) *SubscribeParam {
	return &SubscribeParam{OpSubscribe, args}
}

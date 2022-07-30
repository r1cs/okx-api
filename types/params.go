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
	OpLogin = "login"
	OpError ="error"
)

var ErrorOp =errors.New("err exist: ")

func HandleOpErr(i interface{})error{
	v := reflect.ValueOf(i)
	if v.Kind()==reflect.Ptr{
		v= v.Elem()
	}
	if v.FieldByName("Op").String()==OpError{
		//find out error
		return ErrorOp
	}
	//not exist
	return nil
}

type LoginParams struct {
	Op Op
	Args []LoginArgs
}

type LoginResults struct {
	Event Op
	Code string //错误码
	Msg string //错误消息
	Data []string //失败会返回的apiKey
}

func NewLoginParams(args ...LoginArgs)*LoginParams{
	return &LoginParams{
		OpLogin,
		args,
	}
}

type LoginArgs struct {
	ApiKey  string
	Passphrase string
	Timestamp string
	Sign string
}

func NewLoginArgs(api,passphrase,timestamp,secret string )*LoginArgs{
	ret := &LoginArgs{
		api,
		passphrase,
		timestamp,
		"",
	}
	ret.Sign = base64.StdEncoding.EncodeToString(hmac.New(sha256.New,[]byte(timestamp+"GET/user/self/verify"+secret)).Sum(nil))
	return ret
}

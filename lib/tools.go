/* 
* @Author: dreamzml
* @Date:   2015-02-01 05:57:56
* @Last Modified by:   zoro
* @Last Modified time: 2015-02-01 06:14:59
*/

// Copyright 2013 wetalk authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package lib

import (
    "bytes"
    //"crypto/hmac"
    "crypto/md5"
    "crypto/rand"
    //"crypto/sha1"
    //"crypto/sha256"
    "encoding/hex"
    "fmt"
    //"hash"
    "math/big"
    //"net/url"
    "reflect"
    "strconv"
    //"time"

    //"github.com/astaxie/beego"

   // "github.com/beego/wetalk/setting"
)

func NumberEncode(number string, alphabet []byte) string {
    token := make([]byte, 0, 12)
    x, ok := new(big.Int).SetString(number, 10)
    if !ok {
        return ""
    }
    y := big.NewInt(int64(len(alphabet)))
    m := new(big.Int)
    for x.Sign() > 0 {
        x, m = x.DivMod(x, y, m)
        token = append(token, alphabet[int(m.Int64())])
    }
    return string(token)
}

func NumberDecode(token string, alphabet []byte) string {
    x := new(big.Int)
    y := big.NewInt(int64(len(alphabet)))
    z := new(big.Int)
    for i := len(token) - 1; i >= 0; i-- {
        v := bytes.IndexByte(alphabet, token[i])
        z.SetInt64(int64(v))
        x.Mul(x, y)
        x.Add(x, z)
    }
    return x.String()
}

// Random generate string
func buildhash(n int) string {
    const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, n)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b%byte(len(alphanum))]
    }
    return string(bytes)
}



// Encode string to md5 hex value
func EncodeMd5(str string) string {
    m := md5.New()
    m.Write([]byte(str))
    return hex.EncodeToString(m.Sum(nil))
}




// convert string to specify type

type StrTo string

func (f *StrTo) Set(v string) {
    if v != "" {
        *f = StrTo(v)
    } else {
        f.Clear()
    }
}

func (f *StrTo) Clear() {
    *f = StrTo(0x1E)
}

func (f StrTo) Exist() bool {
    return string(f) != string(0x1E)
}

func (f StrTo) Bool() (bool, error) {
    if f == "on" {
        return true, nil
    }
    return strconv.ParseBool(f.String())
}

func (f StrTo) Float32() (float32, error) {
    v, err := strconv.ParseFloat(f.String(), 32)
    return float32(v), err
}

func (f StrTo) Float64() (float64, error) {
    return strconv.ParseFloat(f.String(), 64)
}

func (f StrTo) Int() (int, error) {
    v, err := strconv.ParseInt(f.String(), 10, 32)
    return int(v), err
}

func (f StrTo) Int8() (int8, error) {
    v, err := strconv.ParseInt(f.String(), 10, 8)
    return int8(v), err
}

func (f StrTo) Int16() (int16, error) {
    v, err := strconv.ParseInt(f.String(), 10, 16)
    return int16(v), err
}

func (f StrTo) Int32() (int32, error) {
    v, err := strconv.ParseInt(f.String(), 10, 32)
    return int32(v), err
}

func (f StrTo) Int64() (int64, error) {
    v, err := strconv.ParseInt(f.String(), 10, 64)
    return int64(v), err
}

func (f StrTo) Uint() (uint, error) {
    v, err := strconv.ParseUint(f.String(), 10, 32)
    return uint(v), err
}

func (f StrTo) Uint8() (uint8, error) {
    v, err := strconv.ParseUint(f.String(), 10, 8)
    return uint8(v), err
}

func (f StrTo) Uint16() (uint16, error) {
    v, err := strconv.ParseUint(f.String(), 10, 16)
    return uint16(v), err
}

func (f StrTo) Uint32() (uint32, error) {
    v, err := strconv.ParseUint(f.String(), 10, 32)
    return uint32(v), err
}

func (f StrTo) Uint64() (uint64, error) {
    v, err := strconv.ParseUint(f.String(), 10, 64)
    return uint64(v), err
}

func (f StrTo) String() string {
    if f.Exist() {
        return string(f)
    }
    return ""
}

// convert any type to string
func ToStr(value interface{}, args ...int) (s string) {
    switch v := value.(type) {
    case bool:
        s = strconv.FormatBool(v)
    case float32:
        s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
    case float64:
        s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
    case int:
        s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
    case int8:
        s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
    case int16:
        s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
    case int32:
        s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
    case int64:
        s = strconv.FormatInt(v, argInt(args).Get(0, 10))
    case uint:
        s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
    case uint8:
        s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
    case uint16:
        s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
    case uint32:
        s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
    case uint64:
        s = strconv.FormatUint(v, argInt(args).Get(0, 10))
    case string:
        s = v
    case []byte:
        s = string(v)
    default:
        s = fmt.Sprintf("%v", v)
    }
    return s
}

// convert any numeric value to int64
func ToInt64(value interface{}) (d int64, err error) {
    val := reflect.ValueOf(value)
    switch value.(type) {
    case int, int8, int16, int32, int64:
        d = val.Int()
    case uint, uint8, uint16, uint32, uint64:
        d = int64(val.Uint())
    default:
        err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
    }
    return
}

type argString []string

func (a argString) Get(i int, args ...string) (r string) {
    if i >= 0 && i < len(a) {
        r = a[i]
    } else if len(args) > 0 {
        r = args[0]
    }
    return
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
    if i >= 0 && i < len(a) {
        r = a[i]
    }
    if len(args) > 0 {
        r = args[0]
    }
    return
}

type argAny []interface{}

func (a argAny) Get(i int, args ...interface{}) (r interface{}) {
    if i >= 0 && i < len(a) {
        r = a[i]
    }
    if len(args) > 0 {
        r = args[0]
    }
    return
}

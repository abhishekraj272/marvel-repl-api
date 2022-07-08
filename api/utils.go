package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GetTimeInNanoSec returns time in nano-second, for use case in api to prevent race condition
func GetTimeInNanoSec() int64 {
  now := time.Now()
  return now.UnixNano()
}

// GetAuthQueryParam is used to get basic auth params used in Marvel API
func (m *Marvel) GetAuthQueryParam() map[string]string {
  var queries = make(map[string]string);

  var ts string = strconv.FormatInt(GetTimeInNanoSec(), 10);

  var token strings.Builder;
  token.WriteString(ts)
  token.WriteString(m.privateKey)
  token.WriteString(m.publicKey)

  queries["apikey"] = m.publicKey;
  queries["ts"] = ts;
  hash := md5.Sum([]byte(token.String()));
  queries["hash"] = hex.EncodeToString(hash[:]);

  return queries;
}

func HandleReqFail(code int) {
  if (code >= 500) {
    fmt.Println("Server Error")
  }
  if (code >= 400) {
    fmt.Println("Bad Request or API Limit reached")
  }
  if (code == 0) {
    fmt.Println("Public or Private Key Invalid")
  }
  fmt.Println("Please try again");
}
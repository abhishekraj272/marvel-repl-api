package api

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

// GetTimeInNanoSec returns time in nano-second, for use case in api to prevent race condition
func GetTimeInNanoSec() int64 {
  now := time.Now()
  return now.UnixNano()
}

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
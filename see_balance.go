package main
import (
	"fmt"
	"net/http"
	"crypto/hmac"
	"hash"
	"crypto/sha256"
	"time"
	"net/http/httputil"
)

// Deus se for da tua vontade que esse projeto dê certo.

const (
	access_ID = "XXXXXXXXXXXXXXXXXXXX"
	secret_Key = "XXXXXXXXXXXXXXXXXXXX"
	base_URL = "https://api.coinex.com/v2"
)


/*
funcoes que eu vou precisar:
Assets: Balance
Futures: Market
Futures: Order
Futures: Position
*/

// base url: https://api.coinex.com/v2


// funcao testada, comparada com a saida em python oficial e aprovada.
func SignedStr_Gen(method, request_path, body, timestamp string) string {
	var prepared_str, signed_str string
	if ( body == "" ) {
		prepared_str = method + request_path + timestamp
	} else {
		prepared_str = method + request_path + body + timestamp
	}
	var hmac_function hash.Hash
	hmac_function = hmac.New(sha256.New, []byte(secret_Key))
	hmac_function.Write([]byte(prepared_str))
	signed_str = fmt.Sprintf("%x", hmac_function.Sum(nil))
	fmt.Println(signed_str)
	return signed_str
}

func BuildAuthRequest(signed_str, method, request_path, timestamp string) {
	client := &http.Client {
		CheckRedirect: nil,
	}
	req, err := http.NewRequest(method, base_URL + request_path, nil)
	req.Header.Add("X-COINEX-KEY", access_ID)
	req.Header.Add("X-COINEX-SIGN", signed_str)
	req.Header.Add("X-COINEX-TIMESTAMP", timestamp)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := httputil.DumpResponse(resp, true)
	_ = err
	fmt.Println(string(body))
	return
}

func GenerateTimestamp() string {
	var umt time.Time
	var timestamp string
	umt = time.Now()
	timestamp = fmt.Sprintf("%d", umt.Unix()*1e3 + int64(umt.Nanosecond())/1e6)
	fmt.Println(timestamp)
	return timestamp
}

func SeeBalance() {
	var signed_str string
	var timestamp string
	timestamp = GenerateTimestamp()
	signed_str = SignedStr_Gen("GET", "/v2/assets/futures/balance", "", timestamp)
	BuildAuthRequest( signed_str, "GET", "/assets/futures/balance", timestamp)
}

func main() {
	SeeBalance()
}

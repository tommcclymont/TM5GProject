package logger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/felixge/httpsnoop"
)

// wrapper to log important information about NF interactions
func LogHandler(h http.Handler, nf string) http.Handler {
	logh := func(rw http.ResponseWriter, r *http.Request) {

		// read request body and preserve it for request handlers
		bd, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bd))

		sn := httpsnoop.CaptureMetrics(h, rw, r)
		var reqbody map[string]interface{}
		var size int

		if (r.Method == "POST" || r.Method == "PUT") && r.Body != http.NoBody {
			// decode request body data to log
			err = json.Unmarshal(bd, &reqbody)
			if err != nil {
				print(bd)
				panic(err)
			}
			// size of request body
			size = len(bd)
		} else {
			size = 0
			reqbody = nil
		}

		info := &LogInfo{
			Method:   r.Method,
			Uri:      r.URL.String(),
			IpAddr:   r.RemoteAddr,
			ReqSize:  size,
			ReqBody:  reqbody,
			RspCode:  sn.Code,
			RspSize:  sn.Written,
			Duration: sn.Duration,
		}

		// store information in json file
		data, _ := json.MarshalIndent(info, "", " ")
		filename := nf + ".json"
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		if _, err = file.Write(data); err != nil {
			panic(err)
		}

		comma := []byte(",")
		if err != nil {
			panic(err)
		}
		if _, err = file.Write(comma); err != nil {
			panic(err)
		}
	}

	return http.HandlerFunc(logh)
}

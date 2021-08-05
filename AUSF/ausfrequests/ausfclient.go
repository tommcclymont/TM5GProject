package ausfrequests

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// structure for the client that will communicate with service
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// create new client
func NewClient(URL string) *Client {
	ausfAddr, err := net.ResolveIPAddr("ip", "127.0.0.80")
	if err != nil {
		panic(err)
	}

	amfTCPAddr := net.TCPAddr{
		IP: ausfAddr.IP,
		//	Port: 9090,
	}

	return &Client{
		BaseURL: URL,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{RootCAs: loadCA("TLS/ca.crt")},
				Proxy:           http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					LocalAddr: &amfTCPAddr,
					Timeout:   time.Minute,
				}).DialContext,
			},
		},
	}
}

// certificates for TLS config
func loadCA(caFile string) *x509.CertPool {
	certPool := x509.NewCertPool()

	if ca, err := ioutil.ReadFile(caFile); err != nil {
		log.Fatal("ReadFile: ", err)
	} else {
		certPool.AppendCertsFromPEM(ca)
	}

	return certPool
}

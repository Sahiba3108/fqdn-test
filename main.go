package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func main() {
	cert := `-----BEGIN CERTIFICATE-----
MIIDbDCCAlSgAwIBAgIJAMaq70Fuw3IsMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
MjMwNTE2MTMzMTI5WhcNMjYwNTE1MTMzMTI5WjBlMQswCQYDVQQGEwJVUzETMBEG
A1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl
ZiBTb2Z0d2FyZSBJbmMxEzARBgNVBAMMCmNoZWZjbGllbnQwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDUo2MH1ktyTJeVm50zO7PH52TrqOI811q6z7m4
JYUONpJOBB5KyxEmPTDtQEIHStjxeAbXZ4DsA8FF0S8hx+do8hMufEePAuCXQAKn
YWZElD3ziGH4PEij4pHpQogfifOA0VJuefbO0QuLWxg1pgXnoPXmKWEa4ahgSx0r
+Z6X6EybkKxVqbb/cyQN7fmMUbhWvINjz8NuR3Jgc0F93+a/IxGa66vQB1T7VpYO
ij4wrj5/E+izGUoz8+SGTOTzkNpHpnwTScbN3TAI5073AycKH4YQSPLpuNeZGlay
+Nguml7suG/QVkbJbkAlf7VUF75lDbXCWweWlAj5qRLbhc5XAgMBAAGjITAfMB0G
A1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATANBgkqhkiG9w0BAQsFAAOCAQEA
iSjdTwlqApi/8yWfFVO9XmVs0FEN4g5HEcKz/rD0GZBesR+6ox1TU+YvGWTP+5VS
ZWwLIc+0OahoRanDcHPMd5chYUwpwJxxXweLvbvRN2iDPXYXas6/hjaA8XBb/Avt
5Ius1e65xYRSF8Ri3OVbScZygJEVWsf4qizt+FY+du84QbmDreTXLknI8OYaDPyL
Y2xvgMxLzo9VrUxH46y1xlJEfWFXkYqRkppJJtbcYiE2GHBxL613Y3i9zEyGaseV
hb+ma8V0fEIkFYjyADJ5HDE/lY/9OF87XiuqCohEfBQJDvLyC3JCdeS3HokMPW/6
geD+u8jHvTft81i6zx14UA==
-----END CERTIFICATE-----`
	key := `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDUo2MH1ktyTJeV
m50zO7PH52TrqOI811q6z7m4JYUONpJOBB5KyxEmPTDtQEIHStjxeAbXZ4DsA8FF
0S8hx+do8hMufEePAuCXQAKnYWZElD3ziGH4PEij4pHpQogfifOA0VJuefbO0QuL
Wxg1pgXnoPXmKWEa4ahgSx0r+Z6X6EybkKxVqbb/cyQN7fmMUbhWvINjz8NuR3Jg
c0F93+a/IxGa66vQB1T7VpYOij4wrj5/E+izGUoz8+SGTOTzkNpHpnwTScbN3TAI
5073AycKH4YQSPLpuNeZGlay+Nguml7suG/QVkbJbkAlf7VUF75lDbXCWweWlAj5
qRLbhc5XAgMBAAECgf967fpSjZEQgzcfRteaxukoqJ2vn/MrOVU6qtyzmP1TOjla
4QaiNsOzV6zxfDSwApVpA/rwjOX7jDAwLL8SyH8ALPLp7O9xvjgZOOhP2/9DpnAq
Wq5I1JEqLdItCEmp8hbUYi9Frz0AU9pYtlFIsrK3NlbyuPlqkQrGoKY8xhLpvlS+
khKPpmpSlkmD2FDL6JH4OhJJIsoSVJIZW/LxrOIc5uhJjKJHbaD4Bubbhrlf2GCc
sCZxVe6tBwQKZDoVmZt6s6tfJY9G4fzGxiMzXYsu3QGytVIhh2aTctpFwBLw+/35
M5a0rSsLYVZfwTmb8VGjMcOMApTsIye2i33YlqECgYEA7Orw7Vvrtn5Qi7UcxT1J
41BYdKZrCAvIOylwmi/nSKdJnB3HWMzvc/igl4txnpjnQ1zQPWhxmEn6h4Ilu2UB
T8qQ+4JhcwXjqVnQ0WYvwb7cs+7HhcxW7dswiOPwYOn1TdqiXXBvQJj0bq0BESxc
DlLXmKNXnt4b24vyriH1c7sCgYEA5cPSOY76lU74Bt+B7gKI7XO3A3mTTg2Ie6lL
OBPPZxuuLj4U5Jszts9qHB4IDLzYtSdXG39vseTawvPhgq3+xuRVPrWMxTXpgFq+
bE/0s+mxMgZSsNou1M290cAnAHM/ILYTXqSiDZdPpOw6JUFjmJTIclNL1c9xH9l5
f7ls8BUCgYEAvo8GNFRwFjwpOwX02yy9xol0bHcCtdkMN5HQUSRgqk5r5ZcKJYHI
xAXh3aK6Q6+Gq+b2U74zSCOQz/e9s6m9UVmEHdgz25mMEUX9sL+5f+Otj0hq0VKX
RP+9XQ3B5aGovfMyD2gFA/dC+9ZJem+sp5S4drxDGLRbwc4h5y/HRksCgYAGzvSt
y/TU89AZtAnPmbehAarMqY+Z4oDG2U9nS/77WTCiIijQVWEE68SDLMikR/xL8ex8
3hvbM6Lf4AYwryM7lYyJHSwcs+pFaWr3Hq9rnWMozlMo9m4o1CfpzT3an4+hUoPk
bBB5QqCTjIR0kil3XgFKkflPTqRVg0Y8aKw3uQKBgQC3bxcHpFXG/IDTVOYR6RGo
3ekWRVJ/0zQ3KsBm0LbPnsAY7DRVmJ/oHFtOIhW8RSNAAnhpDUzSE2f8TAfNPSty
6I2mJngP0lfzBUziQDlJzELohqHqmKF+NOb3+u/j28jB4ovnugEgNjlTti2YiWiB
bTiQ3RlFIKLTO3We2pDh+g==
-----END PRIVATE KEY-----`

	err := StartHTTPSServer(1234, cert, key)
	if err != nil {
		fmt.Printf("temp: %v\n", err)
		return
	}
}

func StartHTTPSServer(port int, cert string, key string) error {

	// Load the TLS certificate and private key
	tlsCert, err := tls.X509KeyPair([]byte(cert), []byte(key))
	if err != nil {
		fmt.Printf("Certificate error: %v\n", err)
		return err
	}

	// Create the TLS configuration for the server
	config := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		MinVersion:   tls.VersionTLS13,
	}

	// Create the HTTPS server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		TLSConfig: config,
	}

	// Handle response
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok\n"))
	})

	// Start the HTTPS server
	go func() {
		err = server.ListenAndServeTLS("", "")
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting HTTPS server: ", err)
			return
		}
		fmt.Printf("HTTPS server started on port %d\n", port)
	}()
	time.Sleep(30 * time.Second)
	return nil
}

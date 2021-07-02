package main
 
import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)
 
 // check the client to the server
func main() {
	 // CertPool represents a set of certificates / certificate pool.
	 // Create a CertPool
	pool := x509.NewCertPool()
	caCertPath := "/GO/src/go-server/ca.crt"
	 // Call ca.crt file
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	 // parse certificate
	pool.AppendCertsFromPEM(caCrt)
 
	tr := &http.Transport{
		 //// the server certificate from the non-leaf pass over, intermediate certificate added to the pool, using the set of intermediate certificate and a root certificate to verify the certificate leaves.
		TLSClientConfig: &tls.Config{RootCAs: pool},
		
		 // TLSClientConfig: & tls.Config {InsecureSkipVerify: true}, // InsecureSkipVerify used to control whether the client certificate and the server host name. If set to true, // // does not check whether the certificate and the certificate host name and server host name consistent.
		
		
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
 
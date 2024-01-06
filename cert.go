package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
)

func getCertificate(host string, port int) (*x509.Certificate, error) {
	config := &tls.Config{ServerName: host}
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	state := conn.ConnectionState()
	if len(state.PeerCertificates) > 0 {
		return state.PeerCertificates[0], nil
	}

	return nil, fmt.Errorf("no peer certificates found")
}

func convertCertificateToJSON(cert *x509.Certificate) (string, error) {
	certJSON, err := json.Marshal(cert)
	if err != nil {
		return "", err
	}
	return string(certJSON), nil
}

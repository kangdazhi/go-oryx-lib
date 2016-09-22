// The MIT License (MIT)
//
// Copyright (c) 2013-2016 Oryx(ossrs)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package https

import (
	"crypto/tls"
	"runtime"
	"strings"
	"strconv"
	"fmt"
)

// The https manager which provides the certificate.
type Manager interface {
	GetCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error)
}

// The cert is sign by ourself.
type selfSignManager struct {
	cert *tls.Certificate
	certFile string
	keyFile  string
}

func NewSelfSignManager(certFile, keyFile string) (m Manager, err error) {
	// check golang version, must 1.6+
	version := strings.Trim(runtime.Version(), "go")
	if versions := strings.Split(version, "."); len(versions) < 1 {
		return nil, fmt.Errorf("invalid version=%v", version)
	} else if minor,err := strconv.Atoi(versions[1]); err != nil {
		return nil, fmt.Errorf("invalid version=%v, err=%v", version, err)
	} else if minor < 6 {
		return nil, fmt.Errorf("requires golang 1.6+, version=%v, minor=%v", version, minor)
	}

	return &selfSignManager{certFile: certFile, keyFile: keyFile},nil
}

func (v *selfSignManager) GetCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	if v.cert != nil {
		return v.cert,nil
	}

	cert, err := tls.LoadX509KeyPair(v.certFile, v.keyFile)
	if err != nil {
		return nil,err
	}

	// cache the cert.
	v.cert = &cert

	return &cert, err
}

// The cert is sign by letsencrypt
type letsencryptManager struct {
}

// Register the email to letsencrypt, cache the certs in cacheFile, set allow hosts.
// @remark set hosts to nil when allow all request hosts, but maybe attack.
// @remark set email to nil to not regiester, use empty email to request cert from letsencrypt.
// @remark set cacheFile to nil to not cache the info and certs.
func NewLetsencryptManager(email string, hosts []string, cacheFile string) (m Manager, err error) {
	return &letsencryptManager{},nil
}

func (v *letsencryptManager) GetCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return nil,nil
}
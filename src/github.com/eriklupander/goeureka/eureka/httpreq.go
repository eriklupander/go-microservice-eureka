/**
The MIT License (MIT)

Copyright (c) 2016 ErikL

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package eureka

import (
    "log"
    "net/http"
    "strings"
    "crypto/tls"
   // "strconv"
)

// Accepts a Httpaction and a one-way channel to write the results to.
func DoHttpRequest(httpAction HttpAction) bool {
    req := buildHttpRequest(httpAction)

    var DefaultTransport http.RoundTripper = &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    resp, err := DefaultTransport.RoundTrip(req)
    if err != nil {
        log.Printf("HTTP request failed: %s", err)
        return false
    } else {
        return true
        defer resp.Body.Close()
    }
    return false
}

func buildHttpRequest(httpAction HttpAction) *http.Request {
    var req *http.Request
    var err error
    if httpAction.Body != "" {
        reader := strings.NewReader(httpAction.Body)
        req, err = http.NewRequest(httpAction.Method, httpAction.Url, reader)
    } else if httpAction.Template != "" {
        reader := strings.NewReader(httpAction.Template)
        req, err = http.NewRequest(httpAction.Method, httpAction.Url, reader)
    } else {
        req, err = http.NewRequest(httpAction.Method, httpAction.Url, nil)
    }
    if err != nil {
        log.Fatal(err)
    }

    // Add headers
    req.Header.Add("Accept", httpAction.Accept)
    if (httpAction.ContentType != "") {
        req.Header.Add("Content-Type", httpAction.ContentType)
    }
    return req
}


/**
 * Trims leading and trailing byte r from string s
 */
func trimChar(s string, r byte) string {
    sz := len(s)

    if sz > 0 && s[sz-1] == r {
        s = s[:sz-1]
    }
    sz = len(s)
    if sz > 0 && s[0] == r {
        s = s[1:sz]
    }
    return s
}
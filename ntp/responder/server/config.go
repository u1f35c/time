/*
Copyright (c) Facebook, Inc. and its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"fmt"
	"net"
	"strings"
)

// DefaultServerIPs is a default list of IPs server will bind to if nothing else is specified
var DefaultServerIPs = MultiIPs{net.ParseIP("127.0.0.1"), net.ParseIP("::1")}

// ListenConfig is a wrapper around multiple IPs and Port to bind to
type ListenConfig struct {
	IPs            MultiIPs
	Port           int
	ShouldAnnounce bool
	Iface          string
}

// MultiIPs is a wrapper allowing to set multiple IPs with flag parser
type MultiIPs []net.IP

// Set adds check to the runlist
func (m *MultiIPs) Set(ipaddr string) error {
	ip := net.ParseIP(ipaddr)
	if ip == nil {
		return fmt.Errorf("invalid ip address %s", ip)
	}
	*m = append([]net.IP(*m), ip)
	return nil
}

// String returns joined list of checks
func (m *MultiIPs) String() string {
	var ips []string
	for _, ip := range *m {
		ips = append(ips, ip.String())
	}
	return strings.Join(ips, ", ")
}

// SetDefault adds all checks to the runlist
func (m *MultiIPs) SetDefault() {
	if len(*m) != 0 {
		return
	}

	*m = DefaultServerIPs
}

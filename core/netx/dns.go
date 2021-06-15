package netx

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
	"strings"
)

type (
	dnsHeader struct {
		Id,
		Bits,
		QdCount,
		AnCount,
		NsCount,
		ArCount uint16
	}
	dnsQuery struct {
		QuestionType  uint16
		QuestionClass uint16
	}
	bits struct {
		QR,
		OperationCode,
		AuthoritativeAnswer,
		Truncation,
		RecursionDesired,
		RecursionAvailable,
		ResponseCode uint16
	}
	IP string
)

func newDNSHeader() dnsHeader {
	return dnsHeader{
		Id:      0x0010,
		QdCount: 1,
	}
}

func newDNSQuery() dnsQuery {
	return dnsQuery{
		QuestionType:  1,
		QuestionClass: 1,
	}
}

func (h *dnsHeader) SetBits(bits bits) {
	h.Bits = bits.QR<<15 + bits.OperationCode<<11 + bits.AuthoritativeAnswer<<10 + bits.Truncation<<9 + bits.RecursionDesired<<8 + bits.RecursionAvailable<<7 + bits.ResponseCode
}

func QueryIP(dnsServer, domain string) (ip IP, err error) {
	requestHeader := newDNSHeader()
	requestHeader.SetBits(bits{RecursionDesired: 1})

	conn, err := net.Dial("udp", dnsServer)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	var buffer bytes.Buffer
	if err := binary.Write(&buffer, binary.BigEndian, requestHeader); err != nil {
		return "", err
	}
	domainName, err := ParseDomainName(domain)
	if err != nil {
		return "", err
	}
	if err := binary.Write(&buffer, binary.BigEndian, domainName); err != nil {
		return "", err
	}
	if err := binary.Write(&buffer, binary.BigEndian, newDNSQuery()); err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return "", err
	}

	var n int
	n, err = conn.Read(buf)
	ipStr := strconv.Itoa(int(buf[n-4])) + "." + strconv.Itoa(int(buf[n-3])) + "." + strconv.Itoa(int(buf[n-2])) + "." + strconv.Itoa(int(buf[n-1]))

	return IP(ipStr), nil
}

func ParseDomainName(domain string) ([]byte, error) {
	var buffer bytes.Buffer
	segments := strings.Split(domain, ".")

	for _, seg := range segments {
		if err := binary.Write(&buffer, binary.BigEndian, byte(len(seg))); err != nil {
			return nil, err
		}
		if err := binary.Write(&buffer, binary.BigEndian, []byte(seg)); err != nil {
			return nil, err
		}
	}
	if err := binary.Write(&buffer, binary.BigEndian, byte(0x00)); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (ip IP) Ipv4() (*net.IPAddr, error) {
	return net.ResolveIPAddr("ip4", string(ip))
}
func (ip IP) Ipv6() (*net.IPAddr, error) {
	return net.ResolveIPAddr("ip6", string(ip))
}

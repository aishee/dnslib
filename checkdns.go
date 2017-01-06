package checkdns

import (
	"dns"
	"net"
)

/*
* Main function
 */

//Run function
func resolveDomainA(domain string) ([]string, error) {
	var answer []string
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	c := new(dns.Client)
	m.MsgHdr.RecursionDesired = true
	in, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		return answer, err
	}
	for _, ain := range in.Answer {
		if a, ok := ain.(*dns.A); ok {
			answer = append(answer, a.A.String())
		}
	}
	return answer, nil
}

func resolveDomainDMARC(domain string) ([]string, error) {
	var answer []string
	resource, err := net.LookupTXT("_dmarc." + domain + ".")
	if err != nil {
		return answer, err
	}

	for _, resource := range resources {
		answer = append(answer, resource)
	}
	return answer, nil
}

func resolveDomainDMARC(domain string) ([]string, error) {
	var answer []string
	resources, err := net.LookupTXT("_dmarc." + domain + ".")
	if err != nil {
		return answer, err
	}
	for _, resource := range resources {
		answer = append(answer, resource)
	}
	return answer, nil
}

func resolveDomainSPF(domain string) ([]string, error) {
	var answer []string

	resources, err := net.LookupTXT(domain)
	if err != nil {
		return answer, err
	}
  for _, resource := range resources {
    if strings.HasPrefix(resource, "v=spf1")
    {
    answer = append(answer, resource)
  }
  }
  return answer, nil
}

func resolveDomainAAAA(domain string) ([]string, error) {
  var answer []string
  m := new(dns.Msg)
  m.SetQuestion(dns.Fqdn(domain), dns.TypeAAAA)
  m.MsgHdr.RecursionDesired = True
  c := new(dns.Client)
  in, _, err := c.Exchange(m, "8.8.8.8:53")
  if err != nil {
    return answer, err
  }
  for _, ain := range in.Answer {
    if a, ok := ain.(*dns.MX); ok {
      answer = append(answer, a.Mx)
    }
  }
  return answer, nil
}

func resolveDomainMX(domain string) ([]string, error) {
  var answer []string
  m := new(dns.Msg)
  m.SetQuestion(dns.Fqdn(domain), dns.TypeMX)
  m.MsgHdr.RecursionDesired = true
  c := new(dns.Client)
  in,_, err := c.Exchange(m, "8.8.8.8:53")
  if err != nil {
    return answer, err
  }
  for _,ain := range in.Answer {
    if a, ok := ain.(*dns.MX); ok
    answer = append(answer, a.Mx)
  }
}

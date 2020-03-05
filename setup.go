package nacos_coredns_plugin

import (
	"fmt"
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/plugin/forward"
	"github.com/coredns/coredns/plugin/pkg/parse"
	"github.com/coredns/coredns/plugin/pkg/transport"
	"os"
	"strconv"
	"strings"
)

func init(){

	caddy.RegisterPlugin("nacos",caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})

	fmt.Println("register nacos plugin")
}

func setup(c *caddy.Controller) error{
	fmt.Println("setup nacos plugin")
	os.Stderr = os.Stdout



	return nil

}

func NacosParse(c *caddy.Controller) (*Nacos, error) {
	fmt.Println("init nacos plugin...")
	nacosImpl := Nacos{}
	var servers = make([]string, 0)
	serverPort := 8848
	for c.Next() {
		nacosImpl.Zones = c.RemainingArgs()

		if c.NextBlock() {
			for {
				switch v :=c.Val();v {
				case "nacos_server":
					servers = strings.Split(c.RemainingArgs()[0], ",")
					/* it is a noop now */
				case "nacos_server_port":
					port, err := strconv.Atoi(c.RemainingArgs()[0])
					if err != nil {
						serverPort = port
					}
				case "cache_ttl":
					ttl, err := strconv.Atoi(c.RemainingArgs()[0])
					if err != nil {
						DNSTTL = uint32(ttl)
					}
				case "upstream":
					args := c.RemainingArgs()
					if len(args) == 0 {
						return &Nacos{}, c.ArgErr()
					}
					ups, err := parse.HostPortOrFile(args...)
					if err != nil {
						return &Nacos{}, err
					}

					var ups1 []string;
					proxys:=forward.New()
					for _, host := range ups {
						if strings.Contains(host, "127.0.0.1"){
							continue
						} else {
							proxys.SetProxy(forward.NewProxy(host,transport.DNS))
							ups1 = append(ups1, host)
							break
						}
					}
					fmt.Println("upstreams: ", ups1)

					nacosImpl.Proxys = *proxys
				case "cache_dir":
					CachePath = c.RemainingArgs()[0]
				case "log_path":
					LogPath = c.RemainingArgs()[0]
				default:
					if c.Val() != "}" {
						return &Nacos{}, c.Errf("unknown property '%s'", c.Val())
					}
				}

				if !c.Next() {
					break
				}
			}

		}


		client := NewNacosClient(servers, serverPort)
		nacosImpl.NacosClientImpl = client
		nacosImpl.DNSCache = NewConcurrentMap()

		return &nacosImpl, nil
	}
	return &Nacos{}, nil
}
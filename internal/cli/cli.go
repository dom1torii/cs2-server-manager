package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dom1torii/cs2-server-manager/internal/api"
	"github.com/dom1torii/cs2-server-manager/internal/config"
	"github.com/dom1torii/cs2-server-manager/internal/fs"
	"github.com/dom1torii/cs2-server-manager/internal/ips"
	"github.com/dom1torii/cs2-server-manager/internal/platform/firewall"
)

func IsCLIMode(cfg *config.Config) bool {
	 return cfg.ListRelays ||
	        len(cfg.SelectRelays) > 0 ||
	        cfg.BlockRelays ||
	        cfg.UnBlockRelays ||
	        cfg.ToBlockCount ||
	        cfg.BlockedCount
}

func HandleFlags(cfg *config.Config) {
	response, err := api.FetchRelays()
	if err != nil {
		log.Fatalln("Failed to fetch relays: ", err)
	}

	if cfg.ListRelays {
		keys := make([]string, 0, len(response.Pops))
		for key, pop := range response.Pops {
	  	keys = append(keys, key+ " - " + pop.Desc)
		}

		fmt.Println("Available relays:\n", strings.Join(keys, "\n"))
	}
	if len(cfg.SelectRelays) > 0 {
		// write relays that are not selected into ips file
		var ipList []string

		selected := make(map[string]struct{})
		for _, s := range cfg.SelectRelays {
		  selected[s] = struct{}{}
		}

		for popName, pop := range response.Pops {
		  if _, isSelected := selected[popName]; !isSelected {
		    for _, relay := range pop.Relays {
		      ipList = append(ipList, relay.Ipv4)
		    }
		  }
		}
		ips.WriteIpsToFile(ipList, cfg)
	}
	if cfg.BlockRelays {
		firewall.BlockIps(cfg, nil)
	}
	if cfg.UnBlockRelays {
		firewall.UnBlockIps(nil)
	}
	if cfg.ToBlockCount {
		fmt.Println(fs.GetFileLineCount(cfg.IpsPath))
	}
	if cfg.BlockedCount {
		fmt.Println(firewall.GetBlockedIpCount())
	}
}

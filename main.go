package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chzyer/readline"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/manifoldco/promptui"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"log"
	"net"
	"os"
)

const menuConnect = "Connect To Other Peer"
const menuShowPeers = "Show Peers"
const menuExit = "Exit"

func main() {
	var iface string
	flag.StringVar(&iface, "iface", "br0", "interface name")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	priv, _, err := crypto.GenerateKeyPair(crypto.Ed25519, 0)
	if err != nil {
		log.Fatalln(err)
	}

	localAddrs, err := GetLinkLocalAddresses(iface)
	if err != nil {
		log.Fatalln(err)
	}

	var listenAddrs []ma.Multiaddr
	for _, addr := range localAddrs {
		listenAddrs = append(listenAddrs, ma.StringCast(addr.String()+"/tcp/0"))
	}
	listenAddrs = append(listenAddrs, ma.StringCast("/ip4/127.0.0.1/tcp/0"))

	h, err := libp2p.New(
		libp2p.Identity(priv),
		libp2p.ListenAddrs(listenAddrs...),
		libp2p.Transport(tcp.NewTCPTransport),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer h.Close()

	_ = ctx
	_ = priv
	_ = localAddrs

	ps := h.Peerstore()

	c := &readline.Config{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
	}
	c.Init()
	c.Stdin = readline.NewCancelableStdin(c.Stdin)

	for {
		prompt := promptui.Select{
			Label: "Command",
			Items: []string{
				menuShowPeers,
				menuConnect,
				menuExit,
			},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Println(err)
			continue
		}

		if result == menuConnect {
			remoteAddrText, err := readline.Line("Enter remote peer address: ")
			if err != nil {
				log.Println(err)
				continue
			}

			addr, err := ma.NewMultiaddr(remoteAddrText)
			if err != nil {
				log.Println(err)
				continue
			}

			addrInfo, err := peer.AddrInfoFromP2pAddr(addr)
			if err != nil {
				log.Println(err)
				continue
			}

			err = h.Connect(context.Background(), *addrInfo)
			if err != nil {
				log.Println("CONNECT FAILED: ", err)
				continue
			} else {
				log.Println("CONNECT SUCCESS")
			}

			log.Printf("AAAA [%s]", remoteAddrText)
		} else if result == menuShowPeers {
			fmt.Println("MY: ", h.ID().String(), " :: ")
			fmt.Println("LOCAL LISTEN ADDRESSES: ")
			for _, multiaddr := range h.Network().ListenAddresses() {
				fmt.Printf("    %s\n", multiaddr.String())
			}

			fmt.Println("==============================================")

			for _, id := range ps.Peers() {
				info := ps.PeerInfo(id)
				fmt.Println("info: ", info.String())
				for _, m := range info.Addrs {
					fmt.Println("    ", m.String())
				}
			}
		} else {
			break
		}
	}
}

func GetLinkLocalAddresses(ifname string) ([]ma.Multiaddr, error) {
	var maddrs []ma.Multiaddr

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		if ifname != "" && ifname != iface.Name {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, a := range addrs {
			maAddr, err := manet.FromNetAddr(a)
			if err != nil {
				continue
			}
			if manet.IsIP6LinkLocal(maAddr) {
				maAddr = ma.StringCast("/ip6zone/" + iface.Name + maAddr.String())
				maddrs = append(maddrs, maAddr)
			}
		}
	}
	return maddrs, nil
}

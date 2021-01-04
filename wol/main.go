package wol

import (
	"fmt"
	"net"
)

// wol唤醒
func Wake(MacAddr string, Bdi string, Bip string, UPort string) {
	if len(MacAddr) <= 0 {
		fmt.Println("No mac address specified to wake command")
	}
	bcastInterface := ""
	BroadcastIP := "255.255.255.255"
	UDPPort := "9"
	if Bdi != "" {
		bcastInterface = Bdi
	}
	if Bip != "" {
		BroadcastIP = Bip
	}
	if UPort != "" {
		UDPPort = UPort
	}
	var localAddr *net.UDPAddr
	if bcastInterface != "" {
		var err error
		localAddr, err = IpFromInterface(bcastInterface)
		if err != nil {
			fmt.Println(err)
		}
	}
	bcastAddr := fmt.Sprintf("%s:%s", BroadcastIP, UDPPort)
	udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
	if err != nil {
		fmt.Println(err)
	}
	mp, err := WakeWol(MacAddr)
	if err != nil {
		fmt.Println(err)
	}
	bs, err := mp.Marshal()
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialUDP("udp", localAddr, udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Printf("Attempting to send a magic packet to MAC %s\n", MacAddr)
	fmt.Printf("... Broadcasting to: %s \n", bcastAddr)
	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent) \n", n)
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Magic packet sent successfully to %s \n", MacAddr)
}

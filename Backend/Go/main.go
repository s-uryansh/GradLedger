package main

import (
	"backend_go/services"
)

func main() {
	services.Connect("wss://rpc.buildbear.io/positive-husk-962d2b1b")

	services.LoadContracts(
		"0xC790576a74a1087D58c1851038b547B0D8E637a6",
		"0x3bc1018E4552Cd56dBe577765578b5dd0CC81aBf",
		"0x056Fc9EDDf99D97Dd4147Ae44Dd4E48dd259D61c",
		"0xF219b8b85a91Cc5D99B280c547E8c340B5a5A753",
		"0x5286264cf7139AEF8F9B1b01f2c66627f2E12784",
	)

	services.WatchUserEvents()
	select {}
}

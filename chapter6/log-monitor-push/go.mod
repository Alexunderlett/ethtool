module log-monitor

go 1.14

require (
	contract v0.0.0
    hubwiz.com/ethtool v0.0.0
	github.com/ethereum/go-ethereum v1.9.11 // indirect
)

replace contract => ../contract
replace hubwiz.com/ethtool =>  ../../ethtool

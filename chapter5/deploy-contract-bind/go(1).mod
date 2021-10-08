module deploy-contract-bind

go 1.14

require (
	contract v0.0.0
	github.com/ethereum/go-ethereum v1.9.11 // indirect
    hubwiz.com/ethtool v0.0.0
)

replace contract => ../contract
replace hubwiz.com/ethtool => ../../ethtool
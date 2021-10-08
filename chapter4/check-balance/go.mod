module check-balance

go 1.14

require github.com/ethereum/go-ethereum v1.9.11 // indirect

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	hubwiz.com/ethtool v0.0.0
)

replace hubwiz.com/ethtool => ../../ethtool

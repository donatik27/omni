package gasprice

// buffer will be the lowest tier that is higher than the live gas price.
var tiers = []uint64{
	gwei(0.0015), // op stack chains generally maintain 0.001+epsilon.
	gwei(0.05),   // arb generally maintains 0.01 - 0.05.
	gwei(0.1),
	gwei(0.5),
	gwei(1),
	gwei(5),
	gwei(10),
	gwei(20),
	gwei(35),
	gwei(50),
	gwei(75),
	gwei(100),
	gwei(200),
	gwei(300),
	gwei(400),
	gwei(500),
	gwei(1000), // sepolia gas prices sometimes get this high
	gwei(1500),
	gwei(2000),
	gwei(2500),
	gwei(3000),
	gwei(5000),
}

func Tiers() []uint64 {
	return tiers
}

func Tier(live uint64) uint64 {
	for _, p := range tiers {
		if p >= live {
			return p
		}
	}

	return tiers[len(tiers)-1]
}

func gwei(p float64) uint64 {
	return uint64(p * 1e9)
}

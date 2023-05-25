package config

type Configurations struct {
	Local			NetworkConfiguration
	AVATestnet		NetworkConfiguration
}

type NetworkConfiguration struct {
	Network			string
	DeployedAddress	string
	ChainId			int
}

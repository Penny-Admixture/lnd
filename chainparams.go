package lnd

import (
	"github.com/btcsuite/btcd/chaincfg"
	bitcoinCfg "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	bitcoinWire "github.com/btcsuite/btcd/wire"
	"github.com/lightningnetwork/lnd/keychain"
	litecoinCfg "github.com/ltcsuite/ltcd/chaincfg"
	litecoinWire "github.com/ltcsuite/ltcd/wire"
	peercoinCfg "github.com/ppcsuite/ppcd/chaincfg"
)

// activeNetParams is a pointer to the parameters specific to the currently
// active bitcoin network.
var activeNetParams = bitcoinTestNetParams

// bitcoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type bitcoinNetParams struct {
	*bitcoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// litecoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type litecoinNetParams struct {
	*litecoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// peercoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type peercoinNetParams struct {
	*peercoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// bitcoinTestNetParams contains parameters specific to the 3rd version of the
// test network.
var bitcoinTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.TestNet3Params,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinMainNetParams contains parameters specific to the current Bitcoin
// mainnet.
var bitcoinMainNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.MainNetParams,
	rpcPort:  "8334",
	CoinType: keychain.CoinTypeBitcoin,
}

// bitcoinSimNetParams contains parameters specific to the simulation test
// network.
var bitcoinSimNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinSimNetParams contains parameters specific to the simulation test
// network.
var litecoinSimNetParams = litecoinNetParams{
	Params:   &litecoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinTestNetParams contains parameters specific to the 4th version of the
// test network.
var litecoinTestNetParams = litecoinNetParams{
	Params:   &litecoinCfg.TestNet4Params,
	rpcPort:  "19334",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinMainNetParams contains the parameters specific to the current
// Litecoin mainnet.
var litecoinMainNetParams = litecoinNetParams{
	Params:   &litecoinCfg.MainNetParams,
	rpcPort:  "9334",
	CoinType: keychain.CoinTypeLitecoin,
}

// litecoinRegTestNetParams contains parameters specific to a local litecoin
// regtest network.
var litecoinRegTestNetParams = litecoinNetParams{
	Params:   &litecoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinRegTestNetParams contains parameters specific to a local bitcoin
// regtest network.
var bitcoinRegTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// peercoinTestNetParams contains parameters specific to Peercoin test network.
var peercoinTestNetParams = litecoinNetParams{
	Params:   &peercoinCfg.TestNetParams,
	rpcPort:  "9904",
	CoinType: keychain.CoinTypeTestnet,
}

// peercoinMainNetParams contains the parameters specific to the current
// Peercoin mainnet.
var peercoinMainNetParams = peercoinNetParams{
	Params:   &peercoinCfg.MainNetParams,
	rpcPort:  "9902",
	CoinType: keychain.CoinTypePeercoin,
}

// applyLitecoinParams applies the relevant chain configuration parameters that
// differ for litecoin to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyLitecoinParams(params *bitcoinNetParams, litecoinParams *litecoinNetParams) {
	params.Name = litecoinParams.Name
	params.Net = bitcoinWire.BitcoinNet(litecoinParams.Net)
	params.DefaultPort = litecoinParams.DefaultPort
	params.CoinbaseMaturity = litecoinParams.CoinbaseMaturity

	copy(params.GenesisHash[:], litecoinParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = litecoinParams.PubKeyHashAddrID
	params.ScriptHashAddrID = litecoinParams.ScriptHashAddrID
	params.PrivateKeyID = litecoinParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = litecoinParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = litecoinParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = litecoinParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], litecoinParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], litecoinParams.HDPublicKeyID[:])

	params.HDCoinType = litecoinParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(litecoinParams.Checkpoints))
	for i := 0; i < len(litecoinParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], litecoinParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: litecoinParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = litecoinParams.rpcPort
	params.CoinType = litecoinParams.CoinType
}

// applyPeercoinParams applies the relevant chain configuration parameters that
// differ for Peercoin to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyPeercoinParams(params *bitcoinNetParams, peercoinParams *peercoinNetParams) {
	params.Name = peercoinParams.Name
	params.Net = bitcoinWire.BitcoinNet(peercoinParams.Net)
	params.DefaultPort = peercoinParams.DefaultPort
	params.CoinbaseMaturity = peercoinParams.CoinbaseMaturity

	copy(params.GenesisHash[:], peercoinParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = peercoinParams.PubKeyHashAddrID
	params.ScriptHashAddrID = peercoinParams.ScriptHashAddrID
	params.PrivateKeyID = peercoinParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = peercoinParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = peercoinParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = peercoinParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], peercoinParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], peercoinParams.HDPublicKeyID[:])

	params.HDCoinType = peercoinParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(peercoinParams.Checkpoints))
	for i := 0; i < len(peercoinParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], peercoinParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: peercoinParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = peercoinParams.rpcPort
	params.CoinType = peercoinParams.CoinType
}

// isTestnet tests if the given params correspond to a testnet
// parameter configuration.
func isTestnet(params *bitcoinNetParams) bool {
	switch params.Params.Net {
	case bitcoinWire.TestNet3, bitcoinWire.BitcoinNet(litecoinWire.TestNet4):
		return true
	default:
		return false
	}
}

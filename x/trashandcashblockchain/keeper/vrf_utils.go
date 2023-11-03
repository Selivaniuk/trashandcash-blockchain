package keeper

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/joho/godotenv"
	"github.com/yoseplee/vrf"
)

func HashRatio(vrfOutput []byte) (result float64, err error) {
	t := &big.Int{}
	t.SetBytes(vrfOutput[:])

	precision := uint(8 * (len(vrfOutput) + 1))
	max, b, err := big.ParseFloat("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0, precision, big.ToNearestEven)
	if b != 16 || err != nil {
		return 0, err
	}

	h := big.Float{}
	h.SetPrec(precision)
	h.SetInt(t)

	ratio := big.Float{}
	cratio, _ := ratio.Quo(&h, max).Float64()

	return cratio, nil
}

func GetEnv(key string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	seed := os.Getenv(key)

	return seed, nil
}

func GetKeys(seed string) (ed25519.PrivateKey, ed25519.PublicKey, error) {
	seedInBytes, err := hex.DecodeString(seed)
	if err != nil {
		return nil, nil, err
	}

	privateKey := ed25519.NewKeyFromSeed(seedInBytes)
	publicKey := privateKey.Public().(ed25519.PublicKey)

	return privateKey, publicKey, nil
}

func GetVRF(seed string, message []byte) (vrfOutput []byte, err error) {
	privateKey, publicKey, err := GetKeys(seed)
	if err != nil {
		return nil, err
	}

	proof, vrfOutput, err := vrf.Prove(publicKey, privateKey, message)
	if err != nil {
		return nil, err
	}

	verify, err := vrf.Verify(publicKey, proof, message)
	if err != nil {
		return nil, err
	}
	if !verify {
		return nil, fmt.Errorf("VRF proof is invalid")
	}

	return vrfOutput, nil
}

func GetRatioByTxBytes(tx []byte) (ratio float64, err error) {
	seed, err := GetEnv("SEED")
	if err != nil {
		return 0, err
	}

	vrfOutput, err := GetVRF(seed, tx)
	if err != nil {
		return 0, err
	}

	ratio, err = HashRatio(vrfOutput)
	if err != nil {
		return 0, err
	}

	return ratio, nil
}

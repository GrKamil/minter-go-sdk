package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type CheckData struct {
	Nonce    uint
	ChainID  ChainID
	DueBlock uint64
	Coin     [10]byte
	Value    *big.Int
	Lock     *big.Int
	V        *big.Int
	R        *big.Int
	S        *big.Int
}

type SignedCheck interface {
	Encode() ([]byte, error)
}

type CheckTODO interface {
	SetPassphrase(passphrase string) CheckTODO
	Sign(prKey string) (SignedCheck, error)
}

type Check struct {
	*CheckData
	passphrase string
}

func NewCheck(nonce uint, chainID ChainID, dueBlock uint64, coin string, value *big.Int) CheckTODO {
	check := &Check{
		CheckData: &CheckData{
			Nonce:    nonce,
			ChainID:  chainID,
			DueBlock: dueBlock,
			Value:    value,
		},
	}
	copy(check.Coin[:], coin)
	return check
}

func (check *Check) SetPassphrase(passphrase string) CheckTODO {
	check.passphrase = passphrase
	return check
}

func (check *Check) Encode() ([]byte, error) {
	src, err := rlp.EncodeToBytes(check.CheckData)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, hex.EncodedLen(len(src))+2)
	dst[0], dst[1] = 'M', 'c'
	hex.Encode(dst[2:], src)
	return dst, err
}

// todo
func (check *Check) Sign(prKey string) (SignedCheck, error) {
	h, err := rlpHash([]interface{}{
		check.Nonce,
		check.ChainID,
		check.DueBlock,
		check.Coin,
		check.Value,
	})
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(prKey)
	if err != nil {
		return nil, err
	}

	passphraseSum256 := sha256.Sum256([]byte(check.passphrase))

	lock, err := secp256k1.Sign(h[:], passphraseSum256[:])
	if err != nil {
		return nil, err
	}

	check.Lock = big.NewInt(0).SetBytes(lock)

	h, err = rlpHash([]interface{}{
		check.Nonce,
		check.ChainID,
		check.DueBlock,
		check.Coin,
		check.Value,
		check.Lock,
	})
	if err != nil {
		return nil, err
	}

	byteBuffer := bytes.NewBuffer(h[:])
	r, s, err := ecdsa.Sign(byteBuffer, privateKey, h[:])
	if err != nil {
		return nil, err
	}

	check.R = r
	check.S = s
	check.V = big.NewInt(0).SetBytes([]byte{27})

	return check, nil
}

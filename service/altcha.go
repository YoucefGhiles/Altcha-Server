package service

import (
	"errors"
	"os"
	"time"

	altcha "github.com/altcha-org/altcha-lib-go/v2"
)

func getSecret() (string, error) {
	secret := os.Getenv("ALTCHA_HMAC_SECRET")

	if secret == "" {
		return "", errors.New("ALTCHA_HMAC_SECRET is required")
	}

	return secret, nil
}

func CreateChallenge() (altcha.Challenge, error) {

	secret, err := getSecret()
	if err != nil {
		return altcha.Challenge{}, err
	}

	expiresAt := time.Now().Add(5 * time.Minute)

	challenge, err := altcha.CreateChallenge(
		altcha.CreateChallengeOptions{
			Algorithm:           "PBKDF2/SHA-256",
			DeriveKey:           altcha.DeriveKeyPBKDF2(),
			HMACSignatureSecret: secret,
			Cost:                5000,
			KeyLength:           32,
			ExpiresAt:           &expiresAt,
		},
	)

	if err != nil {
		return altcha.Challenge{}, err
	}

	return challenge, nil
}

func VerifyChallenge(
	challenge altcha.Challenge,
	solution altcha.Solution,
) (altcha.VerifySolutionResult, error) {

	secret, err := getSecret()
	if err != nil {
		return altcha.VerifySolutionResult{}, err
	}

	return altcha.VerifySolution(
		altcha.VerifySolutionOptions{
			Challenge:           challenge,
			Solution:            solution,
			DeriveKey:           altcha.DeriveKeyPBKDF2(),
			HMACSignatureSecret: secret,
		},
	)
}
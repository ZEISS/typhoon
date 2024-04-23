package main

import (
	"fmt"
	"log"

	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
)

func main() {
	// op, err := nkeys.CreateOperator()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// seed, err := op.Seed()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(seed))
	// fmt.Println(op.PublicKey())

	// op, err = nkeys.FromSeed(seed)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// seed, err = op.Seed()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(seed))
	// fmt.Println(op.PublicKey())

	// // create an operator key pair (private key)
	okp, err := nkeys.CreateOperator()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(okp.PublicKey())

	// extract the public key
	opk, err := okp.PublicKey()
	if err != nil {
		log.Fatal(err)
	}

	// create an operator claim using the public key for the identifier
	oc := jwt.NewOperatorClaims(opk)
	oc.Name = "O"

	// add an operator signing key to sign accounts
	oskp, err := nkeys.CreateOperator()
	if err != nil {
		log.Fatal(err)
	}

	// get the public key for the signing key
	ospk, err := oskp.PublicKey()
	if err != nil {
		log.Fatal(err)
	}

	// add the signing key to the operator - this makes any account
	// issued by the signing key to be valid for the operator
	oc.SigningKeys.Add(ospk)

	fmt.Println(oc)

	// self-sign the operator JWT - the operator trusts itself
	// operatorJWT, err := oc.Encode(okp)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(operatorJWT)

	// create an account keypair
	akp, err := nkeys.CreateAccount()
	if err != nil {
		log.Fatal(err)
	}
	// extract the public key for the account
	apk, err := akp.PublicKey()
	if err != nil {
		log.Fatal(err)
	}

	// create the claim for the account using the public key of the account
	ac := jwt.NewAccountClaims(apk)
	ac.Name = "A"
	// create a signing key that we can use for issuing users
	askp, err := nkeys.CreateAccount()
	if err != nil {
		log.Fatal(err)
	}

	// extract the public key
	aspk, err := askp.PublicKey()
	if err != nil {
		log.Fatal(err)
	}

	// add the signing key (public) to the account
	ac.SigningKeys.Add(aspk)

	fmt.Println(ac)

	// now we could encode an issue the account using the operator
	// key that we generated above, but this will illustrate that
	// the account could be self-signed, and given to the operator
	// who can then re-sign it
	accountJWT, err := ac.Encode(akp)
	if err != nil {
		log.Fatal(err)
	}

	// the operator would decode the provided token, if the token
	// is not self-signed or signed by an operator or tampered with
	// the decoding would fail
	ac, err = jwt.DecodeAccountClaims(accountJWT)
	if err != nil {
		log.Fatal(err)
	}

	// here the operator is going to use its private signing key to
	// re-issue the account
	accountJWT, err = ac.Encode(oskp)
	if err != nil {
		log.Fatal(err)
	}

	// // now back to the account, the account can issue users
	// // need not be known to the operator - the users are trusted
	// // because they will be signed by the account. The server will
	// // look up the account get a list of keys the account has and
	// // verify that the user was issued by one of those keys
	// ukp, err := nkeys.CreateUser()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(ukp)

	// upk, err := ukp.PublicKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(upk)

	// uc := jwt.NewUserClaims(upk)
	// // since the jwt will be issued by a signing key, the issuer account
	// // must be set to the public ID of the account
	// uc.IssuerAccount = apk
	// userJwt, err := uc.Encode(askp)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(userJwt)

	// // the seed is a version of the keypair that is stored as text
	// useed, err := ukp.Seed()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(useed)

	// // generate a creds formatted file that can be used by a NATS client
	// creds, err := jwt.FormatUserConfig(userJwt, useed)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(creds)
}

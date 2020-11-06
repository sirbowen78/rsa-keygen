## Reference
[Build an RSA Asymmetric Cryptography Generator in Go](https://medium.com/better-programming/build-an-rsa-asymmetric-cryptography-generator-in-go-d202b18bcfd0) provides codes to generate RSA key pairs and how to convert the key to PEM then export to a file.

## Usage Examples
| command | Description |
| ------- | ----------- |
| `rsa-keygen` | This generates RSA key pair of 2048-bit key length, the private key is privkey.pem and the public key is pubkey.pem |
| `rsa-keygen -b 4096` | This generates RSA key pair of 4096-bit key length, the private key is privkey.pem, and the public key is pubkey.pem |
| `rsa-keygen -b 4096 --pte private.pem --pub public.pem` | This generates RSA key pair of 4096-bit key length, the private key is private.pem and the public key is public.pem |

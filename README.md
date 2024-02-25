### TODO
- add feature to use yaml config data instead of passing data as os arg (especially for sign-message)
- ~~add feature to get account nonce~~
- receive tx receipts through email (if the email env var is set, receipts will be printed out in os.stdout and also sent to email, else if its "" it'll only be printed out to os.stdout)
- add feature to efficiently manage private key (gast init will create a .gast.yaml file in root folder, gast add --privKey "privKey" will add privKey to the yaml file, gast create will create a new keypair for signing and sending tx at the root .yaml file and return the pubKey to them, a new gast create will override the existing keypair so warn users first when they do it)
- add feature to send normal tx using eth_call
- add feature to trace transaction


```shell
go build .
./gast tx estimate-gas -f=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045 -handleTraceTx=0xbe0eb53f46cd790cd13851d5eff43d12404d33e8 -d=0x -u=https://rpc.mevblocker.io -w=1000000
./gast tx estimate-gas -f=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045 -handleTraceTx=0xbe0eb53f46cd790cd13851d5eff43d12404d33e8 -d=0xde725e890000000000000000000000000f93ae9f3b81c12cbc009e8f0d4a4f4f044df3040000000000000000000000007a250d5630b4cf539739df2c5dacb4c659f2488d0000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000 -u=https://rpc.mevblocker.io -w=0

```

[//]: # (![img.png]&#40;img.png&#41;)

```shell
go run . tx create-raw --url "https://optimism.publicnode.com" --to "0x571B102323C3b8B8Afb30619Ac1d36d85359fb84" --data "eth signed message v2" --private-key "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662" --wei 1000000000000000000 --nonce 0
go run . tx sign-message -d "eth signed message v2" -p "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662"
```


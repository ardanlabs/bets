# The environment has three accounts all using this same passkey (123).
# Geth is started with address 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd and is used as the coinbase address.
# The coinbase address is the account to pay mining rewards to.
# The coinbase address is given a LOT of money to start.
#
# These are examples of what you can do in the attach JS environment.
# 	eth.getBalance("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd") or eth.getBalance(eth.coinbase)
# 	eth.getBalance("0x8e113078adf6888b7ba84967f299f29aece24c55")
# 	eth.getBalance("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7")
#   eth.sendTransaction({from:eth.coinbase, to:"0x8e113078adf6888b7ba84967f299f29aece24c55", value: web3.toWei(0.05, "ether")})
#   eth.sendTransaction({from:eth.coinbase, to:"0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", value: web3.toWei(0.05, "ether")})
#   eth.blockNumber
#   eth.getBlockByNumber(8)
#   eth.getTransaction("0xaea41e7c13a7ea627169c74ade4d5ea86664ff1f740cd90e499f3f842656d4ad")
#
# make geth-deposit
# ./admin -a 1000.00 -f 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd -c 0x531130464929826c57BBBF989e44085a02eeB120
# ./admin -a 1000.00 -f 0x8e113078adf6888b7ba84967f299f29aece24c55 -c 0x531130464929826c57BBBF989e44085a02eeB120
# ./admin -a 1000.00 -f 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7 -c 0x531130464929826c57BBBF989e44085a02eeB120
#
# Web3 API
# https://web3js.readthedocs.io/en/v1.7.4/

# ==============================================================================
# Install dependencies
# https://geth.ethereum.org/docs/install-and-build/installing-geth
# https://docs.soliditylang.org/en/v0.8.11/installing-solidity.html

dev.setup.mac:
	brew update
	brew list ethereum || brew install ethereum
	brew list solidity || brew install solidity
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize
	brew list pgcli || brew install pgcli

dev.update:
	brew update
	brew list ethereum || brew upgrade ethereum
	brew list solidity || brew upgrade solidity

# ==============================================================================
# Building containers

# $(shell git rev-parse --short HEAD)
VERSION := 1.0

all: engine

engine:
	docker build \
		-f zarf/docker/dockerfile.engine \
		-t engine-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Running from within k8s/kind

KIND_CLUSTER := ardan-starter-cluster

# Upgrade to latest Kind: brew upgrade kind
# For full Kind v0.14 release notes: https://github.com/kubernetes-sigs/kind/releases/tag/v0.14.0
# The image used below was copied by the above link and supports both amd64 and arm64.

kind-up:
	kind create cluster \
		--image kindest/node:v1.24.0@sha256:0866296e693efe1fed79d5e6c7af8df71fc73ae45e3679af05342239cdc5bc8e \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml
	#kubectl config set-context --current --namespace=engine-system

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image engine-amd64:$(VERSION) --name $(KIND_CLUSTER)

kind-apply:
	kustomize build zarf/k8s/kind/database-pod | kubectl apply -f -
	kustomize build zarf/k8s/kind/geth-pod | kubectl apply -f -
	kubectl wait --namespace=database-system --timeout=120s --for=condition=Available deployment/database-pod
	kubectl wait --namespace=geth-system --timeout=120s --for=condition=Available deployment/geth-pod
	kustomize build zarf/k8s/kind/engine-api-pod | kubectl apply -f -

kind-restart:
	kubectl rollout restart deployment engine-api-pod --namespace=engine-api-system

kind-update: all kind-load kind-restart

kind-update-apply: all kind-load kind-apply

kind-logs-engine-api:
	kubectl logs -l app=engine-api --all-containers=true -f --tail=100 --namespace=engine-api-system

kind-logs-geth:
	kubectl logs -l app=geth --all-containers=true -f --tail=100 --namespace=geth-system

kind-status:
	kubectl get nodes -o wide --all-namespaces
	kubectl get svc -o wide --all-namespaces
	kubectl get pods -o wide --watch --all-namespaces

kind-status-engine-api:
	kubectl get pods -o wide --watch --namespace=engine-api-system

kind-describe-engine-api:
	kubectl describe pod -l app=engine-api --namespace=engine-api-system

kind-geth-attach:
	$(eval $@_GETH_POD := $(shell kubectl get --namespace=geth-system pod -l app=geth -o jsonpath="{.items[0].metadata.name}"))
	kubectl exec -it --namespace=geth-system $($@_GETH_POD) -- geth attach --datadir /ethereum

kind-geth-new-account:
	$(eval $@_GETH_POD := $(shell kubectl get --namespace=geth-system pod -l app=geth -o jsonpath="{.items[0].metadata.name}"))
	kubectl exec -it --namespace=geth-system $($@_GETH_POD) -- geth account new --datadir /ethereum

# ==============================================================================
# Administration

migrate:
	go run app/tooling/db/main.go migrate

seed: migrate
	go run app/tooling/db/main.go seed

# ==============================================================================
# Game Engine UI

react-install:
	yarn --cwd app/services/game/ install

game-gui: react-install
	yarn --cwd app/services/game/ start

# ==============================================================================
# These commands build and deploy basic smart contract.

# This will compile the smart contract and produce the binary code. Then with the
# abi and binary code, a Go source code file can be generated for Go API access.
contract-build:
	solc --abi business/contract/src/book/book.sol -o business/contract/abi/book --overwrite
	solc --bin business/contract/src/book/book.sol -o business/contract/abi/book --overwrite
	abigen --bin=business/contract/abi/book/book.bin --abi=business/contract/abi/book/book.abi --pkg=book --out=business/contract/go/book/book.go

# This will deploy the smart contract to the locally running Ethereum environment.
admin-build:
	go build -o admin app/tooling/admin/main.go

contract-deploy: contract-build admin-build
	./admin -d

# ==============================================================================
# These are Ethereum commands for attaching, creating a new account and depositing
# and other examples.

# This is start Ethereum in developer mode. Only when a transaction is pending will
# Ethereum mine a block. It provides a minimal environment for development.
geth-up:
	geth --dev --ipcpath zarf/ethereum/geth.ipc --http.corsdomain '*' --http --allow-insecure-unlock --rpc.allow-unprotected-txs --http.vhosts=* --mine --miner.threads 1 --verbosity 5 --datadir "zarf/ethereum/" --unlock 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd --password zarf/ethereum/password

# This will signal Ethereum to shutdown.
geth-down:
	kill -INT $(shell ps | grep "geth " | grep -v grep | sed -n 1,1p | cut -c1-5)

# This will remove the local blockchain and let you start new.
geth-reset:
	rm -rf zarf/ethereum/geth/

# This is a JS console environment for making geth related API calls.
geth-attach:
	geth attach --datadir zarf/ethereum/

# This will add a new account to the keystore. The account will have a zero
# balance until you give it some money.
geth-new-account:
	geth --datadir zarf/ethereum/ account new

# This will deposit 1 ETH into the two extra accounts from the coinbase account.
# Do this if you delete the geth folder and start over or if the accounts need money.
geth-deposit:
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x8E113078ADF6888B7ba84967F299F29AeCe24c55", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7", "value":"0x1000000000000000000"}], "id":1}' localhost:8545

# ==============================================================================
# Running tests within the local computer
# go install honnef.co/go/tools/cmd/staticcheck@latest
# go install golang.org/x/vuln/cmd/govulncheck@latest

test:
	CGO_ENABLED=0 go test -count=1 ./...
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...
	govulncheck ./...

test-gui:
	yarn --cwd app/services/game/ test

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -v ./...
	go mod tidy
	go mod vendor

list:
	go list -mod=mod all

rmds:
	find . -name '.DS_Store' -type f -delete
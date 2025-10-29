build: contracts/
	cd contracts && forge build

anvil:
	killall anvil || true
	anvil &
	@while ! nc -z localhost 8545; do sleep 0.1; done

deploy: build
	cd contracts && \
	forge create \
	  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
	  --rpc-url http://localhost:8545 \
	  --broadcast IdentityRegistry
	# Deployed to 0x5FbDB2315678afecb367f032d93F642f64180aa3

	cd contracts && \
	forge create \
	  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
	  --rpc-url http://localhost:8545 \
	  --broadcast ReputationRegistry \
	  --constructor-args 0x5FbDB2315678afecb367f032d93F642f64180aa3
	# Deployed to 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512

	cd contracts && \
	forge create \
	  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
	  --rpc-url http://localhost:8545 \
	  --broadcast ValidationRegistry \
	  --constructor-args 0x5FbDB2315678afecb367f032d93F642f64180aa3
	# Deployed to 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0

test: anvil deploy
	go test -v

verify-gentx:
	bash -x ./scripts/verify_gentx.sh
	
gen-genesis:
	bash -x ./scripts/gen-genesis.sh

check-amount:
	bash -x ./scripts/check-gentx-amount.sh

.PHONY: verify-gentx
const { ethers } = require("hardhat");

const bscProvider = ethers.getDefaultProvider(
  ""
);

async function extractValidatorSet(extraData) {
  const extra = ethers.utils.arrayify(extraData);

  let currentPosition = 32;
  const endPosition = extra.length - 65;
  const numValidators = (endPosition - currentPosition) / 20;

  const validatorSet = [];
  for (i = 0; i < numValidators; i++) {
    validatorSet.push(
      ethers.utils.hexlify(extra.slice(currentPosition, currentPosition + 20))
    );
    currentPosition += 20;
  }
  return validatorSet;
}

async function main() {
  const blockNumber = await bscProvider.getBlockNumber();
  const block = await bscProvider.getBlock(blockNumber);

  const diff = blockNumber % 200;
  const lastEpochBlockNumber = blockNumber - diff;

  const currentEpochBlock = await bscProvider.getBlock(lastEpochBlockNumber);
  const previousEpochBlock = await bscProvider.getBlock(
    lastEpochBlockNumber - 200
  );

  const previousValidatorSet = await extractValidatorSet(
    previousEpochBlock.extraData
  );

  let validatorSet = await extractValidatorSet(currentEpochBlock.extraData);

  if (diff < previousValidatorSet.length / 2) {
    validatorSet = previousValidatorSet;
  }

  const bscRelayContractFactory = await ethers.getContractFactory("BSCRelay");
  const bscRelay = await bscRelayContractFactory.deploy(
    validatorSet,
    block.hash,
    block.number
  );

  console.log("BSCRelay deployed to:", bscRelay.address);
  console.log("Genesis block:", block.hash);
  console.log("Block number:", block.number);
  console.log("ValidatorSet:", validatorSet);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

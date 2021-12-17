const fs = require("fs");
const path = require("path");

const relayContractArtifacts = path.resolve(
  __dirname,
  "../artifacts/contracts/BSCRelay.sol"
);

const relayContractPath = path.resolve(relayContractArtifacts, "BSCRelay.json");
const relayContractAbiPath = path.resolve(
  relayContractArtifacts,
  "BSCRelay.abi"
);
const relayContractBinPath = path.resolve(
  relayContractArtifacts,
  "BSCRelay.bin"
);

const relayContract = JSON.parse(fs.readFileSync(relayContractPath, "utf8"));

// eslint-disable-next-line consistent-return
fs.writeFile(relayContractAbiPath, JSON.stringify(relayContract.abi), (err) => {
  if (err) {
    return console.error(err);
  }
  console.log("BSCRelay ABI written successfully!");
});

fs.writeFile(
  relayContractBinPath,
  JSON.stringify(relayContract.bytecode),
  // eslint-disable-next-line consistent-return
  (err) => {
    if (err) {
      return console.error(err);
    }
    console.log("BSCRelay BIN written successfully!");
  }
);

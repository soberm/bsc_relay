require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-solhint");
require("@nomiclabs/hardhat-waffle");
require("hardhat-gas-reporter");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (let i = 0; i < accounts.length; i++) {
    const account = accounts[i];
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  defaultNetwork: "hardhat",
  networks: {
    hardhat: {},
    ropsten: {
      url: "",
      accounts: [
        "",
      ],
    },
  },
  solidity: {
    version: "0.7.3",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
};

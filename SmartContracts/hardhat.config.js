require("dotenv").config();
require("@nomicfoundation/hardhat-toolbox");

module.exports = {
  solidity: {
    version: "0.8.28",
    settings: {
      optimizer: { enabled: true, runs: 200 }
    }
  },
  networks: {
    buildbear: {
      url: "https://rpc.buildbear.io/positive-husk-962d2b1b",
      accounts: [process.env.PRIVATE_KEY],
      chainId: 31337
    }
  }
};

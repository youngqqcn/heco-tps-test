require("@nomiclabs/hardhat-waffle");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: "0.8.4",
  networks: {
    private: {
      // chainId: 2285.,
      chainId: 1337.,
      url: "http://localhost:8545",
      accounts: [
        "0xcfe945f87d61aa82e903804bcc32bacdf130ae47268a2f6d7a3d877cbf028ff6",// 0x8284B6412ef6eFA75adDEa85f07E7de5f8F8ec48
        "0x5ea30eea9ba9500f3601f7659f0ccace819c562456e2f745fb2555918ab32277", // 0xf513e4e5Ded9B510780D016c482fC158209DE9AA
        "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
        "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
        "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
      ]
    },
    cluster: {
      chainId: 1874,
      url: "http://10.10.7.30:30545",
      accounts: [
        "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
        "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
        "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
      ]
    }
  }
};

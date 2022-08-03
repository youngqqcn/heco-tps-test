const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("RabbitCollectible", function () {
  it("Should return the symbol", async function () {
    const RabbitCollectible= await ethers.getContractFactory("RabbitCollectible");
    const rabbit = await RabbitCollectible.deploy();
    await rabbit.deployed();

    expect(await rabbit.symbol()).to.equal("RBO");
  });
});

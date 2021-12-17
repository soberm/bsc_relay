// const { expect } = require("chai");
const { ethers } = require("hardhat");
const { expect } = require("chai");
const data = require("./data/data.json");

describe("BSCRelay", () => {
  beforeEach(async () => {
    bscRelayContractFactory = await ethers.getContractFactory("BSCRelay");
    bscRelay = await bscRelayContractFactory.deploy(
      data.validatorSet,
      "0x33630b0b4652353bb0bd4c5bf0863addff2937c47e9142f3929aac0adf35a13e",
      12928775
    );
    await bscRelay.deployed();
  });

  describe("submitBlockHeaderBatch", () => {
    it("should submit block header batch", async () => {
      await bscRelay.submitBlockHeaderBatch(
        data.unsignedHeaders.slice(1, 15),
        data.signedHeaders.slice(1, 15)
      );
    });
    it("should revert with not same amount of headers", async () => {
      const request = bscRelay.submitBlockHeaderBatch(
        data.unsignedHeaders.slice(1, 15),
        data.unsignedHeaders.slice(2, 15)
      );
      await expect(request).to.be.revertedWith("not same amount of headers");
    });
  });

  describe("submitBlockHeader", () => {
    it("should submit header", async () => {
      await bscRelay.submitBlockHeader(
        data.unsignedHeaders[1],
        data.signedHeaders[1]
      );
    });
    it("should revert with unsigned not equals signed", async () => {
      const request = bscRelay.submitBlockHeader(
        data.unsignedHeaders[1],
        data.signedHeaders[2]
      );
      await expect(request).to.be.revertedWith("unsigned not equals signed");
    });
    it("should revert with block already exists", async () => {
      const request = bscRelay.submitBlockHeader(
        data.unsignedHeaders[0],
        data.signedHeaders[0]
      );
      await expect(request).to.be.revertedWith("block already exists");
    });
    it("should revert with no parent", async () => {
      const request = bscRelay.submitBlockHeader(
        data.unsignedHeaders[2],
        data.signedHeaders[2]
      );
      await expect(request).to.be.revertedWith("no parent");
    });
    it("should revert with mixhash not 0", async () => {
      const request = bscRelay.submitBlockHeader(
        data.wrongMixHash[0],
        data.wrongMixHash[1]
      );
      await expect(request).to.be.revertedWith("mixHash not 0");
    });
    it("should revert with difficulty not 1 or 2", async () => {
      const request = bscRelay.submitBlockHeader(
        data.wrongDifficulty[0],
        data.wrongDifficulty[1]
      );
      await expect(request).to.be.revertedWith("difficulty not 1 or 2");
    });
    it("should revert with invalid signature", async () => {
      const request = bscRelay.submitBlockHeader(
        data.wrongSignature[0],
        data.wrongSignature[1]
      );
      await expect(request).to.be.revertedWith("invalid signature");
    });
  });

  describe("verifyTransaction", () => {
    it("should verify valid transaction", async () => {
      for (i = 1; i < data.signedHeaders.length; i++) {
        await bscRelay.submitBlockHeader(
          data.unsignedHeaders[i],
          data.signedHeaders[i]
        );
      }

      const result = await bscRelay.verifyTransaction(
        data.merkleProofTx.header,
        data.merkleProofTx.value,
        data.merkleProofTx.path,
        data.merkleProofTx.nodes
      );
      expect(result).to.be.equal(true);
    });
  });

  describe("verifyReceipt", () => {
    it("should verify valid receipt", async () => {
      for (i = 1; i < data.signedHeaders.length; i++) {
        await bscRelay.submitBlockHeader(
          data.unsignedHeaders[i],
          data.signedHeaders[i]
        );
      }

      const result = await bscRelay.verifyReceipt(
        data.merkleProofReceipt.header,
        data.merkleProofReceipt.value,
        data.merkleProofReceipt.path,
        data.merkleProofReceipt.nodes
      );
      expect(result).to.be.equal(true);
    });
  });
});

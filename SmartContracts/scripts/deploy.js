const { ethers } = require("hardhat");

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with:", deployer.address);

  // 1. UserRegistry
  const UserRegistry = await ethers.getContractFactory("UserRegistry");
  const userRegistry = await UserRegistry.deploy();
  await userRegistry.waitForDeployment();
  console.log("UserRegistry deployed at:", await userRegistry.getAddress());

  // 2. Reputation
  const Reputation = await ethers.getContractFactory("Reputation");
  const reputation = await Reputation.deploy();
  await reputation.waitForDeployment();
  console.log("Reputation deployed at:", await reputation.getAddress());

  // 3. Mentorship
  const Mentorship = await ethers.getContractFactory("Mentorship");
  const mentorship = await Mentorship.deploy(
    await userRegistry.getAddress(),
    await reputation.getAddress()
  );
  await mentorship.waitForDeployment();
  console.log("Mentorship deployed at:", await mentorship.getAddress());

  // Set operator for Reputation
  const tx = await reputation.setOperator(await mentorship.getAddress());
  await tx.wait();
  console.log("Reputation operator set to Mentorship");

  // 4. ContentRegistry
  const ContentRegistry = await ethers.getContractFactory("ContentRegistry");
  const contentRegistry = await ContentRegistry.deploy();
  await contentRegistry.waitForDeployment();
  console.log("ContentRegistry deployed at:", await contentRegistry.getAddress());

  // 5. ContentAccess
  const ContentAccess = await ethers.getContractFactory("ContentAccess");
  const contentAccess = await ContentAccess.deploy(
    await contentRegistry.getAddress()
  );
  await contentAccess.waitForDeployment();
  console.log("ContentAccess deployed at:", await contentAccess.getAddress());

  console.log("All contracts deployed successfully.");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

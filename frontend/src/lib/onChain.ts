import { ethers } from "ethers";

const USER_REGISTRY_ABI = [
  "function verifyUser(address,uint8,string,string,string,string,string)"
];

const USER_REGISTRY_ADDRESS = "0x623E59402bE01B511e373Bb68f67547BfD01b59e"; 

export async function verifyUserOnChain(params: {
  signer?: ethers.Signer; 
  address: string;
  role: number;
  name: string;
  rollNo: string;
  program: string;
  major: string;
  pic: string;
}) {
    console.log(params.role)
  if (!params.signer && typeof window === "undefined") {
    throw new Error("No signer available");
  }

  let provider: any;
  if (params.signer) {
    provider = (params.signer as any).provider ?? (params.signer as any)._provider ?? null;
  }
  if (!provider) {
    if (!(window as any).ethereum) throw new Error("MetaMask not found");
    provider = new ethers.BrowserProvider((window as any).ethereum);
    await provider.send("eth_requestAccounts", []);
  }

  const signer = params.signer ?? (await provider.getSigner());

  const network = await provider.getNetwork();
  if (network.chainId !== 11155111) {
    try {
      await (window as any).ethereum.request({
        method: "wallet_switchEthereumChain",
        params: [{ chainId: "0xAA36A7" }],
      });
      await provider.send?.("eth_requestAccounts", []);
    } catch (err) {
      throw new Error("Please switch MetaMask network to Sepolia");
    }
  }

  const contract = new ethers.Contract(USER_REGISTRY_ADDRESS, USER_REGISTRY_ABI, signer);

//   const roleUint8 = Number(params.role) & 0xff;

  const tx = await contract.verifyUser(
    params.address,
    1,
    params.name,
    params.rollNo,
    params.program,
    params.major,
    params.pic || ""
  );

  return tx; 
}

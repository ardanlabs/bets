require('responses.d')
require('props.d')

export interface appConfig {
  chainId: number
  contractId: string
  network: string
}

export type Falsy = false | 0 | '' | null | undefined

export interface AddEthereumChainParameter {
  chainId: string // A 0x-prefixed hexadecimal string
  chainName: string
  nativeCurrency: {
    name: string
    symbol: string // 2-6 characters long
    decimals: 18
  }
  rpcUrls: string[]
  blockExplorerUrls?: string[]
  iconUrls?: string[] // Currently ignored.
}

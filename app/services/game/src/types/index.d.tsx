import { CSSProperties } from 'react'

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

export interface BetsFilter {
  status: string
  placer: string
  challenger: string
  expirationDate: string
}

export interface Bet {
  id: number
  status: string
  placerAddress: string
  description: string
  terms: string
  name: string
  challengerAddress: string
  expirationDate: string
  amount: number
}

export interface DefaultDoc {
  address: string
  dateTime: string
  betId: number
}

export interface SetWinnerDoc extends DefaultDoc {
  winnerAddress: string
}

export interface SetModeratorDoc extends DefaultDoc {
  moderatorAddress: string
}

export interface StyleObject {
  [key: string]: CSSProperties
}

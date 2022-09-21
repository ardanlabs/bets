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
  player: string
  challenger: string
  expirationDate: string
}

export interface Player {
  address: string
  signed: boolean
}

export interface Bet {
  id: number
  status: string
  players: Player[]
  moderator: string
  description: string
  terms: string
  expirationDate: string
  amount: number
}

export interface DefaultDoc {
  address: string
  dateTime: string
  betId: number
}

export interface SetModeratorDoc extends DefaultDoc {
  moderatorAddress: string
}

export interface StyleObject {
  [key: string]: CSSProperties
}

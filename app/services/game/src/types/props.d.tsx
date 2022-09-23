import {
  AnimationControls,
  Target,
  TargetAndTransition,
  VariantLabels,
} from 'framer-motion'
import { Bet, StyleObject } from './index.d'

export interface GameTableProps {
  timer: number
}

export interface AppHeaderProps {
  show?: boolean
}

export interface SignOutProps {
  disabled: boolean
}

export interface CounterProps {
  timer: number
  show: boolean
}

export interface MainRoomProps {
  timer: number
}

export interface transactionProps {
  buttonText: string
  action: 'Deposit' | 'Withdraw'
  updateBalance: Function
}

export interface ButtonProps {
  clickHandler: Function
  classes?: string
  id?: string
  disabled?: boolean
  children:
    | (JSX.Element[] | JSX.Element | null)[]
    | JSX.Element
    | string
    | null
    | undefined
  style?: React.CSSProperties
  tooltip?: string
}

export interface JoinProps {
  disabled?: boolean
}

export interface PlayersListProps {
  title: string
}

export interface SideBarProps {
  notificationCenterWidth: string
}

export interface SubtitleProps {
  showSearch: boolean
  text: string
}

export interface RoutesWrapperProps {
  children:
    | (JSX.Element[] | JSX.Element | null)[]
    | JSX.Element
    | string
    | null
    | undefined
}

export interface CardProps {
  children:
    | (JSX.Element[] | JSX.Element | null)[]
    | JSX.Element
    | string
    | null
    | undefined
  styleObject?: StyleObject
  initial?: boolean | Target | VariantLabels
  animate?: AnimationControls | TargetAndTransition | VariantLabels | boolean
  exit?: TargetAndTransition | VariantLabels
  closeMethod?: Function
  showClose?: boolean
  subtitle?: string
}

export interface BetCardProps {
  bet: Bet
  styleObject?: StyleObject
  isDetail: boolean
}

export interface EditBetProps {
  bet?: Bet
  hideModalMethod: React.Dispatch<React.SetStateAction<boolean>>
}

export interface BetStatusProps {
  status: string
}

export interface SuccessModalProps {
  show: boolean
  setShow: React.Dispatch<React.SetStateAction<boolean>>
  betId?: string
}

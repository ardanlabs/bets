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
  children: JSX.Element[] | JSX.Element | string
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

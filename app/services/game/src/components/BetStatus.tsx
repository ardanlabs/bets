import { StyleObject } from '../types/index.d'
import { BetStatusProps } from '../types/props.d'

// BetStatus component
function BetStatus(props: BetStatusProps) {
  // Extracts props
  const { status } = props

  const styles: StyleObject = {
    status: {
      borderRadius: '8px',
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '600',
      fontSize: '16px',
      lineHeight: '11px',
      padding: '12px',
      aspectRatio: '3/1',
      color: 'var(--bet-status-text)',
      background: 'var(--bet-status-open)',
    },
    open: {
      background: 'var(--bet-status-open)',
    },
    signing: {
      background: 'var(--bet-status-open)',
    },
    moderate: {
      background: 'var(--bet-status-open)',
    },
    live: {
      background: 'var(--bet-status-live)',
    },
    closed: {
      background: 'var(--bet-status-closed)',
    },
    canceled: {
      background: 'var(--bet-status-canceled)',
    },
    p: {
      margin: '0',
    },
  }
  return status ? (
    <div style={{ ...styles.status, ...styles[status] }}>
      <p style={styles.p}>{status}</p>
    </div>
  ) : null
}

export default BetStatus

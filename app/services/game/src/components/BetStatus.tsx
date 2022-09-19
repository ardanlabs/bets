import { StyleObject } from '../types/index.d'
import { BetStatusProps } from '../types/props.d'

// BetStatus component
function BetStatus(props: BetStatusProps) {
  // Extracts props
  const { status } = props

  const styles: StyleObject = {
    status: {
      background: 'var(--status-open-box)',
      borderRadius: '8px',
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '600',
      fontSize: '16px',
      lineHeight: '11px',
      color: 'var(--status-open-text)',
      padding: '12px',
      aspectRatio: '3/1',
    },
    p: {
      margin: '0',
    },
  }
  return status ? (
    <div style={styles.status}>
      <p style={styles.p}>{status}</p>
    </div>
  ) : null
}

export default BetStatus

import { StyleObject } from '../types/index.d'
import { CardProps } from '../types/props.d'
import { motion } from 'framer-motion'
import Subtitle from './Subtitle'

function Card(props: CardProps) {
  // Extracts props
  const {
    children,
    styleObject,
    initial,
    animate,
    exit,
    showClose,
    closeMethod,
    subtitle,
  } = props

  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    card: {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'flex-start',
      backgroundColor: 'var(--bets-color)',
      boxShadow: '3px 4px 4px rgba(0, 0, 0, 0.25)',
      borderRadius: '8px',
      padding: '8px 16px',
    },
    closeButton: {
      alignSelf: 'flex-end',
    },
    header: {
      width: '100%',
      padding: '8px 0',
      display: 'flex',
      justifyContent: subtitle ? 'space-between' : 'flex-end',
    },
    ...styleObject,
  }
  return (
    <motion.div
      style={styles.card}
      initial={initial}
      animate={animate}
      exit={exit}
    >
      {subtitle || showClose ? (
        <div style={styles.header}>
          {subtitle ? <Subtitle text={subtitle} showSearch={false} /> : null}
          {showClose && closeMethod ? (
            <button
              style={styles.closeButton}
              type="button"
              className="btn-close"
              aria-label="Close"
              onClick={() => closeMethod()}
            ></button>
          ) : null}
        </div>
      ) : null}
      {children}
    </motion.div>
  )
}

export default Card

import React from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import { styles } from './styles'
import Card from '../Card'

interface ModalProps {
  children: (JSX.Element | null)[] | JSX.Element | string | null
  trigger?: JSX.Element | string | null
  subtitle?: string
  show: boolean
  setShow: React.Dispatch<React.SetStateAction<boolean>>
}

// Modal component with 'showModal' prop only
// to manage its state of visibility and
// animated using framer-motion
function Modal(props: ModalProps) {
  // Extracts props
  const { children, trigger, subtitle, setShow, show } = props

  // utility function to set the show value
  // opposite of its last value
  // to toggle modal
  function toggleModal() {
    setShow(!show)
  }
  return (
    <>
      <div
        style={{
          position: 'relative',
          display: 'inline-block',
          cursor: 'pointer',
          width: 'auto',
        }}
        id="btn"
        onClick={toggleModal}
      >
        {trigger}
      </div>
      <AnimatePresence>
        {show ? (
          <motion.div
            style={styles.modalContainer}
            initial={{ opacity: 0 }}
            animate={{
              opacity: 1,
              transition: { ease: 'anticipate', duration: 0.4 },
            }}
            exit={{
              opacity: 0,
              transition: { ease: 'anticipate', duration: 0.4 },
            }}
          >
            <Card
              styleObject={{ modalBox: styles.modalBox }}
              initial={{ opacity: 0, y: 60, scale: 0.5 }}
              showClose={true}
              closeMethod={toggleModal}
              animate={{
                opacity: 1,
                y: 0,
                scale: 1,
                transition: { ease: 'anticipate', duration: 0.4 },
              }}
              exit={{
                opacity: 0,
                scale: 0.5,
                transition: { ease: 'anticipate', duration: 0.4 },
              }}
              subtitle={subtitle}
            >
              {children}
            </Card>
          </motion.div>
        ) : null}
      </AnimatePresence>
    </>
  )
}

export default Modal

import React from 'react'
import { useNavigate } from 'react-router-dom'
import { StyleObject } from '../types/index.d'
import { SuccessModalProps } from '../types/props.d'
import Button from './Button'
import SuccessIcon from './icons/SuccessIcon'
import Modal from './modal/Modal'

// SuccessModal component
function SuccessModal(props: SuccessModalProps) {
  // Extracts props
  const { show, setShow, betId } = props

  // Extracts navigation functionalities
  const navigate = useNavigate()

  const styles: StyleObject = {
    wrapper: {
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      flexFlow: 'column',
      width: '100%',
      padding: '8px 16px',
    },
    text: {
      color: 'var(--success-color)',
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '700',
      fontSize: '36px',
      lineHeight: '38px',
      textAlign: 'center',
    },
  }
  return (
    <Modal show={show} setShow={setShow}>
      <div style={styles.wrapper}>
        <SuccessIcon />
        <p style={styles.text}>You successfully created a bet </p>
        <p
          style={{
            ...styles.text,
            fontSize: '20px',
            lineHeight: '22px',
            marginBottom: '20px',
          }}
        >
          Now, just wait for your challenger to accept it
        </p>
        {betId ? (
          <Button
            classes="btn-outline-secondary"
            style={{
              position: 'relative',
              display: 'inline-block',
              cursor: 'pointer',
              width: 'auto',
              alignSelf: 'flex-end',
            }}
            clickHandler={() => navigate(`/bet/${betId}`)}
          >
            See detail
          </Button>
        ) : null}
      </div>
    </Modal>
  )
}

export default SuccessModal

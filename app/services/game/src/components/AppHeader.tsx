import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import { StyleObject } from '../types/index.d'
import Button from './Button'
import EditBet from './EditBet'
import useEthersConnection from './hooks/useEthersConnection'
import Login from './Login'
import Modal from './modal/Modal'

// AppHeader renders the application header
function AppHeader() {
  // We create a local state to handle if the modals are shown or not
  const [addBetModal, setAddBetModal] = useState(false)
  const [loginModal, setLoginModal] = useState(false)

  // Extracts the account from useEthersConnection hook
  const { account } = useEthersConnection()
  // ===========================================================================

  const headerHeight = '67px'

  const styles: StyleObject = {
    header: {
      display: 'flex',
      justifyContent: 'space-between',
      alignItems: 'center',
      color: 'var(--text-color)',
      height: headerHeight,
      flexDirection: 'row',
      padding: '0 24px',
    },
    firstRow: {
      display: 'flex',
      width: 'auto',
      justifyContent: 'space-evenly',
      alignItems: 'center',
      color: 'var(--text-color)',
      gap: '20px',
      height: headerHeight,
    },
    secondRow: {
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      color: 'var(--text-color)',
      height: headerHeight,
    },
    h1: {
      width: 'auto',
      height: '45px',
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: 700,
      lineHeight: '45px',
      fontSize: '45px',
      color: '#000000',
      flex: 'none',
      order: '0',
      flexGrow: '0',
      padding: '0',
      margin: '0',
    },
    button: {
      margin: '0',
      width: 'auto',
      textDecoration: 'unset',
    },
    text: {
      color: 'var(--text-color)',
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '700',
      fontSize: '20px',
      lineHeight: '16px',
      textAlign: 'center',
    },
  }

  return (
    <header style={styles.header}>
      <div style={styles.firstRow}>
        <h1 style={styles.h1}>Ardan's Bets</h1>
        <Button
          classes={'btn-link'}
          clickHandler={() => {}}
          style={{ ...styles.button, ...styles.text }}
        >
          <Link to={'/'}>
            <span style={styles.text}>Dashboard</span>
          </Link>
        </Button>
        <Modal
          show={addBetModal}
          setShow={setAddBetModal}
          trigger={
            <Button
              classes={'btn-link'}
              style={{ ...styles.button, ...styles.text }}
              clickHandler={() => {}}
            >
              Make a bet
            </Button>
          }
          subtitle="Add bet"
        >
          {account ? (
            <EditBet hideModalMethod={setAddBetModal} />
          ) : (
            <Login hideModalMethod={setAddBetModal} />
          )}
        </Modal>
      </div>
      {account ? null : (
        <div style={styles.secondRow}>
          <Modal
            show={loginModal}
            setShow={setLoginModal}
            trigger={
              <Button
                classes={'btn-link'}
                style={{ ...styles.button, ...styles.text }}
                clickHandler={() => {}}
              >
                Sign in
              </Button>
            }
          >
            <Login hideModalMethod={setAddBetModal} />
          </Modal>
        </div>
      )}
    </header>
  )
}

export default AppHeader

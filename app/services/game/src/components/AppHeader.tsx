import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { StyleObject } from '../types/index.d'
import Button from './Button'
import EditBet from './EditBet'
import useEthersConnection from './hooks/useEthersConnection'
import Modal from './modal/Modal'

// AppHeader renders the application header
function AppHeader() {
  // We create a local state to handle if the modals are shown or not
  const [addBetModal, setAddBetModal] = useState(false)

  // Sets local state to trigger re-render when load is complete.
  const [loading, setLoading] = useState(true)

  // ===========================================================================

  // Extracts functions from useEthersConnection Hook
  // useEthersConnection hook handles all connections to ethers.js
  const { setSigner, account, setAccount, provider } = useEthersConnection()

  // ===========================================================================
  // Prompts user to connect to metamask usign the provider
  // and registers it to useEthersConnection hook
  async function init() {
    const signer = provider.getSigner()

    const signerAddress = await signer.getAddress()

    setAccount(signerAddress)

    setSigner(signer)
  }

  function accountsChangeUEFn() {
    init().then(() => setLoading(false))
  }

  // The Effect Hook lets you perform side effects in function components
  // In this case we use it to handle what happens when metamask accounts change.
  // An empty dependecies array triggers useEffect only on the first render
  // of the component. We disable the next line so eslint doens't complain about
  // missing dependencies.

  // eslint-disable-next-line
  useEffect(accountsChangeUEFn, [])

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
        <Link to={'/'}>
          <h1 style={styles.h1}>Ardan's Bets</h1>
        </Link>
        <Button
          classes={'btn-link'}
          clickHandler={() => {}}
          style={{ ...styles.button, ...styles.text }}
        >
          <Link to={'/'}>
            <span style={styles.text}>Dashboard</span>
          </Link>
        </Button>
        {account && window.sessionStorage.getItem('token') ? (
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
            <EditBet hideModalMethod={setAddBetModal} />
          </Modal>
        ) : null}
      </div>
      {account && window.sessionStorage.getItem('token') ? null : (
        <div style={styles.secondRow}>
          <Link to={'/login'}>
            <span style={styles.text}>Sign in</span>
          </Link>
        </div>
      )}
    </header>
  )
}

export default AppHeader

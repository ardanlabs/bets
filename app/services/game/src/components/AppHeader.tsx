import React, { MouseEvent } from 'react'
import { Link } from 'react-router-dom'
import { StyleObject } from '../types/index.d'
import Button from './Button'

// AppHeader renders the application header
function AppHeader() {
  function handleBetClick(event: MouseEvent<HTMLButtonElement>) {
    console.log(event)
  }
  function handleSingInClick(event: MouseEvent<HTMLButtonElement>) {
    console.log(event)
  }

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
          <Link to={'/dashboard'}>
            <span style={styles.text}>Dashboard</span>
          </Link>
        </Button>
        <Button
          classes={'btn-link'}
          style={{ ...styles.button, ...styles.text }}
          clickHandler={handleBetClick}
        >
          Make a bet
        </Button>
      </div>
      <div style={styles.secondRow}>
        <Button
          classes={'btn-link'}
          style={{ ...styles.button, ...styles.text }}
          clickHandler={handleSingInClick}
        >
          Sign in
        </Button>
      </div>
    </header>
  )
}

export default AppHeader

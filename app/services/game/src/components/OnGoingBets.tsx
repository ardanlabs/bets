import { Bet, StyleObject } from '../types/index.d'
import { shortenIfAddress } from '../utils/address'
import getTimeLeft from '../utils/getTimeLeft'
import Button from './Button'
import ExpandIcon from './icons/ExpandIcon'

// OnGoingBets component. Displays site public bets.
function OnGoingBets() {
  // An array of bets
  const bets: Bet[] = [
    {
      id: 1,
      status: 'open',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      name: 'Bruno',
      placerAddress: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challengerAddress: '0x39249126d90671284cd06495d19C04DD0e54d371',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 2,
      status: 'open',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      name: 'Bruno',
      placerAddress: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challengerAddress: '0x39249126d90671284cd06495d19C04DD0e54d371',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 3,
      status: 'open',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      name: 'Bruno',
      placerAddress: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challengerAddress: '0x39249126d90671284cd06495d19C04DD0e54d371',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 4,
      status: 'open',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      name: 'Bruno',
      placerAddress: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challengerAddress: '0x39249126d90671284cd06495d19C04DD0e54d371',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 1,
      status: 'open',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      name: 'Bruno',
      placerAddress: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challengerAddress: '0x39249126d90671284cd06495d19C04DD0e54d371',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 2,
      status: 'open',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      name: 'Bruno',
      placerAddress: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challengerAddress: '0x39249126d90671284cd06495d19C04DD0e54d371',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 3,
      status: 'open',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      name: 'Bruno',
      placerAddress: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challengerAddress: '0x39249126d90671284cd06495d19C04DD0e54d371',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 4,
      status: 'open',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      name: 'Bruno',
      placerAddress: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challengerAddress: '0x39249126d90671284cd06495d19C04DD0e54d371',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
  ]

  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    publicBets: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'start',
      alignItems: 'flex-start',
      gap: '24px',
      padding: '32px 62px',
    },
    bet: {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'flex-start',
      backgroundColor: 'var(--bets-color)',
      boxShadow: '3px 4px 4px rgba(0, 0, 0, 0.25)',
      borderRadius: '8px',
      padding: '8px 16px',
      aspectRatio: '3/2',
    },
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
    row: {
      display: 'flex',
      flexFlow: 'row',
      justifyContent: 'space-between',
      textAlign: 'center',
      alignItems: 'center',
      width: '100%',
      margin: '8px 0',
      gap: '40px',
    },
    fullRow: {
      display: 'flex',
      flexFlow: 'row',
      justifyContent: 'center',
      alignItems: 'center',
      width: '100%',
      textAlign: 'center',
      marginTop: '32px',
      marginBottom: '20px',
    },
    p: {
      margin: '0',
    },
    text: {
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '600',
      fontSize: '16px',
      lineHeight: '18px',
    },
    title: {
      color: 'var(--onGoingBet-title-color)',
    },
    description: {
      color: 'var(--primary-color)',
      fontWeight: '500',
      fontSize: '16px',
      lineHeight: '18px',
      maxWidth: '280px',
    },
    expandButton: {
      height: '24px',
      width: '24px',
      padding: 0,
    },
  }

  function goToBet(id: string) {
    console.log(id)
  }

  return (
    <section style={styles.publicBets}>
      {bets.map((bet) => (
        <div style={styles.bet} key={bet.id}>
          <div style={styles.row}>
            <div style={styles.status}>
              <p style={styles.p}>{bet.status}</p>
            </div>
            <Button
              style={styles.expandButton}
              classes="btn-link"
              id={`${bet.id}`}
              clickHandler={goToBet}
            >
              <ExpandIcon />
            </Button>
          </div>
          <div style={styles.row}>
            <div style={{ ...styles.text, ...styles.title }}>Placer</div>
            <div style={{ ...styles.text }}>
              {shortenIfAddress(bet.placerAddress)}
            </div>
          </div>
          {bet.challengerAddress ? (
            <div style={styles.row}>
              <div style={{ ...styles.text, ...styles.title }}>Challenger</div>
              <div style={{ ...styles.text }}>
                {shortenIfAddress(bet.challengerAddress)}
              </div>
            </div>
          ) : null}
          <div style={styles.row}>
            <div style={{ ...styles.text, ...styles.title }}>Time Left</div>
            <div style={{ ...styles.text }}>
              {getTimeLeft(bet.expirationDate)}
            </div>
          </div>
          <div style={styles.fullRow}>
            <div style={styles.description}>{bet.description}</div>
          </div>
        </div>
      ))}
    </section>
  )
}

export default OnGoingBets

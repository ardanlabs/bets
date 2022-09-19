import { Bet, StyleObject } from '../types/index.d'
import BetCard from './BetCard'
import Subtitle from './Subtitle'

// YourBets component. Displays your personal bets.
function YourBets() {
  // An array of bets
  const bets: Bet[] = [
    {
      id: 1,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 2,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 3,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 4,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 5,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 6,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
  ]

  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    yourBets: {
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
  return (
    <>
      <Subtitle showSearch text="Your bets" />
      <section style={styles.yourBets}>
        {bets.map((bet) => (
          <BetCard key={bet.id} isDetail={false} bet={bet} />
        ))}
      </section>
    </>
  )
}

export default YourBets

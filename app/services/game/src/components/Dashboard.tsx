import React from 'react'
import { StyleObject } from '../types/index.d'
import AppHeader from './AppHeader'
import OnGoingBets from './OnGoingBets'
import Subtitle from './Subtitle'

// Dashboard component
function Dashboard() {
  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    dashboard: {
      width: '100vw',
      flex: '1 1 auto',
      padding: '28px',
    },
  }
  // Renders this final markup
  return (
    <>
      <AppHeader />
      <main style={styles.dashboard}>
        <section>
          <Subtitle showSearch text="Site Ongoing Bets" />
          <OnGoingBets />
        </section>
      </main>
    </>
  )
}

export default Dashboard

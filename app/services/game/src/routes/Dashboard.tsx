import React from 'react'
import OnGoingBets from '../components/OnGoingBets'
import RoutesWrapper from '../components/RoutesWrapper'
import YourBets from '../components/YourBets'

// Dashboard component
function Dashboard() {
  // Renders this final markup
  return (
    <RoutesWrapper>
      <YourBets />
      <OnGoingBets />
    </RoutesWrapper>
  )
}

export default Dashboard

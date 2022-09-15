import React, { useEffect, useState } from 'react'
import { Bet, BetsFilter } from '../types/index.d'
import useApp from './hooks/useApp'

// MainRoom component
function MainRoom() {
  // Variable to set the notification center width.
  const notificationCenterWidth = '340px'

  // Creates a localState to handle bets with the useState hook.
  const [bet, setBet] = useState<JSX.Element>()

  // Extracts getBets from useApp hook.
  const { getBet } = useApp()

  const [filters, setFilters] = useState({} as BetsFilter)

  const initEFn = () => {
    const getBetsFn = (response: Bet) => {
      setBet(
        <li role="bet" key={response.name}>
          {response.name}
        </li>,
      )
    }

    getBet(1).then(getBetsFn)
  }

  // initial useEffect hook.
  useEffect(initEFn, [])

  // Renders this final markup
  return (
    <div
      className="d-flex align-items-center justify-content-start px-0 flex-column"
      style={{ height: '100%', maxHeight: '100vh' }}
    >
      <div className="d-flex" style={{ width: '100vw' }}>
        <section
          style={{
            width: `calc(100% - ${notificationCenterWidth})`,
            zIndex: '1',
          }}
          className="d-flex flex-column align-items-center justify-content-start"
        >
          <ul role="openBets">{bet}</ul>
        </section>
      </div>
    </div>
  )
}

export default MainRoom

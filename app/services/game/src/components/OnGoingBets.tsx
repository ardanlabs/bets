import axios, { AxiosError, AxiosResponse } from 'axios'
import { useEffect, useState } from 'react'
import { Bet, StyleObject } from '../types/index.d'
import { apiUrl } from '../utils/axiosConfig'
import BetCard from './BetCard'
import Subtitle from './Subtitle'

// OnGoingBets component. Displays site public bets.
function OnGoingBets() {
  // We create an state to display the fetched bets
  const [bets, setBets] = useState<Bet[]>([])
  const [page, setPage] = useState(1)
  const [rows, setRows] = useState(20)

  // Initial function executed by initial effect.
  const initEFn = () => {
    const axiosFn = (response: AxiosResponse) => {
      setBets(response.data)
    }
    const axiosCatchFn = (error: AxiosError) => {
      console.error(error)
    }
    axios
      .get(`http://${apiUrl}/bruno/bets/${page}/${rows}`)
      .then(axiosFn)
      .catch(axiosCatchFn)
  }

  // Initial useEffect to fetch and set the bets.
  useEffect(initEFn, [page, rows])

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
  }
  return (
    <>
      <Subtitle showSearch text="Site Ongoing Bets" />
      <section style={styles.publicBets}>
        {bets.map((bet) => (
          <BetCard key={bet.id} isDetail={false} bet={bet} />
        ))}
      </section>
    </>
  )
}

export default OnGoingBets

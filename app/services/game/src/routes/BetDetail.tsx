import axios, { AxiosError, AxiosResponse } from 'axios'
import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import BetCard from '../components/BetCard'
import RoutesWrapper from '../components/RoutesWrapper'
import Subtitle from '../components/Subtitle'
import { Bet, StyleObject } from '../types/index.d'
import { apiUrl } from '../utils/axiosConfig'

// BetDetail route component
function BetDetail() {
  // Brings the bet passed on load from the userLoaderData hook.
  let params = useParams()

  // Creates local state to display the bet right.
  const [bet, setBet] = useState({} as Bet)

  const initFn = () => {
    const initAxiosFn = (response: AxiosResponse) => {
      setBet(response.data)
    }

    const initAxiosCatch = (response: AxiosError) => {
      console.error(response)
    }

    axios
      .get(`http://${apiUrl}/bruno/bet/${params.betId}`)
      .then(initAxiosFn)
      .catch(initAxiosCatch)
  }

  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    betContainer: {
      maxWidth: '800px',
      margin: '0 auto',
    },
    text: {
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '600',
      fontSize: '32px',
      lineHeight: '34px',
    },
    description: {
      fontFamily: 'Roboto Condensed',
      color: 'var(--primary-color)',
      fontWeight: '600',
      fontSize: '40px',
      lineHeight: '42px',
      maxWidth: '700px',
    },
  }

  useEffect(initFn, [params.betId])

  // Renders this markup
  return (
    <RoutesWrapper>
      <Subtitle text="Bet details" showSearch={false} />
      <div style={styles.betContainer}>
        <BetCard
          bet={bet}
          styleObject={{ text: styles.text, description: styles.description }}
          isDetail
        />
      </div>
    </RoutesWrapper>
  )
}

export default BetDetail

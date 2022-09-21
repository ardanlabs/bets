/* ************useGameHook************

  This hook provides all functions needed to connect to the game engine.

**************************************** */

import axios, { AxiosError } from 'axios'
import { useNavigate } from 'react-router-dom'
import { getAppConfig } from '../..'
import { connectResponse } from '../../types/responses.d'
import { apiUrl } from '../../utils/axiosConfig'

// Create an axios instance to keep the token updated
const axiosInstance = axios.create({
  headers: {
    authorization: window.sessionStorage.getItem('token') as string,
  },
})

function useConnectEngine() {
  const navigate = useNavigate()
  // connectToGameEngine connects to the game engine, and stores the token
  // in the sessionStorage. Takes an object with the following type:
  // { dateTime: string; sig: string }
  function connectToGameEngine(data: {
    address: string
    dateTime: string
    sig: string
  }) {
    const axiosFn = (connectResponse: connectResponse) => {
      window.sessionStorage.setItem(
        'token',
        `bearer ${connectResponse.data.token}`,
      )
      const getAppConfigFn = () => {
        navigate('/')
      }
      getAppConfig.then(getAppConfigFn)
    }

    const axiosErrorFn = (error: AxiosError) => {
      const errorMessage = (error as any).response.data.error.replace(
        / \[.+\]/gm,
        '',
      )

      console.group()
      console.error('Error:', errorMessage)
      console.groupEnd()
    }

    axiosInstance
      .post(`http://${apiUrl}/connect`, { ...data })
      .then(axiosFn)
      .catch(axiosErrorFn)
  }

  return {
    connectToGameEngine,
  }
}

export default useConnectEngine

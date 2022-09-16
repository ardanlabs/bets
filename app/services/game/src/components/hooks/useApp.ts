/* ************useGameHook************

  This hook provides all functions needed to run the game.

**************************************** */
import axios, { AxiosError, AxiosResponse } from 'axios'
import {
  Bet,
  BetsFilter,
  DefaultDoc,
  SetModeratorDoc,
  SetWinnerDoc,
} from '../../types/index.d'
import { apiUrl } from '../../utils/axiosConfig'
import docToUint8Array from '../../utils/docToUint8Array'
import useEthersConnection from './useEthersConnection'

// Create an axios instance to keep the token updated
const axiosWithToken = axios.create({
  baseURL: apiUrl,
  headers: {
    authorization: window.sessionStorage.getItem('token') as string,
  },
})

function useApp() {
  // Extracts functions from useEthersConnection Hook
  // useEthersConnection hook handles all connections to ethers.js
  const { signer } = useEthersConnection()

  // signTransaction returns a signed document and signature
  function signTransaction(doc: Object) {
    const parsedDoc = docToUint8Array(doc)

    // signer.signmessage signs the data.
    return signer?.signMessage(parsedDoc)
  }
  // ===========================================================================

  // Game flow functions
  function getBets(filters: BetsFilter) {
    const axiosFn = (response: AxiosResponse) => {
      return response.data
    }
    const axiosErrorFn = (error: AxiosError) => {
      console.error(error)
    }
    return axios.get(`http://${apiUrl}/bets`).then(axiosFn).catch(axiosErrorFn)
  }

  async function getBet(id: number) {
    const axiosFn = (response: AxiosResponse) => {
      return response.data
    }
    const axiosErrorFn = (error: AxiosError) => {
      console.error(error)
    }
    return await axios
      .get(`http://${apiUrl}/bet/${id}`)
      .then(axiosFn)
      .catch(axiosErrorFn)
  }

  async function postBet(bet: Partial<Bet>) {
    const axiosFn = (response: AxiosResponse) => {
      return response.data
    }
    const axiosErrorFn = (error: AxiosError) => {
      if (error.response) {
        // The request was made and the server responded with a status code
        // that falls out of the range of 2xx
        console.log(error.response.data)
        console.log(error.response.status)
        console.log(error.response.headers)
      } else if (error.request) {
        // The request was made but no response was received
        // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
        // http.ClientRequest in node.js
        console.log(error.request)
      } else {
        // Something happened in setting up the request that triggered an Error
        console.log('Error', error.message)
      }
      return error
    }
    return await axios
      .post(`http://${apiUrl}/bet`)
      .then(axiosFn)
      .catch(axiosErrorFn)
  }

  function addMod(doc: SetModeratorDoc): void {
    signTransaction(doc).then((signerResponse: string) => {
      signBet(signerResponse, doc)
    })
  }

  function acceptMod(modAddress: string, address: string, betId: number): void {
    const axiosFn = (response: AxiosResponse) => {}
    const axiosErrorFn = (error: AxiosError) => {
      console.error(error)
    }
    axiosWithToken
      .post('/acceptMod', { modAddress, address, betId })
      .then(axiosFn)
      .catch(axiosErrorFn)
  }

  function signBet(
    signerResponse: string,
    doc: SetModeratorDoc | DefaultDoc | SetWinnerDoc,
  ): void {
    const axiosFn = (response: AxiosResponse) => {}
    const axiosErrorFn = (error: AxiosError) => {
      console.error(error)
    }
    axiosWithToken
      .post('/signBet', { ...doc, signerResponse })
      .then(axiosFn)
      .catch(axiosErrorFn)
  }

  function setWinner(doc: SetWinnerDoc): void {
    signTransaction(doc).then((signerResponse: string) => {
      signBet(signerResponse, doc)
    })
  }

  return {
    getBets,
    getBet,
    postBet,
    addMod,
    acceptMod,
    signBet,
    setWinner,
  }
}

export default useApp

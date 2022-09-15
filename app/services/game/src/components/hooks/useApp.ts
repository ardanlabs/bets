/* ************useGameHook************

  This hook provides all functions needed to run the game.

**************************************** */
import axios, { AxiosError, AxiosResponse } from 'axios'
import { useNavigate } from 'react-router-dom'
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
const axiosInstance = axios.create({
  baseURL: apiUrl,
  // headers: {
  //   authorization: window.sessionStorage.getItem('token') as string,
  // },
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
      console.log(error)
    }
    return axios.get(`http://${apiUrl}/bets`).then(axiosFn).catch(axiosErrorFn)
  }

  async function getBet(id: number) {
    const axiosFn = (response: AxiosResponse) => {
      return response.data
    }
    return await axiosInstance.get(`http://${apiUrl}/bet/${id}`).then(axiosFn)
  }

  function postBet(bet: Bet) {
    const axiosFn = (response: AxiosResponse) => {
      console.log(response)
    }
    const axiosCatchFn = () => {}
    axiosInstance.post(`/bet`, bet).then(axiosFn).catch(axiosCatchFn)
  }

  function addMod(doc: SetModeratorDoc): void {
    const addModFn = (signerResponse: string) => {
      const axiosFn = (response: AxiosResponse) => {}
      const axiosCatchFn = () => {}
      axiosInstance
        .post('/signBet', { ...doc, signerResponse })
        .then(axiosFn)
        .catch(axiosCatchFn)
    }

    signTransaction(doc).then(addModFn)
  }

  function acceptMod(modAddress: string, address: string, betId: number): void {
    const axiosFn = (response: AxiosResponse) => {}
    const axiosCatchFn = () => {}
    axiosInstance
      .post('/acceptMod', { modAddress, address, betId })
      .then(axiosFn)
      .catch(axiosCatchFn)
  }

  function signBet(doc: DefaultDoc): void {
    const betSignerFn = (signerResponse: string) => {
      const axiosFn = (response: AxiosResponse) => {}
      const axiosCatchFn = () => {}
      axiosInstance
        .post('/signBet', { ...doc, signerResponse })
        .then(axiosFn)
        .catch(axiosCatchFn)
    }

    signTransaction(doc).then(betSignerFn)
  }

  function setWinner(doc: SetWinnerDoc): void {
    const winnerSignerFn = (signerResponse: string) => {
      const axiosFn = (response: AxiosResponse) => {}
      const axiosCatchFn = () => {}
      axiosInstance
        .post('/setWinner', { ...doc, signerResponse })
        .then(axiosFn)
        .catch(axiosCatchFn)
    }

    signTransaction(doc).then(winnerSignerFn)
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

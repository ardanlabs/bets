import axios, { AxiosError, AxiosResponse } from 'axios'
import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Bet, StyleObject } from '../types/index.d'
import { EditBetProps } from '../types/props.d'
import { apiUrl } from '../utils/axiosConfig'
import Button from './Button'
import SuccessModal from '../components/SuccessModal'
import getExpirationDate from '../utils/getExpirationDate'

function EditBet(props: EditBetProps) {
  // Extracts props
  const { bet, hideModalMethod } = props

  // We set a local state to manage input changes
  const [localBet, setLocalBet] = useState<Bet | Object>(
    bet
      ? bet
      : ({
          id: '',
          status: '',
          players: [
            { address: '', signed: false },
            { address: '', signed: false },
          ],
          moderator: '',
          description: '',
          terms: '',
          expirationDate: new Date().getTime(),
          amount: 0,
        } as Bet),
  )

  // We set localstate to manage success modal
  const [show, setShow] = useState(false)

  // We extract navigation functionalitys
  const navigate = useNavigate()

  // ===========================================================================

  function submitBet() {
    const submitBetAxiosFn = (response: AxiosResponse) => {
      if (!bet) {
        navigate(`/bet/${response.data.betId}`)
      }
    }
    const submitBetAxiosCatchFn = (error: AxiosError) => {
      console.error(error)
    }

    axios
      .post(`http://${apiUrl}/${bet ? `editBet/${bet.id}` : 'bet'}`)
      .then(submitBetAxiosFn)
      .catch(submitBetAxiosCatchFn)

    // Add bet submit
    hideModalMethod(false)
  }

  // handlePlayersChange keeps track of the input value change in the local state.
  function handlePlayersChange(
    event: React.ChangeEvent<HTMLInputElement>,
    playerIndex: number,
  ) {
    const inputValue = event.target.value

    setLocalBet((prevState: Bet) => {
      let prevPlayers = prevState.players
        ? prevState.players
        : [
            { address: '', signed: false },
            { address: '', signed: false },
          ]
      let newPlayer = prevPlayers[playerIndex]

      newPlayer.address = inputValue

      const players = [...prevPlayers]
      return { ...prevState, players }
    })
  }

  // handleFormChange keeps track of the input value change in the local state.
  function handleFormChange(event: React.ChangeEvent<HTMLInputElement>) {
    const changedInput = event.target.id
    const inputValue = event.target.value

    let pairKeyValue: any = {}
    pairKeyValue[changedInput] = inputValue

    setLocalBet((prevState) => {
      return { ...prevState, ...pairKeyValue }
    })
  }

  const styles: StyleObject = {
    form: {
      width: '100%',
    },
    formLabel: {
      marginBottom: '8px',
    },
    formInputWrapper: {
      marginBottom: '24px',
      display: 'flex',
      flexFlow: 'column',
      justifyContent: 'center',
      alignItems: 'flex-start',
      margin: '24px 10px 0px 10px',
    },
    formInput: {
      height: '48px',
      aspectRatio: '6/1',
    },
    formPlaceHolder: {
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '500',
      fontSize: '14px',
      lineHeight: '16px',
      color: '#d2d2d2',
    },
    row: {
      display: 'flex',
      flexFlow: 'row',
      justifyContent: 'space-between',
      textAlign: 'center',
      alignItems: 'center',
      width: '100%',
      margin: '8px 0',
    },
    fullRow: {
      display: 'flex',
      flexFlow: 'row',
      justifyContent: 'flex-end',
      alignItems: 'center',
      width: '100%',
      textAlign: 'center',
      marginTop: '32px',
      marginBottom: '20px',
    },
    text: {
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '600',
      fontSize: '16px',
      lineHeight: '18px',
    },
  }
  return (
    <>
      <SuccessModal show={show} setShow={setShow} betId={'1'}></SuccessModal>
      <form style={styles.form}>
        <div style={styles.row}>
          {(localBet as Bet).players.map((player, index) => (
            <div key={index} style={styles.formInputWrapper}>
              <label
                style={{ ...styles.text, ...styles.formLabel }}
                htmlFor={`player-${index}`}
              >
                Player {index + 1}
              </label>
              <input
                onChange={(e) => handlePlayersChange(e, index)}
                value={player.address}
                style={{ ...styles.formInput }}
                type="text"
                className="formInputs form-control"
                placeholder={`Player ${index + 1}`}
                id={`player-${index}`}
              />
            </div>
          ))}
        </div>
        <div style={styles.row}>
          <div style={styles.formInputWrapper}>
            <label
              style={{ ...styles.text, ...styles.formLabel }}
              htmlFor="moderator"
            >
              Moderator address
            </label>
            <input
              onChange={handleFormChange}
              value={
                Object.keys(localBet).length ? (localBet as Bet).moderator : ''
              }
              style={{ ...styles.formInput }}
              type="text"
              className="formInputs form-control"
              placeholder="Moderator address"
              id="moderator"
            />
          </div>
        </div>
        <div style={styles.row}>
          <div style={styles.formInputWrapper}>
            <label
              style={{ ...styles.text, ...styles.formLabel }}
              htmlFor="description"
            >
              Description
            </label>
            <input
              onChange={handleFormChange}
              value={
                Object.keys(localBet).length
                  ? (localBet as Bet).description
                  : ''
              }
              style={{ ...styles.formInput }}
              type="text"
              className="formInputs form-control"
              placeholder="Description"
              id="description"
            />
          </div>
          <div style={styles.formInputWrapper}>
            <label
              style={{ ...styles.text, ...styles.formLabel }}
              htmlFor="terms"
            >
              Terms
            </label>
            <input
              onChange={handleFormChange}
              value={
                Object.keys(localBet).length ? (localBet as Bet).terms : ''
              }
              style={{ ...styles.formInput }}
              type="text"
              className="formInputs form-control"
              placeholder="Terms"
              id="terms"
            />
          </div>
        </div>
        <div style={styles.row}>
          <div style={styles.formInputWrapper}>
            <label
              style={{ ...styles.text, ...styles.formLabel }}
              htmlFor="expirationDate"
            >
              Expiration date
            </label>
            <input
              onChange={handleFormChange}
              value={
                Object.keys(localBet).length
                  ? getExpirationDate((localBet as Bet).expirationDate)
                  : ''
              }
              style={{ ...styles.formPlaceHolder, ...styles.formInput }}
              type="date"
              className="formInputs form-control"
              id="expirationDate"
            />
          </div>
          <div style={styles.formInputWrapper}>
            <label
              style={{ ...styles.text, ...styles.formLabel }}
              htmlFor="amount"
            >
              Amount
            </label>
            <input
              onChange={handleFormChange}
              value={
                Object.keys(localBet).length ? (localBet as Bet).amount : ''
              }
              style={{ ...styles.formInput }}
              type="text"
              className="formInputs form-control"
              placeholder="Amount"
              id="amount"
            />
          </div>
        </div>
        <div style={styles.fullRow}>
          <Button
            classes="btn-link btn-outline-primary"
            style={{
              position: 'relative',
              display: 'inline-block',
              cursor: 'pointer',
              width: 'auto',
            }}
            clickHandler={submitBet}
            id={bet ? `${bet.id}` : 'edit'}
          >
            {bet ? 'Save' : 'Create bet'}
          </Button>
        </div>
      </form>
    </>
  )
}

export default EditBet

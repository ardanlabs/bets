/* ************useWebSocketHook************

  This hook is in charge of providing the websocket connection.
  When you call connect you are connecting directly to the backend websocket endpoint. /events
  On every message send by the backend we will notify all browsers about the changes, aswell as notify them about the actions taken.

  **************************************** */

// import { shortenIfAddress } from '../../utils/address'
import { apiUrl } from '../../utils/axiosConfig'

function useWebSocket(restart: () => void) {
  // Connects to the webscoket.
  async function connect() {
    const ws = new WebSocket(`ws://${apiUrl}/events`)

    // ws.onopen binds an event listener that triggers with the "open" event.
    ws.onopen = () => {
      console.log('opened')
    }

    // ws.onmessage binds an event listener that triggers with "message" event.
    ws.onmessage = (evt: MessageEvent) => {
      if (evt.data) {
        let message = JSON.parse(evt.data)
        // const messageAccount = shortenIfAddress(message.address)
        console.log(message)
      }
      return
    }

    // ws.onclose binds an event listener that triggers with "close" event.
    // If the socket closes we show the user an error and set the game to
    // it's initial state.
    ws.onclose = (evt: CloseEvent) => {
      restart()
    }

    // ws.onerror binds an event listener that triggers with "error" event.
    ws.onerror = function (err) {
      console.error('Socket encountered error: ', err, 'Closing socket')
      ws.close()
    }
  }
  return { connect }
}

export default useWebSocket

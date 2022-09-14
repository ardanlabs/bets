import React, { useEffect } from 'react'
import useWebSocket from './hooks/useWebSocket'
import { token } from '../utils/axiosConfig'
import { useLocation, useNavigate } from 'react-router-dom'
import { appConfig } from '../types/index.d'
import useEthersConnection from './hooks/useEthersConnection'

// MainRoom component
function MainRoom() {
  // Extracts navigate from useNavigate Hook.
  const navigate = useNavigate()

  // Extracts state (a prop send by the router) from useLocation Hook.
  const { state } = useLocation()

  // Extracts account from ethersConnection Hook.
  const { account } = useEthersConnection()

  // Extracts function to connect to ws (connect) from useWebSocket Hook.
  const { connect } = useWebSocket(() => {})

  // Variable to set the notification center width.
  const notificationCenterWidth = '340px'

  const wsStatus = window.sessionStorage.getItem('wsStatus')

  // ===========================================================================

  // initUEFn connects the websocket, clears the round timer and
  // sets Player dice if needed.
  const initUEFn = () => {
    // Connects to websocket depending on status.
    function connectToWs() {
      connect().then(() => {
        window.sessionStorage.setItem('wsStatus', 'open')
      })
    }
    if (wsStatus !== 'open' && wsStatus !== 'attemptingConnection') {
      window.sessionStorage.setItem('wsStatus', 'attemptingConnection')
      connectToWs()
    }
    window.sessionStorage.setItem('wsStatus', 'close')
  }

  // An empty dependecies array triggers useEffect only on the first render of the component
  // We disable the next line so eslint doens't complain about missing dependencies.
  // eslint-disable-next-line
  useEffect(initUEFn, [])

  // ===========================================================================

  const authUEFn = () => {
    // Handles if the user is logged and has a token.
    // If not, we redirect it to the login page. (<Login />)
    function checkAuth() {
      if (!account || !token() || !(state as appConfig)) {
        navigate('/')
      }
    }

    checkAuth()
  }

  // eslint-disable-next-line
  useEffect(authUEFn, [account, state])

  //
  // ===============================Finish Timer================================

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
        ></section>
      </div>
    </div>
  )
}

export default MainRoom

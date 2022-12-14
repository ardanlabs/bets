import React, { useState, useEffect } from 'react'
import './App.scss'
import { appConfig } from './types/index.d'
import { Route, Routes, useLocation, useNavigate } from 'react-router-dom'
import { getAppConfig } from '.'
import { utils } from 'ethers'
import {
  ethersConnectionInterface,
  EthersContext,
  ethersContextInterface,
} from './contexts/ethersContext'
import useEthersConnection from './components/hooks/useEthersConnection'
import WrongNetwork from './routes/WrongNetwork'
import { Network } from '@ethersproject/networks'
import Dashboard from './routes/Dashboard'
import BetDetail from './routes/BetDetail'
import Login from './routes/Login'
import { WalletProvider } from '@viaprotocol/web3-wallets'

// =============================================================================

// App component. First React component after Index.ts
function App() {
  // Extracts provider from useEthersConnection hook
  const { provider, switchNetwork } = useEthersConnection()

  // Extracts router navigation functionality
  const navigate = useNavigate()

  // Gets location from useLocation Hook.
  const location = useLocation()

  // Sets state for the ethersConnection context
  const [ethersConnection, setEthersConnection] =
    useState<ethersConnectionInterface>({} as ethersConnectionInterface)

  // ===========================================================================

  // hideDropdowns handles the clicks outside a dropdown.
  function hideDropdowns(event: React.MouseEvent<HTMLDivElement, MouseEvent>) {
    const dropdown = document.querySelector('.dropdown-menu') as HTMLElement
    const isMenu = (event.target as HTMLElement).classList.contains(
      'dropdown-content',
    )

    if (!isMenu && dropdown) {
      dropdown.style.display = 'none'
    }
  }

  // getEthersContextDefaultValue returns the default context value for the
  // ethers.js support. This context is used for creating global support for
  // ethers.js connection.
  function getEthersContextDefaultValue(): ethersContextInterface {
    return { ethersConnection, setEthersConnection }
  }

  // handleNetworkChange changes of network. When a Provider makes its initial
  // connection,it emits a "network" event with a null oldNetwork along with the
  // newNetwork. So, if the oldNetwork exists, it represents a changing network.
  function handleNetworkChange(newNetwork: Network): void {
    const fn = async (getConfigResponse: appConfig) => {
      if (newNetwork.chainId !== getConfigResponse.chainId) {
        window.sessionStorage.removeItem('token')
        // It navigates to the wrongNetwork component, and switches the network.
        // If the network isn't switched, it will stay at that component until it does.
        navigate('/wrongNetwork', { state: { ...getConfigResponse } })
        switchNetwork({
          chainId: utils.hexValue(getConfigResponse.chainId),
        })
        return
      }
      if (location.pathname === 'wrongNetwork') navigate('/')
    }
    getAppConfig.then(fn)
  }

  // ===========================================================================

  // effectFn handles network changes.
  const effectFn = () => {
    provider.on('network', (newNetwork, _) => handleNetworkChange(newNetwork))
  }
  // We disable the next line so eslint doens't complain about missing dependencies.
  // eslint-disable-next-line
  useEffect(effectFn, [])

  // ===========================================================================

  // Renders this final markup.
  return (
    <div
      className="App"
      style={{ scrollSnapType: 'y mandatory' }}
      onClick={hideDropdowns}
    >
      <WalletProvider>
        <EthersContext.Provider value={getEthersContextDefaultValue()}>
          <Routes>
            <Route path="/" element={<Dashboard />}></Route>
            <Route path="/login" element={<Login />}></Route>
            <Route path="/bet/:betId" element={<BetDetail />}></Route>
            <Route path="/wrongNetwork" element={<WrongNetwork />}></Route>
          </Routes>
        </EthersContext.Provider>
      </WalletProvider>
    </div>
  )
}

export default App

import React, { useContext, useEffect, useState } from 'react'
import Button from '../components/Button'
import MetamaskLogo from '../components/icons/metamask'
import getNowDate from '../utils/getNowDate'
import useEthersConnection from '../components/hooks/useEthersConnection'
import useConnectEngine from '../components/hooks/useConnectEngine'
import { getAppConfig } from '..'
import { useNavigate } from 'react-router-dom'
import { WalletContext } from '@viaprotocol/web3-wallets'
import { Web3Provider } from '@ethersproject/providers'
import WalletConnectLogo from '../components/icons/walletconnect'
import CoinBaseLogo from '../components/icons/coinbase'
import LogOutIcon from '../components/icons/logout'

// Login component.
// Doesn't receives parameters
// Outputs the Html output for login in.
// Handles logic for login in and also connecting to the game engine.
function Login() {
  // ===========================================================================
  // Hooks setup.

  // Extracts navigate from useNavigate Hook
  const navigate = useNavigate()

  // Extracts connection to engine from useConnectEngine hook
  const { connectToGameEngine } = useConnectEngine()

  // Extracts functions from useEthersConnection Hook
  // useEthersConnection hook handles all connections to ethers.js
  const { setSigner, account, setAccount } = useEthersConnection()

  const { connect, isConnected, disconnect, address, signMessage, provider } =
    useContext(WalletContext)

  // Sets local state to trigger re-render when load is complete.
  const [loading, setLoading] = useState(true)

  // Prompts user to connect to metamask usign the provider
  // and registers it to useEthersConnection hook
  async function init() {
    if (provider instanceof Web3Provider) {
      const signer = provider.getSigner()
      setSigner(signer)
    }

    setAccount(address)
  }

  // ===========================================================================

  function accountsChangeUEFn() {
    // handleAccountsChanged handles what happen when metamask account changes
    // If the array of accounts is non-empty, you're already connected.
    function handleAccountsChanged(accounts: string[]) {
      if (accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        setLoading(true)
        return
      }
      window.sessionStorage.removeItem('token')
    }

    function connectFn() {
      init().then(() => setLoading(false))
    }
    // Note that this event is emitted on page load.
    window.ethereum.on('accountsChanged', handleAccountsChanged)
    // This event checks when the provider becomes able to submit
    // RPC requests to a chain.
    window.ethereum.on('connect', connectFn)
    init().then(() => setLoading(false))
  }

  // The Effect Hook lets you perform side effects in function components
  // In this case we use it to handle what happens when metamask accounts change.
  // An empty dependecies array triggers useEffect only on the first render
  // of the component. We disable the next line so eslint doens't complain about
  // missing dependencies.

  // eslint-disable-next-line
  useEffect(accountsChangeUEFn, [])

  // ===========================================================================

  // loggedUEFn handles what happen if you're already log when you enter the app
  const loggedUEFn = () => {
    if (window.sessionStorage.getItem('token') && account) {
      getAppConfig.then((response) => {
        navigate(-1)
      })
    }
  }

  // Next line is disabled so eslint doens't complain about missing dependencies.

  // eslint-disable-next-line
  useEffect(loggedUEFn, [account])

  // ===========================================================================
  //
  // End of hooks.
  //
  // ===========================================================================

  // handleConnectAccount takes care of the connection to the browser wallet.
  async function handleConnectAccount(wallet: { name: any; chainId: any }) {
    await connect(wallet).then(() => {
      init().then(() => setLoading(false))
    })
  }

  // signTransaction handles click on sign transaction
  // Creates a document to sign, signs the document and connects to game engine.
  function signTransaction() {
    const date = getNowDate()

    const doc = { address: address as string, dateTime: date }

    const parsedDoc = JSON.stringify(doc)

    // signerFn connects to the game Engine sending the signature and the signed document.
    const signerFn = (signerResponse: any) => {
      const data = { ...doc, sig: signerResponse }
      connectToGameEngine(data)
    }

    // signer.signmessage signs the data. The underlying code will apply the Ardan stamp and
    // ID to the signature thanks to changes made to the ether.js api.
    signMessage(parsedDoc).then(signerFn)
  }

  // ===========================================================================

  // Renders this final markup.
  return (
    <div
      className="container-fluid d-flex align-items-center justify-content-center px-0 flex-column"
      style={{
        display: 'flex',
        alignItems: 'center',
        height: 'calc(100vh - 70px)',
      }}
    >
      <div
        id="login__wrapper"
        className="d-flex align-items-start justify-content-center flex-column mt-10"
      >
        <h2>
          <strong> Connect your wallet </strong>
        </h2>
        Or you can also select a provider to create one.
        <div id="wallets__wrapper" className="mt-4">
          {isConnected && !loading ? (
            <div className="d-flex">
              <span className="ml-2 px-2 py-2">Wallet {address} connected</span>
              <div
                onClick={() => disconnect()}
                className="mx-2 px-2 py-2"
                style={{ cursor: 'pointer' }}
              >
                <LogOutIcon />
              </div>
            </div>
          ) : (
            <div>
              <Button
                {...{
                  id: 'metamask__wrapper',
                  clickHandler: () =>
                    handleConnectAccount({ name: 'MetaMask', chainId: 1 }),
                  classes: 'd-flex align-items-center',
                }}
              >
                <MetamaskLogo {...{ width: '50px', height: '50px' }} />
                <span className="ms-4"> Metamask </span>
              </Button>
              <Button
                {...{
                  id: 'coinbase__wrapper',
                  clickHandler: () =>
                    handleConnectAccount({ name: 'Coinbase', chainId: 1337 }),
                  classes: 'd-flex align-items-center',
                }}
              >
                <CoinBaseLogo {...{ width: '50px', height: '50px' }} />
                <span className="ms-4"> Coinbase </span>
              </Button>
              <Button
                {...{
                  id: 'walletConnect__wrapper',
                  clickHandler: () =>
                    handleConnectAccount({ name: 'WalletConnect', chainId: 1 }),
                  classes: 'd-flex align-items-center',
                }}
              >
                <WalletConnectLogo />
                <span className="ms-4"> Wallet Connect</span>
              </Button>
            </div>
          )}
        </div>
        <div id="wallets__wrapper" className="mt-4">
          <Button
            {...{
              id: 'metamask__wrapper',
              clickHandler: signTransaction,
              classes: 'd-flex align-items-center',
            }}
          >
            <>Sign into app</>
          </Button>
        </div>
      </div>
    </div>
  )
}

export default Login

import React, { useEffect, useState } from 'react'
import Button from '../components/Button'
import MetamaskLogo from '../components/icons/metamask'
import getNowDate from '../utils/getNowDate'
import { useNavigate } from 'react-router-dom'
import { getAppConfig } from '..'
import useEthersConnection from '../components/hooks/useEthersConnection'
import docToUint8Array from '../utils/docToUint8Array'
import useConnectEngine from '../components/hooks/useConnectEngine'

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
  const { setSigner, account, signer, setAccount, provider } =
    useEthersConnection()

  // Sets local state to trigger re-render when load is complete.
  const [loading, setLoading] = useState(true)

  // Prompts user to connect to metamask usign the provider
  // and registers it to useEthersConnection hook
  async function init() {
    const signer = provider.getSigner()

    const signerAddress = await signer.getAddress()

    setAccount(signerAddress)

    setSigner(signer)
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
        navigate('/', {
          state: { ...response, reload: true },
          replace: true,
        })
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
  async function handleConnectAccount() {
    await provider.send('eth_requestAccounts', []).then((accounts: string) => {
      init().then(() => setLoading(false))
    })
  }

  // signTransaction handles click on sign transaction
  // Creates a document to sign, signs the document and connects to game engine.
  function signTransaction() {
    const date = getNowDate()

    const doc = { address: account as string, dateTime: date }

    const parsedDoc = docToUint8Array(doc)

    // signerFn connects to the game Engine sending the signature and the signed document.
    const signerFn = (signerResponse: any) => {
      const data = { ...doc, sig: signerResponse }
      connectToGameEngine(data)
    }

    // signer.signmessage signs the data. The underlying code will apply the Ardan stamp and
    // ID to the signature thanks to changes made to the ether.js api.
    signer?.signMessage(parsedDoc).then(signerFn)
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
          {account && !loading ? (
            <div className="d-flex">
              <span className="ml-2 px-2 py-2">Wallet {account} connected</span>
            </div>
          ) : (
            <Button
              {...{
                id: 'metamask__wrapper',
                clickHandler: handleConnectAccount,
                classes: 'd-flex align-items-center',
              }}
            >
              <MetamaskLogo {...{ width: '50px', height: '50px' }} />
              <span className="ml-4"> Metamask </span>
            </Button>
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
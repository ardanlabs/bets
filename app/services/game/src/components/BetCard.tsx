import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { StyleObject } from '../types/index.d'
import { BetCardProps } from '../types/props.d'
import { shortenIfAddress } from '../utils/address'
import getTimeLeft from '../utils/getTimeLeft'
import Button from './Button'
import Card from './Card'
import EditBet from './EditBet'
import useEthersConnection from './hooks/useEthersConnection'
import EditIcon from './icons/EditIcon'
import ExpandIcon from './icons/ExpandIcon'
import Modal from './modal/Modal'
import BetStatus from './BetStatus'

function BetCard(props: BetCardProps) {
  // Extracts props
  const { bet, styleObject, isDetail } = props

  // Extracts account from useEthersConnection hook
  const { account } = useEthersConnection()

  // Extracts navigate functionality from React Router useNavigate hook.
  const navigate = useNavigate()

  // We create 2 states to handle what modals are open
  const [addModeratorModal, setAddModeratorModal] = useState(false)
  const [editBetModal, setEditBetModal] = useState(false)

  // ===========================================================================

  // Navigates to BetDetail route
  function handleAction(id: string) {
    if (isDetail) {
      navigate(`/editBet/${id}`)
      return
    }
    navigate(`/bet/${id}`)
  }

  // ===========================================================================

  // Logic for showing the add mod action button.
  const showAddModeratorButton =
    isDetail &&
    !bet.moderator &&
    (bet.placer === account || bet.challenger === account)

  // Logic for showing the edit action button.
  const showEditButton =
    isDetail &&
    bet.status !== 'live' &&
    bet.status !== 'canceled' &&
    bet.status !== 'moderate' &&
    (bet.placer === account || bet.challenger === account)

  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    column: {
      display: 'flex',
      flexFlow: 'column',
      justifyContent: 'flex-start',
      alignItems: 'flex-start',
      margin: '8px 0',
    },
    row: {
      display: 'flex',
      flexFlow: 'row',
      justifyContent: 'space-between',
      textAlign: 'center',
      alignItems: 'center',
      width: '100%',
      margin: '8px 0',
      gap: '20px',
    },
    fullRow: {
      display: 'flex',
      flexFlow: 'row',
      justifyContent: 'center',
      alignItems: 'center',
      width: '100%',
      textAlign: 'center',
      marginTop: '32px',
      marginBottom: '20px',
    },
    p: {
      margin: '0',
    },
    text: {
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: '600',
      fontSize: '16px',
      lineHeight: '18px',
    },
    title: {
      color: 'var(--onGoingBet-title-color)',
    },
    description: {
      fontFamily: 'Roboto Condensed',
      color: 'var(--primary-color)',
      fontWeight: '600',
      fontSize: '16px',
      lineHeight: '18px',
      maxWidth: '200px',
    },
    iconButton: {
      height: '24px',
      width: '24px',
      padding: 0,
      margin: 0,
    },
    addModButton: {
      width: 'fit-content',
      borderRadius: '8px',
    },
    terms: {
      fontStyle: 'normal',
      fontWeight: '500',
      fontSize: '16px',
      lineHeight: '19px',
      color: '#515151',
      paddingLeft: '16px',
      margin: '8px 0',
    },
    card: {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'flex-start',
      backgroundColor: 'var(--bets-color)',
      boxShadow: '3px 4px 4px rgba(0, 0, 0, 0.25)',
      borderRadius: '8px',
      padding: '0 16px',
      aspectRatio: '3/2',
    },
    ...styleObject,
  }

  return bet ? (
    <Card styleObject={{ card: styles.card }}>
      <div style={styles.row}>
        <BetStatus status={bet.status} />
        <div style={{ ...styles.row, justifyContent: 'flex-end' }}>
          {showAddModeratorButton ? (
            <Modal
              show={addModeratorModal}
              setShow={setAddModeratorModal}
              trigger={
                <Button
                  classes="btn-link btn-outline-primary"
                  style={{
                    position: 'relative',
                    display: 'inline-block',
                    cursor: 'pointer',
                    width: 'auto',
                  }}
                  clickHandler={() => {}}
                  id="btn"
                >
                  Add moderator
                </Button>
              }
              subtitle="Add Moderator"
            >
              <EditBet hideModalMethod={setAddModeratorModal} />
            </Modal>
          ) : null}

          {isDetail ? (
            showEditButton ? (
              <Modal
                show={editBetModal}
                setShow={setEditBetModal}
                trigger={
                  <Button
                    style={styles.iconButton}
                    classes="btn-link"
                    id={`${bet.id}`}
                    clickHandler={() => {}}
                  >
                    <EditIcon />
                  </Button>
                }
                subtitle="Edit bet"
              >
                <EditBet hideModalMethod={setEditBetModal} />
              </Modal>
            ) : null
          ) : (
            <Button
              style={styles.iconButton}
              classes="btn-link"
              id={`${bet.id}`}
              clickHandler={handleAction}
            >
              <ExpandIcon />
            </Button>
          )}
        </div>
      </div>
      {bet.placer ? (
        <div style={styles.row}>
          <div style={{ ...styles.text, ...styles.title }}>Placer</div>
          <div style={{ ...styles.text }}>{shortenIfAddress(bet.placer)}</div>
        </div>
      ) : null}
      {bet.challenger && isDetail ? (
        <div style={styles.row}>
          <div style={{ ...styles.text, ...styles.title }}>Challenger</div>
          <div style={{ ...styles.text }}>
            {shortenIfAddress(bet.challenger)}
          </div>
        </div>
      ) : null}
      {bet.moderator && isDetail ? (
        <div style={styles.row}>
          <div style={{ ...styles.text, ...styles.title }}>Moderator</div>
          <div style={{ ...styles.text }}>
            {bet.moderator ? shortenIfAddress(bet.moderator) : 'Not assigned'}
          </div>
        </div>
      ) : null}
      {bet.amount ? (
        <div style={styles.row}>
          <div style={{ ...styles.text, ...styles.title }}>Stakes</div>
          <div style={{ ...styles.text }}>{bet.amount}</div>
        </div>
      ) : null}
      {bet.expirationDate ? (
        <div style={styles.row}>
          <div style={{ ...styles.text, ...styles.title }}>Time Left</div>
          <div style={{ ...styles.text }}>
            {getTimeLeft(bet.expirationDate)}
          </div>
        </div>
      ) : null}
      {bet.expirationDate && isDetail ? (
        <div style={styles.row}>
          <div style={{ ...styles.text, ...styles.title }}>Expiration Date</div>
          <div style={{ ...styles.text }}>{bet.expirationDate}</div>
        </div>
      ) : null}
      {bet.terms && isDetail ? (
        <div style={styles.column}>
          <div style={{ ...styles.text, ...styles.title }}>Terms</div>
          <div style={{ ...styles.text, ...styles.terms }}>{bet.terms}</div>
        </div>
      ) : null}
      <div style={styles.fullRow}>
        <div style={styles.description}>{bet.description}</div>
      </div>
    </Card>
  ) : null
}

export default BetCard

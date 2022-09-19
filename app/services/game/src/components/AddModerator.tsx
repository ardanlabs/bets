import { StyleObject } from '../types/index.d'
import { EditBetProps } from '../types/props.d'
import Button from './Button'

function EditBet(props: EditBetProps) {
  // Extracts props
  const { bet, hideModalMethod } = props

  // ===========================================================================

  function submitBet() {
    hideModalMethod(false)
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
      <form style={styles.form}>
        <div style={styles.row}>
          <div style={styles.formInputWrapper}>
            <label
              style={{ ...styles.text, ...styles.formLabel }}
              htmlFor="placer"
            >
              Placer Address
            </label>
            <input
              value={bet?.placer}
              style={{ ...styles.formPlaceHolder, ...styles.formInput }}
              type="text"
              className="formInputs form-control"
              placeholder="Placer Address"
              id="placer"
            />
          </div>
          <div style={styles.formInputWrapper}>
            <label
              style={{ ...styles.text, ...styles.formLabel }}
              htmlFor="challenger"
            >
              Challenger address
            </label>
            <input
              value={bet?.challenger}
              style={{ ...styles.formPlaceHolder, ...styles.formInput }}
              type="text"
              className="formInputs form-control"
              placeholder="Challenger address"
              id="challenger"
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
              value={bet?.description}
              style={{ ...styles.formPlaceHolder, ...styles.formInput }}
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
              value={bet?.terms}
              style={{ ...styles.formPlaceHolder, ...styles.formInput }}
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
              value={bet?.expirationDate}
              style={{ ...styles.formPlaceHolder, ...styles.formInput }}
              type="date"
              className="formInputs form-control"
              placeholder="Expiration date"
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
              value={bet?.amount}
              style={{ ...styles.formPlaceHolder, ...styles.formInput }}
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
            Create bet
          </Button>
        </div>
      </form>
    </>
  )
}

export default EditBet

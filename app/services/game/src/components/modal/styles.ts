import { StyleObject } from '../../types/index.d'

export const styles: StyleObject = {
  modalBox: {
    position: 'relative',
    zIndex: '2',
    width: '400px',
    height: '200px',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
  },
  modalContent: {
    padding: '5px',
  },
  modalContainer: {
    position: 'absolute',
    margin: '0 auto',
    height: '100vh',
    width: '100vw',
    background: 'rgb(0 0 0 / 54%)',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    top: '0',
    left: '0',
  },
}

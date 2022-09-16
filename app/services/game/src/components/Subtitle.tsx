import { StyleObject } from '../types/index.d'
import { SubtitleProps } from '../types/props.d'
import SearchIcon from './icons/SearchIcon'

function Subtitle(props: SubtitleProps) {
  // Extracts props
  const { showSearch, text } = props

  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    subtitle: {
      fontFamily: 'Roboto Condensed',
      fontStyle: 'normal',
      fontWeight: 700,
      fontSize: '24px',
      lineHeight: '28px',
      color: '#000000',
      display: 'flex',
      justifyContent: 'space-between',
      gap: '8px',
      alignContent: 'center',
      width: 'fit-content',
    },
  }

  return (
    <div style={styles.subtitle}>
      {text} {showSearch ? <SearchIcon /> : null}
    </div>
  )
}

export default Subtitle

import React from 'react'
import { StyleObject } from '../types/index.d'
import { RoutesWrapperProps } from '../types/props.d'
import AppHeader from './AppHeader'

// RoutesWrapper, component to wrapp all routes to mantain design
function RoutesWrapper(props: RoutesWrapperProps) {
  // Extracts children from props/
  const { children } = props

  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    mainWrapper: {
      width: '100vw',
      flex: '1 1 auto',
      padding: '28px',
    },
  }
  // Renders this final markup
  return (
    <>
      <AppHeader />
      <main style={styles.mainWrapper}>
        <section>{children}</section>
      </main>
    </>
  )
}

export default RoutesWrapper

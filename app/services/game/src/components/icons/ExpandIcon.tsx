import React from 'react'

interface ExpandIconProps {
  width?: string
  height?: string
}

function ExpandIcon(props: ExpandIconProps) {
  let { width, height } = props
  width = width ? width : '24px'
  height = height ? height : '24px'
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 15 14"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M6.52589 11.2V10.163H4.67761L6.99439 7.82963L6.26846 7.09852L3.95168 9.43185V7.57037H2.922L2.922 11.2H6.52589ZM8.84267 5.96815L11.1595 3.63481V5.4963H12.1891V1.86667L8.58525 1.86667V2.9037H10.4335L8.11674 5.23704L8.84267 5.96815Z"
        fill="#515151"
      />
    </svg>
  )
}

export default ExpandIcon

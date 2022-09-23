// getTimeLeft returns how many days, months and years are left
function getTimeLeft(expirationDate: number) {
  const date = new Date(expirationDate)
  const yearsLeft = date.getFullYear() - new Date().getFullYear()
  const daysLeft = date.getDate() - new Date().getDate()
  const monthsLeft = date.getMonth() - new Date().getMonth()

  return `${daysLeft > 0 ? `${daysLeft} days` : '-'} | ${
    monthsLeft > 0 ? `${monthsLeft} months` : '-'
  } | ${yearsLeft > 0 ? `${yearsLeft} years` : '-'}`
}

export default getTimeLeft

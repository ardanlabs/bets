// getTimeLeft returns how many days, months and years are left
function getTimeLeft(expirationDate: string) {
  const yearsLeft =
    new Date(expirationDate).getFullYear() - new Date().getFullYear()
  const daysLeft = new Date(expirationDate).getDate() - new Date().getDate()
  const monthsLeft = new Date(expirationDate).getMonth() - new Date().getMonth()

  return `${daysLeft} days | ${monthsLeft} months | ${yearsLeft} years`
}

export default getTimeLeft

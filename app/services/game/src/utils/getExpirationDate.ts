// getExpirationDate returns a date in YYYY-MM-DD format
export default function getExpirationDate(expirationDate: number) {
  const date = new Date(expirationDate)
  const YYYY = date.getFullYear()
  let MM = (date.getMonth() + 1).toString()
  MM = MM.length <= 1 ? `0${MM}` : MM
  let DD = date.getDate().toString()
  DD = DD.length <= 1 ? `0${DD}` : DD

  return `${YYYY}-${MM}-${DD}`
}

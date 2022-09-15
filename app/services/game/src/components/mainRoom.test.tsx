import React from 'react'
import { rest } from 'msw'
import { setupServer } from 'msw/node'
import { render, screen, waitFor } from '@testing-library/react'
import MainRoom from './mainRoom'
import { apiUrl } from '../utils/axiosConfig'
import { Bet } from '../types/index.d'

const bets: Bet[] = [
  {
    id: 1,
    description: 'In 2022 there will be 2000 electric cars accidents',
    terms: 'Has to be in the us.',
    name: 'Bruno',
    challengerAddress: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
    expirationDate: '20221231000000',
    amount: 30,
  },
]

const handlers = [
  // Get bets mock request
  rest.get(`http://${apiUrl}/bets`, (req, res, ctx) => {
    return res(ctx.json(bets))
  }),
  // Get bet mock request
  rest.get(`http://${apiUrl}/bet/${bets[0].id}`, (req, res, ctx) => {
    return res(ctx.json(bets[0]))
  }),
]

const server = setupServer(handlers[1])
beforeAll(() => {
  // Establish requests interception layer before all tests.
  server.listen()
})
afterEach(() => server.resetHandlers())
afterAll(() => {
  // Clean up after all tests are done, preventing this
  // interception layer from affecting irrelevant tests.
  server.close()
})

test('Gets bets from mocked API', async () => {
  render(<MainRoom />)

  const out = await waitFor(() => screen.getByRole('openBets'))
  const li = await waitFor(() => screen.getByRole('bet'))

  expect(out).toContainElement(li)
  expect(li).toHaveTextContent(bets[0].name)
})

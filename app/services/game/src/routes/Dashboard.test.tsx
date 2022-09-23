import React from 'react'
import { rest } from 'msw'
import { setupServer } from 'msw/node'
import { render, screen, waitFor } from '@testing-library/react'
import Dashboard from './Dashboard'
import { apiUrl } from '../utils/axiosConfig'
import { Bet } from '../types/index.d'

const bets: Bet[] = [
  {
    id: '1',
    status: 'open',
    players: [
      { address: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC', signed: false },
      { address: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7', signed: false },
    ],
    moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
    description: 'In 2022 there will be 2000 electric cars accidents',
    terms: 'Has to be in the us.',
    expirationDate: new Date().getTime(),
    amount: 30,
  },
]

const resolver = jest.fn()
const handlers = [
  // Get bets mock request
  rest.get(`http://${apiUrl}/bets`, (_, res, ctx) => {
    return res(ctx.json(bets))
  }),
  // Get bet mock request
  rest.get(`http://${apiUrl}/bet/${bets[0].id}`, (_, res, ctx) => {
    return res(ctx.status(200), ctx.json(bets[0]))
  }),
  // Post bet mock request
  rest.post(`http://${apiUrl}/bet`, (req, res, ctx) => {
    resolver
    return res(ctx.status(200), ctx.json({ betId: bets[0].id }))
  }),
]

const server = setupServer(...handlers)
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
  render(<Dashboard />)

  const out = await waitFor(() => screen.getByRole('betId'))
  await waitFor(resolver)

  expect(out).toHaveTextContent('1')
  expect(resolver).toBeCalledTimes(1)
})

import 'isomorphic-fetch'

const getHeaders = () => ({
  Accept: 'application/json',
  'Content-Type': 'application/json',
})

const _fetch = method => (url, body) =>
  fetch(url, {
    method,
    headers: getHeaders(),
    body: JSON.stringify(body),
  })

export const post = _fetch('POST')

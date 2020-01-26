const API_URL = process.env.REACT_APP_API_URL || "http://localhost:3000"

const opt = {
  cache: "no-cache",
  mode: "cors",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json"
  }
}

export const post = async (body) => {
  const response = await fetch(`${API_URL}/secret`, {
    ...opt,
    method: "POST",
    body: JSON.stringify(body)
  })

  const { status } = response
  if (status === 404) {
    return { status }
  }

  const data = await response.json()
  return { status, data }
}

export const get = async (hash) => {
  const response = await fetch(`${API_URL}/secret/${hash}`, {
    ...opt,
    method: "GET"
  })

  const { status } = response
  if (status === 404) {
    return { status }
  }

  const data = await response.json()
  return { status, data }
}

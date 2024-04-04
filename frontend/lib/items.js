
let SERVER_URL = `http://host.docker.internal:${process.env.SERVER_PORT}`

if (!process.env.SERVER_PORT) {
  SERVER_URL = "localhost:5006"
}

console.log(SERVER_URL)

export async function getFeaturedReview() {
  const reviews = await getGmarketItems()
  return reviews[0]
}

export async function getGmarketItems(pageSize) {
  const response = await fetchGmarketItems()

  return response
}

export async function getTmonItems(pageSize) {
  const response = await fetchTmonItems()

  return response
}

export async function getWeMakePriceItems(pageSize) {
  const response = await fetchWeMakePriceItems()

  return response
}

export async function getElevenStItems(pageSize) {
  const response = await fetchElevenStItems()

  return response
}

async function fetchGmarketItems() {
  const url = `${SERVER_URL}/gmarket/items`

  const response = await fetch(url)

  if (!response.ok) {
    throw new Error(`Server returned ${response.status} for ${url}`)
  }

  return await response.json()
}

async function fetchTmonItems() {
  const url = `${SERVER_URL}/tmon/items`

  const response = await fetch(url)

  if (!response.ok) {
    throw new Error(`Server returned ${response.status} for ${url}`)
  }

  return await response.json()
}

async function fetchElevenStItems() {
  const url = `${SERVER_URL}/11st/items`

  const response = await fetch(url)

  if (!response.ok) {
    throw new Error(`Server returned ${response.status} for ${url}`)
  }

  return await response.json()
}

async function fetchWeMakePriceItems() {
  const url = `${SERVER_URL}/wemakeprice/items`

  const response = await fetch(url)

  if (!response.ok) {
    throw new Error(`Server returned ${response.status} for ${url}`)
  }

  return await response.json()
}


export async function getSlugs() {
  const { data } = await fetchGmarketItems({
    fields: ['slug'],
    pagination: { pageSize: 100 },
    sort: ['publishedAt:desc']
  })

  return data.map(({ attributes }) => {
    return {
      slug: attributes.slug,
    }
  })
}

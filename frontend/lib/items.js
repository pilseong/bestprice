let SERVER_URL = process.env.SERVER_URL

if (!process.env.SERVER_URL) {
  SERVER_URL = "http://host.docker.internal:5006"
}

// console.log(SERVER_URL)

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

export async function getDealItems(slug) {
  console.log("getdeal", slug)
  switch (slug) {
    case 'gmarket':
      return await getGmarketItems()
    case 'wemakeprice':
      return await getWeMakePriceItems()
    case '11st':
      return await getElevenStItems()
    case 'tmon':
      console.log("ttttmon")
      return await getTmonItems()
    default:
  }
}

export async function getFrontItems() {
  console.log("getFrontItems")


  const url = `${SERVER_URL}/front/items`
  console.log(url)

  const response = await fetch(url)

  if (!response.ok) {
    throw new Error(`Server returned ${response.status} for ${url}`)
  }

  return await response.json()
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

async function fetchElevenStItems(pageNo, pageSize) {
  const params = new URLSearchParams({
    pageNo: pageNo || 1,
    pageSize: pageSize || 200
  })

  const url = `${SERVER_URL}/11st/items?${params.toString()}`

  console.log(url)


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

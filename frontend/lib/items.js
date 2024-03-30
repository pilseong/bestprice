
const SERVER_URL = 'http://localhost:5006'

export async function getFeaturedReview() {
  const reviews = await getGmarketItems()
  return reviews[0]
}

export async function getReview(slug) {
  const { data } = await fetchGmarketItems({
    filters: { slug: { $eq: slug } },
    fields: ['slug', 'title', 'subtitle', 'publishedAt', 'body'],
    populate: { image: { fields: ['url'] } },
    pagination: { pageSize: 1, withCount: false },
  })


  if (data.length === 0) return null

  const { attributes: { title, subtitle, publishedAt, image, body } } = data[0]

  return {
    slug,
    title,
    subtitle,
    date: publishedAt.slice(0, 'yyyy-mm-dd'.length),
    image: `${SERVER_URL}${image.data.attributes.url}`,
    body: marked(body)
  }
}

export async function getGmarketItems(pageSize) {
  const response = await fetchGmarketItems()

  return response
}

export async function getTmonItems(pageSize) {
  const response = await fetchTmonItems()

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

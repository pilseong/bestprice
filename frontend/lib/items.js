
const SERVER_URL = 'http://localhost:5006'

export async function getFeaturedReview() {
  const reviews = await getItems()
  return reviews[0]
}

export async function getReview(slug) {
  const { data } = await fetchItems({
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

export async function getItems(pageSize) {
  const response = await fetchItems()

  return response
}

async function fetchItems() {
  const url = `${SERVER_URL}/gmarket/items`

  const response = await fetch(url)

  if (!response.ok) {
    throw new Error(`Server returned ${response.status} for ${url}`)
  }

  return await response.json()
}


export async function getSlugs() {
  const { data } = await fetchItems({
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

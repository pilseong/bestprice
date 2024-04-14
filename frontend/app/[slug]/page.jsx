import React from 'react'
import Heading from "@/components/Heading"
import Link from "next/link"
// import { getFeaturedReview, getItems } from "@/lib/reviews"
import Image from "next/image"
import { getDealItems } from "@/lib/items"

// export const dynamic = 'force-dynamic'
export const revalidate = 600

async function ShoppingMallPage({ params: { slug } }) {

  const items = await getDealItems(slug)
  // const featuredReview = await getFeaturedReview()
  // const items = await getGmarketItems()


  return (
    <>
      <div className="mb-10 text-center">
        <div className="flex flex-row justify-center my-2">
          <div>
            <a href="https://link.coupang.com/a/bxObvY" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://image14.coupangcdn.com/image/affiliate/banner/3a8d95746b930663fa4df8fdc9608e1d@2x.jpg" alt="웨스트우드 남성 로고 그래픽 반팔 라운드 티셔츠_WL2MCCT759" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bxOcO4" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://image14.coupangcdn.com/image/affiliate/banner/cc30c536e023ae5cbf2e54ced4c0344a@2x.jpg" alt="아이더 Safety 친환경 원단 기능성 티셔츠" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bxOfhS" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://image13.coupangcdn.com/image/affiliate/banner/044bb9eb32ba0d92d42b9e4808eed892@2x.jpg" alt="맥심 카누 라떼, 13.5g, 50개입, 2개" width="240" height="240" /></a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bxOgrd" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://image8.coupangcdn.com/image/affiliate/banner/c1b8a1c75c59ead28909e47ccc1c17c1@2x.jpg" alt="앙블랑 세이프 인디핑크 아기 물티슈 캡형 63평량, 20매, 12개" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bxOgUI" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://image3.coupangcdn.com/image/affiliate/banner/065e614d0fc31caec6f6e52790450c0f@2x.jpg" alt="지리산 물하나, 2L, 12개" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bxOho7" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://image14.coupangcdn.com/image/affiliate/banner/891aeba62995293211fd9945352c0197@2x.jpg" alt="앱솔루트 명작 2FL 분유 2단계, 800g, 2개" width="240" height="240" /></a>
          </div>
        </div>
        <h2 className="text-xs sm:text-base">위의 포스팅은 쿠팡 파트너스 활동의 일환으로, 이에 따른 일정액의 수수료를 제공받습니다.</h2>
      </div>
      <div className="flex flex-row flex-wrap gap-4 justify-center">
        {
          items && items.map((item, index) => (
            <>
              <div key={item.id} className="list-none mb-4 border w-72 bg-white shadow rounded hover:shadow-xl flex flex-col">
                <a href={item.linkUrl} target="_blank">
                  <div>
                    <Image src={item.imageUrl} alt=""
                      width={320} height={180} className="rounded-t" />
                    <h2 className="p-2 border-t text-sm font-semibold font-notosanskr">{item.title}</h2>
                  </div>
                </a>
                <div className="mt-auto">
                  <div className="flex justify-around items-center bg-pink-100 text-gray-900 pt-1">
                    <div className="font-exo2">
                      <span className="text-red-900">할인가 {item.afterPrice.toLocaleString()}</span>{' '}
                      <span className="text-sm">
                        <span className="line-through text-sm">{item.beforePrice.toLocaleString()}</span> ({item.discountRate}% off)
                      </span>
                    </div>
                  </div>
                  <div className="flex justify-around items-center bg-pink-100 text-gray-900 pb-1 text-sm font-exo2">
                    <div><span className="text-gray-700">{item.shippingInfo}</span></div>
                  </div>
                </div>
              </div>
            </>
          ))
        }
      </div >
    </>
  )
}

export default ShoppingMallPage
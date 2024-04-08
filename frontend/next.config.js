// /** @type {import('next').NextConfig} */
// const nextConfig = {

// };



// export default nextConfig;modu
// images: {
//   // unoptimized: true
//   remotePatterns: [
//     {
//       protocol: 'https',
//       hostname: 'gdimg.gmarket.co.kr',
//       port: '',
//       pathname: '/**'
//     }
//   ]
// }

module.exports = {
  output: 'standalone',
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'gdimg.gmarket.co.kr',
        port: '',
        pathname: '/**'
      },
      {
        protocol: 'http',
        hostname: '*.tmon.kr',
        port: '',
        pathname: '/**'
      },
      {
        protocol: 'http',
        hostname: '*.ticketmonster.co.kr',
        port: '',
        pathname: '/**'
      },
      {
        protocol: 'https',
        hostname: '*.wemep.co.kr',
        port: '',
        pathname: '/**'
      },
      {
        protocol: 'https',
        hostname: '*.wemakeprice.com',
        port: '',
        pathname: '/**'
      },
      {
        protocol: 'https',
        hostname: '*.alicdn.com',
        port: '',
        pathname: '/**'
      },
      {
        protocol: 'https',
        hostname: '*.011st.com',
        port: '',
        pathname: '/**'
      },
      {
        protocol: 'https',
        hostname: '*.coupangcdn.com',
        port: '',
        pathname: '/**'
      }
    ]
  },
}
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
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'gdimg.gmarket.co.kr',
        port: '',
        pathname: '/**'
      }
    ]
  },
}
/** @type {import('next').NextConfig} */

module.exports = {
  compiler:{
    removeConsole: true,
  },
  reactStrictMode: true,
  async redirects() {
    return [
      {
        source: '/problem',
        destination: '/',
        permanent: false,
      },
    ]
  },
}

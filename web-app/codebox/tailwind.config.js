/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",

    // Or if using `src` directory:
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        purple : '#8870FF',
        darkBlue : '#3A3457',
        dullWhite: '#F0F1F5',
        hard: '#FF7F7E',
        unAttemptedStrip: '#5F88FE',
        AttemptedStrip:'#FEB721',
        Accepted: '#2EB65E',
        Rejected: '#ef4743',
        AcceptedStrip:'#2BE2D0',
        gray: {
          900: '#202225',
          800: '#2f3136',
          700: '#36393f',
          600: '#4f545c',
          400: '#d4d7dc',
          300: '#e3e5e8',
          200: '#ebedef',
          100: '#f2f3f5',
        },
      },
      spacing: {
        88: '22rem',
      },
    },
    fontFamily: {
      nunito: ["Nunito Sans","sans-serif"]
    }
  },
  plugins: [
    require('@tailwindcss/line-clamp'),
  ],
};
import { defineConfig } from 'cypress'

export default defineConfig({
  chromeWebSecurity: false,
  experimentalFetchPolyfill: true,
  screenshotOnRunFailure: false,
  video: false,
  e2e: {
    baseUrl: 'http://localhost:8444',
  },
})

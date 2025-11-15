/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    turbo: {
      rules: {
        // Disable Turbopack for tesseract
        '*.ts': { loaders: ['default'], environment: 'node' },
      },
    },
  },
  webpack(config) {
    config.resolve.alias['tesseract.js'] = require.resolve('tesseract.js');
    return config;
  },
};

export default nextConfig;
